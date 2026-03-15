# go-jules

go-jules is a Go client library for accessing the Google Jules API (v1alpha). It provides a unified, strongly-typed interface to interact with Jules sessions, activities, and sources.

## Installation

```bash
go get github.com/deepnor/jules-sdk
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deepnor/jules-sdk"
)

func main() {
	// Initialize the unified client
	client := jules.NewClient(os.Getenv("JULES_API_KEY"))
	ctx := context.Background()

	// Construct request options
	req := &jules.CreateSessionRequest{
		Prompt: "Implement a concurrent rate limiter",
		Title:  "Rate Limiter Implementation",
	}

	// Create a new Jules session
	session, err := client.Sessions.Create(ctx, req)
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	fmt.Printf("Session Name: %s\n", session.Name)
	fmt.Printf("Session URL: %s\n", session.URL)
}
```

## Pagination

The Jules API uses cursor-based pagination. Use the `PageToken` and `PageSize` arguments to iterate through resources.

```go
pageSize := 50
pageToken := ""

for {
	resp, err := client.Sources.List(ctx, "", pageSize, pageToken)
	if err != nil {
		log.Fatalf("Error listing sources: %v", err)
	}

	for _, source := range resp.Sources {
		fmt.Printf("Source: %s\n", source.Name)
	}

	pageToken = resp.NextPageToken
	if pageToken == "" {
		break
	}
}
```

## Error Handling

API errors are surfaced as `*jules.APIError`. Use `errors.As` from the standard library to extract HTTP status codes and API error messages.

```go
import (
	"errors"
	"log"
	
	"github.com/deepnor/jules-sdk"
)

// ...

session, err := client.Sessions.Get(ctx, "sessions/invalid-id")
if err != nil {
	var apiErr *jules.APIError
	if errors.As(err, &apiErr) {
		log.Printf("HTTP Error: %d", apiErr.HTTPStatusCode)
		log.Printf("Status: %s", apiErr.Status)
		log.Printf("Message: %s", apiErr.Message)
		// Handle specific HTTP Status Codes
		if apiErr.HTTPStatusCode == 404 {
			log.Println("Session not found")
		}
	} else {
		log.Fatalf("Unexpected error: %v", err)
	}
}
```
