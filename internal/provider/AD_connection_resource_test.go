package provider

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

type ADConnectorTestData struct {
	ConnectionName string
	ConnectionType string
	URL            string
	Password       string
	Username       string
	// SearchFilter         string
	// Base                 string
	// GroupSearchBaseDN    string
	// LdapOrAd             string
	// ObjectFilter         string
	// AccountAttribute     string
	// EntitlementAttribute string
	// PageSize             string
	// UserAttribute        string
	// EndpointsFilter      string
	// CreateAccountJson    string
	// UpdateAccountJson    string
	// UpdateUserJson       string
	EnableAccountJson string
	// AccountNameRule      string
	// RemoveAccountAction  string
	// SetRandomPassword    string
}

func loadADConnectorTestData(csvPath string) ([]ADConnectorTestData, error) {
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

	var data []ADConnectorTestData
	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}
		data = append(data, ADConnectorTestData{
			ConnectionType: row[0],
			ConnectionName: row[1],
			URL:            row[2],
			Password:       row[3],
			Username:       row[4],
			// SearchFilter:         row[8],
			// Base:                 row[9],
			// GroupSearchBaseDN:    row[10],
			// LdapOrAd:             row[11],
			// ObjectFilter:         row[12],
			// AccountAttribute:     row[13],
			// EntitlementAttribute: row[14],
			// PageSize:             row[15],
			// UserAttribute:        row[16],
			// EndpointsFilter:      row[17],
			// CreateAccountJson:    row[18],
			// UpdateAccountJson:    row[19],
			// UpdateUserJson:       row[20],
			EnableAccountJson: row[5],
			// AccountNameRule:      row[22],
			// RemoveAccountAction:  row[23],
			// SetRandomPassword:    row[24],
		})
	}

	return data, nil
}

func TestAccSaviyntADConnectorResource(t *testing.T) {
	testData, err := loadADConnectorTestData("AD_connection_resource_test_data.csv")
	if err != nil {
		t.Fatalf("failed to load test data: %s", err)
	}

	for _, row := range testData {
		resourceName := "saviynt_ad_connection_resource." + row.ConnectionName

		t.Run(row.ConnectionName, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: generateADConnectorConfigFromRow(row),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(row.ConnectionName)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(row.ConnectionType)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(row.URL)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact(row.Password)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact(row.Username)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("searchfilter"), knownvalue.StringExact(row.SearchFilter)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("base"), knownvalue.StringExact(row.Base)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("group_search_base_dn"), knownvalue.StringExact(row.GroupSearchBaseDN)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ldap_or_ad"), knownvalue.StringExact(row.LdapOrAd)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("objectfilter"), knownvalue.StringExact(row.ObjectFilter)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_attribute"), knownvalue.StringExact(row.AccountAttribute)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entitlement_attribute"), knownvalue.StringExact(row.EntitlementAttribute)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("page_size"), knownvalue.StringExact(row.PageSize)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_attribute"), knownvalue.StringExact(row.UserAttribute)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoints_filter"), knownvalue.StringExact(row.EndpointsFilter)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_account_json"), knownvalue.StringExact(row.CreateAccountJson)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(row.UpdateAccountJson)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_user_json"), knownvalue.StringExact(row.UpdateUserJson)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_account_json"), knownvalue.StringExact(row.EnableAccountJson)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_name_rule"), knownvalue.StringExact(row.AccountNameRule)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_account_action"), knownvalue.StringExact(row.RemoveAccountAction)),
							// statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("set_random_password"), knownvalue.StringExact(row.SetRandomPassword)),
						},
					},
				},
			})
		})
	}
}

func generateADConnectorConfigFromRow(row ADConnectorTestData) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

resource "saviynt_ad_connection_resource" "%s" {
  connection_type     = "%s"
  connection_name     = "%s"
  url				  ="%s"
  password            = "%s"
  username            = "%s"
  enable_account_json=jsonencode(%s)
}
`,
		os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		row.ConnectionName, row.ConnectionType, row.ConnectionName, row.URL,
		row.Password, row.Username, row.EnableAccountJson,
	)
}
