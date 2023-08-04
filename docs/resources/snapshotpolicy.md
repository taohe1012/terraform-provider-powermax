---
# Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://mozilla.org/MPL/2.0/
#
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

title: "powermax_snapshotpolicy resource"
linkTitle: "powermax_snapshotpolicy"
page_title: "powermax_snapshotpolicy Resource - terraform-provider-powermax"
subcategory: ""
description: |-
  Resource for a specific Snapshot Policy in PowerMax array.
---

# powermax_snapshotpolicy (Resource)

Resource for a specific Snapshot Policy in PowerMax array.


## Example Usage

```terraform
/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.
Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://mozilla.org/MPL/2.0/
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

resource "powermax_snapshotpolicy" "terraform_sp" {
  # Required Field
  snapshot_policy_name = "terraform_sp"

  interval = "2 Hours"

  // Default values defined for some of the optional Fields
  # interval             = "1 Hour"
  # snapshot_count       = "48"
  # compliance_count_critical = 46
  # compliance_count_warning  = 47
  # offset_minutes            = 420
  # secure = false

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `snapshot_policy_name` (String) Name of the snapshot policy. Only alphanumeric characters, underscores ( _ ), and hyphens (-) are allowed and max length can be 32 characters

### Optional

- `compliance_count_critical` (Number) The threshold of good snapshots which are not failed/bad for compliance to change from warning to critical
- `compliance_count_warning` (Number) The threshold of good snapshots which are not failed/bad for compliance to change from normal to warning.
- `interval` (String) The interval between snapshots Enumeration values: 10 Minutes, 12 Minutes, 15 Minutes, 20 Minutes, 30 Minutes, 1 Hour, 2 Hours, 3 Hours, 4 Hours, 6 Hours, 8 Hours, 12 Hours, 1 Day, 7 Days
- `interval_minutes` (Number) Number of minutes between each policy execution
- `last_time_used` (String) The last time that the snapshot policy was run
- `offset_minutes` (Number) Number of minutes after 00:00 on Monday morning that the policy will execute
- `provider_name` (String) The name of the cloud provider associated with this policy. Only applies to cloud policies
- `retention_days` (Number) The number of days that snapshots will be retained in the cloud for. Only applies to cloud policies
- `secure` (Boolean) Set if the snapshot policy creates secure snapshots
- `snapshot_count` (Number) Number of snapshots that will be taken before the oldest ones are no longer required
- `suspended` (Boolean) Set if the snapshot policy has been suspended

### Read-Only

- `id` (String) Identifier
- `storage_group_count` (Number) The total number of storage groups that this snapshot policy is associated with
- `type` (String) The type of Snapshots that are created with the policy, local or cloud

## Import

Import is supported using the following syntax:

```shell
# Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://mozilla.org/MPL/2.0/
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# The command is
# terraform import powermax_storagegroup.test <id>
# Example:
terraform import powermax_snapshotpolicy.terraform_sp terraform_sp
# after running this command, populate the name field in the config file to start managing this resource
```