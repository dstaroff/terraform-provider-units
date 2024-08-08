// Copyright (c) Dmitrii Starov
// SPDX-License-Identifier: MPL-2.0

package datasource_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/dstaroff/terraform-provider-units/internal/testutils"
)

func TestAccDataSizeBase1000(t *testing.T) {
	for _, config := range []string{
		// language=hcl-terraform
		`
		data "units_data_size" "test" {
		  bytes = 1000000000000000
		}
		`,

		// language=hcl-terraform
		`
		data "units_data_size" "test" {
		  petabytes = 1
		}
		`,
	} {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.units_data_size.test", "bytes", "1000000000000000"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "kilobytes", "1000000000000"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "megabytes", "1000000000"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "gigabytes", "1000000"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "terabytes", "1000"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "petabytes", "1"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "kibibytes", "976562500000"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "mebibytes", "953674316.40625"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "gibibytes", "931322.5746154785"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "tebibytes", "909.4947017729282"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "pebibytes", "0.8881784197001252"),
				),
			}},
		})
	}
}

func TestAccDataSizeBase1024(t *testing.T) {
	for _, config := range []string{
		// language=hcl-terraform
		`
		data "units_data_size" "test" {
		  bytes = 1125899906842624
		}
		`,

		// language=hcl-terraform
		`
		data "units_data_size" "test" {
		  pebibytes = 1
		}
		`,
	} {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.units_data_size.test", "bytes", "1125899906842624"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "kilobytes", "1125899906842.624"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "megabytes", "1125899906.842624"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "gigabytes", "1125899.906842624"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "terabytes", "1125.899906842624"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "petabytes", "1.125899906842624"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "kibibytes", "1099511627776"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "mebibytes", "1073741824"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "gibibytes", "1048576"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "tebibytes", "1024"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "pebibytes", "1"),
				),
			}},
		})
	}
}

func TestAccDataSizeFromZero(t *testing.T) {
	for _, config := range []string{
		// language=hcl-terraform
		`
		data "units_data_size" "test" {
		  bytes = 0
		}
		`,

		// language=hcl-terraform
		`
		data "units_data_size" "test" {
		  pebibytes = 0
		}
		`,
	} {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.units_data_size.test", "bytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "kilobytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "megabytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "gigabytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "terabytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "petabytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "kibibytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "mebibytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "gibibytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "tebibytes", "0"),
					resource.TestCheckResourceAttr("data.units_data_size.test", "pebibytes", "0"),
				),
			}},
		})
	}
}

func TestAccDataSizeMultipleAttributesProvided(t *testing.T) {
	const config =
	// language=hcl-terraform
	`
	data "units_data_size" "test" {
	  bytes = 0
	  pebibytes = 0
	}
	`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config:      config,
			ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
		}},
	})
}

func TestAccDataSizeNoAttributesProvided(t *testing.T) {
	const config =
	// language=hcl-terraform
	`
	data "units_data_size" "test" {}
	`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config:      config,
			ExpectError: regexp.MustCompile(`Missing Attribute Configuration`),
		}},
	})
}
