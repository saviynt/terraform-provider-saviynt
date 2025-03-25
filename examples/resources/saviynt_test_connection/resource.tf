terraform {
  required_providers {
    saviynt = {
      source  = "registry.terraform.io/local/saviynt"
      version = "1.0.0"
    }
  }
}
provider "saviynt" {
  server_url = "SAVIYNT_SERVER_URL"
  username   = "SAVIYNT_USERNAME"
  password   = "SAVIYNT_PASSWORD"
}

resource "saviynt_test_connection" "example3" {
  connection_type         = "2"
  connection_name         = "YOUR_CONNECTION_NAME"
  url                     = format("%s://%s:%d", "LSAP_PROTOCOL", "IP_ADDRESS", "LDAP_PORT")
  password                = "PASSWORD"
  username                = "BIND_USER"
  vault_connection        = "VAULT_CONNECTION"
  vault_configuration     = "VAULT_CONFIG"
  save_in_vault           = "SAVE_IN_VAULT"
  ldap_or_ad              = "AD"
  status_threshold_config = "{\"statusAndThresholdConfig\":{\"statusColumn\":\"customproperty30\",\"activeStatus\":[\"512\",\"544\",\"66048\"],\"inactiveStatus\":[\"546\",\"514\",\"66050\"],\"deleteLinks\":false,\"accountThresholdValue\":1000,\"correlateInactiveAccounts\":true,\"inactivateAccountsNotInFile\":false,\"lockedStatusColumn\":\"\",\"lockedStatusMapping\":{\"Locked\":[\"\"],\"Unlocked\":[\"\"]}}}"
  entitlement_attribute   = "ENTITLEMENT_ATTRIBUTE"
  group_search_base_dn    = "BASE_CONTAINER"
  page_size               = "PAGE_SIZE"
  base                    = "BASE_CONTAINER"
}










