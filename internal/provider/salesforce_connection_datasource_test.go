package provider

import (
	"fmt"
	"os"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSaviyntSalesforceConnectionDataSource(t *testing.T) {
	filePath := "salesforce_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_salesforce_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSalesforceConnectionDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.client_id", createCfg["client_id"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.redirect_uri", createCfg["redirect_uri"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.instance_url", createCfg["instance_url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.object_to_be_imported", createCfg["object_to_be_imported"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.createaccountjson", createCfg["createaccountjson"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.status_threshold_config", createCfg["status_threshold_config"]),
				),
			},
		},
	})
}

func testAccSalesforceConnectionDataSourceConfig() string {
	jsonPath := "/Users/shaleen.shukla/terraform-provider-saviynt/internal/provider/salesforce_connection_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_salesforce_connection_resource" "salesforce" {
  connection_type    = local.cfg.connection_type
  connection_name    = local.cfg.connection_name
  client_id             = local.cfg.client_id
  redirect_uri          = local.cfg.redirect_uri
  instance_url          = local.cfg.instance_url
  object_to_be_imported = local.cfg.object_to_be_imported
  createaccountjson = jsonencode(local.cfg.createaccountjson)
  status_threshold_config=jsonencode(local.cfg.status_threshold_config)
}
  
data "saviynt_salesforce_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_salesforce_connection_resource.salesforce]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
