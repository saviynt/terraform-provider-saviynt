terraform {
	required_providers {
	  saviynt = {
		source  = "registry.terraform.io/local/saviynt"
		version = "1.0.0"
	  }
	}
  }
  provider "saviynt" {
	server_url = var.SAVIYNT_SERVER_URL
	username   = var.SAVIYNT_USERNAME
	password   = var.SAVIYNT_PASSWORD
  }
  resource "saviynt_adsi_connection_resource" "example" {
	connection_type = "ADSI"
  	connection_name = "Shaleen_testing_ADSI_terraform_3490"
	
	status_threshold_config=<<EOF
	{
  "statusAndThresholdConfig": {
    "statusColumn": "customproperty24",
    "activeStatus": [
      "512",
      "544",
      "66048"
    ],
    "deleteLinks": false,
    "accountThresholdValue": 50000,
    "correlateInactiveAccounts": true,
    "inactivateAccountsNotInFile": false,
    "deleteAccEntForActiveAccounts": false
	}
	}
	EOF
	
	updategroupjson=jsonencode({
	"objects": [
		{
		"objectClasses": [
			"group"
		],
		"distinguishedName": "$${role.role_name}",
		"attributes": {
			"description": "$${role.description}",
			"proxyAddresses": "$${role.customproperty20}"
		}
		}
	]
	})
}