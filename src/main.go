package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

const (
	AIFace     = "🤖"
	ArrowRight = "➡️"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Provide command for OpenAI \n %s ", ArrowRight)
	text, _ := reader.ReadString('\n')
	processAICommand(text)
}

// processAICommand sends the command to OpenAI and prints the response
func processAICommand(prompt string) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	// or
	//apiKey := "sk-"
	// for local-only deployment
	aiClient := openai.NewClient(apiKey)
	resp, err := aiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	if len(resp.Choices) > 0 && resp.Choices[0].Message.Content != "" {
		fmt.Printf("%s: %s", AIFace, resp.Choices[0].Message.Content)
	} else {
		fmt.Println("No response from AI")
	}
}
