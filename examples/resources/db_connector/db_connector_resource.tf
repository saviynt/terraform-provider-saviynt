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
  connection_type = "DB"
  connection_name = "Shaleen_testing_db_terraform"
  username="connadmin"
  description="description"
  password=var.PASSWORD
  driver_name="com.mysql.jdbc.Driver"
  create_account_json=<<EOF
{"CreateAccountQry":["INSERT INTO mysqllocal.users(username, firstname, lastname, statuskey, departmentname, displayname, manager, enabled, updatedate) VALUES ('$${user.username}', '$${user.firstname}', '$${user.lastname}', '$${user.statuskey}', '$${user.departmentname}', '$${user.displayname}', '$${user.manager}', '$${user.enabled}', CURRENT_TIMESTAMP)","INSERT INTO mysqllocal.accounts(AccountName, EntitlementType, EntitlementValue, Status, UPDATEDATE) VALUES ('$${user.username}', 'DEFAULT', 'BASE_ACCESS', 1, CURRENT_TIMESTAMP)"]}
EOF
  update_account_json=jsonencode({
    "UpdateAccountQry": [
    "UPDATE mysqllocal.users SET firstname =$${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
]
  })
  grant_access_json=jsonencode({
    "ENTITLEMENT_TYPE1": [
    "INSERT INTO mysqllocal.accounts(AccountName, EntitlementType, EntitlementValue, Status, UPDATEDATE) VALUES($${user.username}, $${entitlement.type}, $${entitlement.value}, 1, CURRENT_TIMESTAMP)"
],
"ROLE_ACCESS": [
    "INSERT INTO mysqllocal.roles(role_name, description, displayname, status, updatedate) VALUES($${role.name}, $${role.description}, $${role.displayname}, 1, CURRENT_TIMESTAMP)"
]
  })
  revoke_access_json=jsonencode({
    "ENTITLEMENT_TYPE1": [
    "UPDATE mysqllocal.accounts SET Status = 0, UPDATEDATE = CURRENT_TIMESTAMP WHERE AccountName = $${user.username} AND EntitlementType = $${entitlement.type} AND EntitlementValue = $${entitlement.value}"
]
  })
  change_pass_json=jsonencode({
    "ChangePasswordQry": [
    "UPDATE mysqllocal.users SET password = $${user.password}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
]
  })
  delete_account_json=jsonencode({
    "DeleteAccountQry": [
    "DELETE FROM mysqllocal.users WHERE username = $${user.username}",
    "DELETE FROM mysqllocal.accounts WHERE AccountName = $${user.username}"
]
  })
  enable_account_json=jsonencode({
    "EnableAccountQry": [
    "UPDATE mysqllocal.users SET enabled = 1, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}",
    "UPDATE mysqllocal.accounts SET Status = 1, UPDATEDATE = CURRENT_TIMESTAMP WHERE AccountName = $${user.username}"
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
}

check "instance_health" {
  assert {
    condition     = saviynt_db_connection_resource.example.error_code != "1"
    error_message = "The error is: ${saviynt_db_connection_resource.example.msg}"
  }
}
