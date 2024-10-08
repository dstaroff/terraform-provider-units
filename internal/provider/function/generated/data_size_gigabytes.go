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
	_ function.Function = &FromGigabytesModel{}
	_ function.Function = &ToGigabytesModel{}
)

func NewFromGigabytesModel() function.Function {
	return &FromGigabytesModel{}
}

type FromGigabytesModel struct{}

func (f *FromGigabytesModel) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "from_gb"
}

func (f *FromGigabytesModel) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Converts gigabytes to bytes",
		Description:         "Given data size in gigabytes, converts it to bytes.",
		MarkdownDescription: "Given data size in **gigabytes**, converts it to **bytes**.",

		Parameters: []function.Parameter{
			function.NumberParameter{
				Name:                "gigabytes",
				Description:         "Data size in gigabytes",
				MarkdownDescription: "Data size in **gigabytes**",
			},
		},
		Return: function.NumberReturn{},
	}
}

func (f *FromGigabytesModel) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var gigabytes types.Number

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &gigabytes))
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, converter.GigabytesToBytes(gigabytes)))
}

func NewToGigabytesModel() function.Function {
	return &ToGigabytesModel{}
}

type ToGigabytesModel struct{}

func (f *ToGigabytesModel) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "to_gb"
}

func (f *ToGigabytesModel) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Converts bytes to gigabytes",
		Description:         "Given data size in bytes, converts it to gigabytes.",
		MarkdownDescription: "Given data size in **bytes**, converts it to **gigabytes**.",

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

func (f *ToGigabytesModel) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var bytes types.Number

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &bytes))
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, converter.GigabytesFromBytes(bytes)))
}
