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
  endpointname = "sample-103"
  display_name = "sample 103"
  security_system  = "shaleenhuddle"
  mapped_endpoints =[
    {
      security_system = "Shaleen_testing_terraform"
      endpoint        = "Shaleen_testing_terraform"
      requestable     = "true"
      operation       = "ADD"
    }
  ]
}
