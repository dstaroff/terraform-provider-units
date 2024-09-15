/*
 * Copyright (c) 2024. Dmitry Starov
 * SPDX-License-Identifier: MPL-2.0
 */

package function_test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/dstaroff/terraform-provider-units/internal/testutils"
)

func TestAccDataSizeFunctions(t *testing.T) {
	type testCaseType struct {
		config string
		result string
	}
	var testCases []testCaseType

	for _, base := range []int{1000, 1024} {
		var abbr string
		if base == 1024 {
			abbr = "i"
		}

		for i, unitPrefix := range []string{"k", "m", "g", "t", "p"} {
			result := base
			for j := 1; j <= i; j++ {
				result *= base
			}

			testCases = append(testCases, testCaseType{
				config: fmt.Sprintf(
					// language=hcl-terraform
					`
				output "test" {
					value = provider::units::from_%s%sb(1)
				}
				`, unitPrefix, abbr,
				),
				result: strconv.Itoa(result),
			}, testCaseType{
				config: fmt.Sprintf(
					// language=hcl-terraform
					`
				output "test" {
					value = provider::units::to_%s%sb(%d)
				}
				`, unitPrefix, abbr, result,
				),
				result: strconv.Itoa(1),
			})
		}
	}

	for _, tc := range testCases {
		resource.UnitTest(t, resource.TestCase{
			TerraformVersionChecks: []tfversion.TerraformVersionCheck{
				tfversion.SkipBelow(tfversion.Version1_8_0),
			},
			ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: tc.config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckOutput("test", tc.result),
					),
				},
			},
		})
	}
}

func TestAccDataSizeFunctions_0(t *testing.T) {
	type testCaseType struct {
		config string
	}
	var testCases []testCaseType

	result := strconv.Itoa(0)

	for _, base := range []int{1000, 1024} {
		var abbr string
		if base == 1024 {
			abbr = "i"
		}

		for _, unitPrefix := range []string{"k", "m", "g", "t", "p"} {
			testCases = append(testCases, testCaseType{
				config: fmt.Sprintf(
					// language=hcl-terraform
					`
				output "test" {
					value = provider::units::from_%s%sb(0)
				}
				`, unitPrefix, abbr,
				),
			}, testCaseType{
				config: fmt.Sprintf(
					// language=hcl-terraform
					`
				output "test" {
					value = provider::units::to_%s%sb(0)
				}
				`, unitPrefix, abbr,
				),
			})
		}
	}

	for _, tc := range testCases {
		resource.UnitTest(t, resource.TestCase{
			TerraformVersionChecks: []tfversion.TerraformVersionCheck{
				tfversion.SkipBelow(tfversion.Version1_8_0),
			},
			ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: tc.config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckOutput("test", result),
					),
				},
			},
		})
	}
}

func TestAccDataSizeFunctions_null(t *testing.T) {
	type testCaseType struct {
		config string
	}
	var testCases []testCaseType

	for _, base := range []int{1000, 1024} {
		var abbr string
		if base == 1024 {
			abbr = "i"
		}

		for _, unitPrefix := range []string{"k", "m", "g", "t", "p"} {
			testCases = append(testCases, testCaseType{
				config: fmt.Sprintf(
					// language=hcl-terraform
					`
				output "test" {
					value = provider::units::from_%s%sb(null)
				}
				`, unitPrefix, abbr,
				),
			}, testCaseType{
				config: fmt.Sprintf(
					// language=hcl-terraform
					`
				output "test" {
					value = provider::units::to_%s%sb(null)
				}
				`, unitPrefix, abbr,
				),
			})
		}
	}

	for _, tc := range testCases {
		resource.UnitTest(t, resource.TestCase{
			TerraformVersionChecks: []tfversion.TerraformVersionCheck{
				tfversion.SkipBelow(tfversion.Version1_8_0),
			},
			ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config:      tc.config,
					ExpectError: regexp.MustCompile(`argument must not be null`),
				},
			},
		})
	}
}
