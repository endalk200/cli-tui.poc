package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/endalk200/bgit/internal/config"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type ErrAPIKeyNotFound struct {
	Code    int
	Message string
}

func (e ErrAPIKeyNotFound) Error() string {
	return fmt.Sprintf("API key not found: %d %s", e.Code, e.Message)
}

type ErrAIProviderCallFailed struct {
	Code    int
	Message string
}

func (e ErrAIProviderCallFailed) Error() string {
	return fmt.Sprintf("AI provider call failed: %d %s", e.Code, e.Message)
}

type ErrUnkownAIProvider struct {
	Code    int
	Message string
}

func (e ErrUnkownAIProvider) Error() string {
	return fmt.Sprintf("unknown AI provider: %d %s", e.Code, e.Message)
}

type ErrUnknownIssue struct {
	Code    int
	Message string
}

func (e ErrUnknownIssue) Error() string {
	return fmt.Sprintf("unknown issue: %d %s", e.Code, e.Message)
}

func GenerateCommitMessage(diff string, provider config.Provider) (string, error) {
	prompt := fmt.Sprintf("Generate a concise conventional commit style message summarizing changes made in this git diff. \n%s", diff)

	switch provider.Name {
	case "OpenAI":
		API_KEY, err := getOpenAIAPIKey(provider.EnvName)
		if err != nil {
			return "", err
		}

		commitMessage, err := OpenAIChatCompletion(prompt, API_KEY)
		if err != nil {
			return "", err
		}

		return commitMessage, nil
	case "OpenRouter":
		API_KEY, err := getOpenRouterAPIKey(provider.EnvName)
		if err != nil {
			return "", err
		}

		commitMessage, err := OpenRouterChatCompletion(prompt, API_KEY)
		if err != nil {
			return "", err
		}

		return commitMessage, nil
	default:
		return "", ErrUnkownAIProvider{
			Code:    400,
			Message: provider.Name + " is not a valid AI provider",
		}
	}
}

func OpenAIChatCompletion(prompt string, API_KEY string) (string, error) {
	context := context.Background()

	client := openai.NewClient(option.WithAPIKey(API_KEY))
	response, err := client.Chat.Completions.New(context, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT5Mini,
	})
	if err != nil {
		return "", ErrAIProviderCallFailed{
			Code:    500,
			Message: err.Error(),
		}
	}

	if len(response.Choices) == 0 || response.Choices[0].Message.Content == "" {
		return "", ErrAIProviderCallFailed{
			Code:    500,
			Message: "no AI response content",
		}
	}
	return strings.TrimSpace(response.Choices[0].Message.Content), nil
}

func OpenRouterChatCompletion(prompt string, API_KEY string) (string, error) {
	context := context.Background()

	header := http.Header{}
	header.Set("X-Title", "bgit")

	client := openai.NewClient(
		option.WithAPIKey(API_KEY),
		option.WithBaseURL("https://openrouter.ai/api/v1"),
	)
	response, err := client.Chat.Completions.New(context, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT5Mini,
	})
	if err != nil {
		return "", ErrAIProviderCallFailed{
			Code:    500,
			Message: err.Error(),
		}
	}

	if len(response.Choices) == 0 || response.Choices[0].Message.Content == "" {
		return "", ErrAIProviderCallFailed{
			Code:    500,
			Message: "no AI response content",
		}
	}
	return strings.TrimSpace(response.Choices[0].Message.Content), nil
}

func AntropicChatCompletion(prompt string, API_KEY string) (string, error) {
	return "", nil
}

func getOpenAIAPIKey(keyName string) (string, error) {
	OPENAI_API_KEY, exists := os.LookupEnv(keyName)
	if exists {
		return OPENAI_API_KEY, nil
	}

	return "", ErrAPIKeyNotFound{
		Code:    401,
		Message: keyName + " not set",
	}
}

func getOpenRouterAPIKey(keyName string) (string, error) {
	OPENROUTER_API_KEY, exists := os.LookupEnv(keyName)
	if exists {
		return OPENROUTER_API_KEY, nil
	}

	return "", ErrAPIKeyNotFound{
		Code:    401,
		Message: keyName + " not set",
	}
}
