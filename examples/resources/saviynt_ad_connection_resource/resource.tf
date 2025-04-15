resource "saviynt_ad_connection_resource" "example" {
  connection_type     = "AD"
  connection_name     = "namefortheconnection"
  url                 = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  password            = var.PASSWORD
  username            = var.BIND_USER
  vault_connection    = var.VAULT_CONNECTION
  vault_configuration = var.VAULT_CONFIG
  save_in_vault       = var.SAVE_IN_VAULT
}