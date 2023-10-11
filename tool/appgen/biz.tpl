package biz

import (
	"context"

	pb "github.com/ydssx/morphix/api/{{.appName}}/v1"
    {{range .Imports}}"{{.}}"
	{{end }}
)

type {{.appName | Title}}UseCase struct {
}

func New{{.appName | Title}}UseCase() *{{.appName | Title}}UseCase {
	return &{{.appName | Title}}UseCase{}
}
{{range $m := .RpcMeths}}
//{{$m.Comment}}
func (*{{$m.AppName | Title}}UseCase) {{$m.MethName}}(ctx context.Context, req *{{$m.Param}}) (res *pb.{{$m.Return}}, err error) {
	res = new(pb.{{$m.Return}})

	// TODO:ADD logic here and delete this line.

	return
}
{{end -}}

