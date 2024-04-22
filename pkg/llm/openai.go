package llm

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type ModelName string

const (
	Llama3 ModelName = "llama3"
)

type LLM struct {
	model llms.Model
}

func New(model ModelName) *LLM {
	llm, err := ollama.New(ollama.WithModel(string(model)))
	if err != nil {
		log.Fatal(err)
	}

	return &LLM{model: llm}
}

func (o *LLM) Generate(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	opts := llms.CallOptions{}
	for _, opt := range options {
		opt(&opts)
	}

	return llms.GenerateFromSinglePrompt(ctx, o.model, prompt, options...)
}

func (o *LLM) Name() string {
	return "openai"
}
