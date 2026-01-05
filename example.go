package main

import (
	"context"
	"fmt"
	// "log"
	// "os"
	// "net/http"
	// "net/url"
	openai "github.com/sashabaranov/go-openai"
)

func main() {

	apiKey := "ollama"
	config := openai.DefaultConfig(apiKey) 
	config.BaseURL = "http://10.88.88.180:11434/v1"

	// Optional: Configure a proxy if needed
	// proxyUrl, err := url.Parse("")
	// if err == nil {
	// 	config.HTTPClient = &http.Client{
	// 		Transport: &http.Transport{
	// 			Proxy: http.ProxyURL(proxyUrl),
	// 		},
	// 	}
	// }

	client := openai.NewClientWithConfig(config)

	// client := openai.NewClient(os.Getenv("OPENAI_API_KEY"), "http://10.88.88.180:11434")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// Model: openai.GPT3Dot5Turbo,
			Model: "gpt-oss",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

