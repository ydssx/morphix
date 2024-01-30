package llm

import (
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type LLM struct{}

func New() *LLM {
	return &LLM{}
}

func (o *LLM) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	opts := llms.CallOptions{}
	for _, opt := range options {
		opt(&opts)
	}
	llm, err := openai.New()
	if err != nil {
		return "", err
	}
	llm.Call(ctx, prompt, options...)
	return "", nil
}

func (o *LLM) Name() string {
	return "openai"
}
