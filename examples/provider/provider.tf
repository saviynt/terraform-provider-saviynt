
terraform {
  required_providers {
    saviynt = {
      source  = "saviynt/saviynt"
      version = "0.2.1"
    }
  }
}

provider "saviynt" {
  server_url = "www.example.com"
  username   = "username"
  password   = "password"
}
