//go:build generate

/*
 * Copyright (c) 2024. Dmitry Starov
 * SPDX-License-Identifier: MPL-2.0
 */

package main

import (
	"github.com/dstaroff/terraform-provider-units/internal/generator"
	"github.com/dstaroff/terraform-provider-units/internal/generator/datasize"
)

var (
	generators = []generator.Generator{
		datasize.NewGenerator(),
	}
)

func main() {
	var functionConstructorNames []string

	for _, g := range generators {
		functionConstructorNames = append(functionConstructorNames, g.GenerateFunctions()...)
		g.GenerateConverterNames()
	}

	generator.NewBase().GenerateGeneratedFunctions(functionConstructorNames)
}

//go:generate go run ./${GOFILE}
