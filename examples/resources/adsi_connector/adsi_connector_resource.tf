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
resource "saviynt_adsi_connection_resource" "example" {
  connection_type = "ADSI"
  connection_name     = "namefortheconnection"
}