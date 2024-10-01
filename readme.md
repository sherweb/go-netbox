# go-netbox

[![GoDoc](https://pkg.go.dev/badge/github.com/sherweb/go-netbox/v4)](https://pkg.go.dev/github.com/netbox-community/sherweb/v4)

This is a fork of [netbox-community/go-netbox](https://github.com/netbox-community/go-netbox) as it's not very maintainted.

## Installation

Use `go get` to add the library as a dependency to your project. Do not forget to run `go mod init` first if necessary.

```shell
go get github.com/sherweb/go-netbox/v4

```

## Usage

### Instantiate the client

The package has a constructor for creating a client by providing a URL and an authentication token.

```golang
package main

import (
	"context"
	"log"

	"github.com/sherweb/go-netbox/v4/netbox"
)

func main() {
	cfg := netbox.NewConfiguration()
	cfg.Servers[0].URL = "Netbox URL"
	cfg.AddDefaultHeader("Authorization", "Token sometoken")

	client := netbox.NewAPIClient(cfg)
	
}

```

# Important !

We'd recommend you look at the tests/ folder to look at what we've tested works...

Contribs welcome.