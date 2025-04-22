resource "saviynt_ad_connection_resource" "example" {
  connection_type               = "AD"
  connection_name               = "namefortheconnection"
  url                           = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  password                      = var.PASSWORD
  username                      = var.BIND_USER
  vault_connection              = var.VAULT_CONNECTION
  vault_configuration           = var.VAULT_CONFIG
  save_in_vault                 = var.SAVE_IN_VAULT
  object_filter                 = "example"
  entitlement_attribute         = "example"
  group_search_base_dn          = var.BASE_CONTAINER
  page_size                     = "1000"
  base                          = var.BASE_CONTAINER
  account_name_rule             = format("CN=task.accountName,%s###CN=task.accountName1,%s###CN=task.accountName2,%s", var.BASE_CONTAINER, var.BASE_CONTAINER, var.BASE_CONTAINER)
  set_random_password           = "false"
  reuse_inactive_account        = "false"
  check_for_unique              = "example"
  reset_and_change_passwrd_json = "example"
  password_min_length           = "2"
  password_max_length           = "2"
  password_no_of_caps_alpha     = "2"
  password_no_of_digits         = "2"
  password_no_of_spl_chars      = "2"
  unlock_account_json           = "example"
  pam_config                    = "example"
}
