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

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSnapshotResource(t *testing.T) {
	var snapshotTerraformName = "powermax_snapshot.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: ProviderConfig + SnapshotResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(snapshotTerraformName, "name", "tfacc_snapshot_1"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "linked_storage_group.#", "0"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "linked", "false"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "restored", "false"),
					resource.TestCheckNoResourceAttr(snapshotTerraformName, "secure_expiry_date"),
				),
			},
			// ImportState testing
			{
				ResourceName:      snapshotTerraformName,
				ImportStateId:     "tfacc_sg_snapshot.tfacc_snapshot_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update Link and Rename
			{
				Config: ProviderConfig + SnapshotResourceLinkConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(snapshotTerraformName, "name", "tfacc_snapshot_2"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "linked_storage_group.#", "2"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "linked", "true"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "restored", "false"),
					resource.TestCheckNoResourceAttr(snapshotTerraformName, "secure_expiry_date"),
				),
			},
			// Update Unlink
			{
				Config: ProviderConfig + SnapshotResourceUnlinkConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(snapshotTerraformName, "name", "tfacc_snapshot_3"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "linked_storage_group.#", "0"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "linked", "false"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "restored", "false"),
					resource.TestCheckNoResourceAttr(snapshotTerraformName, "secure_expiry_date"),
				),
			},
			// Update Restore
			{
				Config: ProviderConfig + SnapshotResourceRestore,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(snapshotTerraformName, "name", "tfacc_snapshot_4"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "linked_storage_group.#", "0"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "linked", "false"),
					resource.TestCheckResourceAttr(snapshotTerraformName, "restored", "true"),
					resource.TestCheckNoResourceAttr(snapshotTerraformName, "secure_expiry_date"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

var SnapshotResourceConfig = `
resource "powermax_snapshot" "test" {
	storage_group {
		name = "tfacc_sg_snapshot"
	}
	snapshot_actions {
		# Required, name of new snapshot
		name = "tfacc_snapshot_1"
	}
}
`
var SnapshotResourceLinkConfig = `
resource "powermax_snapshot" "test" {
	storage_group {
		name = "tfacc_sg_snapshot"
	}
	snapshot_actions {
		# Required, name of new snapshot
		name = "tfacc_snapshot_2"
		link = {
		   enable = true
		   target_storage_group = "tfacc_test_target_snapshot_sg"
		   no_compression = true
		   remote = false
		   copy = false
		}
		time_to_live = {
		   enable = true
		   time_in_hours = true
		   time_to_live = 1
		}
	}
}
`

var SnapshotResourceUnlinkConfig = `
resource "powermax_snapshot" "test" {
	storage_group {
		name = "tfacc_sg_snapshot"
	}
	snapshot_actions {
		# Required, name of new snapshot
		name = "tfacc_snapshot_3"
		link = {
		   enable = false
		   target_storage_group = "tfacc_test_target_snapshot_sg"
		   no_compression = true
		   remote = false
		   copy = false
		}
		time_to_live = {
		   enable = true
		   time_in_hours = true
		   time_to_live = 1
		}
	}
}
`

var SnapshotResourceRestore = `
resource "powermax_snapshot" "test" {
	storage_group {
		name = "tfacc_sg_snapshot"
	}
	snapshot_actions {
		# Required, name of new snapshot
		name = "tfacc_snapshot_4"

		restore = {
		   enable = true
		   remote = false
		}
		time_to_live = {
			enable = true
			time_in_hours = true
			time_to_live = 1
		 }
	}
}
`
