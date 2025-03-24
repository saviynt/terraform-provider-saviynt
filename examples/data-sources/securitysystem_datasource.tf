terraform {
  required_providers {
    saviynt = {
      source  = "registry.terraform.io/local/saviynt"
      version = "1.0.0"
    }
  }
}


provider "saviynt" {
  server_url = var.saviynt_server_url
  username   = var.saviynt_username
  password   = var.saviynt_password
}


data "saviynt_security_systems_datasource" "example" {
  systemname="withoutaccessgetaccess5"
}
output "security_systems_data" {
  value = {
    msg          = data.saviynt_security_systems_datasource.example.msg
    display_count = data.saviynt_security_systems_datasource.example.display_count
    error_code   = data.saviynt_security_systems_datasource.example.error_code
    total_count  = data.saviynt_security_systems_datasource.example.total_count
    results      = data.saviynt_security_systems_datasource.example.results
  }
}
