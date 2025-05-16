// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccSaviyntUnixConnectionResource(t *testing.T) {
	filePath := "unix_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	updateCfg := util.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_unix_connection_resource.unix"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Step
			{
				Config: testAccUnixConnectionResourceConfig("create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("host_name"), knownvalue.StringExact(createCfg["host_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("port_number"), knownvalue.StringExact(createCfg["port_number"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact(createCfg["username"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("groups_file"), knownvalue.StringExact(createCfg["groups_file"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_threshold_config"), knownvalue.StringExact(createCfg["status_threshold_config"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ssh_key"), knownvalue.StringExact(createCfg["ssh_key"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Import
			{
				ResourceName:      resourceName,
				ImportStateId:     createCfg["connection_name"],
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore:  []string {"msg", "ssh_key", "password"},
			},
			// Update Step
			{
				Config: testAccUnixConnectionResourceConfig("update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("host_name"), knownvalue.StringExact(updateCfg["host_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("port_number"), knownvalue.StringExact(updateCfg["port_number"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact(updateCfg["username"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("groups_file"), knownvalue.StringExact(updateCfg["groups_file"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_threshold_config"), knownvalue.StringExact(updateCfg["status_threshold_config"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ssh_key"), knownvalue.StringExact(updateCfg["ssh_key"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccUnixConnectionResourceConfig("update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
			// Update the Connectiontype to a new value
			{
				Config:      testAccUnixConnectionResourceConfig("update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot by updated`),
			},
		},
	})
}

func testAccUnixConnectionResourceConfig(operation string) string {
	jsonPath:="{filepath}/unix_connection_test_data.json"
	return fmt.Sprintf(`
	provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}
  locals {
  cfg = jsondecode(file("%s"))["%s"]
}

  resource "saviynt_unix_connection_resource" "unix" {
  connection_type                    = local.cfg.connection_type
  connection_name                    = local.cfg.connection_name
  host_name       = local.cfg.host_name
  port_number     = local.cfg.port_number
  username        = local.cfg.username
  password        = local.cfg.password

  groups_file   = local.cfg.groups_file
  status_threshold_config = jsonencode(local.cfg.status_threshold_config)
  ssh_key = local.cfg.ssh_key
}
  `, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"), 
		jsonPath, 
    operation,
	)
}
