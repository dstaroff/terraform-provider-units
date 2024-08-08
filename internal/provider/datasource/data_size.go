// Copyright (c) Dmitrii Starov
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"
	"math/big"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/dstaroff/terraform-provider-units/internal/converter"
)

var _ datasource.DataSource = &DataSize{}

func NewDataSize() datasource.DataSource {
	return &DataSize{}
}

// DataSize defines the data source implementation for data size conversion.
type DataSize struct{}

var dataSizeDescription = strings.Join([]string{
	"Container for data sizes.",
	"This data source is capable of taking data size in one unit (e.g. MiB) and convert it to other units (e.g. KB).",
	"This is done by converting input size to bytes and then converting it back to other units.",
	"NOTE: Specify exactly one of provided attributes to get others converted.",
}, " ")

const dataSizeDescriptionMd =
// language=markdown
`
## Container for data sizes

This data source is capable of taking data size in one unit (e.g. ` + "`MiB`" + `) and convert it to other units (e.g. ` + "`KB`" + `).

This is done by converting input size to bytes and then converting it back to other units.

**NOTE**:
Specify exactly one of provided attributes to get others converted.
`

var _ converter.Converter = &DataSizeModel{}

// DataSizeModel describes the data source data model.
type DataSizeModel struct {
	Bytes types.Number `tfsdk:"bytes"`

	Kibibytes types.Number `tfsdk:"kibibytes"`
	Mebibytes types.Number `tfsdk:"mebibytes"`
	Gibibytes types.Number `tfsdk:"gibibytes"`
	Tebibytes types.Number `tfsdk:"tebibytes"`
	Pebibytes types.Number `tfsdk:"pebibytes"`

	Kilobytes types.Number `tfsdk:"kilobytes"`
	Megabytes types.Number `tfsdk:"megabytes"`
	Gigabytes types.Number `tfsdk:"gigabytes"`
	Terabytes types.Number `tfsdk:"terabytes"`
	Petabytes types.Number `tfsdk:"petabytes"`
}

// Convert performs the conversion of data size.
func (m *DataSizeModel) Convert() {
	bytes := types.NumberValue(big.NewFloat(0))
	if !m.Bytes.IsNull() {
		bytes = m.Bytes
	} else if !m.Kibibytes.IsNull() {
		bytes = converter.KibibytesToBytes(m.Kibibytes)
	} else if !m.Mebibytes.IsNull() {
		bytes = converter.MebibytesToBytes(m.Mebibytes)
	} else if !m.Gibibytes.IsNull() {
		bytes = converter.GibibytesToBytes(m.Gibibytes)
	} else if !m.Tebibytes.IsNull() {
		bytes = converter.TebibytesToBytes(m.Tebibytes)
	} else if !m.Pebibytes.IsNull() {
		bytes = converter.PebibytesToBytes(m.Pebibytes)
	} else if !m.Kilobytes.IsNull() {
		bytes = converter.KilobytesToBytes(m.Kilobytes)
	} else if !m.Megabytes.IsNull() {
		bytes = converter.MegabytesToBytes(m.Megabytes)
	} else if !m.Gigabytes.IsNull() {
		bytes = converter.GigabytesToBytes(m.Gigabytes)
	} else if !m.Terabytes.IsNull() {
		bytes = converter.TerabytesToBytes(m.Terabytes)
	} else if !m.Petabytes.IsNull() {
		bytes = converter.PetabytesToBytes(m.Petabytes)
	}

	m.Bytes = bytes

	m.Kibibytes = converter.BytesToKibibytes(bytes)
	m.Mebibytes = converter.BytesToMebibytes(bytes)
	m.Gibibytes = converter.BytesToGibibytes(bytes)
	m.Tebibytes = converter.BytesToTebibytes(bytes)
	m.Pebibytes = converter.BytesToPebibytes(bytes)

	m.Kilobytes = converter.BytesToKilobytes(bytes)
	m.Megabytes = converter.BytesToMegabytes(bytes)
	m.Gigabytes = converter.BytesToGigabytes(bytes)
	m.Terabytes = converter.BytesToTerabytes(bytes)
	m.Petabytes = converter.BytesToPetabytes(bytes)
}

func (d *DataSize) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_data_size"
}

func (d *DataSize) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         dataSizeDescription,
		MarkdownDescription: dataSizeDescriptionMd,

		Attributes: map[string]schema.Attribute{
			"bytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in bytes.",
				Optional:            true,
				Computed:            true,
			},
			"kibibytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in kibibytes.",
				Optional:            true,
				Computed:            true,
			},
			"mebibytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in mebibytes.",
				Optional:            true,
				Computed:            true,
			},
			"gibibytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in gibibytes.",
				Optional:            true,
				Computed:            true,
			},
			"tebibytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in tebibytes.",
				Optional:            true,
				Computed:            true,
			},
			"pebibytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in pebibytes.",
				Optional:            true,
				Computed:            true,
			},
			"kilobytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in kilobytes.",
				Optional:            true,
				Computed:            true,
			},
			"megabytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in megabytes.",
				Optional:            true,
				Computed:            true,
			},
			"gigabytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in gigabytes.",
				Optional:            true,
				Computed:            true,
			},
			"terabytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in terabytes.",
				Optional:            true,
				Computed:            true,
			},
			"petabytes": schema.NumberAttribute{
				MarkdownDescription: "Data size in petabytes.",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (d *DataSize) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DataSizeModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "converting data size")
	data.Convert()

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *DataSize) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("bytes"),
			path.MatchRoot("kibibytes"),
			path.MatchRoot("mebibytes"),
			path.MatchRoot("gibibytes"),
			path.MatchRoot("tebibytes"),
			path.MatchRoot("pebibytes"),
			path.MatchRoot("kilobytes"),
			path.MatchRoot("megabytes"),
			path.MatchRoot("gigabytes"),
			path.MatchRoot("terabytes"),
			path.MatchRoot("petabytes"),
		),
	}
}
