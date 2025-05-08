package provider

import (
	"fmt"
	"os"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSaviyntSapConnectionDataSource(t *testing.T) {
	filePath := "sap_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_sap_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSapConnectionDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.message_server", createCfg["message_server"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.jco_snc_mode", createCfg["jco_snc_mode"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.system_name", createCfg["system_name"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.user_import_json", createCfg["user_import_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.password_max_length", createCfg["password_max_length"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.update_account_json", createCfg["update_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.set_cua_system", createCfg["set_cua_system"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.saptable_filter_lang", createCfg["saptable_filter_lang"]),
				),
			},
		},
	})
}

func testAccSapConnectionDataSourceConfig() string {
	jsonPath := "/Users/shaleen.shukla/terraform-provider-saviynt/internal/provider/sap_connection_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_sap_connection_resource" "sap" {
  connection_type    				 = local.cfg.connection_type
  connection_name    				 = local.cfg.connection_name
  message_server                     = local.cfg.message_server
  jco_snc_mode                       = local.cfg.jco_snc_mode
  system_name                        = local.cfg.system_name
  user_import_json                   = jsonencode(local.cfg.user_import_json)
  password_max_length                = local.cfg.password_max_length
  update_account_json                = jsonencode(local.cfg.update_account_json)
  set_cua_system                     = local.cfg.set_cua_system
  saptable_filter_lang               = local.cfg.saptable_filter_lang
}
  
data "saviynt_sap_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_sap_connection_resource.sap]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
