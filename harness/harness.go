package harness

import (
	"context"


	"google.golang.org/genai"
	"fmt"
)

type Harness struct {
	Client *genai.Client
	Model string

}

func New(client *genai.Client, model string) *Harness {
	return &Harness{
		Client: client,
		Model: model,
	}
}

func (h *Harness) AskLLM(ctx context.Context, prompt string) (string, error) {
	response, err := h.Client.Models.GenerateContent(
		ctx,
		h.Model,
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			SystemInstruction: &genai.Content{
				Parts: []*genai.Part{
					{
						Text: systemPrompt,
					},
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("LLM error: %w", err)
	}

	return response.Text(), nil
}
