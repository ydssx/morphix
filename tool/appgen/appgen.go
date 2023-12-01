package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"

	"github.com/cheggaaa/pb/v3"
	"github.com/emicklei/proto"
	"github.com/fatih/color"
	apigen "github.com/ydssx/api-gen/gen"
	"github.com/ydssx/morphix/pkg/util"
)

var (
	//go:embed template/dockerfile.tpl
	Dockerfile string

	//go:embed template/cmd/main.tpl
	mainFile string

	//go:embed template/cmd/wire.tpl
	wireFile string

	//go:embed template/server/server.tpl
	serverFile string

	//go:embed template/server/grpc.tpl
	grpcFile string

	//go:embed template/service/service.tpl
	serviceFile string

	//go:embed template/service/service_set.tpl
	serviceSetFile string

	//go:embed template/biz/biz.tpl
	bizFile string

	//go:embed template/biz/biz_set.tpl
	bizSetFile string
)

var (
	appName   = flag.String("app", "aiart", "app name")
	protoFile = flag.String("proto", "../../api/aiart/v1/aiart.proto", "proto file")
	port      = flag.Int("port", 0, "app service port")
)

func main() {
	flag.Parse()
	if *appName == "" || *protoFile == "" {
		log.Fatal("app and proto must be set.")
	}
	if *port == 0 {
		*port = util.GenerateRandomNumber(9000, 10000)
	}
	gen(*appName, *protoFile, *port)
}

func gen(appName, protoFile string, port int) {
	baseDir := "app/" + appName
	cmdDir := baseDir + "/cmd/" + appName
	internalDir := baseDir + "/internal"
	serverDir := internalDir + "/server"
	serviceDir := internalDir + "/service"
	bizDir := internalDir + "/biz"
	paths := []string{baseDir, cmdDir, internalDir, serverDir, bizDir, serviceDir}
	bar := pb.StartNew(len(paths))
	bar.SetMaxWidth(150)
	tmpl := `{{ red "generating:" }} {{ bar . "<" "▇" (cycle . "=" ) "." ">"}} {{speed . | rndcolor }} {{percent .}} {{string . "my_green_string" | green}}`
	bar.SetTemplateString(tmpl)
	bar.Set("my_green_string", "green")
	for _, v := range paths {
		bar.Set("my_green_string", v)
		bar.Increment()
		err := os.MkdirAll(v, 0o644)
		if err != nil {
			log.Fatal(err)
		}
	}
	bar.Finish()
	serviceInfo := parseProto(protoFile)
	data := map[string]interface{}{
		"port":        port,
		"appName":     appName,
		"serviceInfo": serviceInfo,
		"module":      parseGoModule(),
	}

	mkFile(data, baseDir+"/Dockerfile", Dockerfile)
	mkFile(data, cmdDir+"/main.go", mainFile)
	mkFile(data, cmdDir+"/wire.go", wireFile)
	mkFile(data, serverDir+"/server.go", serverFile)
	mkFile(data, serverDir+"/grpc.go", grpcFile)
	mkFile(data, serviceDir+"/service.go", serviceSetFile)
	mkFile(data, serviceDir+"/"+appName+".go", serviceFile)
	mkFile(data, bizDir+"/biz.go", bizSetFile)
	mkFile(data, bizDir+"/"+appName+".go", bizFile)

	x := fmt.Sprintf("make api && cd app/%s/cmd/%s/ && wire", appName, appName)
	cmd := exec.Command("cmd.exe", "/C", x)
	if runtime.GOOS != "windows" {
		cmd = exec.Command("sh", "-c", x)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(string(output))
	}
	fmt.Println(string(output))
}

