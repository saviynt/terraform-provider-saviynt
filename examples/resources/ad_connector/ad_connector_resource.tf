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
  connection_name     = "dellwireles3490123"
  url                 = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  username            = var.BIND_USER
  password=var.PASSWORD
  vault_connection    = var.VAULT_CONNECTION
  vault_configuration = var.VAULT_CONFIG
  save_in_vault       = var.SAVE_IN_VAULT


   account_name_rule="CN=$${user.username},CN=Users,DC=corp,DC=SaviyntAD,DC=com"
   account_attribute="[CUSTOMPROPERTY1::samaccountname#String,CUSTOMPROPERTY2::userprincipalname#String,LASTLOGONDATE::lastLogon#millisec,DISPLAYNAME::cn#String,CUSTOMPROPERTY25::company#String,CUSTOMPROPERTY3::sn#String,COMMENTS::distinguishedname#String,CUSTOMPROPERTY4::homedirectory#String,LASTPASSWORDCHANGE::pwdlastset#millisec,CUSTOMPROPERTY5::co#String,CUSTOMPROPERTY6::cn#String,CUSTOMPROPERTY7::givenname#String,CUSTOMPROPERTY8::title#String,CUSTOMPROPERTY9::telephonenumber#String,CUSTOMPROPERTY10::c#String,DESCRIPTION::description#String,CUSTOMPROPERTY11::usncreated#String,VALIDTHROUGH::accountexpires#millisec,CUSTOMPROPERTY12::logoncount#String,CUSTOMPROPERTY13::physicaldeliveryofficename#String,UPDATEDATE::whenchanged#date,CUSTOMPROPERTY14::extensionattribute1#String,CUSTOMPROPERTY15::extensionattribute2#String,CUSTOMPROPERTY16::streetaddress#String,CUSTOMPROPERTY17::mailnickname#String,CUSTOMPROPERTY18::department#String,CUSTOMPROPERTY19::countrycode#String,NAME::distinguishedname#String,CUSTOMPROPERTY21::manager#String,CUSTOMPROPERTY22::homephone#String,CUSTOMPROPERTY23::mobile#String,ACCOUNTCLASS::objectclass#String,ACCOUNTID::distinguishedname#String,CUSTOMPROPERTY24::useraccountcontrol#String,status::useraccountcontrol#Number,CUSTOMPROPERTY26::objectguid#String,CUSTOMPROPERTY28::forest#String,CUSTOMPROPERTY29::domain#string,CUSTOMPROPERTY30::objectclass#String]"
   entitlement_attribute="memberOf"
   user_attribute="[FIRSTNAME::givenname#String,LASTNAME::sn#String,CUSTOMPROPERTY1::samaccountname#String,USERNAME::distinguishedname#String,DISPLAYNAME::cn#String,CUSTOMPROPERTY25::description#String,CUSTOMPROPERTY3::sn#String,COMMENTS::distinguishedname#String,CUSTOMPROPERTY4::homedirectory#String,CUSTOMPROPERTY5::co#String,CUSTOMPROPERTY6::cn#String,CUSTOMPROPERTY7::givenname#String,CUSTOMPROPERTY8::title#String,CUSTOMPROPERTY9::telephonenumber#String,CUSTOMPROPERTY10::c#String,CUSTOMPROPERTY11::uSNCreated#String,ENDDATE::accountExpires#millisec,CUSTOMPROPERTY12::logonCount#String,CUSTOMPROPERTY13::physicaldeliveryofficename#String,UPDATEDATE::whenchanged#date,CUSTOMPROPERTY14::extensionattribute1#String,CUSTOMPROPERTY15::extensionattribute2#String,CUSTOMPROPERTY16::streetaddress#String,CUSTOMPROPERTY17::mailnickname#String,CUSTOMPROPERTY18::department#String,CUSTOMPROPERTY19::countrycode#String,CUSTOMPROPERTY2::samaccountname#String,CUSTOMPROPERTY20::userprincipalname#String,CUSTOMPROPERTY21::manager#String,CUSTOMPROPERTY22::homephone#String,CUSTOMPROPERTY23::mobile#String,CREATEDATE::whencreated#date,customproperty24::useraccountcontrol#String,CUSTOMPROPERTY26::distinguishedname#String,statuskey::useraccountcontrol#String,CUSTOMPROPERTY27::objectguid#String,RECONCILATION_FIELD::CUSTOMPROPERTY27,CUSTOMPROPERTY28::forest#String,CUSTOMPROPERTY29::domain#string]"
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
  
  create_update_mappings = jsonencode({
    "cn": "$${role?.customproperty27}",
"objectCategory": "CN=Group,CN=Schema,CN=Configuration,DC=corp,DC=SaviyntAD,DC=com",
"distinguishedName": "$${role?.role_name}",
"displayName": "$${role?.displayname}",
"SamAccountName": "$${role?.customproperty27}",
"description": "$${role?.description}",
"objectClass": "group",
"name": "$${role?.customproperty27}",
"managedBy":"$${user?.comments}"
  }
  )
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
objectfilter=jsonencode(
  {"full":"(&(objectCategory=person)(objectClass=user)(sAMAccountName=*))","incremental":"(&(objectCategory=person)(objectClass=user)(sAMAccountName=*))"}
)
searchfilter="OU=CONNQA,OU=SaviyntTeams,DC=saviyntlabs,DC=org"
create_account_json=jsonencode(
 {"samaccountname":"$${task.accountName}","sn":"$${user.lastname}","displayName":"$${user.displayname}","cn":"$${cn}","objectclass":["top","person","organizationalPerson","user"],  "userAccountControl":"544","givenName":"$${user.firstname}","name":"$${user.displayname}"}
)
status_threshold_config = jsonencode({
  statusAndThresholdConfig = {
    statusColumn               = "customproperty24"
    activeStatus              = ["512", "544"]
    deleteLinks               = false
    accountThresholdValue     = 1000
    correlateInactiveAccounts = true
    inactivateAccountsNotInFile = false
  }
})


}