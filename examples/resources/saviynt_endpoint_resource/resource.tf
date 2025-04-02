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
  endpointname = "sample-102"
  display_name = "sample 102"
  security_system  = "shaleenhuddle"
  # email_template_type="3"
  task_type="1"
  email_template="Add Password Expiry Email"
}
