# Units Terraform provider

This provider gives a possibility to use data sources as containers for measurement units and converting them in an interoperable manner.

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

Simply use:

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

## Requirements

| Component                                                        | Version   |
|:-----------------------------------------------------------------|:----------|
| [Terraform](https://developer.hashicorp.com/terraform/downloads) | `>= 1.0`  |
| [Go](https://golang.org/doc/install)                             | `>= 1.21` |

## Liability

> This provider is not intended to do automatic rounding and outputs conversion results as is.
> Since results are `number`s, they can be both `int`s and `float`s.

Do not forget checking computed values and provide additional handling logic.
