package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ydssx/morphix/app/chat/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewHTTPServer(c *conf.Bootstrap, svc *service.ChatService) *http.Server {
	srv := common.NewHTTPServer(c.ServiceSet.Chat.Server)

	return srv
}
