package biz

import (
	"context"

	{{.serviceInfo.PkgName}} "{{.serviceInfo.PkgPath}}"
    {{range .serviceInfo.Pkgs}}"{{.}}"
	{{end }}
)

type {{.appName | Title}}UseCase struct {
}

func New{{.appName | Title}}UseCase() *{{.appName | Title}}UseCase {
	return &{{.appName | Title}}UseCase{}
}
{{$appName := .appName}}
{{range $m := .serviceInfo.RpcMeths}}
{{if .Comment}}//{{$m.Comment}}{{end}}
func (b *{{$appName| Title}}UseCase) {{$m.MethName}}(ctx context.Context, req *{{$m.Param}}) (res *{{$m.Return}}, err error) {
	res = new({{$m.Return}})

	// TODO:ADD logic here and delete this line.

	return
}
{{end -}}
