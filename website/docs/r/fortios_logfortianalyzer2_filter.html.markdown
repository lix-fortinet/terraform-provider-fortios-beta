---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzer2_filter"
description: |-
  Filters for FortiAnalyzer.
---

# fortios_logfortianalyzer2_filter
Filters for FortiAnalyzer.

## Example Usage

```hcl
resource "fortios_logfortianalyzer2_filter" "trname" {
  anomaly           = "enable"
  dlp_archive       = "enable"
  dns               = "enable"
  filter_type       = "include"
  forward_traffic   = "enable"
  gtp               = "enable"
  local_traffic     = "enable"
  multicast_traffic = "enable"
  severity          = "information"
  sniffer_traffic   = "enable"
  ssh               = "enable"
  voip              = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `severity` - Log every message above and including this severity level.
* `forward_traffic` - Enable/disable forward traffic logging.
* `local_traffic` - Enable/disable local in or out traffic logging.
* `multicast_traffic` - Enable/disable multicast traffic logging.
* `sniffer_traffic` - Enable/disable sniffer traffic logging.
* `anomaly` - Enable/disable anomaly logging.
* `netscan_discovery` - Enable/disable netscan discovery event logging. (`ver <= 6.2.0`)
* `netscan_vulnerability` - Enable/disable netscan vulnerability event logging. (`ver <= 6.2.0`)
* `voip` - Enable/disable VoIP logging.
* `dlp_archive` - Enable/disable DLP archive logging.
* `gtp` - Enable/disable GTP messages logging.
* `free_style` - Free Style Filters The structure of `free_style` block is documented below. (`ver >= 6.6.0`)
* `dns` - Enable/disable detailed DNS event logging. (`ver <= 6.2.0`)
* `ssh` - Enable/disable SSH logging. (`ver <= 6.2.0`)
* `filter` - FortiAnalyzer 2 log filter. (`ver <= 6.4.2`)
* `filter_type` - Include/exclude logs that match the filter. (`ver <= 6.4.2`)

The `free_style` block supports (`ver >= 6.6.0`):

* `id` - Entry ID.
* `category` - Log category.
* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortianalyzer2 Filter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logfortianalyzer2_filter.labelname LogFortianalyzer2Filter
$ unset "FORTIOS_IMPORT_TABLE"
```
