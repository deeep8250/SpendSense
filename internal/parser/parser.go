package parses

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/deeep8250/SpendSense/models"
	"github.com/sashabaranov/go-openai"
)

func AiParser(message string, categories []string) (string, error) {

	apiKey := os.Getenv("GROQ_API_KEY")
	fmt.Println("KEY:", apiKey)
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.groq.com/openai/v1"
	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "llama-3.3-70b-versatile",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
	})

	if err != nil {
		fmt.Println("Error :", err.Error())
		return "", err
	}
	fmt.Println(resp.Choices[0].Message.Content)

	prompt := fmt.Sprintf(`You are an expense parser. Extract expense details from the user's message.

Return a JSON object with these fields:
- amount: number
- merchant: string
- category: must be one from this list: [%s]
- description: short summary
- date: YYYY-MM-DD format

Today's date is %s.
If the user says "yesterday", "last Monday", etc., calculate the actual date.
If no date is mentioned, use today's date.
If the category is unclear, use "Other".

Example:
Input: "spent 250 at Starbucks for coffee yesterday"
Output: {"amount": 250, "merchant": "Starbucks", "category": "Dining", "description": "coffee", "date": "2026-04-03"}

Return ONLY valid JSON. No explanation, no markdown, no code blocks.`, strings.Join(categories, ", "), time.Now().Format("2006-01-02"))

	resp, err = client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "llama-3.3-70b-versatile",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: message,
			},
		},
	})

	if err != nil {
		fmt.Println("Error:", err.Error())
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func ValidateParsedExpense(values models.AiParserResponseHolder, categories []string) error {
	_, err := time.Parse("2006-01-02", values.Date)
	if err != nil {
		return errors.New("invalid date")
	}

	if values.Amount <= 0 {
		if values.Amount == 0 {
			return errors.New("amount should be greater than 0")
		}
		return fmt.Errorf("invalid amount format %v", values.Amount)
	}

	if !containsCategory(categories, values.Category) {
		return errors.New("invalid categories")
	}

	if values.Merchant == "" {
		return errors.New("please enter the valid merchant detail")
	}

	if values.Description == "" {
		return errors.New("please enter the valid description detail")
	}
	return nil
}

func containsCategory(categories []string, name string) bool {
	for _, c := range categories {
		if c == name {
			return true
		}
	}
	return false
}
