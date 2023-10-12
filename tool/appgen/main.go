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
	"strings"
	"text/template"

	"github.com/emicklei/proto"
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
	appName   = flag.String("app", "", "")
	protoFile = flag.String("proto", "", "proto file")
)

func main() {
	flag.Parse()
	if *appName == "" || *protoFile == "" {
		log.Fatal("app and proto must be set.")
	}
	gen(*appName, *protoFile)
}

type Generator struct {
}

func gen(appName, protoFile string) {
	baseDir := "app/" + appName
	cmdDir := baseDir + "/cmd/" + appName
	internalDir := baseDir + "/internal"
	serverDir := internalDir + "/server"
	serviceDir := internalDir + "/service"
	bizDir := internalDir + "/biz"
	paths := []string{baseDir, cmdDir, internalDir, serverDir, bizDir, serviceDir}
	for _, v := range paths {
		err := os.MkdirAll(v, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	serviceInfo := parseProto(protoFile, appName)
	data := map[string]interface{}{
		"port":        9005,
		"appName":     appName,
		"serviceName": serviceInfo.Name,
		"RpcMeths":    serviceInfo.RpcMeths,
		"Imports":     serviceInfo.pkgs,
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

	cmd := exec.Command("sh", "-c", fmt.Sprintf("make api && cd app/%s/cmd/%s/ && wire", appName, appName))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func mkFile(data map[string]interface{}, outFile string, text string) {
	funcs := template.FuncMap{"Title": strings.Title}
	tmpl, err := template.New("tmp").Funcs(funcs).Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}
	var codes = buf.Bytes()
	if strings.HasSuffix(outFile, ".go") {
		codes, _ = format.Source(buf.Bytes())
	}
	err = os.WriteFile(outFile, codes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("generate file [%s] succeed.\n", outFile)
}

type ServiceInfo struct {
	Name     string
	pkgs     []string
	RpcMeths []MethInfo
}

type MethInfo struct {
	MethName    string
	Param       string
	Return      string
	Comment     string
	ServiceName string
	AppName     string
}

func parseProto(protoFile, appName string) (info ServiceInfo) {
	reader, err := os.Open(protoFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	proto.Walk(definition,
		proto.WithService(func(s *proto.Service) { info.Name = s.Name }),
		proto.WithRPC(func(r *proto.RPC) {
			req, pkg := convertRequest(r.RequestType)
			info.RpcMeths = append(info.RpcMeths, MethInfo{
				MethName:    r.Name,
				Param:       req,
				Return:      r.ReturnsType,
				Comment:     r.Comment.Message(),
				ServiceName: info.Name,
				AppName:     appName,
			})
			if !util.SliceContain(info.pkgs, pkg) && pkg != "" {
				info.pkgs = append(info.pkgs, pkg)
			}
		}),
	)
	return
}

func convertRequest(reqType string) (string, string) {
	switch reqType {
	case "google.protobuf.Empty":
		return "emptypb.Empty", "google.golang.org/protobuf/types/known/emptypb"
	default:
		return "pb." + reqType, ""
	}
}
