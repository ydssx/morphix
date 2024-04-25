package llm

import (
	"context"
	"errors"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
)

type ModelName string

const (
	Llama3  ModelName = "llama3:instruct"
	Mistral ModelName = "mistral:instruct"
)

type LLM struct {
	model llms.Model
}

func New(model ModelName) *LLM {
	var llm llms.Model
	var err error
	switch model {
	case Llama3, Mistral:
		llm, err = ollama.New(ollama.WithModel(string(model)), ollama.WithSystemPrompt("用简体中文回答所有问题"))
	}
	if err != nil {
		log.Fatal(err)
	}

	return &LLM{model: llm}
}

// GenerateFromText 使用给定的提示生成文本。
// 它接受一个上下文和一个或多个选项来配置生成过程。
// 返回生成的文本和任何错误。
func (o *LLM) GenerateFromText(ctx context.Context, prompt string, options ...chains.ChainCallOption) (string, error) {
	pt := prompts.NewPromptTemplate("用简体中文回答以下问题，要尽可能详细完整：{{.input}}", []string{"input"})
	llmChain := chains.NewLLMChain(o.model, pt)

	return chains.Run(ctx, llmChain, prompt, options...)
}

// GenerateFromContent 使用给定的提示生成文本。
// 它接受一个上下文和一个或多个选项来配置生成过程。
// 返回生成的文本和任何错误。
func (o *LLM) GenerateFromContent(ctx context.Context, prompt, imgUrl string, options ...llms.CallOption) (string, error) {
	content := []llms.MessageContent{
		{Role: llms.ChatMessageTypeHuman, Parts: []llms.ContentPart{llms.TextContent{Text: prompt}}},
		{Role: llms.ChatMessageTypeHuman, Parts: []llms.ContentPart{llms.ImageURLContent{imgUrl}}},
		{Role: llms.ChatMessageTypeSystem, Parts: []llms.ContentPart{llms.TextContent{"请用简体中文回答所有问题。"}}},
	}
	resp, err := o.model.GenerateContent(ctx, content, options...)
	if err != nil {
		return "", err
	}

	choices := resp.Choices
	if len(choices) < 1 {
		return "", errors.New("empty response from model")
	}
	c1 := choices[0]
	return c1.Content, nil
}
