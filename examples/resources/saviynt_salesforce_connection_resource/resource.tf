resource "saviynt_salesforce_connection_resource" "ss" {
  connection_type  = "SalesForce"
  connection_name  = "namefortheconnection"
  client_id="XXXXXX"
  client_secret="XXXXXX"
  refresh_token="XXXXXX.THpP30pYSpM738toa_9PLv8B5rYw6bCHJ6rMR143AOHvVvm"
  redirect_uri="https://salesforce.com"
  instance_url="https://salesforce.com"
  object_to_be_imported="Profile,Group,PermissionSet,Role"
  account_field_query="Id,Username,LastName,FirstName,Name,CompanyName,Email,IsActive,UserRoleId,ProfileId,UserType,ManagerId,LastLoginDate,LastPasswordChangeDate,CreatedDate,CreatedById,LastModifiedDate,LastModifiedById,SystemModstamp,ContactId,AccountId,FederationIdentifier,UserPermissionsSupportUser"
  field_mapping_json=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  createaccountjson=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  modifyaccountjson = jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  status_threshold_config=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
  customconfigjson=jsonencode({
        "UpdateUserQry": [
        "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username =$${user.username}"
    ]
  })
}
