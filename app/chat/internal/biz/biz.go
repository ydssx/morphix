package biz

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/pkg/llm"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewChatUseCase, NewLlm)

func NewLlm() *llm.LLM {
	return llm.New(llm.Llama3)
}
