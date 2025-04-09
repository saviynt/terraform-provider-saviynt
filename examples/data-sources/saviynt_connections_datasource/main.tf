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

data "saviynt_ad_connection_datasource" "ad"{
    # connection_key=""
    connection_name="dellwireles349012"
}

data "saviynt_adsi_connection_datasource" "adsi"{
    connection_key="4061"
}

data "saviynt_rest_connection_datasource" "rest"{
    connection_key="4060"
}