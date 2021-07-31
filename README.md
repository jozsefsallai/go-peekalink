# Peekalink Go API Client

This Go package is an API client for the [Peekalink][peekalink-url] service.

## Usage

**1. Get the module:**

```
go get github.com/jozsefsallai/go-peekalink
```

**2. Import it:**

```go
package main

import (
  "github.com/jozsefsallai/go-peekalink"
)
```

**3. Instantiate a client:**

```go
client := peekalink.NewClient("your-api-key")
```

**Get link preview data:**

```go
preview, err := client.Preview("https://okuna.io")
if err != nil {
  panic(err)
}

fmt.Println(preview)
```

**Get link availability status:**

```go
status, err := client.IsAvailable("https://okuna.io")
if err != nil {
  panic(err)
}

fmt.Println(status)
```

Please read the [documentation][docs-url] of the module for more information
about all the available methods, types, and constants.

## Contribution

The Peekalink Go API client library is an open-source project and any
contribution is welcome! When contributing, we kindly ask you to follow the
[Go style guide][style-guide] and use `gofmt` whenever possible.

For major changes, we also expect that you update/write unit tests and make sure
they pass. Most unit tests will not require direct interaction with the real
API, however, some of them do need to check against actual data. By default,
these tests will **not** run when running `go test`, unless you provide a
Peekalink API key through the `PEEKALINK_API_KEY` environment variable:

```
PEEKALINK_API_KEY=your-api-key go test -v ./...
```

Please note that running the tests with your API key **will count towards your
usage quota**.

## License

This Go module is licensed under the MIT license.

[peekalink-url]: https://peekalink.io
[docs-url]: https://pkg.go.dev/github.com/jozsefsallai/go-peekalink
[style-guide]: https://golang.org/doc/effective_go
