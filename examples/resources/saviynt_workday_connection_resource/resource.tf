resource "saviynt_workday_connection_resource" "ss" {
  connection_type    = "Workday"
  connection_name    = "namefortheconnection"
  base_url           = "https://example.com/"
  api_version        = "v34.0"
  tenant_name        = "example"
  report_owner       = "example"
  use_oauth          = "TRUE"
  username           = "example"
  password           = "XXXXX"
  client_id          = "XXXXX"
  client_secret      = "XXXXX"
  refresh_token      = "XXXXX"
  access_import_list = "Security Group, Domain Security Policy, Business Process Security Policy"
  status_key_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  user_import_payload = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  user_import_mapping = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
}