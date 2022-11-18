package powermax

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	TestAccHostName1   = "test_acc_chost11"
	TestAccHostName2   = "test_acc_chost1"
	TestAccHostName3   = "test_acc_chost2"
	TestAccHostName4   = "test_acc_chost3"
	TestAccHostName5   = "test_acc_chost4"
	InvalidInitiatorID = "0000000000000000"
)

func TestAccHost_CreateHost(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: CreateHostParams,
				Check:  resource.ComposeTestCheckFunc(checkCreateHost(t, testProvider, TestAccHostName1)),
			},
		},
	})
}

func TestAccHost_CreateHostWithOptionalFlags(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: CreateHostParamsWithOptionalFlags,
				Check:  resource.ComposeTestCheckFunc(checkCreateHost(t, testProvider, TestAccHostName2)),
			},
		},
	})
}

func TestAccHost_CreateHostWithOptionalFlagsOverride(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: CreateHostParamsWithOptionalFlagsOverride,
				Check:  resource.ComposeTestCheckFunc(checkCreateHost(t, testProvider, TestAccHostName3)),
			},
		},
	})
}

func TestAccHost_CreateHostWithInvalidInitiator(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      CreateHostParamsWithInvalidInitiator,
				ExpectError: regexp.MustCompile(CreateHostDetailErrorMsg),
			},
		},
	})
}

func TestAccHost_UpdateHostRename(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: HostParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName4),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.#", "2"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.0", InitiatorID),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.1", InitiatorID2)),
			},
			{
				Config: HostParamsRename,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName5)),
			},
		},
	})
}

func TestAccHost_UpdateHostInitiatorsRemove(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: HostParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName4),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.#", "2"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.0", InitiatorID),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.1", InitiatorID2)),
			},
			{
				Config: HostParamsChangeInitiatorRemove,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName4),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.#", "1"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.0", InitiatorID)),
			},
		},
	})
}

func TestAccHost_UpdateHostAddInitiatorsAndNameChange(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: HostParamsForUpdate,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName5),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.#", "1"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.0", InitiatorID),
				),
			},
			{
				Config: HostParamsChangeAddInitiatorAndNameChange,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName4),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.#", "2"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.0", InitiatorID),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.1", InitiatorID2)),
			},
		},
	})
}

func TestAccHost_UpdateHostFlags(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: HostParamsForUpdate,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName5),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.#", "1"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.0", InitiatorID),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.volume_set_addressing.enabled", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.volume_set_addressing.override", "true")),
			},
			{
				Config: HostParamsChangeForUpdateFlags,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName5),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.#", "1"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.0", InitiatorID),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.volume_set_addressing.enabled", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.volume_set_addressing.override", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.avoid_reset_broadcast.enabled", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.avoid_reset_broadcast.override", "true")),
			},
		},
	})
}

func TestAccHost_UpdateHostFlagsInitiatorAndName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: HostParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName4),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.#", "2"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.0", InitiatorID),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.1", InitiatorID2),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.volume_set_addressing.enabled", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.volume_set_addressing.override", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.openvms.enabled", "false"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.openvms.override", "true")),
			},
			{
				Config: HostParamsChangeForUpdateFlagsInitiatorAndName,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "name", TestAccHostName5),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.#", "1"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "initiators.0", InitiatorID),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.volume_set_addressing.enabled", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.volume_set_addressing.override", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.avoid_reset_broadcast.enabled", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.avoid_reset_broadcast.override", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.openvms.enabled", "false"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.openvms.override", "true"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.scsi_3.enabled", "false"),
					resource.TestCheckResourceAttr("powermax_host.host_create_rename_test", "host_flags.scsi_3.override", "true")),
			},
		},
	})
}

func checkCreateHost(t *testing.T, p tfsdk.Provider, hostID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providers := p.(*provider)
		_, err := providers.client.PmaxClient.GetHostByID(context.Background(), serialno, hostID)
		if err != nil {
			return fmt.Errorf("failed to fetch host")
		}
		if !providers.configured {
			return fmt.Errorf("provider not configured")
		}

		if providers.client.PmaxClient == nil {
			return fmt.Errorf("provider not configured")
		}
		return nil
	}
}

