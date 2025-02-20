package goopenai

import (
	"context"
	"encoding/json"
)

type LlamaCppParams struct {
	CachePrompt *bool `json:"cache_prompt,omitempty"` // default: false, pass true to request caching.
	SlotId      *int  `json:"slot_id,omitempty"`      // default: 0, pass a specific slot id to use a cached prompt.
}

type CreateChatCompletionsRequest struct {
	Model            string            `json:"model,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             *float64          `json:"top_p,omitempty"`
	N                *int              `json:"n,omitempty"`
	Stream           *bool             `json:"stream,omitempty"`
	Stop             StrArray          `json:"stop,omitempty"`
	MaxTokens        *int              `json:"max_tokens,omitempty"`
	PresencePenalty  *float64          `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64          `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	ResponseFormat   *ResponseFormat   `json:"response_format,omitempty"`
	Seed             *int              `json:"seed,omitempty"`
	Tools            []Tools           `json:"tools,omitempty"`
	ToolChoice       *ToolChoice       `json:"tool_choice,omitempty"`
	User             *string           `json:"user,omitempty"`
	// The following struct is specific to the llama.cpp API and should be left empty for the OpenAI Chat API.
	CachePrompt *bool   `json:"cache_prompt,omitempty"` // default: false, pass true to request caching.
	SlotId      *int    `json:"slot_id,omitempty"`      // default: 0, pass a specific slot id to use a cached prompt.
	Prompt      *string `json:"prompt,omitempty"`       // used when asking for simple text completions
	// LlamaCpp *LlamaCppParams `json:"params,omitempty"`

	// FunctionCall is deprecated in favor of Tools
	FunctionCall *string `json:"function_call,omitempty"`
	// Functions is deprecated in favor of Tools
	Functions []CompletionFunciton `json:"functions,omitempty"`
}

type ToolChoice struct {
	String *string           `json:"string,omitempty"`
	Object *ToolChoiceObject `json:"object,omitempty"`
}

func (c *ToolChoice) MarshalJSON() ([]byte, error) {
	if c == nil {
		return nil, nil
	}
	if c.String != nil && *c.String != "" {
		return json.Marshal(c.String)
	}
	return json.Marshal(c.Object)
}

type ToolChoiceObject struct {
	Type     string        `json:"type"`
	Function ToolsFunction `json:"function"`
}

type ToolsFunction struct {
	Name string `json:"name"`
}

type Tools struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Function struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  any    `json:"parameters"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type CompletionFunciton struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Parameters  any    `json:"parameters,omitempty"`
}

type CreateChatCompletionsResponse struct {
	ID                string                        `json:"id,omitempty"`
	Choices           []CreateChatCompletionsChoice `json:"choices,omitempty"`
	Created           int                           `json:"created,omitempty"`
	Model             string                        `json:"model,omitempty"`
	SystemFingerprint string                        `json:"system_fingerprint,omitempty"`
	Object            string                        `json:"object,omitempty"`
	Usage             CreateChatCompletionsUsave    `json:"usage,omitempty"`
}

type CreateChatCompletionsChoice struct {
	Index        int      `json:"index,omitempty"`
	Message      *Message `json:"message,omitempty"`
	Delta        *Message `json:"delta,omitempty"`
	FinishReason string   `json:"finish_reason,omitempty"`
}

type CreateChatCompletionsUsave struct {
	CompletionTokens int `json:"completion_tokens,omitempty"`
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

// CreateChatCompletionsRaw creates a new completion. To use the OpenAI Chat
// API, pass "" as the url. Otherwise pass a valide endpoint for any server that
// mimics the API.
func (c *Client) CreateChatCompletionsRaw(ctx context.Context, r *CreateChatCompletionsRequest, url string) ([]byte, error) {
	if url == "" {
		return c.Post(ctx, completionsUrl, r)
	}
	// else
	return c.Post(ctx, url, r)

}

func (c *Client) CreateChatCompletions(ctx context.Context, r *CreateChatCompletionsRequest, url string) (response *CreateChatCompletionsResponse, err error) {
	raw, err := c.CreateChatCompletionsRaw(ctx, r, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
