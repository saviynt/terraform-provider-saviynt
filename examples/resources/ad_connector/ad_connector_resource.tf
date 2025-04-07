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
  connection_name     = "dellwireless"
  url                 = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  username            = var.BIND_USER
  password=var.PASSWORD
  vault_connection    = var.VAULT_CONNECTION
  vault_configuration = var.VAULT_CONFIG
  save_in_vault       = var.SAVE_IN_VAULT

  enable_account_json=jsonencode(
{
  "USEDNFROMACCOUNT": "YES",
  "MOVEDN": "NO",
  "REMOVEGROUPS": "NO",
  "ENABLEACCOUNTOU": "CN=Users,DC=corp,DC=SaviyntAD,DC=com",
  "AFTERMOVEACTIONS": {
    "userAccountControl": "544",
    "otherMailbox": [
      "$${user?.customproperty15.toString().replace(',','\",\"')}"
    ]
  }
})
   account_name_rule="CN=$${user.username},CN=Users,DC=corp,DC=SaviyntAD,DC=com"
   account_attribute="[CUSTOMPROPERTY1::CN#String,CUSTOMPROPERTY2::userPrincipalName#String,LASTLOGONDATE::lastLogon#millisec,DISPLAYNAME::displayName#String,CUSTOMPROPERTY25::company#String,CUSTOMPROPERTY3::sn#String,COMMENTS::distinguishedName#String,CUSTOMPROPERTY4::homeDirectory#String,LASTPASSWORDCHANGE::pwdLastSet#millisec,CUSTOMPROPERTY5::co#String,CUSTOMPROPERTY6::cn#String,CUSTOMPROPERTY7::givenName#String,CUSTOMPROPERTY8::title#String,CUSTOMPROPERTY9::telephoneNumber#String,CUSTOMPROPERTY10::c#String,DESCRIPTION::description#String,CUSTOMPROPERTY11::uSNCreated#String,VALIDTHROUGH::accountExpires#millisec,CUSTOMPROPERTY12::logonCount#String,CUSTOMPROPERTY13::physicalDeliveryOfficeName#String,UPDATEDATE::whenChanged#date,CUSTOMPROPERTY14::extensionAttribute1#String,CUSTOMPROPERTY15::extensionAttribute2#String,CUSTOMPROPERTY16::streetAddress#String,CUSTOMPROPERTY17::mailNickname#String,CUSTOMPROPERTY18::department#String,CUSTOMPROPERTY19::countryCode#String,NAME::name#String,CUSTOMPROPERTY21::manager#String,CUSTOMPROPERTY22::homePhone#String,CUSTOMPROPERTY23::mobile#String,CREATED_ON::whenCreated#date,ACCOUNTCLASS::objectClass#String,ACCOUNTID::distinguishedName#String,CUSTOMPROPERTY24::userAccountControl#String,STATUS::userAccountControl#Number,CUSTOMPROPERTY26::objectGUID#Binary,CUSTOMPROPERTY27::objectSid#Binary]"
   entitlement_attribute="memberOf"
   user_attribute="[CUSTOMPROPERTY1::CN#String,USERNAME::name#String,DISPLAYNAME::displayName#String,CUSTOMPROPERTY25::company#String,CUSTOMPROPERTY3::sn#String,COMMENTS::distinguishedName#String,CUSTOMPROPERTY4::homeDirectory#String,CUSTOMPROPERTY5::co#String,CUSTOMPROPERTY6::cn#String,CUSTOMPROPERTY7::givenName#String,CUSTOMPROPERTY8::title#String,CUSTOMPROPERTY9::telephoneNumber#String,CUSTOMPROPERTY10::c#String,CUSTOMPROPERTY11::uSNCreated#String,ENDDATE::accountExpires#millisec,CUSTOMPROPERTY12::logonCount#String,CUSTOMPROPERTY13::physicalDeliveryOfficeName#String,UPDATEDATE::whenChanged#date,CUSTOMPROPERTY14::extensionAttribute1#String,CUSTOMPROPERTY15::extensionAttribute2#String,CUSTOMPROPERTY16::streetAddress#String,CUSTOMPROPERTY17::mailNickname#String,CUSTOMPROPERTY18::department#String,CUSTOMPROPERTY19::countryCode#String,CUSTOMPROPERTY2::sAMAccountName#String,CUSTOMPROPERTY20::userPrincipalName#String,CUSTOMPROPERTY21::manager#String,CUSTOMPROPERTY22::homePhone#String,CUSTOMPROPERTY23::mobile#String,CREATEDATE::whenCreated#date,customproperty24::userAccountControl#String,STATUSKEY::userAccountControl#Number]"
   endpoints_filter=jsonencode({"AD_Child":[{"memberOf":["CN=QAGroup01,DC=corp,DC=SaviyntAD,DC=com"]}]})
   update_user_json=jsonencode({ "sn": "$${user.lastname}"})
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
   group_search_base_dn="DC=corp,DC=SaviyntAD,DC=com"
  create_update_mappings = "chinatown"
  group_import_mapping=jsonencode({
  "importGroupHierarchy": "true",
  "entitlementTypeName": "",
  "performGroupAccountLinking": "true",
  "incrementalTimeField": "whenChanged",
  "groupObjectClass": "(objectclass=group)",
  "mapping": "memberHash:member_char,customProperty1:sAMAccountType_char,customProperty2:instanceType_char,customProperty3:uSNCreated_char,customProperty4:groupType_char,customProperty5:dSCorePropagationData_char,customProperty12:dn_char,customProperty13:cn_char,lastscandate:whenCreated_date,customProperty15:managedBy_char,entitlement_glossary:description_char,description:description_char,displayname:name_char,customProperty9:name_char,customProperty10:objectCategory_char,customProperty11:sAMAccountName_char,entitlement_value:distinguishedName_char,entitlementid:distinguishedName_char,customProperty14:objectClass_char,updatedate:whenChanged_date,customProperty17:distinguishedName_char,RECONCILATION_FIELD:customproperty18,customProperty18:objectGUID_Binary,customProperty19:managedBy_char",
  "entitlementOwnerAttribute": "managedBy",
  "tableFieldAttribute": "accountID"
})
objectfilter="(&(objectCategory=person)(objectClass=user))"
searchfilter="CN=Users,DC=corp,DC=SaviyntAD,DC=com"
}