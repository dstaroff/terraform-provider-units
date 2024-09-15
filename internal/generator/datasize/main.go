/*
 * Copyright (c) 2024. Dmitry Starov
 * SPDX-License-Identifier: MPL-2.0
 */

package datasize

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Masterminds/goutils"

	"github.com/dstaroff/terraform-provider-units/internal/generator"
)

var (
	units = []struct {
		Full  string
		Short string
	}{{
		Full:  "kibibytes",
		Short: "kib",
	}, {
		Full:  "mebibytes",
		Short: "mib",
	}, {
		Full:  "gibibytes",
		Short: "gib",
	}, {
		Full:  "tebibytes",
		Short: "tib",
	}, {
		Full:  "pebibytes",
		Short: "pib",
	}, {
		Full:  "kilobytes",
		Short: "kb",
	}, {
		Full:  "megabytes",
		Short: "mb",
	}, {
		Full:  "gigabytes",
		Short: "gb",
	}, {
		Full:  "terabytes",
		Short: "tb",
	}, {
		Full:  "petabytes",
		Short: "pb",
	}}
)

var _ generator.Generator = &Generator{}

type Generator struct {
	generator.Base
}

func NewGenerator() *Generator {
	return &Generator{
		Base: generator.NewBase(),
	}
}

func (g *Generator) GenerateFunctions() (functionConstructorNames []string) {
	var directions []generator.ConversionDirection
	{
		dirFrom := generator.ConversionDirection{
			Title: "From",
			Name:  "from",
		}
		dirTo := generator.ConversionDirection{
			Title: "To",
			Name:  "to",
		}
		dirFrom.Opposite = &dirTo
		dirTo.Opposite = &dirFrom

		directions = append(directions, dirFrom, dirTo)
	}

	var functions []generator.Function
	for _, unit := range units {
		functions = append(functions, generator.Function{
			Conversion: generator.Conversion{
				Unit: generator.ConversionUnit{
					Title: goutils.CapitalizeFully(unit.Full),
					Name:  unit.Full,
					Short: strings.ToLower(unit.Short),
				},
				Directions: directions,
			},
			CopyrightInfo: g.CopyrightInfo,
		})
	}

	for _, function := range functions {
		g.Generate(
			filepath.Join(generator.PathDirFunctionsGenerated, fmt.Sprintf("data_size_%s.go", function.Conversion.Unit.Name)),
			filepath.Join(generator.PathDirTemplates, "data_size_function.go.gotmpl"),
			function,
		)
		for _, direction := range function.Conversion.Directions {
			functionConstructorNames = append(
				functionConstructorNames,
				fmt.Sprintf("New%s%sModel", direction.Title, function.Conversion.Unit.Title),
			)
		}

		for _, direction := range directions {
			function.Conversion.Directions = []generator.ConversionDirection{
				direction,
			}
			g.Generate(
				filepath.Join(generator.PathDirFunctionExamples, fmt.Sprintf("%s_%s", direction.Name, function.Conversion.Unit.Short), "function.tf"),
				filepath.Join(generator.PathDirTemplates, "data_size_function_example.tf.gotmpl"),
				function,
			)
		}
	}

	return functionConstructorNames
}
