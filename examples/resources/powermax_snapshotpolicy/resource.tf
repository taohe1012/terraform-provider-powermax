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