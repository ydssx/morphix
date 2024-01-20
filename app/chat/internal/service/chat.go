package service

import (
	"context"

	chat "github.com/ydssx/morphix/api/chat/v1"
	"github.com/ydssx/morphix/app/chat/internal/biz"
)

var _ = context.Background()

type ChatService struct {
	uc *biz.ChatUseCase

	chat.UnimplementedChatServiceServer
}

func NewChatService(uc *biz.ChatUseCase) *ChatService {
	return &ChatService{uc: uc}
}

// 客户端到服务器的流，用于发送消息
func (s *ChatService) SendMessage(stream chat.ChatService_SendMessageServer) (err error) {
	return s.uc.SendMessage(stream)
}

// 双向流，用于实现聊天
func (s *ChatService) Chat(stream chat.ChatService_ChatServer) (err error) {
	return s.uc.Chat(stream)
}
