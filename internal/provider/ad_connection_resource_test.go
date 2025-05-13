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

func TestAccSaviyntADConnectionResource(t *testing.T) {
	filePath := "ad_connection_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	updateCfg := util.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_ad_connection_resource.ad"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccADConnectionResourceConfig("create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(createCfg["url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact(createCfg["password"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact(createCfg["username"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("searchfilter"), knownvalue.StringExact(createCfg["searchfilter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("base"), knownvalue.StringExact(createCfg["base"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("group_search_base_dn"), knownvalue.StringExact(createCfg["group_search_base_dn"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ldap_or_ad"), knownvalue.StringExact(createCfg["ldap_or_ad"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("objectfilter"), knownvalue.StringExact(createCfg["objectfilter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_attribute"), knownvalue.StringExact(createCfg["account_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entitlement_attribute"), knownvalue.StringExact(createCfg["entitlement_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("page_size"), knownvalue.StringExact(createCfg["page_size"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_attribute"), knownvalue.StringExact(createCfg["user_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoints_filter"), knownvalue.StringExact(createCfg["endpoints_filter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_account_json"), knownvalue.StringExact(createCfg["create_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(createCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_user_json"), knownvalue.StringExact(createCfg["update_user_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportStateId:           createCfg["connection_name"],
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"msg", "error_code", "password"},
			},
			// Update
			{
				Config: testAccADConnectionResourceConfig("update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(updateCfg["url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact(updateCfg["password"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact(updateCfg["username"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("searchfilter"), knownvalue.StringExact(updateCfg["searchfilter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("base"), knownvalue.StringExact(updateCfg["base"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("group_search_base_dn"), knownvalue.StringExact(updateCfg["group_search_base_dn"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ldap_or_ad"), knownvalue.StringExact(updateCfg["ldap_or_ad"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("objectfilter"), knownvalue.StringExact(updateCfg["objectfilter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_attribute"), knownvalue.StringExact(updateCfg["account_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entitlement_attribute"), knownvalue.StringExact(updateCfg["entitlement_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("page_size"), knownvalue.StringExact(updateCfg["page_size"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_attribute"), knownvalue.StringExact(updateCfg["user_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoints_filter"), knownvalue.StringExact(updateCfg["endpoints_filter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_account_json"), knownvalue.StringExact(updateCfg["create_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(updateCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_user_json"), knownvalue.StringExact(updateCfg["update_user_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccADConnectionResourceConfig("update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
			// Update the Connectiontype to a new value
			{
				Config:      testAccADConnectionResourceConfig("update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot by updated`),
			},
		},
	})
}

func testAccADConnectionResourceConfig(operation string) string {
	jsonPath := "${filepath}/ad_connection_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_ad_connection_resource" "ad" {
  connection_type     = local.cfg.connection_type
  connection_name     = local.cfg.connection_name
  url                 = local.cfg.url
  password            = local.cfg.password
  username            = local.cfg.username
  searchfilter        = local.cfg.searchfilter
  base                = local.cfg.base
  group_search_base_dn= local.cfg.group_search_base_dn
  ldap_or_ad          = local.cfg.ldap_or_ad
  objectfilter        = local.cfg.objectfilter
  account_attribute   = local.cfg.account_attribute
  entitlement_attribute = local.cfg.entitlement_attribute
  page_size           = local.cfg.page_size
  user_attribute      = local.cfg.user_attribute

  endpoints_filter    = jsonencode(local.cfg.endpoints_filter)
  create_account_json = jsonencode(local.cfg.create_account_json)
  update_account_json = jsonencode(local.cfg.update_account_json)
  update_user_json    = jsonencode(local.cfg.update_user_json)
  enable_account_json = jsonencode(local.cfg.enable_account_json)

  account_name_rule   = local.cfg.account_name_rule
  remove_account_action = jsonencode(local.cfg.remove_account_action)
  set_random_password = local.cfg.set_random_password
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
