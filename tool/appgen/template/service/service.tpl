package service

import (
	"context"

	{{.serviceInfo.PkgName}} "{{.serviceInfo.PkgPath}}"
	"{{.module}}/app/{{.appName}}/internal/biz"
	{{range .serviceInfo.Pkgs}}"{{.}}"
	{{end }}
)

type {{.serviceInfo.Name}} struct {
	uc *biz.{{.appName | Title}}UseCase

	{{.serviceInfo.PkgName}}.Unimplemented{{.serviceInfo.Name}}Server
}

func New{{.serviceInfo.Name | Title}}(uc *biz.{{.appName | Title}}UseCase) *{{.serviceInfo.Name}} {
	return &{{.serviceInfo.Name}}{uc: uc}
}
{{$serviceName := .serviceInfo.Name}}
{{range .serviceInfo.RpcMeths}}
{{if .Comment}}//{{.Comment}}{{end}}
func (s *{{$serviceName}}) {{.MethName}}(ctx context.Context,req *{{.Param}}) (res *{{.Return}}, err error) {
	return s.uc.{{.MethName}}(ctx, req)
}
{{end -}}
