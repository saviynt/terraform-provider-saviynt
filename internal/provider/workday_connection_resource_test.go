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

func TestAccSaviyntWorkdayConnectionResource(t *testing.T) {
  filePath := "workday_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	updateCfg := util.LoadConnectorData(t, filePath, "update")
	resourceName:="saviynt_workday_connection_resource.w"
  t.Logf("Status key json: %q", createCfg["status_key_json"])
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Step
			{
				Config: testAccWorkdayConnectionResourceConfig("create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("base_url"), knownvalue.StringExact(createCfg["base_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("api_version"), knownvalue.StringExact(createCfg["api_version"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("use_oauth"), knownvalue.StringExact(createCfg["use_oauth"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact(createCfg["username"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.StringExact(createCfg["client_id"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_import_list"), knownvalue.StringExact(createCfg["access_import_list"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_key_json"), knownvalue.StringExact(createCfg["status_key_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_payload"), knownvalue.StringExact(createCfg["user_import_payload"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_mapping"), knownvalue.StringExact(createCfg["user_import_mapping"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Import
			{
				ResourceName:      resourceName,
				ImportStateId: createCfg["connection_name"],
				ImportState:       true,
				ImportStateVerify: true,
        		ImportStateVerifyIgnore: []string{"msg", "client_secret", "password", "refresh_token"},
			},
			// Update Step
			{
				Config: testAccWorkdayConnectionResourceConfig("update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("base_url"), knownvalue.StringExact(updateCfg["base_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("api_version"), knownvalue.StringExact(updateCfg["api_version"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("use_oauth"), knownvalue.StringExact(updateCfg["use_oauth"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact(updateCfg["username"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.StringExact(updateCfg["client_id"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_import_list"), knownvalue.StringExact(updateCfg["access_import_list"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_key_json"), knownvalue.StringExact(updateCfg["status_key_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_payload"), knownvalue.StringExact(updateCfg["user_import_payload"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_mapping"), knownvalue.StringExact(updateCfg["user_import_mapping"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
      // Update the Connectionname to a new value
			{
				Config:      testAccWorkdayConnectionResourceConfig("update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
      // Update the Connectiontype to a new value
			{
				Config:      testAccWorkdayConnectionResourceConfig("update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot by updated`),
			},
		},
	},
)
}

func testAccWorkdayConnectionResourceConfig(operation string) string {
  jsonPath:="{path}/workday_connection_test_data.json"
	return fmt.Sprintf(`
	provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}
  locals {
  cfg = jsondecode(file("%s"))["%s"]
}

  resource "saviynt_workday_connection_resource" "w" {
  connection_type    = local.cfg.connection_type
  connection_name    = local.cfg.connection_name
  base_url           = local.cfg.base_url
  api_version        = local.cfg.api_version
  use_oauth          = local.cfg.use_oauth
  username           = local.cfg.username
  client_id          = local.cfg.client_id
  access_import_list = local.cfg.access_import_list
  status_key_json = jsonencode(local.cfg.status_key_json)

  user_import_payload = local.cfg.user_import_payload

  user_import_mapping = jsonencode(local.cfg.user_import_mapping)
  }`,
 os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"), 
    jsonPath, 
    operation,
	)
}

