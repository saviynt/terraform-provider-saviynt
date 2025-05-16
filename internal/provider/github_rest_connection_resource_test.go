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

func TestAccSaviyntGithubRestConnectionResource(t *testing.T) {
	filePath := "github_rest_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	updateCfg := util.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_github_rest_connection_resource.gr"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Step
			{
				Config: testAccGithubRestConnectionResourceConfig("create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_json"), knownvalue.StringExact(createCfg["connection_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("import_account_ent_json"), knownvalue.StringExact(createCfg["import_account_ent_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("organization_list"), knownvalue.StringExact(createCfg["organization_list"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Import
			{
				ResourceName:      resourceName,
				ImportStateId:     createCfg["connection_name"],
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore:  []string {"msg", "connection_json"},
			},
			// Update Step
			{
				Config: testAccGithubRestConnectionResourceConfig("update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("import_account_ent_json"), knownvalue.StringExact(updateCfg["import_account_ent_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("organization_list"), knownvalue.StringExact(updateCfg["organization_list"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccGithubRestConnectionResourceConfig("update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
			// Update the Connectiontype to a new value
			{
				Config:      testAccGithubRestConnectionResourceConfig("update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot by updated`),
			},
		},
	})
}

func testAccGithubRestConnectionResourceConfig(operation string) string {
	jsonPath:="{filepath}/github_rest_connection_test_data.json"
	return fmt.Sprintf(`
	provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}
  locals {
  cfg = jsondecode(file("%s"))["%s"]
}

  resource "saviynt_github_rest_connection_resource" "gr" {
  connection_type                    = local.cfg.connection_type
  connection_name                    = local.cfg.connection_name
  connection_json                    = jsonencode(local.cfg.connection_json)
  import_account_ent_json			 = jsonencode(local.cfg.import_account_ent_json)
  organization_list                  = local.cfg.organization_list
}
  `, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"), 
		jsonPath, 
    operation,
	)
}

