package provider

import (
	"fmt"
	"os"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSaviyntGithubRestConnectionDataSource(t *testing.T) {
	filePath := "github_rest_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_github_rest_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRestConnectionDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.import_account_ent_json", createCfg["import_account_ent_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.organization_list", createCfg["organization_list"]),
				),
			},
		},
	})
}

func testAccGithubRestConnectionDataSourceConfig() string {
	jsonPath := "${filepath}/github_rest_connection_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_github_rest_connection_resource" "github_rest" {
  connection_type    				 = local.cfg.connection_type
  connection_name   				 = local.cfg.connection_name
  connection_json                    = jsonencode(local.cfg.connection_json)
  import_account_ent_json			 = jsonencode(local.cfg.import_account_ent_json)
  organization_list                  = local.cfg.organization_list
}
  
data "saviynt_github_rest_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_github_rest_connection_resource.github_rest]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
