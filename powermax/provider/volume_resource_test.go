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
	"dell/powermax-go-client"
	"fmt"
	"regexp"
	"terraform-provider-powermax/powermax/helper"
	"testing"

	. "github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccVolumeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: ProviderConfig + VolumeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "sg_name", resourceVolSGName),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "vol_name", resourceVolName),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "size", "2.45"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "cap_unit", "GB"),

					resource.TestCheckResourceAttr("powermax_volume.volume_test", "type", "TDEV"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "emulation", "FBA"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "allocated_percent", "0"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "status", "Ready"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "reserved", "false"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "pinned", "false"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "reserved", "false"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "encapsulated", "false"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "num_of_storage_groups", "1"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "num_of_front_end_paths", "0"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "snapvx_source", "false"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "snapvx_target", "false"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "has_effective_wwn", "false"),
				),
			},
			// ImportState testing
			{
				ResourceName: "powermax_volume.volume_test",
				ImportState:  true,
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					assert.Equal(t, resourceVolName, states[0].Attributes["vol_name"])
					assert.Equal(t, "2.45", states[0].Attributes["size"])
					assert.Equal(t, "GB", states[0].Attributes["cap_unit"])
					assert.Equal(t, "", states[0].Attributes["sg_name"])
					return nil
				},
			},
			// Update Name, Size, Mobility and Read testing
			{
				Config: ProviderConfig + VolumeUpdateNameSizeMobility,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "vol_name", resourceVolName+"_2"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "size", "1"),
					resource.TestCheckResourceAttr("powermax_volume.volume_test", "cap_unit", "TB"),
				),
			},
		},
	})
}

func TestAccVolumeResourceReadError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetVolume).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfig + VolumeUpdateNameSizeMobility,
				ExpectError: regexp.MustCompile(`.*Error creating volume*.`),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.UpdateVolResourceState).Return(fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfig + VolumeResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error creating volume*.`),
			},
		},
	})
}

func TestAccVolumeResource_Invalid_Config(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Config with no SG
			{
				Config:      ProviderConfig + VolumeConfigNoSG,
				ExpectError: regexp.MustCompile("Error creating volume"),
			},
			// Config with invalid unit
			{
				Config:      ProviderConfig + VolumeConfigInvalidCYL,
				ExpectError: regexp.MustCompile("Error creating volume"),
			},
			// Config with invalid SG name
			{
				Config:      ProviderConfig + VolumeConfigInvalidSG,
				ExpectError: regexp.MustCompile("Error creating volume"),
			},
		},
	})
}

func TestAccVolumeResource_Error_Updating(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Normal Config
			{
				Config:      ProviderConfig + VolumeConfigWithCYL,
				ExpectError: nil,
			},
			// Invalid SG name
			{
				Config:      ProviderConfig + VolumeConfigInvalidSG,
				ExpectError: regexp.MustCompile("Error updating volume"),
			},
			// Invalid name
			{
				Config:      ProviderConfig + VolumeConfigInvalidName,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Match"),
			},
			// Normal Config
			{
				Config:      ProviderConfig + VolumeConfigWithCYL,
				ExpectError: nil,
			},
		},
	})
}

func TestAccVolumeResourceCreateError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.CreateVolume).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfig + VolumeResourceConfig,
				ExpectError: regexp.MustCompile(`.*mock error*.`),
			},
		},
	})
}

func TestAccVolumeResourceListError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.ListVolumes).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfig + VolumeResourceConfig,
				ExpectError: regexp.MustCompile(`.*mock error*.`),
			},
		},
	})
}

func TestAccVolumeResourceNewVolumeMissingError(t *testing.T) {
	missingPmax := powermax.Iterator{
		ResultList: *powermax.NewResultListWithDefaults(),
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.ListVolumes).Return(&missingPmax, nil, nil).Build()
				},
				Config:      ProviderConfig + VolumeResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error creating volume*.`),
			},
		},
	})
}

func TestAccVolumeResourceVolumeDetailstError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetVolume).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfig + VolumeResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error creating volume*.`),
			},
		},
	})
}

func TestAccVolumeResourceVolumeMappertError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.UpdateVolResourceState).Return(fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfig + VolumeResourceConfig,
				ExpectError: regexp.MustCompile(`.*mock error*.`),
			},
		},
	})
}

var resourceVolSGName = "tfacc_ds_vol_sg_sOCTK"
var resourceVolName = "tfacc_res_vol"

var VolumeResourceConfig = fmt.Sprintf(`
resource "powermax_volume" "volume_test" {
	vol_name = "%s"
	size = 2.45
	cap_unit = "GB"
	sg_name = "%s"
}
`, resourceVolName, resourceVolSGName)

var VolumeUpdateNameSizeMobility = fmt.Sprintf(`
resource "powermax_volume" "volume_test" {
	vol_name = "%s_2"
	size = 1
	cap_unit = "TB"
	sg_name = "%s"
	mobility_id_enabled = "true"
}
`, resourceVolName, resourceVolSGName)

var VolumeConfigNoSG = fmt.Sprintf(`
resource "powermax_volume" "volume_test" {
	vol_name = "%s"
	size = 5
	cap_unit = "CYL"
}
`, resourceVolName)

var VolumeConfigInvalidCYL = fmt.Sprintf(`
resource "powermax_volume" "volume_test" {
	vol_name = "%s"
	sg_name = "%s"
	size = 0.5
	cap_unit = "CYL"
}
`, resourceVolName, resourceVolSGName)

var VolumeConfigWithCYL = fmt.Sprintf(`
resource "powermax_volume" "volume_test" {
	vol_name = "%s-modify"
	size = 1
	cap_unit = "CYL"
	sg_name = "%s"
}
`, resourceVolName, resourceVolSGName)

var VolumeConfigInvalidSG = fmt.Sprintf(`
resource "powermax_volume" "volume_test" {
	vol_name = "%s"
	sg_name = "invalid#SG"
	size = 0.5
	cap_unit = "CYL"
}
`, resourceVolName)

var VolumeConfigInvalidName = `
resource "powermax_volume" "volume_test" {
	vol_name = "!@#$%"
	sg_name = "tfacc_ds_vol_sg_volume_resource"
	size = 5
	cap_unit = "CYL"
}
`
