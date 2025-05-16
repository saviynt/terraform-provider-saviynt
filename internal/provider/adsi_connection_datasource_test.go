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

func TestAccSaviyntADSIConnectionDataSource(t *testing.T) {
	filePath := "adsi_connection_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_adsi_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccADSIConnectionDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.connection_url", createCfg["connection_url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.provisioning_url", createCfg["provisioning_url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.forest_list", createCfg["forestlist"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.update_account_json", createCfg["updateaccountjson"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.add_access_json", createCfg["addaccessjson"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.search_filter", createCfg["searchfilter"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.object_filter", createCfg["objectfilter"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.account_attribute", createCfg["account_attribute"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.status_threshold_config", createCfg["status_threshold_config"]),
				),
			},
		},
	})
}

func testAccADSIConnectionDataSourceConfig() string {
	jsonPath := "${filepath}/adsi_connection_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_adsi_connection_resource" "adsi" {
  connection_type              = local.cfg.connection_type
  connection_name              = local.cfg.connection_name
  url                          = local.cfg.url
  password                     = local.cfg.password
  username                     = local.cfg.username
  connection_url               = local.cfg.connection_url
  provisioning_url             = local.cfg.provisioning_url
  forestlist                   = local.cfg.forestlist
  searchfilter                 = local.cfg.searchfilter
  group_search_base_dn         = local.cfg.group_search_base_dn
  objectfilter                 = local.cfg.objectfilter
  account_attribute            = local.cfg.account_attribute
  entitlement_attribute        = local.cfg.entitlement_attribute
  user_attribute               = local.cfg.user_attribute
  page_size                    = tostring(local.cfg.page_size)
  import_nested_membership     = tostring(local.cfg.import_nested_membership)
  statuskeyjson                = jsonencode(local.cfg.statuskeyjson)
  status_threshold_config      = jsonencode(local.cfg.status_threshold_config)
  checkforunique               = jsonencode(local.cfg.checkforunique)
  group_import_mapping         = jsonencode(local.cfg.group_import_mapping)
  createaccountjson          = jsonencode(local.cfg.createaccountjson)
  updateaccountjson          = jsonencode(local.cfg.updateaccountjson)
  enableaccountjson          = jsonencode(local.cfg.enableaccountjson)
  disableaccountjson         = jsonencode(local.cfg.disableaccountjson)
  removeaccountjson          = jsonencode(local.cfg.removeaccountjson)
  addaccessjson                = jsonencode(local.cfg.addaccessjson)
  removeaccessjson             = jsonencode(local.cfg.removeaccessjson)
  resetandchangepasswrdjson    = jsonencode(local.cfg.resetandchangepasswrdjson)
  creategroupjson            = jsonencode(local.cfg.creategroupjson)
  updategroupjson            = jsonencode(local.cfg.updategroupjson)
  removegroupjson            = jsonencode(local.cfg.removegroupjson)
  addaccessentitlementjson     = jsonencode(local.cfg.addaccessentitlementjson)
  removeaccessentitlementjson  = jsonencode(local.cfg.removeaccessentitlementjson)
}
  
data "saviynt_adsi_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_adsi_connection_resource.adsi]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
