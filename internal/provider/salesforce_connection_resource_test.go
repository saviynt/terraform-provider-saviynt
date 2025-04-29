package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var testSalesforceConnectionName = util.GenerateRandomName("salesforce")

func TestAccSaviyntSalesforceConnectionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Step
			{
				Config: testAccSalesforceConnectionResourceConfig("SalesForce", testSalesforceConnectionName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("saviynt_salesforce_connection_resource.ss", tfjsonpath.New("connection_name"), knownvalue.StringExact(testSalesforceConnectionName)),
					statecheck.ExpectKnownValue("saviynt_salesforce_connection_resource.ss", tfjsonpath.New("connection_type"), knownvalue.StringExact("SalesForce")),
					statecheck.ExpectKnownValue("saviynt_salesforce_connection_resource.ss", tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Import
			{
				ResourceName:      "saviynt_salesforce_connection_resource.ss",
        ImportStateId: testSalesforceConnectionName,
				ImportState:       true,
				ImportStateVerify: true,
        ImportStateVerifyIgnore: []string{"msg", "client_secret", "refresh_token"},
			},
			// Update Step
			{
				Config: testAccSalesforceConnectionResourceObjImpConfig(testSalesforceConnectionName, "Profile,Group,PermissionSet"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("saviynt_salesforce_connection_resource.ss", tfjsonpath.New("object_to_be_imported"), knownvalue.StringExact("Profile,Group,PermissionSet")),
					statecheck.ExpectKnownValue("saviynt_salesforce_connection_resource.ss", tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},

			{
				Config:      testAccSalesforceConnectionResourceConfig("SalesForce", "new_"+testSalesforceConnectionName),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},

      {
				Config:      testAccSalesforceConnectionResourceConfig("AD", testSalesforceConnectionName),
				ExpectError: regexp.MustCompile(`Connection type cannot be updated`),
			},
		},
	})
}

func testAccSalesforceConnectionResourceConfig(connType string, connName string) string {
	return fmt.Sprintf(`
	provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

  resource "saviynt_salesforce_connection_resource" "ss" {
  connection_type       = "%s"
  connection_name       = "%s"
  client_id             = "<client_id>"
  client_secret         = "<client_secret>"
  refresh_token         = "<refresh_token>"
  redirect_uri          = "https://<domain>/services/oauth2/success"
  instance_url          = "https://<domain>/services/oauth2/success"
  object_to_be_imported = "Profile,Group,PermissionSet,Role"
  account_field_query   = "Id,Username,LastName,FirstName,Name,CompanyName,Email,IsActive,UserRoleId,ProfileId,UserType,ManagerId,LastLoginDate,LastPasswordChangeDate,CreatedDate,CreatedById,LastModifiedDate,LastModifiedById,SystemModstamp,ContactId,AccountId,FederationIdentifier,UserPermissionsSupportUser"
  field_mapping_json = jsonencode({
    accountfield_mapping = {
        accountID                = "Id~#~char"
        name                     = "Username~#~char"
        customproperty2          = "LastName~#~char"
        customproperty1          = "FirstName~#~char"
        displayName              = "Name~#~char"
        customproperty3          = "CompanyName~#~char"
        customproperty4          = "Email~#~char"
        status                   = "IsActive~#~bool"
        customproperty12         = "IsActive~#~bool"
        customproperty5          = "UserRoleId~#~char"
        customproperty6          = "ProfileId~#~char"
        accounttype              = "UserType~#~char"
        customproperty7          = "ManagerId~#~char"
        lastlogondate            = "LastLoginDate~#~date"
        lastpasswordchange       = "LastPasswordChangeDate~#~date"
        CREATED_ON               = "CreatedDate~#~date"
        creator                  = "CreatedById~#~char"
        customproperty8          = "LastModifiedDate~#~date"
        updateUser               = "LastModifiedById~#~char"
        updatedate               = "SystemModstamp~#~date"
        customproperty9          = "ContactId~#~char"
        customproperty10         = "AccountId~#~char"
        customproperty13         = "FederationIdentifier~#~char"
        customproperty21         = "UserPermissionsSupportUser~#~bool"
        customproperty20         = "CreatedDate~#~char"
    }
    })

  createaccountjson = jsonencode({
    Alias                 = "$${user?.getFirstname()}"
    Email                 = "$${user?.getEmail()}"
    Username              = "$${user?.getEmail()}"
    CommunityNickname     = "$${user?.getFirstname()}"
    FirstName             = "$${user?.getFirstname()}"
    LastName              = "$${user?.getLastname()}"
    TimeZoneSidKey        = "America/Los_Angeles"
    LocaleSidKey          = "en_US"
    EmailEncodingKey      = "ISO-8859-1"
    ProfileId             = "$${profileId}"
    LanguageLocaleKey     = "en_US"
    IsActive              = true
    FederationIdentifier  = "$${user?.getEmail()}"
    })

  modifyaccountjson = jsonencode({
    Username             = "$${user?.customproperty16 + \".company\"}"
    FirstName            = "$${user?.getFirstname()}"
    LastName             = "$${user?.getLastname()}"
    FederationIdentifier = "$${user?.customproperty16}"
    })
  status_threshold_config = jsonencode({
    statusAndThresholdConfig = {
        accountThresholdValue = 100
        statusColumn          = "customproperty12"
        activeStatus          = [
        "true"
        ]
        deleteLinks           = true
        lockedStatusColumn    = "customproperty28"
        lockedStatusMapping   = {
        Locked   = ["1"]
        Unlocked = ["0"]
        }
    }
})

  customconfigjson = jsonencode({
        disableAccountForRevokeTask = false
        defaultEntitlementId       = "<entitlement_id>"
    })

}`,
 os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"), connType, connName,
	)
}


func testAccSalesforceConnectionResourceObjImpConfig(connName string, objImp string) string {
	return fmt.Sprintf(`
	provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

  resource "saviynt_salesforce_connection_resource" "ss" {
  connection_type       = "SalesForce"
  connection_name       = "%s"
  client_id             = "<client_id>"
  client_secret         = "<client_secret>"
  refresh_token         = "<refresh_token>"
  redirect_uri          = "https://<domain>/services/oauth2/success"
  instance_url          = "https://<domain>/services/oauth2/success"
  object_to_be_imported = "%s"
  account_field_query   = "Id,Username,LastName,FirstName,Name,CompanyName,Email,IsActive,UserRoleId,ProfileId,UserType,ManagerId,LastLoginDate,LastPasswordChangeDate,CreatedDate,CreatedById,LastModifiedDate,LastModifiedById,SystemModstamp,ContactId,AccountId,FederationIdentifier,UserPermissionsSupportUser"
  field_mapping_json = jsonencode({
    accountfield_mapping = {
        accountID                = "Id~#~char"
        name                     = "Username~#~char"
        customproperty2          = "LastName~#~char"
        customproperty1          = "FirstName~#~char"
        displayName              = "Name~#~char"
        customproperty3          = "CompanyName~#~char"
        customproperty4          = "Email~#~char"
        status                   = "IsActive~#~bool"
        customproperty12         = "IsActive~#~bool"
        customproperty5          = "UserRoleId~#~char"
        customproperty6          = "ProfileId~#~char"
        accounttype              = "UserType~#~char"
        customproperty7          = "ManagerId~#~char"
        lastlogondate            = "LastLoginDate~#~date"
        lastpasswordchange       = "LastPasswordChangeDate~#~date"
        CREATED_ON               = "CreatedDate~#~date"
        creator                  = "CreatedById~#~char"
        customproperty8          = "LastModifiedDate~#~date"
        updateUser               = "LastModifiedById~#~char"
        updatedate               = "SystemModstamp~#~date"
        customproperty9          = "ContactId~#~char"
        customproperty10         = "AccountId~#~char"
        customproperty13         = "FederationIdentifier~#~char"
        customproperty21         = "UserPermissionsSupportUser~#~bool"
        customproperty20         = "CreatedDate~#~char"
    }
    })

  createaccountjson = jsonencode({
    Alias                 = "$${user?.getFirstname()}"
    Email                 = "$${user?.getEmail()}"
    Username              = "$${user?.getEmail()}"
    CommunityNickname     = "$${user?.getFirstname()}"
    FirstName             = "$${user?.getFirstname()}"
    LastName              = "$${user?.getLastname()}"
    TimeZoneSidKey        = "America/Los_Angeles"
    LocaleSidKey          = "en_US"
    EmailEncodingKey      = "ISO-8859-1"
    ProfileId             = "$${profileId}"
    LanguageLocaleKey     = "en_US"
    IsActive              = true
    FederationIdentifier  = "$${user?.getEmail()}"
    })

  modifyaccountjson = jsonencode({
    Username             = "$${user?.customproperty16 + \".company\"}"
    FirstName            = "$${user?.getFirstname()}"
    LastName             = "$${user?.getLastname()}"
    FederationIdentifier = "$${user?.customproperty16}"
    })
  status_threshold_config = jsonencode({
    statusAndThresholdConfig = {
        accountThresholdValue = 100
        statusColumn          = "customproperty12"
        activeStatus          = [
        "true"
        ]
        deleteLinks           = true
        lockedStatusColumn    = "customproperty28"
        lockedStatusMapping   = {
        Locked   = ["1"]
        Unlocked = ["0"]
        }
    }
})

  customconfigjson = jsonencode({
        disableAccountForRevokeTask = false
        defaultEntitlementId       = "<entitlement_id>"
    })

}`,
 os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"), connName, objImp,
	)
}

