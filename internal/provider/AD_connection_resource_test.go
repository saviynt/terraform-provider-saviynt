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
	ConnectionName       string
	ConnectionType       string
	URL                  string
	Password             string
	Username             string
	VaultConnection      string
	VaultConfiguration   string
	SaveInVault          string
	SearchFilter         string
	Base                 string
	GroupSearchBaseDN    string
	LdapOrAd             string
	ObjectFilter         string
	AccountAttribute     string
	EntitlementAttribute string
	PageSize             string
	UserAttribute        string
	EndpointsFilter      string
	CreateAccountJson    string
	UpdateAccountJson    string
	UpdateUserJson       string
	EnableAccountJson    string
	AccountNameRule      string
	RemoveAccountAction  string
	SetRandomPassword    string
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
			ConnectionName:       row[0],
			ConnectionType:       row[1],
			URL:                  row[2],
			Password:             row[3],
			Username:             row[4],
			VaultConnection:      row[5],
			VaultConfiguration:   row[6],
			SaveInVault:          row[7],
			SearchFilter:         row[8],
			Base:                 row[9],
			GroupSearchBaseDN:    row[10],
			LdapOrAd:             row[11],
			ObjectFilter:         row[12],
			AccountAttribute:     row[13],
			EntitlementAttribute: row[14],
			PageSize:             row[15],
			UserAttribute:        row[16],
			EndpointsFilter:      row[17],
			CreateAccountJson:    row[18],
			UpdateAccountJson:    row[19],
			UpdateUserJson:       row[20],
			EnableAccountJson:    row[21],
			AccountNameRule:      row[22],
			RemoveAccountAction:  row[23],
			SetRandomPassword:    row[24],
		})
	}

	return data, nil
}

func TestAccSaviyntADConnectorResource(t *testing.T) {
	testData, err := loadADConnectorTestData("ad_connector_test_data_20250430_000441.csv")
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
						},
					},
				},
			})
		})
	}
}

func generateADConnectorConfigFromRow(row ADConnectorTestData) string {
	return fmt.Sprintf(` + "`" + `
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

resource "saviynt_ad_connection_resource" "%s" {
  connection_name = "%s"
  connection_type = "%s"
  url = "%s"
  password = "%s"
  username = "%s"
  vault_connection = "%s"
  vault_configuration = "%s"
  save_in_vault = "%s"
  searchfilter = "%s"
  base = "%s"
  group_search_base_dn = "%s"
  ldap_or_ad = "%s"
  objectfilter = "%s"
  account_attribute = "%s"
  entitlement_attribute = "%s"
  page_size = "%s"
  user_attribute = "%s"
  endpoints_filter = %s
  create_account_json = %s
  update_account_json = %s
  update_user_json = %s
  enable_account_json = %s
  account_name_rule = "%s"
  remove_account_action = %s
  set_random_password = "%s"
  group_import_mapping = %s
}
` + "`" + `,
		os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		row.ConnectionName, row.ConnectionName, row.ConnectionType,
		row.URL, row.Password, row.Username, row.VaultConnection,
		row.VaultConfiguration, row.SaveInVault, row.SearchFilter,
		row.Base, row.GroupSearchBaseDN, row.LdapOrAd, row.ObjectFilter,
		row.AccountAttribute, row.EntitlementAttribute, row.PageSize,
		row.UserAttribute, row.EndpointsFilter, row.CreateAccountJson,
		row.UpdateAccountJson, row.UpdateUserJson, row.EnableAccountJson,
		row.AccountNameRule, row.RemoveAccountAction, row.SetRandomPassword,
		row.GroupImportMapping,
	)
}
