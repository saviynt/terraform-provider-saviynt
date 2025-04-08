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
resource "saviynt_security_system_resource" "example" {
  systemname   = "read_16"
  display_name = "read_16"
  instant_provision="true"
  recon_application="true"
  use_open_connector="true"
}
