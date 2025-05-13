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

func TestAccSaviyntSalesforceConnectionResource(t *testing.T) {
	filePath := "salesforce_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	updateCfg := util.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_salesforce_connection_resource.ss"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Step
			{
				Config: testAccSalesforceConnectionResourceConfig("create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.StringExact(createCfg["client_id"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("redirect_uri"), knownvalue.StringExact(createCfg["redirect_uri"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("instance_url"), knownvalue.StringExact(createCfg["instance_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("object_to_be_imported"), knownvalue.StringExact(createCfg["object_to_be_imported"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("createaccountjson"), knownvalue.StringExact(createCfg["createaccountjson"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_threshold_config"), knownvalue.StringExact(createCfg["status_threshold_config"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportStateId:           createCfg["connection_name"],
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"msg", "client_secret", "refresh_token"},
			},
			// Update Step
			{
				Config: testAccSalesforceConnectionResourceConfig("update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.StringExact(updateCfg["client_id"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("redirect_uri"), knownvalue.StringExact(updateCfg["redirect_uri"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("instance_url"), knownvalue.StringExact(updateCfg["instance_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("object_to_be_imported"), knownvalue.StringExact(updateCfg["object_to_be_imported"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("createaccountjson"), knownvalue.StringExact(updateCfg["createaccountjson"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_threshold_config"), knownvalue.StringExact(createCfg["status_threshold_config"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccSalesforceConnectionResourceConfig("update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
			// Update the Connectiontype to a new value
			{
				Config:      testAccSalesforceConnectionResourceConfig("update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot be updated`),
			},
		},
	})
}

func testAccSalesforceConnectionResourceConfig(operation string) string {
	jsonPath := "{filepath}/salesforce_connection_test_data.json"
	return fmt.Sprintf(`
	provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

  resource "saviynt_salesforce_connection_resource" "ss" {
  connection_type       = local.cfg.connection_type
  connection_name       = local.cfg.connection_name
  client_id             = local.cfg.client_id
  redirect_uri          = local.cfg.redirect_uri
  instance_url          = local.cfg.instance_url
  object_to_be_imported = local.cfg.object_to_be_imported
  createaccountjson = jsonencode(local.cfg.createaccountjson)
  status_threshold_config=jsonencode(local.cfg.status_threshold_config)
}`,
		os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
