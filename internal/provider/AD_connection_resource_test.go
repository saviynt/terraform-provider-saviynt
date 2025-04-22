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

const (
	testADConnectionName = "ad_connection_test_2"
)

func TestAccSaviyntADConnectionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Step
			{
				Config: testAccADConnectionConfig(),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("saviynt_ad_connection_resource.example", tfjsonpath.New("connection_name"), knownvalue.StringExact(testADConnectionName)),
					statecheck.ExpectKnownValue("saviynt_ad_connection_resource.example", tfjsonpath.New("connection_type"), knownvalue.StringExact("AD")),
					statecheck.ExpectKnownValue("saviynt_ad_connection_resource.example", tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},

			// Update Step
			{
				Config: testAccADConnectionConfigWithURL("newpass"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("saviynt_ad_connection_resource.example", tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
		},
	})
}

func testAccADConnectionConfig() string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

resource "saviynt_ad_connection_resource" "example" {
  connection_type = "AD"
  connection_name = "%s"
  password        = "%s"
}
`,
		os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		testADConnectionName,
		// os.Getenv("LDAP_PROTOCOL"),
		// os.Getenv("IP_ADDRESS"),
		// getEnvInt("LDAP_PORT", 389),
		// os.Getenv("BIND_USER"),
		os.Getenv("PASSWORD"),
	)
}

func testAccADConnectionConfigWithURL(pass string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

resource "saviynt_ad_connection_resource" "example" {
  connection_type = "AD"
  connection_name = "%s"
  password        = "%s"
}
`,
		os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		testADConnectionName,
		pass,
	)
}

