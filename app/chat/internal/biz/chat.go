package biz

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms"
	chat "github.com/ydssx/morphix/api/chat/v1"
	"github.com/ydssx/morphix/pkg/llm"
)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

type ChatUseCase struct {
	llm *llm.LLM
}

func NewChatUseCase(llm *llm.LLM) *ChatUseCase {
	return &ChatUseCase{llm: llm}
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

// 服务器到客户端的流，用于接收消息
func (uc *ChatUseCase) ReceiveMessage(req *chat.ClientMessage, stream chat.ChatService_ReceiveMessageServer) (err error) {
	resp, err := uc.llm.Generate(
		stream.Context(),
		req.MessageText,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			return stream.Send(&chat.ServerMessage{MessageText: string(chunk)})
		}),
	)
	log.Printf("resp: %s", resp)
	
	return err
}
