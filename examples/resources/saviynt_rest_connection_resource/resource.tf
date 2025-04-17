resource "saviynt_rest_connection_resource" "example" {
  connection_type         = "REST"
  connection_name         = "namefortheconnection"
  vault_connection        = var.VAULT_CONNECTION
  vault_configuration     = var.VAULT_CONFIG
  save_in_vault           = var.SAVE_IN_VAULT
  config_json             = jsonencode({"showLogs":true})
  connection_json = jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  create_account_json = jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  import_account_ent_json =jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
update_account_json=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
enable_account_json=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
disable_account_json=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
add_access_json=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
remove_access_json=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
remove_account_json=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
change_pass_json=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
}