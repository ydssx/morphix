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
{{$pkgName := .serviceInfo.PkgName}}
{{$svcName := .serviceInfo.Name}}
{{range $m := .serviceInfo.RpcMeths}}
{{if .Comment}}//{{$m.Comment}}{{end -}}
{{if and .StreamsReturns .StreamsRequest}}
func (uc *{{$appName| Title}}UseCase) {{$m.MethName}}(stream {{$pkgName}}.{{$svcName}}_{{$m.MethName}}Server) (err error) {
	// TODO:ADD logic here and delete this line.

	return
}
{{else if .StreamsReturns }}
func (uc *{{$appName| Title}}UseCase) {{$m.MethName}}(req *{{$m.Param}}, stream {{$pkgName}}.{{$svcName}}_{{$m.MethName}}Server) (err error) {
	res := new({{$m.Return}})

	// TODO:ADD logic here and delete this line.

	err = stream.Send(res)
	if err != nil {
		return err
	}

	return
}
{{else if .StreamsRequest}}
func (uc *{{$appName | Title}}UseCase) {{$m.MethName}}(stream {{$pkgName}}.{{$svcName}}_{{$m.MethName}}Server) (err error) {
	resp := new(aiartv1.GetGeneratedImageResponse)

	// TODO:ADD logic here and delete this line.

	err = stream.SendAndClose(resp)
	
	return
}
{{else}}
func (uc *{{$appName | Title}}UseCase) {{$m.MethName}}(ctx context.Context, req *{{$m.Param}}) (res *{{$m.Return}}, err error) {
	res = new({{$m.Return}})

	// TODO:ADD logic here and delete this line.

	return
}
{{end -}}
{{end -}}
