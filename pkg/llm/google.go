package llm

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AI interface {
	TextToText(ctx context.Context, prompt string) (string, error)
	TextToImage(ctx context.Context, prompt string) (string, error)
	ImageToText(ctx context.Context, image string) (string, error)
	ImageToImage(ctx context.Context, prompt string) (string, error)
}

// GoogleModel defines the available Google AI models.
type GoogleModel string

const (
	GoogleModel_GenAI        GoogleModel = "gemini-pro"
	GoogleModel_GenAI_Vision GoogleModel = "gemini-pro-vision"
)

type GoogleAI struct {
	*genai.Client
}

type Chat struct {
	session *genai.ChatSession
}

func init() {
	os.Setenv("GENAI_API_KEY", "AIzaSyAyLhOvzO69qBr5zav3DpobtznSxQPDnuI")
}

func NewGoogleAI() *GoogleAI {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(os.Getenv("GENAI_API_KEY")))
	if err != nil {
		panic(err)
	}
	return &GoogleAI{client}
}

func (g *GoogleAI) Close() error {
	return g.Client.Close()
}

func (g *GoogleAI) GenerativeModel(name GoogleModel) *genai.GenerativeModel {
	return g.Client.GenerativeModel(string(name))
}

func (g *GoogleAI) StartChat() *Chat {
	model := g.GenerativeModel(GoogleModel_GenAI)
	return &Chat{session: model.StartChat()}
}

func (c *Chat) TextToText(ctx context.Context, prompt ...string) (interface{}, error) {
	if len(prompt) == 0 {
		return "", nil
	}

	res, err := c.session.SendMessage(ctx, c.convertPrompt(prompt...)...)
	if err != nil {
		return "", err
	}
	return res.Candidates, nil
}

func (c *Chat) TextToImage(ctx context.Context, prompt string) (interface{}, error) {
	res, err := c.session.SendMessage(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	return res.Candidates, nil
}

func (c *Chat) convertPrompt(prompt ...string) []genai.Part {
	var texts []genai.Part
	for _, p := range prompt {
		texts = append(texts, genai.Text(p))
	}
	return texts
}
