/*
 * Copyright (c) 2024. Dmitry Starov
 * SPDX-License-Identifier: MPL-2.0
 */

package generator

import (
	"bytes"
	"crypto/sha256"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type (
	Function struct {
		Conversion    Conversion
		CopyrightInfo copyrightInfo
	}
	Conversion struct {
		Unit       ConversionUnit
		Directions []ConversionDirection
	}
	ConversionUnit struct {
		Title string
		Name  string
		Short string
	}
	ConversionDirection struct {
		Title    string
		Name     string
		Opposite *ConversionDirection
	}

	GeneratedFunctions struct {
		Names         []string
		CopyrightInfo copyrightInfo
	}
)

type Generator interface {
	GenerateFunctions() (functionConstructorNames []string)
}

type copyrightInfo struct {
	Author string
	Year   int
}

type Base struct {
	CopyrightInfo copyrightInfo
}

func NewBase() Base {
	return Base{
		CopyrightInfo: copyrightInfo{
			Author: "Dmitry Starov",
			Year:   time.Now().Year(),
		},
	}
}

func (b Base) Generate(filename string, templatePath string, data any) {
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		log.Fatal(err)
	}

	checksumBefore := b.getFileChecksum(filename)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(file, data)
	if err != nil {
		log.Fatal(err)
	}

	checksumAfter := b.getFileChecksum(filename)
	if !bytes.Equal(checksumBefore, checksumAfter) {
		_, p, _ := strings.Cut(filename, "terraform-provider-units")
		log.Println("Generated", p)
	}
}

func (_ Base) getFileChecksum(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		return []byte{}
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err = io.Copy(hasher, file); err != nil {
		log.Fatal(err)
	}

	return hasher.Sum(nil)
}

func (b Base) GenerateGeneratedFunctions(functionConstructorNames []string) {
	b.Generate(
		filepath.Join(PathDirFunctions, "generated.go"),
		filepath.Join(PathDirTemplates, "generated_functions.go.gotmpl"),
		GeneratedFunctions{
			Names:         functionConstructorNames,
			CopyrightInfo: b.CopyrightInfo,
		},
	)
}
