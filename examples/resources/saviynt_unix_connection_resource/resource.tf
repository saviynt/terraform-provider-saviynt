
terraform {
  required_providers {
    saviynt = {
      source  = "registry.terraform.io/local/saviynt"
      version = "1.0.0"
    }
  }
}

provider "saviynt" {
  server_url      = var.SAVIYNT_SERVER_URL
  username = var.SAVIYNT_USERNAME
  password = var.SAVIYNT_PASSWORD
}

resource "saviynt_unix_connection_resource" "example" {
    connection_type               = "Unix"
    connection_name               = "shaleen_unix_4"
    host_name="18.211.227.33"
    port_number="22"
    username="sav-e"
    shadow_file="{\"UNIX\":\"sudo cat /etc/shadow\"}"
    accounts_file="{\"UNIX\":\"sudo cat /etc/passwd\"}"
    provision_account_command="sudo useradd $${username} -p $${password} -c \"$${user?.country}/$${user?.employeeType}/$${user?.employeeid}/$${user?.lastname}.$${user?.firstname}/$${user?.email}\" -g users"
}