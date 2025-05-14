// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"os"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSaviyntADConnectionDataSource(t *testing.T) {
	filePath := "ad_connection_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_ad_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccADConnectionDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						res, ok := s.RootModule().Resources[datasource]
						if !ok {
							t.Fatalf("Resource %s not found in state", datasource)
						}
						t.Logf("Full data source attributes:\n%+v", res.Primary.Attributes)
						return nil
					},
					// Now assert values
					resource.TestCheckResourceAttr(datasource, "msg", "success"),
					resource.TestCheckResourceAttr(datasource, "error_code", "0"),
					resource.TestCheckResourceAttr(datasource, "connection_name", createCfg["connection_name"]),
					resource.TestCheckResourceAttr(datasource, "connection_type", createCfg["connection_type"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.url", createCfg["url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.username", createCfg["username"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.search_filter", createCfg["searchfilter"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.base", createCfg["base"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.group_search_base_dn", createCfg["group_search_base_dn"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.ldap_or_ad", createCfg["ldap_or_ad"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.object_filter", createCfg["objectfilter"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.account_attribute", createCfg["account_attribute"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.entitlement_attribute", createCfg["entitlement_attribute"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.page_size", createCfg["page_size"]),
				),
			},
		},
	})
}

func testAccADConnectionDataSourceConfig() string {
	jsonPath := "${filepath}/ad_connection_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
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
  
data "saviynt_ad_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_ad_connection_resource.ad]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
