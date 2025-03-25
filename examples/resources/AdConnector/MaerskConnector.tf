terraform {
  required_providers {
    saviynt = {
      source  = "registry.terraform.io/local/saviynt"
      version = "1.0.0"
    }
  }
}
provider "saviynt" {
  server_url      = var.SAVIYNT_SERVER_URL
  username = var.SAVIYNT_USERNAME
  password = var.SAVIYNT_PASSWORD
}

resource "saviynt_test_connection" "example3" {
  connection_type               = "2"
  connection_name               = "Maersk_AD_Connector_1"
  url                           = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  password                      = var.PASSWORD
  username                      = var.BIND_USER
  vault_connection              = var.VAULT_CONNECTION
  vault_configuration           = var.VAULT_CONFIG
  save_in_vault                 = var.SAVE_IN_VAULT
  ldap_or_ad="AD"
  status_threshold_config       = "{\"statusAndThresholdConfig\":{\"statusColumn\":\"customproperty30\",\"activeStatus\":[\"512\",\"544\",\"66048\"],\"inactiveStatus\":[\"546\",\"514\",\"66050\"],\"deleteLinks\":false,\"accountThresholdValue\":1000,\"correlateInactiveAccounts\":true,\"inactivateAccountsNotInFile\":false,\"lockedStatusColumn\":\"\",\"lockedStatusMapping\":{\"Locked\":[\"\"],\"Unlocked\":[\"\"]}}}"
  entitlement_attribute         = "memberOf"
  group_search_base_dn          = var.BASE_CONTAINER
  page_size                     = "1000"
  base                          = var.BASE_CONTAINER
}










