# This Fork
The original library from Francisco Escher has the OpenAI endpoints hard-coded as constants.  This fork will add an alternate completion call that allows you to specify any server endpoint that mimics that the OpenAI API, e.g. the llama.cpp server interface.

# Go OpenAI

This is a Go client library for the OpenAI API.

It implements the methods described in the docs: https://platform.openai.com/docs/api-reference/introduction

Implemented methods can be found in the Interface.go file.

## Installation

    go get github.com/franciscoescher/goopenai/v3

## Usage

First, you need to create a client with the api key and organization id.

```
client := goopenai.NewClient(apiKey, organization)
```

Then, you can use the client to call the api.

Example:

```
package main

import (
	"context"
	"fmt"

	"github.com/franciscoescher/goopenai"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	organization := os.Getenv("API_ORG")

	client := goopenai.NewClient(apiKey, organization)

	r := &goopenai.CreateChatCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []goopenai.Message{
			{
				Role:    "user",
				Content: "Say this is a test!",
			},
		},
		Temperature: 0.7,
	}

	completions, err := client.CreateChatCompletions(context.Background(), r)
	if err != nil {
		panic(err)
	}

	fmt.Println(completions)
}

```

Run this code using:

`API_KEY=<your-api-key> API_ORG=<your-org-id> go run .`

## Note

This library is not complete and not fully tested.

Feel free to contribute.
