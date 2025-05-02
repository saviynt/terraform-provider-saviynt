// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

type SecuritySystemTestData struct {
	Systemname                   string
	DisplayName                  string
	AccessAddWorkflow            string
	AccessRemoveWorkflow         string
	AddServiceAccountWorkflow    string
	RemoveServiceAccountWorkflow string
	AutomatedProvisioning        string
}

func loadSecuritySystemTestData(csvPath string) ([]SecuritySystemTestData, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	var data []SecuritySystemTestData
	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}
		data = append(data, SecuritySystemTestData{
			Systemname:                   row[0],
			DisplayName:                  row[1],
			AccessAddWorkflow:            row[2],
			AccessRemoveWorkflow:         row[3],
			AddServiceAccountWorkflow:    row[4],
			RemoveServiceAccountWorkflow: row[5],
			AutomatedProvisioning:        row[6],
		})
	}
	return data, nil
}
func TestAccSaviyntSecuritySystemResource(t *testing.T) {
	testData, err := loadSecuritySystemTestData("securitysystem_resource_test_data.csv")
	if err != nil {
		t.Fatalf("failed to load test data: %s", err)
	}

	for _, row := range testData {
		resourceName := "saviynt_security_system_resource." + row.DisplayName

		t.Run(row.Systemname, func(t *testing.T) {
			updatedConfig, updatedRow := generateSecuritySystemWithNewValue(row)
			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					// Create the resource
					{
						Config: generateSecuritySystemConfigFromRow(row),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("systemname"), knownvalue.StringExact(row.Systemname)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(row.DisplayName)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_add_workflow"), knownvalue.StringExact(row.AccessAddWorkflow)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_remove_workflow"), knownvalue.StringExact(row.AccessRemoveWorkflow)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_service_account_workflow"), knownvalue.StringExact(row.AddServiceAccountWorkflow)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_service_account_workflow"), knownvalue.StringExact(row.RemoveServiceAccountWorkflow)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("automated_provisioning"), knownvalue.StringExact(row.AutomatedProvisioning)),
						},
					},
					// Update the resource with new values
					{
						Config: updatedConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("systemname"), knownvalue.StringExact(updatedRow.Systemname)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(updatedRow.DisplayName)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_add_workflow"), knownvalue.StringExact(updatedRow.AccessAddWorkflow)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_remove_workflow"), knownvalue.StringExact(updatedRow.AccessRemoveWorkflow)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_service_account_workflow"), knownvalue.StringExact(updatedRow.AddServiceAccountWorkflow)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_service_account_workflow"), knownvalue.StringExact(updatedRow.RemoveServiceAccountWorkflow)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("automated_provisioning"), knownvalue.StringExact(updatedRow.AutomatedProvisioning)),
						},
					},
					// Import the resource
					{
						ResourceName:      resourceName,
						ImportStateId:     row.Systemname,
						ImportState:       true,
						ImportStateVerify: true,
					},
					// Update the Systemname to a new value
					{
						Config:      generateSecuritySystemRenameConfig(row),
						ExpectError: regexp.MustCompile(`System name cannot be updated`),
					},
					// Create a new resource with the same Systemname
					{
						Config:      generateSecuritySystemWithSameNameConfig(row),
						ExpectError: regexp.MustCompile(`Security System Already Exists`),
					},
				},
			})
		})
	}

}

func generateSecuritySystemWithNewValue(row SecuritySystemTestData) (string, SecuritySystemTestData) {
	row.AccessAddWorkflow = "Manager Approval"
	row.AccessRemoveWorkflow = "Manager Approval"
	row.AddServiceAccountWorkflow = "Manager Approval"
	row.RemoveServiceAccountWorkflow = "Manager Approval"
	return generateSecuritySystemConfigFromRow(row), row
}

func generateSecuritySystemRenameConfig(row SecuritySystemTestData) string {
	row.Systemname = "renamed_" + row.Systemname
	return generateSecuritySystemConfigFromRow(row)
}
func generateSecuritySystemWithSameNameConfig(row SecuritySystemTestData) string {
	row.DisplayName = "new" + row.DisplayName
	return generateSecuritySystemConfigFromRow(row)
}

func generateSecuritySystemConfigFromRow(row SecuritySystemTestData) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

resource "saviynt_security_system_resource" "%s" {
  systemname                         = "%s"
  display_name                       = "%s"
  access_add_workflow                = "%s"
  access_remove_workflow             = "%s"
  add_service_account_workflow       = "%s"
  remove_service_account_workflow    = "%s"
  automated_provisioning             = "%s"
}
`, os.Getenv("SAVIYNT_URL"), os.Getenv("SAVIYNT_USERNAME"), os.Getenv("SAVIYNT_PASSWORD"),
		row.DisplayName, row.Systemname, row.DisplayName,
		row.AccessAddWorkflow, row.AccessRemoveWorkflow,
		row.AddServiceAccountWorkflow, row.RemoveServiceAccountWorkflow,
		row.AutomatedProvisioning,
	)
}
