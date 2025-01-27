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

title: "powermax_volume data source"
linkTitle: "powermax_volume"
page_title: "powermax_volume Data Source - terraform-provider-powermax"
subcategory: ""
description: |-
  Data source for reading Volumes in PowerMax array. PowerMax volumes is an identifiable unit of data storage. Storage groups are sets of volumes.
---

# powermax_volume (Data Source)

Data source for reading Volumes in PowerMax array. PowerMax volumes is an identifiable unit of data storage. Storage groups are sets of volumes.

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

# This terraform DataSource is used to query the existing volume from PowerMax array.
# The information fetched from this data source can be used for getting the details / for further processing in resource block.

# Returns all of the PowerMax volumes and their details
# NOTE: PowerMax can have many volumes, running this command unfiltered can take several minutes
data "powermax_volume" "volume_datasource_all" {
}

output "volume_datasource_output" {
  value = data.powermax_volume.volume_datasource_all
}

# Returns a subset of the PowerMax volumes based on the filtered properties in the block
# All filter values are optional
# If you use more then one filter at a time, it will only show the subset of volumes which both of those filters satisfies 
data "powermax_volume" "volume_datasource_test" {
  filter {

    # Optional Volume ids from a single Storage Group only
    storage_group_name = "terraform_vol_sg"

    # Optional Volume ids that contain the specified volume wwn
    wwn = "wwn_num"

    # Optiona lVolume ids that contain the specified volume encapsulated_wwn
    encapsulated_wwn = "encapsulated_wwn_num"

    # Optional Volume ids that contain greater than(">1"), Less than("<1") or equal to the specified symmlun
    symmlun = "0"

    # Optional Volume ids that contain the specified volume status
    status = "Ready"

    # Optional Volume ids that contain the specified volume physical_name
    physical_name = "physical_name"

    # Optional Volume ids that contain the specified volume volume_identifier
    volume_identifier = "test_acc_create_volume"

    # Optional Volume ids that contain greater than(">1"), Less than("<1") or equal to the specified cap_tb
    cap_tb = "0"

    # Optional Volume ids that contain greater than(">1"), Less than("<1") or equal to the specified cap_gb
    cap_gb = "0"

    # Optional Volume ids that contain greater than(">1"), Less than("<1") or equal to the specified cap_mb
    cap_mb = "0"

    # Optional Volume ids that contain greater than(">1"), Less than("<1") or equal to the specified cap_cyl
    cap_cyl = "0"

    # Optional Volume ids that contain greater than(">1"), Less than("<1") or equal to the specified allocated_percent
    allocated_percent = "0"

    # Optional Volume ids that contain greater than(">1"), Less than("<1") or equal to the specified num_of_storage_groups
    num_of_storage_groups = "1"

    # Optional Volume ids that contain greater than(">1"), Less than("<1") or equal to the specified num_of_masking_views
    num_of_masking_views = "0"

    # Optional Volume ids that contain greater than(">1"), Less than("<1") or equal to the specified num_of_front_end_paths
    num_of_front_end_paths = "0"

    # Optional Volume ids that are mobility ID enabled (true/false)
    mobility_id_enabled = false

    # Optional Volume ids that are virtual_volumes (true/false)
    virtual_volumes = true

    # Optional Volume ids that are private_volumes (true/false)
    private_volumes = false

    # Optional Volume ids that are tdev (true/false)
    tdev = true

    # Optional Volume ids that are vdev (true/false)
    vdev = false

    # Optional Volume ids that are available_thin_volumes (true/false)
    available_thin_volumes = false

    # Optional Volume ids that are gatekeeper (true/false)
    gatekeeper = false

    # Optional Volume ids that are data_volume (true/false)
    data_volume = false

    # Optional Volume ids that are dld (true/false)
    dld = false

    # Optional Volume ids that are drv (true/false)
    drv = false

    # Optional Volume ids that are encapsulated (true/false)
    encapsulated = false

    # Optional Volume ids that are associated (true/false)
    associated = false

    # Optional Volume ids that are reserved (true/false)
    reserved = false

    # Optional	Volume ids that are pinned (true/false)
    pinned = false

    # Optional Volume ids that are mapped (true/false)
    mapped = false

    # Optional Volume ids that are bound_tdev (true/false)
    bound_tdev = true

    # Optional Volume ids that are of the specified emulation
    emulation = "FBA"

    # Optional Volume ids that are of the specified emulation.
    has_effective_wwn = false

    # Optional Volume ids that contain the specified effective_wwn
    effective_wwn = "effective_wwn"

    # Optional Volume ids that are mapped to CU images associated to the specified FICON split
    split_name = "split_name"

    # Optional Volume ids that contain the specified volume type
    type = "TDEV"

    # Optional Volume ids that contain greater than("unreducible_data_gb=>1"),Less than("unreducible_data_gb=<1") or equal to the unreducible_data_gb
    unreducible_data_gb = "0"

    # Optional Volume ids that are mapped to a CU image with the specified CU image number
    cu_image_num = "0"

    # Optional Volume ids that are mapped to a CU image with the specified CU SSID
    cu_image_ssid = "cu_image_ssid"

    # Optional Volume ids that are part of the specified rdf group
    rdf_group_number = "0"

    # Optional Volume ids that contain the specified Oracle Instance Name
    oracle_instance_name = "oracle_instance_name"

    # Optional Volumes Ids that correspond to Namespace Globally Unique Identifier that uses the EUI64 16-byte designator format. Used in conjunction with NVMe volumes
    nguid = "nguid"
  }
}

