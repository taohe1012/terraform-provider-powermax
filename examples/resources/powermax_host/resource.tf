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

resource "powermax_host" "host_1" {
  name      = "host_1"
  initiator = ["10000000c9fc4b7e"]
  host_flags = {
    volume_set_addressing = {
      override = true
      enabled  = true
    }
    openvms = {
      override = true
      enabled  = false
    }
  }
}