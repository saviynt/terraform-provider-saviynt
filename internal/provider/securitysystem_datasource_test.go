package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccSaviyntSecuritySystemDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSaviyntSecuritySystemDataSourceConfig(),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.saviynt_security_systems_datasource.test",
						tfjsonpath.New("msg"),
						knownvalue.StringExact("Success"),
					),
					statecheck.ExpectKnownValue(
						"data.saviynt_security_systems_datasource.test",
						tfjsonpath.New("error_code"),
						knownvalue.StringExact("0"),
					),
				},
			},
		},
	})
}

func testAccSaviyntSecuritySystemDataSourceConfig() string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

data "saviynt_security_systems_datasource" "test" {
}
`, os.Getenv("SAVIYNT_URL"), os.Getenv("SAVIYNT_USERNAME"), os.Getenv("SAVIYNT_PASSWORD"))
}
