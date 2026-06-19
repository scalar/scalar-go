# Scalar API

Generated Go SDK for Scalar API.
API for managing Scalar platform resources.

## TypeScript SDK

For TypeScript, we provide a SDK that makes using our API even easier.

### Install

```bash
npm add @scalar/sdk
```

### Get a Scalar API key

Create an API key in your Scalar account:

- Dashboard: https://dashboard.scalar.com/account
- Store it in `.env`, for example:

```bash
SCALAR_API_KEY=your_personal_token
```

### Exchange your API key for an access token

The personal token is not an access token. Exchange it first with `postv1AuthExchange`.

If you use the personal token directly for authenticated API calls, the API returns `401 Invalid authentication token`.

```ts
import { Scalar } from '@scalar/sdk'

const scalar = new Scalar()

const exchange = await scalar.auth.postv1AuthExchange({
  personalToken: process.env.SCALAR_API_KEY!,
})

const accessToken = exchange.accessToken
```

### Use the access token

Construct a second client with bearer auth. Use this authenticated client for API calls.

```ts
import { Scalar } from '@scalar/sdk'

const scalar = new Scalar()

const exchange = await scalar.auth.postv1AuthExchange({
  personalToken: process.env.SCALAR_API_KEY!,
})

const authedScalar = new Scalar({
  bearerAuth: exchange.accessToken,
})
```

### Notes

- The exchange request itself can be made from a client constructed with no arguments (`new Scalar()`).
- The exchanged access token is valid for 12 hours.
- Timestamps are Unix seconds.

### Read more

- [@scalar/sdk on npm](https://www.npmjs.com/package/@scalar/sdk)

<br />

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Reference](./api.md)
- [Authentication](#authentication)
- [Errors](#errors)
- [Client Options](#client-options)
- [Request Options](#request-options)
- [Retries and Timeouts](#retries-and-timeouts)
- [Helpers](#helpers)
- [Logging](#logging)
- [Requirements](#requirements)

<br />

## Installation

```sh
go get github.com/scalar/scalar-go
```

<br />

## Usage

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

The examples in the following sections assume a `client` configured as shown above.

See the [API reference](./api.md) for every available operation.

<br />

## Authentication

Pass credentials to the generated client constructor. Environment variables are read automatically when supported by the target runtime.

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `option.WithBearerAuth` | `string \| provider` | - | Credential for the BearerAuth client option. Defaults to BEARER_AUTH. |

Declared schemes:

- `BearerAuth` bearer token

<br />

## Errors

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

Documented error statuses: `400`, `401`, `403`, `404`, `422`, `500`.

<br />

## Client Options

Configure the generated client by setting any of these options when you create it.

```go
client := sdk.NewClient(
	option.WithBaseURL("https://api.example.com"),
	option.WithMaxRetries(2),
	option.WithTimeout(60*time.Second),
)

// imports: sdk "github.com/scalar/scalar-go", "github.com/scalar/scalar-go/option", "time"
```

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `option.WithBearerAuth` | `func(string) option.RequestOption` | `os.Getenv("BEARER_AUTH")` | Credential for the BearerAuth client option. |
| `option.WithEnvironmentProduction` | `func() option.RequestOption` | - | Select the production API environment. |
| `option.WithEnvironmentLocal` | `func() option.RequestOption` | - | Select the local API environment. |
| `option.WithBaseURL` | `func(string) option.RequestOption` | `os.Getenv("SCALAR_BASE_URL")` | Override the default API base URL. |
| `option.WithTimeout` | `func(time.Duration) option.RequestOption` | - | Maximum time to wait for each request attempt. |
| `option.WithMaxRetries` | `func(int) option.RequestOption` | `2` | Number of retries for temporary failures. |
| `option.WithHTTPClient` | `func(option.HTTPClient) option.RequestOption` | - | Custom HTTP client or transport implementation. |

<br />

## Request Options

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `option.WithHeader` | `func(string, string) option.RequestOption` | - | Set a per-request header. |
| `option.WithQuery` | `func(string, string) option.RequestOption` | - | Set a per-request query parameter. |
| `option.WithRequestBody` | `func(string, any) option.RequestOption` | - | Override the serialized request body and content type. |
| `option.WithResponseInto` | `func(**http.Response) option.RequestOption` | - | Capture the raw HTTP response. |
| `option.WithResponseBodyInto` | `func(any) option.RequestOption` | - | Override the response deserialization target. |

<br />

## Retries and Timeouts

Generated clients support request timeouts and retry temporary failures such as network errors, 408, 409, 429, and 5xx responses. Retry delays honor `Retry-After` headers when present. Tune the retry and timeout client options shown above, or override them per request.

<br />

## Helpers

- Pass `option.WithResponseInto(&raw)` to capture the underlying `*http.Response` for a request.
- Use the generated `String`, `Int`, `Bool`, `Float`, `Time`, `Opt`, and `Ptr` helpers when setting optional params.

<br />

## Logging

- Wrap the HTTP client with `option.WithMiddleware(...)` to add request logging or tracing.

<br />

## Requirements

- Go 1.22 or newer

Powered by Scalar.


## Contributions

This SDK is generated programmatically. Manual edits to generated files will be
overwritten on the next build.

### SDK created by [Scalar](https://www.scalar.com/?utm_source=scalar-typescript-sdk-go&utm_campaign=sdk)
