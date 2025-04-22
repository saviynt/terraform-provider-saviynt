// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	testSystemName  = "security_system_test_8"
	testDisplayName = "security_system_test_8"
)

func TestAccSaviyntSecuritySystemResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Step
			{
				Config: testAccSecuritySystemConfig(),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("saviynt_security_system_resource.example", tfjsonpath.New("systemname"), knownvalue.StringExact(testSystemName)),
					statecheck.ExpectKnownValue("saviynt_security_system_resource.example", tfjsonpath.New("display_name"), knownvalue.StringExact(testDisplayName)),
					statecheck.ExpectKnownValue("saviynt_security_system_resource.example", tfjsonpath.New("port"), knownvalue.StringExact("443")),
				},				
			},

			// Update Step: Modify the `port` to a new value (e.g., 8080)
			{
				Config: testAccSecuritySystemConfigWithPort("8080"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("saviynt_security_system_resource.example", tfjsonpath.New("systemname"), knownvalue.StringExact(testSystemName)),
					statecheck.ExpectKnownValue("saviynt_security_system_resource.example", tfjsonpath.New("display_name"), knownvalue.StringExact(testDisplayName)),
					statecheck.ExpectKnownValue("saviynt_security_system_resource.example", tfjsonpath.New("port"), knownvalue.StringExact("8080")),
				},
			},
		},
	})
}

func testAccSecuritySystemConfig() string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

resource "saviynt_security_system_resource" "example" {
  systemname   = "%s"
  display_name = "%s"
  port         = "443"
}
`, os.Getenv("SAVIYNT_URL"), os.Getenv("SAVIYNT_USERNAME"), os.Getenv("SAVIYNT_PASSWORD"), testSystemName, testDisplayName)
}

func testAccSecuritySystemConfigWithPort(port string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

resource "saviynt_security_system_resource" "example" {
  systemname   = "%s"
  display_name = "%s"
  port         = "%s"
}
`, os.Getenv("SAVIYNT_URL"), os.Getenv("SAVIYNT_USERNAME"), os.Getenv("SAVIYNT_PASSWORD"), testSystemName, testDisplayName, port)
}
