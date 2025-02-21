package chatgpt

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient(ChatGPT_gateway, apiKey string) *OpenAIClient {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = ChatGPT_gateway
	return &OpenAIClient{
		client: openai.NewClientWithConfig(config),
	}
}

func (client *OpenAIClient) SendMessage(messages []openai.ChatMessagePart, model string) (string, error) {
	req := openai.ChatCompletionRequest{
		TopP:        1.0,
		Temperature: 0.5,
		Model:       model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:         openai.ChatMessageRoleUser,
				MultiContent: messages,
			},
		},
		Stream: false,
	}

	resp, err := client.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
