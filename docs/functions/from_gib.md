---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "from_gib function - units"
subcategory: ""
description: |-
  Converts gibibytes to bytes
---

# function: from_gib

Given data size in **gibibytes**, converts it to **bytes**.

## Example Usage

```terraform
output "example" {
  size_in_bytes = provider::units::from_gib(42)
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
from_gib(gibibytes number) number
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `gibibytes` (Number) Data size in **gibibytes**

