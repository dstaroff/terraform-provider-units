/*
 * Copyright (c) 2024. Dmitry Starov
 * SPDX-License-Identifier: MPL-2.0
 */

package generator

import (
	"path/filepath"
)

var (
	PathDirRoot      string
	PathDirTemplates string

	PathDirExamples         string
	PathDirFunctionExamples string

	PathDirInternal           string
	PathDirConverter          string
	PathDirProvider           string
	PathDirDataSources        string
	PathDirFunctions          string
	PathDirFunctionsGenerated string
)

func init() {
	{
		var err error

		PathDirRoot, err = filepath.Abs("../../")
		if err != nil {
			panic(err)
		}

		PathDirTemplates, err = filepath.Abs("./templates")
		if err != nil {
			panic(err)
		}
	}

	PathDirExamples = filepath.Join(PathDirRoot, "examples")
	PathDirFunctionExamples = filepath.Join(PathDirExamples, "functions")

	PathDirInternal = filepath.Join(PathDirRoot, "internal")
	PathDirConverter = filepath.Join(PathDirInternal, "converter")
	PathDirProvider = filepath.Join(PathDirInternal, "provider")
	PathDirDataSources = filepath.Join(PathDirProvider, "datasource")
	PathDirFunctions = filepath.Join(PathDirProvider, "function")
	PathDirFunctionsGenerated = filepath.Join(PathDirFunctions, "generated")
}
