//go:build tools
// +build tools

package main

import (
	_ "ariga.io/atlas-provider-gorm"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)
