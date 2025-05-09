package provider

import (
	"fmt"
	"os"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSaviyntUnixConnectionDataSource(t *testing.T) {
	filePath := "unix_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_unix_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUnixConnectionDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.host_name", createCfg["host_name"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.port_number", createCfg["port_number"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.username", createCfg["username"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.groups_file", createCfg["groups_file"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.status_threshold_config", createCfg["status_threshold_config"]),
				),
			},
		},
	})
}

func testAccUnixConnectionDataSourceConfig() string {
	jsonPath := "${filepath}/unix_connection_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_unix_connection_resource" "unix" {
  connection_type    				 = local.cfg.connection_type
  connection_name    				 = local.cfg.connection_name
  host_name       = local.cfg.host_name
  port_number     = local.cfg.port_number
  username        = local.cfg.username
  password        = local.cfg.password
  groups_file   = local.cfg.groups_file
  status_threshold_config = jsonencode(local.cfg.status_threshold_config)
  ssh_key = local.cfg.ssh_key
}
  
data "saviynt_unix_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_unix_connection_resource.unix]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
