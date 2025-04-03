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

resource "saviynt_endpoint_resource" "example" {
  endpointname    = "sample"
  display_name    = "sample"
  security_system = "samplesystem"
}

# data "saviynt_rest_connection_datasource" "ss"{
#   # connection_key="123"
# }

# output "endpoints_data" {
#   value = {
#     msg           = data.saviynt_rest_connection_datasource.ss.msg
#     error_code    = data.saviynt_rest_connection_datasource.ss.error_code
#   }
# }
