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
    connection_key="123"
}

data "saviynt_adsi_connection_datasource" "adsi"{
    connection_key="123"
}

data "saviynt_rest_connection_datasource" "rest"{
    connection_key="123"
}

data "saviynt_db_connection_datasource" "db" {
  connection_key="123"
}

data "saviynt_workday_connection_datasource" "w"{
    connection_key="123"
}

data "saviynt_salesforce_connection_datasource" "s"{
    connection_key="123"
}

data "saviynt_entraid_connection_datasource" "en"{
    connection_key="123"
}

data "saviynt_sap_connection_datasource" "sap" {
  connection_key="123"
}