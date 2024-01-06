package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	aiartv1 "github.com/ydssx/morphix/api/aiart/v1"
	"github.com/ydssx/morphix/app/aiart/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewHTTPServer(bootstrap *conf.Bootstrap, artService *service.ArtService) *http.Server {
	srv := common.NewHTTPServer(bootstrap.ServiceSet.Aiart.Server)
	aiartv1.RegisterArtServiceHTTPServer(srv, artService)
	return srv
}
