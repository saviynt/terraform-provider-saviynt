package provider

import (
	"fmt"
	"os"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSaviyntDBConnectionDataSource(t *testing.T) {
	filePath := "db_connection_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_db_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDBConnectionDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.drivername", createCfg["driver_name"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.createaccount_json", createCfg["create_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.updateaccount_json", createCfg["update_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.grantaccess_json", createCfg["grant_access_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.revokeaccess_json", createCfg["revoke_access_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.deleteaccount_json", createCfg["delete_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.enableaccount_json", createCfg["enable_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.disableaccount_json", createCfg["disable_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.accountexists_json", createCfg["account_exists_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.updateuser_json", createCfg["update_user_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.accounts_import", createCfg["accounts_import"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.entitlementvalue_import", createCfg["entitlement_value_import"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.user_import", createCfg["user_import"]),
				),
			},
		},
	})
}

func testAccDBConnectionDataSourceConfig() string {
	jsonPath := "${filepath}/db_connection_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_db_connection_resource" "db" {
  connection_type           = local.cfg.connection_type
  connection_name           = local.cfg.connection_name
  url                       = local.cfg.url
  username                  = local.cfg.username
  password                  = local.cfg.password
  driver_name               = local.cfg.driver_name

  create_account_json       = jsonencode(local.cfg.create_account_json)
  update_account_json       = jsonencode(local.cfg.update_account_json)
  grant_access_json         = jsonencode(local.cfg.grant_access_json)
  revoke_access_json        = jsonencode(local.cfg.revoke_access_json)
  change_pass_json          = jsonencode(local.cfg.change_pass_json)
  delete_account_json       = jsonencode(local.cfg.delete_account_json)
  enable_account_json       = jsonencode(local.cfg.enable_account_json)
  disable_account_json      = jsonencode(local.cfg.disable_account_json)
  account_exists_json       = jsonencode(local.cfg.account_exists_json)
  update_user_json          = jsonencode(local.cfg.update_user_json)

  accounts_import           = local.cfg.accounts_import
  entitlement_value_import  = local.cfg.entitlement_value_import
  user_import               = local.cfg.user_import
}
  
data "saviynt_db_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_db_connection_resource.db]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
