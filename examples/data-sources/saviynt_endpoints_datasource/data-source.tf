# List all endpoints
data "saviynt_endpoints_datasource" "all" {}

# Get endpoint by name
data "saviynt_endpoints_datasource" "by_name" {
  endpoint_name = "sample"
}