---
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_private_dns_a_record"
sidebar_current: "docs-azurerm-resource-private-dns-a-record"
description: |-
  Manages a Private DNS A Record.
---

# azurerm_dns_a_record

Enables you to manage DNS A Records within Azure Private DNS.

## Example Usage

```hcl
resource "azurerm_resource_group" "test" {
  name     = "acceptanceTestResourceGroup1"
  location = "West US"
}

resource "azurerm_private_dns_zone" "test" {
  name                = "mydomain.com"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_private_dns_a_record" "test" {
  name                = "test"
  zone_name           = "${azurerm_dns_zone.test.name}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  ttl                 = 300
  records             = ["10.0.180.17"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the DNS A Record.

* `resource_group_name` - (Required) Specifies the resource group where the resource exists. Changing this forces a new resource to be created.

* `zone_name` - (Required) Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.

* `TTL` - (Required) The Time To Live (TTL) of the DNS record in seconds.

* `records` - (Required) List of IPv4 Addresses.

* `tags` - (Optional) A mapping of tags to assign to the resource.

## Attributes Reference

The following attributes are exported:

* `id` - The Private DNS A Record ID.

## Import

Private DNS A Records can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_private_dns_a_record.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1/A/myrecord1
```
