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

resource "saviynt_db_connection_resource" "example" {
  connection_type     = "DB"
  connection_name     = "shaleen_db_12_april"
  username            = "connadmin"
  url="jdbc:mysql://34.139.69.20:3306"
  password            = "MyOffice12#"
  driver_name         = "com.mysql.jdbc.Driver"
}  