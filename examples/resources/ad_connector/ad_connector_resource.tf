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

resource "saviynt_ad_connection_resource" "ss" {
  connection_type     = "AD"
  connection_name     = "shaleentestingaftercommit"
  url                 = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  username            = var.BIND_USER
  password=var.PASSWORD
  vault_connection    = var.VAULT_CONNECTION
  vault_configuration = var.VAULT_CONFIG
  save_in_vault       = var.SAVE_IN_VAULT
  ldap_or_ad="AD"
  group_search_base_dn="DC=corp,DC=SaviyntAD,DC=com"
  entitlement_attribute="memberOf"
  account_attribute="[customproperty1::cn#String,customproperty30::userAccountControl#String,customproperty2::userPrincipalName#String,customproperty28::primaryGroupID#String,lastlogondate::lastLogon#millisec,displayname::name#String,customproperty25::company#String,customproperty20::employeeID#String,customproperty3::sn#String,comments::distinguishedName#String,customproperty4::homeDirectory#String,lastpasswordchange::pwdLastSet#millisec,customproperty5::co#String,customproperty6::employeeNumber#String,customproperty7::givenName#String,customproperty8::title#String,customproperty9::telephoneNumber#String,customproperty10::c#String,description::description#String,customproperty11::uSNCreated#String,validthrough::accountExpires#millisec,customproperty12::logonCount#String,customproperty13::physicalDeliveryOfficeName#String,updatedate::whenChanged#date,customproperty14::extensionAttribute1#String,customproperty15::extensionAttribute2#String,customproperty16::streetAddress#String,customproperty17::mailNickname#String,customproperty18::department#String,customproperty19::countryCode#String,name::sAMAccountName#String,customproperty21::manager#String,customproperty22::homePhone#String,customproperty23::mobile#String,created_on::whenCreated#date,accountclass::objectClass#String,accountid::objectGUID#Binary,customproperty24::userAccountControl#String,customproperty27::objectSid#Binary,RECONCILATION_FIELD::customproperty26,customproperty26::objectGUID#Binary,customproperty29::st#String]"
  status_threshold_config=jsonencode({"statusAndThresholdConfig":{"statusColumn":"customproperty30","activeStatus":["512","544","66048"],"inactiveStatus":["546","514","66050"],"deleteLinks":false,"accountThresholdValue":1000,"correlateInactiveAccounts":true,"inactivateAccountsNotInFile":false,"lockedStatusColumn":"","lockedStatusMapping":{"Locked":[""],"Unlocked":[""]}}})
  account_name_rule="CN=$${user.username},CN=Users,DC=corp,DC=SaviyntAD,DC=com"
  remove_account_action=jsonencode({ "removeAction":"SUSPEND","userAccountControl":"546" })
  set_random_password="FALSE"
  create_account_json=jsonencode(
    {"cn":"$${user.username}","displayname":"$${user.displayname}","givenname":"$${user.firstname}","mail":"$${user.email}","name":"$${user.displayname}","objectClass":["top","person","organizationalPerson","user"],"userAccountControl":"544","sAMAccountName":"$${task.accountName}","sn":"$${user.lastname}","title":"$${user.title}"}
  )
  update_account_json =jsonencode(
   {"department":"$${user.departmentname}","streetAddress":"$${user.street}","title":"$${user.title}","sn":"$${user.lastname}","displayName":"$${user.firstname}","cn":"$${user.username}","manager":"$${managerAccount?.accountID}"}
  )
  status_key_json=jsonencode(
    {
         "STATUS_ACTIVE": [
                 "1", "ACTIVE", "true", "512", "544"
         ],
         "STATUS_INACTIVE": [
                 "0", "INACTIVE", "false", "546", "514"
         ]
 }
  )
  searchfilter="CN=Users,DC=corp,DC=SaviyntAD,DC=com"
  group_import_mapping=jsonencode(
    {"importGroupHierarchy":"true","entitlementTypeName":"memberOf","performGroupAccountLinking":"true","incrementalTimeField":"whenChanged","groupObjectClass":"(objectclass=group)","mapping":"memberHash:member_char,customproperty1:sAMAccountType_char,customproperty2:instanceType_char,customproperty3:uSNCreated_char,customproperty4:groupType_char,customproperty5:dSCorePropagationData_char,customproperty12:dn_char,customproperty13:cn_char,lastscandate:whenCreated_date,customproperty15:managedBy_char,entitlement_glossary:description_char,description:description_char,displayname:name_char,customproperty9:name_char,customproperty10:objectCategory_char,customproperty11:sAMAccountName_char,entitlement_value:distinguishedName_char,entitlementid:distinguishedName_char,customproperty14:objectClass_char,updatedate:whenChanged_date,customproperty17:distinguishedName_char,RECONCILATION_FIELD:customproperty17,customproperty18:objectGUID_Binary","activeGroupPossibleValues":[],"entitlementOwnerAttribute":"managedBy","tableFieldAttribute":"comments"}
  )
}