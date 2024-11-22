# Go library for accessing the freemyip.com API

[![Build Status](https://github.com/nrdcg/freemyip/actions/workflows/main.yml/badge.svg)](https://github.com/nrdcg/freemyip/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nrdcg/freemyip)](https://pkg.go.dev/github.com/nrdcg/freemyip)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/freemyip)](https://goreportcard.com/report/github.com/nrdcg/freemyip)

A Go client library for accessing the [freemyip.com](https://freemyip.com) API.

## Examples

```go
package main

import (
	"context"
	"fmt"

	"github.com/nrdcg/freemyip"
)

func main() {
	client := freemyip.New("secret", true)

	ctx := context.Background()

	resp, err := client.UpdateDomain(ctx, "example", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
```

## API Documentation

- [API docs](https://freemyip.com/help)
