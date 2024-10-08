// Code generated by go generate; DO NOT EDIT.
/*
 * Copyright (c) 2024. Dmitry Starov
 * SPDX-License-Identifier: MPL-2.0
 */

package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/dstaroff/terraform-provider-units/internal/converter"
)

var (
	_ function.Function = &FromMegabytesModel{}
	_ function.Function = &ToMegabytesModel{}
)

func NewFromMegabytesModel() function.Function {
	return &FromMegabytesModel{}
}

type FromMegabytesModel struct{}

func (f *FromMegabytesModel) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "from_mb"
}

func (f *FromMegabytesModel) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Converts megabytes to bytes",
		Description:         "Given data size in megabytes, converts it to bytes.",
		MarkdownDescription: "Given data size in **megabytes**, converts it to **bytes**.",

		Parameters: []function.Parameter{
			function.NumberParameter{
				Name:                "megabytes",
				Description:         "Data size in megabytes",
				MarkdownDescription: "Data size in **megabytes**",
			},
		},
		Return: function.NumberReturn{},
	}
}

func (f *FromMegabytesModel) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var megabytes types.Number

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &megabytes))
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, converter.MegabytesToBytes(megabytes)))
}

func NewToMegabytesModel() function.Function {
	return &ToMegabytesModel{}
}

type ToMegabytesModel struct{}

func (f *ToMegabytesModel) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "to_mb"
}

func (f *ToMegabytesModel) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Converts bytes to megabytes",
		Description:         "Given data size in bytes, converts it to megabytes.",
		MarkdownDescription: "Given data size in **bytes**, converts it to **megabytes**.",

		Parameters: []function.Parameter{
			function.NumberParameter{
				Name:                "bytes",
				Description:         "Data size in bytes",
				MarkdownDescription: "Data size in **bytes**",
			},
		},
		Return: function.NumberReturn{},
	}
}

func (f *ToMegabytesModel) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var bytes types.Number

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &bytes))
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, converter.MegabytesFromBytes(bytes)))
}
