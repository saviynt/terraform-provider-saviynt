# List all the security systems
data "saviynt_security_systems_datasource" "all" {}


# Get a security system by its name
data "saviynt_security_systems_datasource" "by_name" {
  systemname = "sample"
}