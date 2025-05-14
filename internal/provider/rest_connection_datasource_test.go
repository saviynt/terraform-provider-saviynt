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

func TestAccSaviyntRESTConnectionDataSource(t *testing.T) {
	filePath := "rest_connection_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_rest_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRESTConnectionDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.import_user_json", createCfg["import_user_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.create_account_json", createCfg["create_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.update_account_json", createCfg["update_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.add_access_json", createCfg["add_access_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.remove_access_json", createCfg["remove_access_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.remove_account_json", createCfg["remove_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.import_account_ent_json", createCfg["import_account_ent_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.change_pass_json", createCfg["change_pass_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.disable_account_json", createCfg["disable_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.enable_account_json", createCfg["enable_account_json"]),
				),
			},
		},
	})
}

func testAccRESTConnectionDataSourceConfig() string {
	jsonPath := "${filepath}/rest_connection_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_rest_connection_resource" "rest" {
  connection_type            = local.cfg.connection_type
  connection_name            = local.cfg.connection_name
  connection_json            = jsonencode(local.cfg.connection_json)
  import_user_json           = jsonencode(local.cfg.import_user_json)
  create_account_json        = jsonencode(local.cfg.create_account_json)
  update_account_json        = jsonencode(local.cfg.update_account_json)
  add_access_json            = jsonencode(local.cfg.add_access_json)
  remove_access_json         = jsonencode(local.cfg.remove_access_json)
  remove_account_json        = jsonencode(local.cfg.remove_account_json)
  import_account_ent_json    = jsonencode(local.cfg.import_account_ent_json)
  change_pass_json           = jsonencode(local.cfg.change_pass_json)
  disable_account_json       = jsonencode(local.cfg.disable_account_json)
  enable_account_json        = jsonencode(local.cfg.enable_account_json)
}
  
data "saviynt_rest_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_rest_connection_resource.rest]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
