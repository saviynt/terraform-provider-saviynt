# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

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

resource "saviynt_workday_connection_resource" "ss" {
  connection_type    = "Workday"
  connection_name    = "hehehe_workday91"
  use_oauth          = "TRUE"
}

# resource "saviynt_workday_connection_resource" "ss1" {
#   connection_type    = "Workday"
#   connection_name    = "hehehe_workday99"
#   use_oauth          = "TRUE"
# }