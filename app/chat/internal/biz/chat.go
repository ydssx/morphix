package biz

import (
	"context"

	chat "github.com/ydssx/morphix/api/chat/v1"
)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

type ChatUseCase struct{}

func NewChatUseCase() *ChatUseCase {
	return &ChatUseCase{}
}

// 客户端到服务器的流，用于发送消息
func (uc *ChatUseCase) SendMessage(stream chat.ChatService_SendMessageServer) (err error) {
	resp := new(chat.ServerMessage)

	// TODO:ADD logic here and delete this line.

	err = stream.SendAndClose(resp)

	return
}

// 双向流
func (uc *ChatUseCase) Chat(stream chat.ChatService_ChatServer) (err error) {
	// TODO:ADD logic here and delete this line.

	return
}
