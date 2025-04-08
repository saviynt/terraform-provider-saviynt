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

resource "saviynt_rest_connection_resource" "example" {
  connection_type = "REST"
  connection_name = "namefortheconnection"
  config_json     = jsonencode({"showLogs":true})
}