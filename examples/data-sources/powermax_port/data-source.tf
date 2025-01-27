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

# This terraform DataSource is used to query the existing port from PowerMax array.
# The information fetched from this data source can be used for getting the details / for further processing in resource block.

# Returns all of the PowerMax ports and their details
data "powermax_port" "all" {}

output "all" {
  value = data.powermax_port.all
}

# Returns a subset of the PowerMax ports based on the names provided in the `names` filter block and their details
data "powermax_port" "portFilter" {
  # Optional list of names to filter upon
  filter {
    # Should be in the format ["directorId:portId"]
    port_ids = ["OR-1C:2"]
  }
}

output "portFilter" {
  value = data.powermax_port.portFilter
}

# After the successful execution of above said block, We can see the output value by executing 'terraform output' command.
# Also, we can use the fetched information by the variable data.powermax_port.example