var CreateHostParams = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_test" {
	name = "` + TestAccHostName1 + `"
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = true
		}
		openvms = {
			override = true
			enabled = false
		}
		avoid_reset_broadcast = {
			enabled = true
			override = true
		}
		scsi_3 = {
			enabled = true
			override = true
		}
		spc2_protocol_version = {
			enabled = false
			override = true
		}
		scsi_support1 = {
			enabled = false
			override = true
		}
		consistent_lun = false
	}
	initiators = ["` + InitiatorID + `"]
}
`

var CreateHostParamsWithOptionalFlags = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_test" {
	name = "` + TestAccHostName2 + `"
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = false
		}
		disable_q_reset_on_ua = {
			enabled = false
			override = false
		}
		environ_set = {
			enabled = false
			override = false
		}
		avoid_reset_broadcast = {
			enabled = false
			override = false
		}
		scsi_3 = {
			enabled = false
			override = false
		}
		openvms = {
			override = true
			enabled = true
		}
		spc2_protocol_version = {
			enabled = true
			override = true
		}
		scsi_support1 = {
			enabled = true
			override = true
		}
		consistent_lun = true
	}
	initiators = ["` + InitiatorID + `"]
}
`

var CreateHostParamsWithOptionalFlagsOverride = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_test" {
	name = "` + TestAccHostName3 + `"
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = false
		}
		disable_q_reset_on_ua = {
			override = true
			enabled = false
		}
		environ_set = {
			override = true
			enabled = false
		}
		avoid_reset_broadcast = {
			override = true
			enabled = false
		}
		scsi_3 = {
			override = true
			enabled = false
		}
		openvms = {
			override = false
			enabled = false
		}
		spc2_protocol_version = {
			enabled = false
			override = false
		}
		scsi_support1 = {
			enabled = false
			override = false
		}
		consistent_lun = true
	}
	initiators = ["` + InitiatorID + `"]
}
`

var CreateHostParamsWithInvalidInitiator = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_test" {
	name = "` + TestAccHostName2 + `"
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = false
		}
		disable_q_reset_on_ua = {
			enabled = false
			override = false
		}
		environ_set = {
			enabled = false
			override = false
		}
		avoid_reset_broadcast = {
			enabled = false
			override = false
		}
		scsi_3 = {
			enabled = false
			override = false
		}
		openvms = {
			override = true
			enabled = true
		}
		spc2_protocol_version = {
			enabled = true
			override = true
		}
		scsi_support1 = {
			enabled = true
			override = true
		}
		consistent_lun = true
	}
	initiators = ["` + InvalidInitiatorID + `"]	    
}
`
var HostParams = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_rename_test" {
	name = "` + TestAccHostName4 + `"
	initiators = ["` + InitiatorID + `","` + InitiatorID2 + `"]
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = true
		}
		openvms = {
			override = true
			enabled = false
		}
	}
}
`

var HostParamsRename = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_rename_test" {
	name = "` + TestAccHostName5 + `"
	initiators = ["` + InitiatorID + `"]
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = true
		}
		openvms = {
			override = true
			enabled = false
		}
	}
}
`

var HostParamsChangeInitiatorRemove = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_rename_test" {
	name = "` + TestAccHostName4 + `"
	initiators = ["` + InitiatorID + `"]
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = true
		}
		openvms = {
			override = true
			enabled = false
		}
	}
}
`

var HostParamsForUpdate = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_rename_test" {
	name = "` + TestAccHostName5 + `"
	initiators = ["` + InitiatorID + `"]
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = true
		}
		openvms = {
			override = true
			enabled = false
		}
	}
}
`

var HostParamsChangeAddInitiatorAndNameChange = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_rename_test" {
	name = "` + TestAccHostName4 + `"
	initiators = ["` + InitiatorID + `","` + InitiatorID2 + `"]
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = true
		}
		openvms = {
			override = true
			enabled = false
		}
	}
}
`

var HostParamsChangeForUpdateFlags = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_rename_test" {
	name = "` + TestAccHostName5 + `"
	initiators = ["` + InitiatorID + `"]
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = true
		}
		openvms = {
			override = true
			enabled = false
		}
		avoid_reset_broadcast = {
			override = true
			enabled = true
		}
	}
}
`

var HostParamsChangeForUpdateFlagsInitiatorAndName = `
provider "powermax" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	serial_number = "` + serialno + `"
	insecure = true
}

resource "powermax_host" "host_create_rename_test" {
	name = "` + TestAccHostName5 + `"
	initiators = ["` + InitiatorID + `"]
	host_flags = {
		volume_set_addressing = {
			override = true
			enabled = true
		}
		openvms = {
			override = true
			enabled = false
		}
		avoid_reset_broadcast = {
			override = true
			enabled = true
		}
		scsi_3 = {
			enabled = false
			override = true
		}
	}
}
`