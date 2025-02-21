package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// OpenAIClient 封装了与 OpenAI API 交互的逻辑
type OpenAIClient struct {
	apiKey string
	apiURL string
}

// Message 定义了消息的结构体
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest 用于构建发送给 OpenAI 的请求体
type ChatRequest struct {
	Model    string     `json:"model"`
	Messages []Message `json:"messages"`
}

// ChatResponseChoice 用于解析响应中的每一项
type ChatResponseChoice struct {
	Message Message `json:"message"`
}

// ChatResponse 是 API 返回的整体响应结构
type ChatResponse struct {
	Choices []ChatResponseChoice `json:"choices"`
}

// NewOpenAIClient 创建一个新的 OpenAIClient 实例
func NewOpenAIClient(ChatGPT_gateway,apiKey string) *OpenAIClient {
	return &OpenAIClient{
		apiKey: apiKey,
		apiURL: ChatGPT_gateway,
	}
}

// SendMessage 发送消息和图片链接给 OpenAI API，并返回响应
func (client *OpenAIClient) SendMessage(messages []Message, model string, imageURL string) (string, error) {
	// 构造请求体
	chatRequest := ChatRequest{
		Model:    model,
		Messages: messages,
	}

	// 如果有图片链接，添加到消息中
	if imageURL != "" {
		imageMessage := Message{
			Role: "system",
			Content: fmt.Sprintf(`{
				"type": "image_url",
				"image_url": {
					"url": "%s"
				}
			}`, imageURL),
		}

		chatRequest.Messages = append(chatRequest.Messages, imageMessage)
	}

	// 转换为 JSON 格式
	requestBody, err := json.Marshal(chatRequest)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}
	fmt.Println(string(requestBody))
	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", client.apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.apiKey)

	// 发送请求
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// 解析响应
	var chatResponse ChatResponse
	err = json.Unmarshal(respBody, &chatResponse)
	if err != nil {
		log.Println(string(respBody))
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	// 如果有回复，返回第一个回复内容
	if len(chatResponse.Choices) > 0 {
		return chatResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from ChatGPT")
}
