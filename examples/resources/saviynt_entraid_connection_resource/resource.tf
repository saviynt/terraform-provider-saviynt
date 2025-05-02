resource "saviynt_entraid_connection_resource" "ss" {
  connection_type           = "AzureAD"
  connection_name           = "namefortheconnection"
  client_id                 = "XXXXXX"
  client_secret             = "XXXXXX"
  aad_tenant_id             = "XXXXXX"
  authentication_endpoint   = "https://login.microsoft.com/"
  microsoft_graph_endpoint  = "https://graph.microsoft.com/"
  azure_management_endpoint = "https://azure.com/"
  account_attributes = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  account_import_fields = "accountEnabled,mail,businessPhone,surname,givenName,displayName,userPrincipalName,id"
  status_threshold_config = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  endpoints_filter = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  connection_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  import_user_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  create_account_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  update_account_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  enable_account_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  disable_account_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  add_access_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  remove_access_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  remove_account_json = jsonencode({
    "UpdateUserQry" : [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
}
