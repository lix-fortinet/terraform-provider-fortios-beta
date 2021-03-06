---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerptp_policy"
description: |-
  PTP policy configuration.
---

# fortios_switchcontrollerptp_policy
PTP policy configuration. Applies to FortiOS Version `>= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Policy name.
* `status` - Enable/disable PTP policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerPtp Policy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerptp_policy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
