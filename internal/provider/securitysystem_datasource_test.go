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

func TestAccSaviyntSecuritySystemDataSource(t *testing.T) {
	filePath := "securitysystem_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_security_systems_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSaviyntSecuritySystemDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "msg", "Success"),
					resource.TestCheckResourceAttr(datasource, "error_code", "0"),
					resource.TestCheckResourceAttr(datasource, "results.0.systemname1", createCfg["systemname"]),
					resource.TestCheckResourceAttr(datasource, "results.0.display_name", createCfg["display_name"]),
					resource.TestCheckResourceAttr(datasource, "results.0.access_add_workflow", createCfg["access_add_workflow"]),
					resource.TestCheckResourceAttr(datasource, "results.0.access_remove_workflow", createCfg["access_remove_workflow"]),
					resource.TestCheckResourceAttr(datasource, "results.0.add_service_account_workflow", createCfg["add_service_account_workflow"]),
					resource.TestCheckResourceAttr(datasource, "results.0.remove_service_account_workflow", createCfg["remove_service_account_workflow"]),
					resource.TestCheckResourceAttr(datasource, "results.0.automated_provisioning", createCfg["automated_provisioning"]),
				),
			},
		},
	})
}

func testAccSaviyntSecuritySystemDataSourceConfig() string {
	jsonPath := "${filepath}/securitysystem_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_security_system_resource" "ss" {
  systemname                         = local.cfg.systemname
  display_name                       = local.cfg.display_name
  access_add_workflow                = local.cfg.access_add_workflow
  access_remove_workflow             = local.cfg.access_remove_workflow
  add_service_account_workflow       = local.cfg.add_service_account_workflow
  remove_service_account_workflow    = local.cfg.remove_service_account_workflow
  automated_provisioning             = local.cfg.automated_provisioning
}
  
data "saviynt_security_systems_datasource" "test" {
	systemname = local.cfg.systemname
	depends_on = [saviynt_security_system_resource.ss]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
