resource "saviynt_db_connection_resource" "example" {
  connection_type = "DB"
  connection_name = "namefortheconnection"
  url="jdbc:mysql:3490"
  username="connadmin"
  password=var.PASSWORD
  driver_name="com.mysql.jdbc.Driver"
  create_account_json=jsonencode({
    "UpdateAccountQry": [
    "UPDATE mysqllocal.users SET firstname =$${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
]
  })
  update_account_json=jsonencode({
    "UpdateAccountQry": [
    "UPDATE mysqllocal.users SET firstname =$${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
]
  })
  grant_access_json=jsonencode({
    "UpdateAccountQry": [
    "UPDATE mysqllocal.users SET firstname =$${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
]
  })
  revoke_access_json=jsonencode({
    "UpdateAccountQry": [
    "UPDATE mysqllocal.users SET firstname =$${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
]
  })
  change_pass_json=jsonencode({
    "UpdateAccountQry": [
    "UPDATE mysqllocal.users SET firstname =$${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
]
  })
  delete_account_json=jsonencode({
    "DeleteAccountQry": [
    "DELETE FROM mysqllocal.users WHERE username = $${user.username}",
    "DELETE FROM mysqllocal.accounts WHERE AccountName = $${user.username}"
]
  })
  enable_account_json=jsonencode({
    "UpdateAccountQry": [
    "UPDATE mysqllocal.users SET firstname =$${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
]
  })
  disable_account_json=jsonencode({
    "DisableAccountQry": [
    "UPDATE mysqllocal.users SET enabled = 0, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}",
    "UPDATE mysqllocal.accounts SET Status = 0, UPDATEDATE = CURRENT_TIMESTAMP WHERE AccountName = $${user.username}"
]
  })
  account_exists_json=jsonencode({
        "AccountExistQry": "SELECT username FROM mysqllocal.users WHERE username = $${user.username}"
  })
  update_user_json=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  accounts_import= "NAME::sAMAccountName#String"
}