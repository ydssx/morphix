package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	chat "github.com/ydssx/morphix/api/chat/v1"
	"github.com/ydssx/morphix/app/chat/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, svc *service.ChatService) *grpc.Server {
	srv := common.NewGRPCServer(c.ServiceSet.Chat.Server)

	chat.RegisterChatServiceServer(srv, svc)

	return srv
}