// mkFile generates a file with the provided data, template, and file path.
// It checks if the file already exists and if the file extension is ".go".
// If the file exists and has the same name as the application,
// it writes the formatted code to the file using apigen.WriteDecl.
// If the file exists and does not have the same name as the application,
// it returns without doing anything.
// If the file does not exist, it writes the code to the file.
// It uses the "Title" function from the strings package as a custom template function.
// It returns an error if any error occurs during the process.
func mkFile(data map[string]interface{}, outFile string, text string) error {
	// Define custom template function
	funcs := template.FuncMap{"Title": strings.Title}

	// Parse the template
	tmpl, err := template.New("tmp").Funcs(funcs).Parse(text)
	if err != nil {
		return err
	}

	// Execute the template with the provided data
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return err
	}

	// Get the bytes of the generated code
	codes := buf.Bytes()

	// Format the code if the file extension is ".go"
	if strings.HasSuffix(outFile, ".go") {
		formattedCodes, err := format.Source(codes)
		if err != nil {
			return err
		}
		codes = formattedCodes
	}

	// Check if the file exists
	if fileExists(outFile) {
		// If the file has the same name as the application, write the code to the file
		if strings.HasSuffix(outFile, *appName+".go") {
			apigen.WriteDecl(outFile, string(codes))
		}
		return nil
	}

	// Write the code to the file
	err = os.WriteFile(outFile, codes, 0o644)
	if err != nil {
		return err
	}

	// Print success message
	color.Green("generate file [%s] succeed.\n", outFile)

	return nil
}

type ServiceInfo struct {
	Name     string
	Pkgs     []string
	RpcMeths []MethInfo
	PkgPath  string
	PkgName  string
}

type MethInfo struct {
	MethName       string
	Param          string
	Return         string
	Comment        string
	StreamsRequest bool
	StreamsReturns bool
}

//	parseProto parses a proto file and extracts information about the service, RPC methods, and package.
//
// It returns a ServiceInfo struct containing the extracted information.
func parseProto(protoFile string) (info ServiceInfo) {
	// Open the proto file
	reader, err := os.Open(protoFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	// Create a parser and parse the proto file
	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// Walk through the proto definition and extract information
	proto.Walk(definition,
		proto.WithPackage(func(p *proto.Package) {
			info.PkgName = p.Name
		}),
		proto.WithService(func(s *proto.Service) {
			info.Name = s.Name
		}),
		proto.WithRPC(func(r *proto.RPC) {
			// Convert the request type and return type
			req, pkg := convertRequest(info.PkgName, r.RequestType)
			returnType := convertReturnType(info.PkgName, r.ReturnsType)

			// Create a MethInfo struct and populate it with the extracted information
			x := MethInfo{
				MethName:       r.Name,
				Param:          req,
				Return:         returnType,
				StreamsRequest: r.StreamsRequest,
				StreamsReturns: r.StreamsReturns,
			}

			// Add the comment to the MethInfo struct if it exists
			if r.Comment != nil {
				x.Comment = strings.Join(r.Comment.Lines, "\n//")
			}

			// Add the MethInfo struct to the RpcMeths slice in the ServiceInfo struct
			info.RpcMeths = append(info.RpcMeths, x)

			// Add the package to the Pkgs slice in the ServiceInfo struct if it doesn't already exist
			if !util.SliceContain(info.Pkgs, pkg) && pkg != "" {
				info.Pkgs = append(info.Pkgs, pkg)
			}
		}),
		proto.WithOption(func(o *proto.Option) {
			// Check if the constant source exists
			if o.Constant.Source == "" {
				return
			}

			// Split the constant source and set the package path in the ServiceInfo struct
			x := strings.Split(o.Constant.Source, ";")
			if len(x) > 0 {
				info.PkgPath = x[0]
			} else {
				info.PkgPath = o.Constant.Source
			}
		}),
	)

	return
}

func convertRequest(pkgName, reqType string) (string, string) {
	switch reqType {
	case "google.protobuf.Empty":
		return "emptypb.Empty", "google.golang.org/protobuf/types/known/emptypb"
	default:
		return pkgName + "." + reqType, ""
	}
}

func convertReturnType(pkgName, returnType string) string {
	switch returnType {
	case "google.protobuf.Empty":
		return "emptypb.Empty"
	default:
		return pkgName + "." + returnType
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func parseGoModule() string {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	x := strings.Split(string(content), "\n")[0]
	return strings.TrimSpace(strings.Split(x, " ")[1])
}
