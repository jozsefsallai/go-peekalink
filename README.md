# Peekalink Go API Client

This Go package is an API client for the [Peekalink][peekalink-url] service.

## Usage

**1. Get the module:**

```
go get github.com/ankida/go-peekalink
```

**2. Import it:**

```go
package main

import (
  "github.com/ankida/go-peekalink"
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

## License

This Go module is licensed under the MIT license.

[peekalink-url]: https://peekalink.io
[docs-url]: https://pkg.go.dev/github.com/ankida/go-peekalink
