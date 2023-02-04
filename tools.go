//go:build tools
// +build tools

// go:build tools
// ref: https://marcofranssen.nl/manage-go-tools-via-go-modules

package gocron

import (
	_ "github.com/cweill/gotests/gotests"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt"
	_ "golang.org/x/tools/cmd/goimports"
	_ "mvdan.cc/gofumpt"
)
