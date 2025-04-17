resource "saviynt_rest_connection_resource" "example" {
  connection_type         = "REST"
  connection_name         = "namefortheconnection"
  vault_connection        = var.VAULT_CONNECTION
  vault_configuration     = var.VAULT_CONFIG
  save_in_vault           = var.SAVE_IN_VAULT
  connection_json         = jsonencode({
      "example_key":"connection_value"
    })
}