output "volume_datasource_output" {
  value = data.powermax_volume.volume_datasource_test
}

# After the successful execution of above said block, We can see the output value by executing 'terraform output' command.
# Also, we can use the fetched information by the variable data.powermax_volume.example
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Block, Optional) (see [below for nested schema](#nestedblock--filter))

### Read-Only

- `id` (String) Placeholder for acc testing
- `volumes` (Attributes List) List of volumes. (see [below for nested schema](#nestedatt--volumes))

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Optional:

- `allocated_percent` (String) Greater than, Less than or equal to the allocated percent.
- `associated` (Boolean) Volumes that are associated (true/false).
- `available_thin_volumes` (Boolean) Volumes that are available thin volumes (true/false).
- `bound_tdev` (Boolean) Volumes that are bound tdev (true/false).
- `cap_cyl` (String) Greater than, Less than or equal to the cap CYL.
- `cap_gb` (String) Greater than, Less than or equal to the cap gb.
- `cap_mb` (String) Greater than, Less than or equal to the cap mb.
- `cap_tb` (String) Greater than, Less than or equal to the cap tb.
- `cu_image_num` (String) Volumes that are mapped to a CU image with the specified CU image number.
- `cu_image_ssid` (String) Volumes that are mapped to a CU image with the specified CU SSID.
- `data_volume` (Boolean) Volumes that are data volume (true/false).
- `dld` (Boolean) Volumes that are dld (true/false).
- `drv` (Boolean) Volumes that are drv (true/false).
- `effective_wwn` (String) Volumes that contain the specified effective_wwn.
- `emulation` (String) Volumes that are of the specified emulation.
- `encapsulated` (Boolean) Volumes that are encapsulated (true/false).
- `encapsulated_wwn` (String) The specified volume encapsulated_wwn.
- `gatekeeper` (Boolean) Volumes that are gatekeeper (true/false).
- `has_effective_wwn` (Boolean) Volumes that have effective wwns (true/false)
- `mapped` (Boolean) Volumes that are mapped (true/false).
- `mobility_id_enabled` (Boolean) Volumes that are mobility ID enabled (true/false).
- `nguid` (String) Volumes that correspond to Namespace Globally Unique Identifier that uses the EUI64 16-byte designator format.
- `num_of_front_end_paths` (String) Greater than, Less than or equal to the number of front end paths.
- `num_of_masking_views` (String) Greater than, Less than or equal to the number of masking views.
- `num_of_storage_groups` (String) Greater than, Less than or equal to the number of storage groups.
- `oracle_instance_name` (String) Volumes that contain the specified Oracle Instance Name.
- `physical_name` (String) The specified volume physical name.
- `pinned` (Boolean) Volumes that are pinned (true/false).
- `private_volumes` (Boolean) Volumes that are private (true/false).
- `rdf_group_number` (String) Volumes that are part of the specified rdf group.
- `reserved` (Boolean) Volumes that are reserved (true/false).
- `split_name` (String) Volumes that are mapped to CU images associated to the specified FICON split.
- `status` (String) The specified volume status.
- `storage_group_name` (String) The name of the storage group.
- `symmlun` (String) Greater than, Less than or equal to the specified symmlun.
- `tdev` (Boolean) Volumes that are tdev (true/false).
- `thin_bcv` (Boolean) Volumes that are thin bcv (true/false).
- `type` (String) Volumes that contain the specified volume type.
- `unreducible_data_gb` (String) Greater than,Less than or equal to the unreducible data gb.
- `vdev` (Boolean) Volumes that are vdev (true/false).
- `virtual_volumes` (Boolean) Volumes that are virtual volumes (true/false).
- `volume_identifier` (String) The specified volume volume identifier.
- `wwn` (String) The specified volume wwn.


<a id="nestedatt--volumes"></a>
### Nested Schema for `volumes`

Optional:

- `mobility_id_enabled` (Boolean) States whether mobility ID is enabled on the volume.
- `volume_identifier` (String) The identifier of the volume.

Read-Only:

- `allocated_percent` (Number) The allocated percentage of the volume.
- `cap_cyl` (Number) The capability of volume in the unit of CYL.
- `cap_gb` (Number) The capability of volume in the unit of GB.
- `cap_mb` (Number) The capability of volume in the unit of MB.
- `effective_wwn` (String) Effective WWN of the volume.
- `emulation` (String) The emulation of the volume Enumeration values.
- `encapsulated` (Boolean) States whether the volume is encapsulated.
- `encapsulated_wwn` (String) Encapsulated  WWN of the volume.
- `has_effective_wwn` (Boolean) States whether volume has effective WWN.
- `id` (String) The ID of the volume.
- `nguid` (String) The NGUID of the volume.
- `num_of_front_end_paths` (Number) The number of front end paths of the volume.
- `num_of_storage_groups` (Number) The number of storage groups associated with the volume.
- `oracle_instance_name` (String) Oracle instance name associated with the volume.
- `physical_name` (String) The physical name of the volume.
- `pinned` (Boolean) States whether the volume is pinned.
- `rdf_group_ids` (Attributes List) The RDF groups associated with the volume. (see [below for nested schema](#nestedatt--volumes--rdf_group_ids))
- `reserved` (Boolean) States whether the volume is reserved.
- `snapvx_source` (Boolean) States whether the volume is a snapvx source.
- `snapvx_target` (Boolean) States whether the volume is a snapvx target.
- `ssid` (String) The ssid of the volume.
- `status` (String) The status of the volume.
- `storage_groups` (Attributes List) List of storage groups which are associated with the volume. (see [below for nested schema](#nestedatt--volumes--storage_groups))
- `symmetrix_port_key` (Attributes List) The symmetrix ports associated with the volume. (see [below for nested schema](#nestedatt--volumes--symmetrix_port_key))
- `type` (String) The type of the volume.
- `unreducible_data_gb` (Number) The amount of unreducible data in Gb.
- `wwn` (String) The WWN of the volume.

<a id="nestedatt--volumes--rdf_group_ids"></a>
### Nested Schema for `volumes.rdf_group_ids`

Read-Only:

- `label` (String) The label of the rdf group.
- `rdf_group_number` (Number) The number of rdf group.


<a id="nestedatt--volumes--storage_groups"></a>
### Nested Schema for `volumes.storage_groups`

Read-Only:

- `parent_storage_group_name` (String) The ID of the storage group parents.
- `storage_group_name` (String) The ID of the storage group.


<a id="nestedatt--volumes--symmetrix_port_key"></a>
### Nested Schema for `volumes.symmetrix_port_key`

Read-Only:

- `director_id` (String) The ID of the director.
- `port_id` (String) The ID of the symmetrix port.