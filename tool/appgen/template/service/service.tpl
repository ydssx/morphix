package service

import (
	"context"

	pb "github.com/ydssx/morphix/api/{{.appName}}/v1"
	"github.com/ydssx/morphix/app/{{.appName}}/internal/biz"
	{{range .Imports}}"{{.}}"
	{{end }}
)

type {{.serviceName}} struct {
	uc *biz.{{.appName | Title}}UseCase

	pb.Unimplemented{{.serviceName}}Server
}

func New{{.serviceName | Title}}(uc *biz.{{.appName | Title}}UseCase) *{{.serviceName}} {
	return &{{.serviceName}}{uc: uc}
}
{{range .RpcMeths}}
//{{.Comment}}
func (s *{{.ServiceName}}) {{.MethName}}(ctx context.Context,req *{{.Param}}) (res *pb.{{.Return}}, err error) {
	return s.uc.{{.MethName}}(ctx, req)
}
{{end -}}
