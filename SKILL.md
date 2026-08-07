---
name: scalar-api-go-sdk
description: "Go SDK for Scalar API. Use when writing Go code that calls Scalar API with the github.com/scalar/scalar-go package: installing it, constructing and authenticating the client, and calling API operations."
---

# Scalar API Go SDK

Generated Go client for Scalar API, published as `github.com/scalar/scalar-go`. Use the generated client instead of hand-writing HTTP requests.

## Install

```sh
go get github.com/scalar/scalar-go
```

## Client setup and authentication

```go
import (
	"context"
	"fmt"

	sdk "github.com/scalar/scalar-go"
)

client := sdk.NewClient()
```

Provide credentials using the options below. Environment variables are read automatically when the target runtime supports them:

- `option.WithBearerAuth` (env: `BEARER_AUTH`) — Credential for the BearerAuth client option.

## Calling operations

```go
package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/scalar/scalar-go"
	"github.com/scalar/scalar-go/option"
)

func main() {
	client := sdk.NewClient(
		option.WithBearerAuth(os.Getenv("BEARER_AUTH")),
	)

	registry, err := client.Registry.ListAllAPIDocuments(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(registry)
}
```

Method names, parameter shapes, and response types are generated from the API description — do not guess them. Look up the exact call signature in [api.md](./api.md) before writing a call.

## Error handling

Non-success responses return generated API errors. Error objects expose status, headers, response body, and request metadata where the target runtime supports it.

```go
registry, err := client.Registry.ListAllAPIDocuments(context.Background())
if err != nil {
	var apiErr *sdk.Error
	if errors.As(err, &apiErr) {
		fmt.Println(apiErr.StatusCode, apiErr.RawJSON())
	}
	panic(err)
}

// imports: sdk "github.com/scalar/scalar-go", "errors", "fmt"
```

## Requirements

- Go 1.22 or newer

## Reference files

- [README.md](./README.md) — full feature tour: client options, request options, retries and timeouts, logging.
- [api.md](./api.md) — complete catalogue of every operation with request and response types.
