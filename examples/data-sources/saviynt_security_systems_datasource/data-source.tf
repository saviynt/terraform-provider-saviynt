# List all the security systems
data "saviynt_security_systems_datasource" "all" {}


# Get a security system by its name
data "saviynt_security_systems_datasource" "oneResource" {
  systemname = "SYSTEM_NAME"
}

# Output of data
output "security_systems_data" {
  value = {
    msg           = data.saviynt_security_systems_datasource.oneResource.msg
    display_count = data.saviynt_security_systems_datasource.oneResource.display_count
    error_code    = data.saviynt_security_systems_datasource.oneResource.error_code
    total_count   = data.saviynt_security_systems_datasource.oneResource.total_count
    results       = data.saviynt_security_systems_datasource.oneResource.results
  }
}