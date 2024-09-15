# Units Terraform provider

This provider gives you a possibility to convert categorized units in an interoperable manner.
Use [data sources](https://developer.hashicorp.com/terraform/language/data-sources) as containers for measurement units and converting them.
Or, convert them using [provider-defined functions](https://www.hashicorp.com/blog/terraform-1-8-improves-extensibility-with-provider-defined-functions).

## Problem to solve

- Tired of lacking possibility of an easy definition of quantities?
- One resource asks for disk size in GiB and other resource outputs it in MB?
- Tired of writing code like this?

    ```terraform
    resource "cloud_provider_disk" "this" {
      size = var.disk_size_gib * 1024 * 1024 * 1024
    }

    resource "another_cloud_provider_disk" "that" {
      size_gb = ceil((var.disk_size_gib * (1024 * 1024 * 1024)) / (1000 * 1000 * 1000))
    }
    ```

## Solution

### Data source

> With data sources, you can store converted values in a container, which will be stored in your state.

```terraform
data "units_data_size" "disk" {
  gibibytes = var.disk_size_gib
}

resource "cloud_provider_disk" "this" {
  size = data.units_data_size.disk.bytes
}

resource "another_cloud_provider_disk" "that" {
  size_gb = ceil(data.units_data_size.disk.gigabytes)
}
```

### Functions

> Converter function results are being computed during `plan`, and won't be stored in the state.

```terraform
resource "cloud_provider_disk" "this" {
  size = provider::units::from_gib(var.disk_size_gib)
}

resource "another_cloud_provider_disk" "that" {
  size_gb = ceil(provider::units::to_gb(provider::units::from_gib(var.disk_size_gib)))
}
```

## Requirements

| Component                                                        | Version    |
|:-----------------------------------------------------------------|:-----------|
| [Terraform](https://developer.hashicorp.com/terraform/downloads) | `>= 1.8.0` |
| [Go](https://golang.org/doc/install)                             | `>= 1.21`  |

## Liability

> This provider is not intended to do automatic rounding and outputs conversion results as is.
> Since results are `number`s, they can be both `int`s and `float`s.

Do not forget checking computed values and provide additional handling logic.
