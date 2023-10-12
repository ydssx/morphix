package service

import (
	"context"

	{{.PkgName}} "{{.PkgPath}}"
	"github.com/ydssx/morphix/app/{{.appName}}/internal/biz"
	{{range .Imports}}"{{.}}"
	{{end }}
)

type {{.serviceName}} struct {
	uc *biz.{{.appName | Title}}UseCase

	{{.PkgName}}.Unimplemented{{.serviceName}}Server
}

func New{{.serviceName | Title}}(uc *biz.{{.appName | Title}}UseCase) *{{.serviceName}} {
	return &{{.serviceName}}{uc: uc}
}
{{range .RpcMeths}}
{{if .Comment}}//{{.Comment}}{{end}}
func (s *{{.ServiceName}}) {{.MethName}}(ctx context.Context,req *{{.Param}}) (res *{{.PkgName}}.{{.Return}}, err error) {
	return s.uc.{{.MethName}}(ctx, req)
}
{{end -}}
