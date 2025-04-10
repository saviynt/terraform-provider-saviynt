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
  connection_name = "shaleen_adsi_1000"
  vault_connection="Hashicorp"
vault_configuration ="{\"path\":\"/secrets/data/kv-dev-intgn1/-AD_Credential\",\"keyMapping\":{\"PASSWORD\":\"AD_Credential_PASSWORD~#~None\"}}"
save_in_vault="false"
      url="LDAP://saviyntdc01.saviyntlabs.org"
    status_threshold_config=jsonencode(
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
    "inactivateAccountsNotInFile": true,
    "deleteAccEntForActiveAccounts": false
  }
}
)
    username="saviyntlabs\\Administrator"
    connection_url="http://saviyntdc01.saviyntlabs.org:8090/api/v1/discovery"
    forestlist="saviyntlabs.org"
    password="MyOffice12#"
    searchfilter="DC=saviyntlabs,DC=org"
    objectfilter="(&(objectCategory=person)(objectClass=user))"
}