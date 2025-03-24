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
