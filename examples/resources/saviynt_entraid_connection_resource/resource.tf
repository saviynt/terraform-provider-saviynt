resource "saviynt_entraid_connection_resource" "ss" {
  connection_type = "AzureAD"
  connection_name = "namefortheconnection"
  client_secret   = "XXXX"
  aad_tenant_id   = "XXXX"
  client_id       = "XXXX"
}
