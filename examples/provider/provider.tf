terraform {
  required_providers {
    saviynt = {
      source  = "saviynt/saviynt"
      version = "x.x.x"
    }
  }
}

provider "saviynt" {
  server_url = "https://example.saviyntcloud.com"
  username   = "username"
  password   = "password"
}