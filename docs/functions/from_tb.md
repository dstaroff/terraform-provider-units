---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "from_tb function - units"
subcategory: ""
description: |-
  Converts terabytes to bytes
---

# function: from_tb

Given data size in **terabytes**, converts it to **bytes**.

## Example Usage

```terraform
output "example" {
  size_in_bytes = provider::units::from_tb(42)
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
from_tb(terabytes number) number
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `terabytes` (Number) Data size in **terabytes**

