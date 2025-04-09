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
    connection_key="4064"
}

data "saviynt_adsi_connection_datasource" "adsi"{
    connection_key="4061"
}

data "saviynt_rest_connection_datasource" "rest"{
    connection_key="4060"
}

data "saviynt_db_connection_datasource" "db" {
  connection_key="4057"
}

data "saviynt_workday_connection_datasource" "w"{
    connection_key="276"
}

data "saviynt_salesforce_connection_datasource" "s"{
    connection_key="3563"
}

data "saviynt_entraid_connection_datasource" "en"{
    connection_key="2881"
}