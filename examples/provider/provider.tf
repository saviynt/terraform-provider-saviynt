terraform {
  required_providers {
    saviynt = {
      source  = "registry.terraform.io/local/saviynt"
      version = "1.0.0"
    }
  }
}

provider "saviynt" {
  server_url = "https://example.saviyntcloud.com"
  username   = "username"
  password   = "password"
}
