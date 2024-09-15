/*
 * Copyright (c) HashiCorp, Inc.
 * Copyright (c) 2024. Dmitry Starov
 * SPDX-License-Identifier: MPL-2.0
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	mydatasource "github.com/dstaroff/terraform-provider-units/internal/provider/datasource"
	myfuncs "github.com/dstaroff/terraform-provider-units/internal/provider/function"
)

var _ provider.Provider = &Units{}
var _ provider.ProviderWithFunctions = &Units{}

// Units defines the provider implementation.
type Units struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

var unitsDescription = `This provider gives you a possibility to convert categorized units in an interoperable manner. ` +
	`Use data sources as containers for measurement units and converting them. ` +
	`Or, convert them using provider-defined functions. `

const unitsDescriptionMd =
// language=markdown
`
This provider gives you a possibility to convert categorized units in an interoperable manner.
Use [data sources](https://developer.hashicorp.com/terraform/language/data-sources) as containers for measurement units and converting them.
Or, convert them using [provider-defined functions](https://www.hashicorp.com/blog/terraform-1-8-improves-extensibility-with-provider-defined-functions).

## Problem to solve

- Tired of lacking possibility of an easy definition of quantities?
- One resource asks for disk size in GiB and other resource outputs it in MB?
- Tired of writing code like this?

` + "```terraform" + `
resource "cloud_provider_disk" "this" {
  size = var.disk_size_gib * 1024 * 1024 * 1024
}

resource "another_cloud_provider_disk" "that" {
  size_gb = ceil((var.disk_size_gib * (1024 * 1024 * 1024)) / (1000 * 1000 * 1000))
}
` + "```" + `

## Solution

### Data source

> With data sources, you can store converted values in a container, which will be stored in your state.

` + "```terraform" + `
data "units_data_size" "disk" {
  gibibytes = var.disk_size_gib
}

resource "cloud_provider_disk" "this" {
  size = data.units_data_size.disk.bytes
}

resource "another_cloud_provider_disk" "that" {
  size_gb = ceil(data.units_data_size.disk.gigabytes)
}
` + "```" + `

### Functions

> Converter function results are being computed during ` + "`plan`" + `, and won't be stored in the state.

` + "```terraform" + `
resource "cloud_provider_disk" "this" {
  size = provider::units::from_gib(var.disk_size_gib)
}

resource "another_cloud_provider_disk" "that" {
  size_gb = ceil(provider::units::to_gb(provider::units::from_gib(var.disk_size_gib)))
}
` + "```" + `

## Liability

This provider is not intended to do automatic rounding and outputs conversion results as is.
Since results are ` + "`number`s" + `, they can be both ` + "`int`s" + ` and ` + "`float`s." + `

Do not forget checking computed values and provide additional handling logic.
`

// UnitsModel describes the provider data model.
type UnitsModel struct{}

func (p *Units) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "units"
	resp.Version = p.version
}

func (p *Units) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         unitsDescription,
		MarkdownDescription: unitsDescriptionMd,
		Attributes:          map[string]schema.Attribute{},
	}
}

func (p *Units) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data UnitsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *Units) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *Units) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		mydatasource.NewDataSize,
	}
}

func (p *Units) Functions(_ context.Context) []func() function.Function {
	var res []func() function.Function
	res = append(res, myfuncs.GeneratedFunctions...)
	return res
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Units{
			version: version,
		}
	}
}
