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

resource "saviynt_ad_connection_resource" "example" {
  connection_type     = "AD"
  connection_name     = "Maersk_AD_Connector_3"
  url                 = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  password            = var.PASSWORD
  username            = var.BIND_USER
  vault_connection    = var.VAULT_CONNECTION
  vault_configuration = var.VAULT_CONFIG
  save_in_vault       = var.SAVE_IN_VAULT
}




