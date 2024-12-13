package services

import (
	"context"

	"github.com/ollama/ollama/api"
)

type ChatService struct {
	client *api.Client
}

func NewChatService() (*ChatService, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, err
	}
	return &ChatService{client: client}, nil
}

func (s *ChatService) ListModels(ctx context.Context) (*api.ListResponse, error) {
	return s.client.List(ctx)
}

func (s *ChatService) Generate(ctx context.Context, prompt string) (string, error) {
	req := api.GenerateRequest{
		Model:  "tinyllama",
		Prompt: prompt,
		Stream: new(bool), // false by default
	}

	var response string
	err := s.client.Generate(ctx, &req, func(resp api.GenerateResponse) error {
		response += resp.Response
		return nil
	})

	if err != nil {
		return "", err
	}

	return response, nil
} 