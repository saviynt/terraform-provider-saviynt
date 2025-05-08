package provider

import (
	"fmt"
	"os"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSaviyntWorkdayConnectionDataSource(t *testing.T) {
	filePath := "workday_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_workday_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccWorkdayConnectionDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.base_url", createCfg["base_url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.api_version", createCfg["api_version"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.use_oauth", createCfg["use_oauth"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.username", createCfg["username"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.client_id", createCfg["client_id"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.access_import_list", createCfg["access_import_list"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.status_key_json", createCfg["status_key_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.user_import_payload", createCfg["user_import_payload"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.user_import_mapping", createCfg["user_import_mapping"]),
				),
			},
		},
	})
}

func testAccWorkdayConnectionDataSourceConfig() string {
	jsonPath := "${filepath}/workday_connection_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_workday_connection_resource" "workday" {
  connection_type    = local.cfg.connection_type
  connection_name    = local.cfg.connection_name
  base_url           = local.cfg.base_url
  api_version        = local.cfg.api_version
  use_oauth          = local.cfg.use_oauth
  username           = local.cfg.username
  client_id          = local.cfg.client_id
  access_import_list = local.cfg.access_import_list
  status_key_json = jsonencode(local.cfg.status_key_json)
  user_import_payload = local.cfg.user_import_payload
  user_import_mapping = jsonencode(local.cfg.user_import_mapping)
}
  
data "saviynt_workday_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_workday_connection_resource.workday]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
