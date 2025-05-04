// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"os"
	"regexp"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccSaviyntSecuritySystemResource(t *testing.T) {
	filePath := "securitysystem_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	updateCfg := util.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_security_system_resource.ss"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccSecuritySystemConnectionResourceConfig("create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("systemname"), knownvalue.StringExact(createCfg["systemname"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(createCfg["display_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_add_workflow"), knownvalue.StringExact(createCfg["access_add_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_remove_workflow"), knownvalue.StringExact(createCfg["access_remove_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_service_account_workflow"), knownvalue.StringExact(createCfg["add_service_account_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_service_account_workflow"), knownvalue.StringExact(createCfg["remove_service_account_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("automated_provisioning"), knownvalue.StringExact(createCfg["automated_provisioning"])),
				},
			},
			// Import
			{
				ResourceName:      resourceName,
				ImportStateId:     createCfg["systemname"],
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: testAccSecuritySystemConnectionResourceConfig("update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("systemname"), knownvalue.StringExact(updateCfg["systemname"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(updateCfg["display_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_add_workflow"), knownvalue.StringExact(updateCfg["access_add_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_remove_workflow"), knownvalue.StringExact(updateCfg["access_remove_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_service_account_workflow"), knownvalue.StringExact(updateCfg["add_service_account_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_service_account_workflow"), knownvalue.StringExact(updateCfg["remove_service_account_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("automated_provisioning"), knownvalue.StringExact(updateCfg["automated_provisioning"])),
				},
			},
			// Update the Systemname to a new value
			{
				Config:      testAccSecuritySystemConnectionResourceConfig("update_security_system_name"),
				ExpectError: regexp.MustCompile(`System name cannot be updated`),
			},
			// Create a new resource with the same Systemname
			{
				Config:      testAccSecuritySecuritySystemWithSameNameConfig("create_duplicate_security_system"),
				ExpectError: regexp.MustCompile(`Security System Already Exists`),
			},
		},
	})
}

func testAccSecuritySystemConnectionResourceConfig(operation string) string {
	jsonPath := "${filepath}/securitysystem_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_security_system_resource" "ss" {
  systemname                         = local.cfg.systemname
  display_name                       = local.cfg.display_name
  access_add_workflow                = local.cfg.access_add_workflow
  access_remove_workflow             = local.cfg.access_remove_workflow
  add_service_account_workflow       = local.cfg.add_service_account_workflow
  remove_service_account_workflow    = local.cfg.remove_service_account_workflow
  automated_provisioning             = local.cfg.automated_provisioning
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}

func testAccSecuritySecuritySystemWithSameNameConfig(operation string) string {
	jsonPath := "/Users/shaleen.shukla/terraform-provider-saviynt/internal/provider/securitysystem_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_security_system_resource" "ss1" {
  systemname                         = local.cfg.systemname
  display_name                       = local.cfg.display_name
  access_add_workflow                = local.cfg.access_add_workflow
  access_remove_workflow             = local.cfg.access_remove_workflow
  add_service_account_workflow       = local.cfg.add_service_account_workflow
  remove_service_account_workflow    = local.cfg.remove_service_account_workflow
  automated_provisioning             = local.cfg.automated_provisioning
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
