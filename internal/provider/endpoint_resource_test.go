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

func TestAccSaviyntEndpointResource(t *testing.T) {
	filePath := "endpoint_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	updateCfg := util.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_endpoint_resource.endpoint"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccEndpointResourceConfig("create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoint_name"), knownvalue.StringExact(createCfg["endpoint_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(createCfg["display_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("security_system"), knownvalue.StringExact(createCfg["security_system"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("owner_type"), knownvalue.StringExact(createCfg["owner_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_property1"), knownvalue.StringExact(createCfg["custom_property1"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_custom_property_1_label"), knownvalue.StringExact(createCfg["account_custom_property_1_label"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_property31_label"), knownvalue.StringExact(createCfg["custom_property31_label"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_copy_access"), knownvalue.StringExact(createCfg["enable_copy_access"])),
				},
			},
			// Import
			{
				ResourceName:      resourceName,
				ImportStateId:     createCfg["endpoint_name"],
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string {"endpoint_config"},
			},
			// Update
			{
				Config: testAccEndpointResourceConfig("update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoint_name"), knownvalue.StringExact(updateCfg["endpoint_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(updateCfg["display_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("security_system"), knownvalue.StringExact(updateCfg["security_system"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("owner_type"), knownvalue.StringExact(updateCfg["owner_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_property1"), knownvalue.StringExact(updateCfg["custom_property1"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_custom_property_1_label"), knownvalue.StringExact(updateCfg["account_custom_property_1_label"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_property31_label"), knownvalue.StringExact(updateCfg["custom_property31_label"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_copy_access"), knownvalue.StringExact(updateCfg["enable_copy_access"])),
				},
			},
			// Update the Endpoint name to a new value
			{
				Config:      testAccEndpointResourceConfig("update_endpoint_name"),
				ExpectError: regexp.MustCompile(`Endpoint name cannot be updated`),
			},
			// Create a new resource with the same Endpoint name
			{
				Config:      testAccEndpointWithSameNameConfig("create_duplicate_endpoint"),
				ExpectError: regexp.MustCompile(`Endpoint name already exists`),
			},
		},
	})
}

func testAccEndpointResourceConfig(operation string) string {
	jsonPath := "{filepath}/endpoint_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_endpoint_resource" "endpoint" {
  endpoint_name                      = local.cfg.endpoint_name
  display_name                       = local.cfg.display_name
  security_system                    = local.cfg.security_system
  owner_type					   	 = local.cfg.owner_type	
  custom_property1            		 = local.cfg.custom_property1
  account_custom_property_1_label 	 = local.cfg.account_custom_property_1_label
  custom_property31_label 			 = local.cfg.custom_property31_label
  enable_copy_access				 = local.cfg.enable_copy_access
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}

func testAccEndpointWithSameNameConfig(operation string) string {
	jsonPath := "{filepath}/endpoint_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_endpoint_resource" "endpoint_1" {
  endpoint_name                      = local.cfg.endpoint_name
  display_name                       = local.cfg.display_name
  security_system                    = local.cfg.security_system
  owner_type					   	 = local.cfg.owner_type	
  custom_property1            		 = local.cfg.custom_property1
  account_custom_property_1_label 	 = local.cfg.account_custom_property_1_label
  custom_property31_label 			 = local.cfg.custom_property31_label
  enable_copy_access				 = local.cfg.enable_copy_access
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}