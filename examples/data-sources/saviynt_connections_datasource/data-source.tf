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

# data "saviynt_connections_datasource" "oneResource" {}
# data "saviynt_endpoints_datasource" "data"{}
data "saviynt_security_systems_datasource" "ss"{}

# Output of data
output "endpoints_data" {
  value = {
    msg           = data.saviynt_security_systems_datasource.ss.msg
    display_count = data.saviynt_security_systems_datasource.ss.display_count
    error_code    = data.saviynt_security_systems_datasource.ss.error_code
    total_count   = data.saviynt_security_systems_datasource.ss.total_count
    results       = data.saviynt_security_systems_datasource.ss.results
  }
}

# resource "saviynt_ad_connection_resource" "example" {
#   connection_type     = "AD"
#   connection_name     = "Maersk_AD_Connector_3"
#   url                 = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
#   password            = var.PASSWORD
#   username            = var.BIND_USER
# }
