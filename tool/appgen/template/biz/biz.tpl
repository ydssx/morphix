package biz

import (
	"context"

	{{.PkgName}} "{{.PkgPath}}"
    {{range .Imports}}"{{.}}"
	{{end }}
)

type {{.appName | Title}}UseCase struct {
}

func New{{.appName | Title}}UseCase() *{{.appName | Title}}UseCase {
	return &{{.appName | Title}}UseCase{}
}
{{range $m := .RpcMeths}}
{{if .Comment}}//{{$m.Comment}}{{end}}
func (b *{{$m.AppName | Title}}UseCase) {{$m.MethName}}(ctx context.Context, req *{{$m.Param}}) (res *{{$m.Return}}, err error) {
	res = new({{$m.Return}})

	// TODO:ADD logic here and delete this line.

	return
}
{{end -}}

