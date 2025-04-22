resource "saviynt_adsi_connection_resource" "sample" {
  connection_type = "ADSI"
  connection_name = "sample-adsi-connector"
  url             = "<URL>"
  status_threshold_config = jsonencode({
    "statusAndThresholdConfig" : {
      "statusColumn" : "customproperty24",
      "activeStatus" : [
        "512",
        "544",
        "66048"
      ],
      "deleteLinks" : false,
      "accountThresholdValue" : 50000,
      "correlateInactiveAccounts" : true,
      "inactivateAccountsNotInFile" : false,
      "deleteAccEntForActiveAccounts" : false
    }
  })
  username       = "<USERNAME>"
  connection_url = "<CONNECTION_URL>"
  forestlist     = "<FORESTLIST>"
  password       = "<PASSWORD>"
  searchfilter   = "<SEARCH_FILTER>"
  objectfilter   = "OBJECT_FILTER>"
  user_attribute = "<USER_ATTRIBUTE>"

  account_attribute        = "<ACCOUNT_ATTRIBUTE>"
  entitlement_attribute    = "<ENTITLEMENT_ATTRIBUTE>"
  accountnamerule          = "<ACCOUNT_NAME_RULE>"
  checkforunique           = "<CHECK_FOR_UNIQUE>"
  group_search_base_dn     = "<GROUP_SEARCH_BASE_DN>"
  group_import_mapping     = <<EOF
{
  "importGroupHierarchy": "true",
  "entitlementTypeName": "memberOf",
  "performGroupAccountLinking": "true",
  "groupObjectClass": "(objectclass=group)",
  "entitlementOwnerAttribute": "managedby",
  "tableFieldAttribute": "accountID",
  "mapping": "memberHash:memberof_char,customProperty1:samaccounttype_char,customProperty2:instancetype_char,customProperty3:usncreated_char,customProperty4:grouptype_char,customProperty5:dscorepropagationdata_char,customProperty12:dn_char,customProperty13:cn_char,lastscandate:whencreated_date,customProperty15:managedBy_char,entitlement_glossary:description_char,description:description_char,displayname:name_char,customProperty9:name_char,customProperty10:objectcategory_char,customProperty11:samaccounttype_char,entitlement_value:distinguishedname_char,entitlementid:distinguishedname_char,customProperty14:objectclass_char,updatedate:whenchanged_date,customProperty17:distinguishedname_char,RECONCILATION_FIELD:customProperty18,customProperty18:objectguid_char"
}
EOF
  import_nested_membership = "false"
  page_size                = "1000"
  provisioning_url         = "<URL>"
  createaccountjson = jsonencode({
    "objects" : [
      {
        "objectClasses" : [
          "user",
          "top",
          "Person",
          "OrganizationalPerson"
        ],
        "baseDn" : "<...>",
        "password" : "<...>",
        "attributes" : {
          "sn" : "$${user.lastname}",
          "sAMAccountName" : "$${task.accountName}",
          "cn" : "$${task.accountName}",
          "userAccountControl" : 512,
          "co" : "",
          "department" : "$${user.departmentname}",
          "displayName" : "$${user.displayname}",
          "employeeID" : "$${user.employeeid}",
          "employeeNumber" : "1",
          "employeeType" : "$${user.employeeType}",
          "givenName" : "$${user.firstname}",
          "l" : "$${user.city}",
          "mail" : "$${user.email}",
          "proxyAddresses" : [
            "SMTP:test@test.org",
            "SIP:Test@test1.org"
          ]
        }
      }
    ]
  })
  updateaccountjson = jsonencode(
    {
      "objects" : [
        {
          "objectClasses" : [
            "user"
          ],
          "distinguishedName" : "$${account.accountID?.replace('\\', '\\\\')?.replace('/', '\\/')}",
          "attributes" : {
            "displayName" : "$${user.displayname}",
            "streetAddress" : "$${user.street}",
            "additionalAttributes" : {
              "departmentName" : "PM",
              "companyName" : "<...>"
            }
          }
        }
      ]
  })
  addaccessjson             = <<EOF
{
  "objects": [
    {
      "objectClasses": [
        "user"
      ],
      "distinguishedName": "$${accountID}",
      "addGroup": "$${entitlement_values}"
    }
  ],
  "requestConfiguration": {
    "grpMemExistenceChk": {
      "enable": true
    }
  }
}
EOF
  removeaccessjson          = <<EOF
{
  "objects": [
    {
      "objectClasses": [
        "group"
      ],
      "distinguishedName": "$${accountID}",
      "removeGroup": "$${entitlement_values}"
    }
  ],
  "requestConfiguration": {
    "grpMemExistenceChk": {
      "enable": true
    }
  }
}
EOF
  enableaccountjson         = <<EOF
{"objects":[{"objectClasses":["user"],"distinguishedName":"$${account.accountID}","deleteAllGroups":false,"attributes":{"userAccountControl":512}}]}
EOF
  disableaccountjson        = <<EOF
{"objects":[{"objectClasses":["user"],"distinguishedName":"$${account.accountID}","deleteAllGroups":false,"attributes":{"userAccountControl":514}}]}
EOF
  removeaccountjson         = <<EOF
{"objects": [{"distinguishedName": "$${account.accountID}","removeAction": "DELETE","deleteChildObjects": false}]}
EOF
  resetandchangepasswrdjson = <<EOF
{ "objects":[ { "objectClasses":[ "user" ], "password":"$${password}", "distinguishedName":"", "attributes":{ "pwdLastSet":"$${pwdLastSet}" } } ]}
EOF
  creategroupjson = jsonencode(
    {
      "objects" : [
        {
          "objectClasses" : [
            "group"
          ],
          "baseDn" : "$${role.customproperty24}",
          "attributes" : {
            "cn" : "$${role.displayname}",
            "name" : "$${role.displayname}",
            "description" : "$${role.description}",
            "displayName" : "$${role.displayname}",
            "groupType" : "$${role?.customproperty21 == 'Security' && role?.customproperty22 == 'Global'?'-2147483646' : role?.customproperty21=='Security'&&role?.customproperty22=='Universal'?'-2147483640' : role?.customproperty21== 'Security'&&role?.customproperty22=='Domain Local' ? '-2147483644':role?.customproperty21=='Distribution'&&role?.customproperty22=='Global' ? '2':role?.customproperty21== 'Distribution'&&role?.customproperty22=='Universal'?'8':role?.customproperty21=='Distribution'&& role?.customproperty22=='Domain Local'?'4':''}"
          }
        }
      ]
  })
  updategroupjson = jsonencode({
    "objects" : [
      {
        "objectClasses" : [
          "group"
        ],
        "distinguishedName" : "$${role.role_name}",
        "attributes" : {
          "description" : "$${role.description}",
          "proxyAddresses" : "$${role.customproperty20}"
        }
      }
    ]
  })
  removegroupjson             = <<EOF
{
  "objects": [
    {
      "distinguishedName": "$${role.role_name}",
      "removeAction": "DELETE",
      "deleteChildObjects": false
    }
  ]
}
EOF
  addaccessentitlementjson    = <<EOF
{
   "objects":[
      {
         "objectClasses":[
            "group"
         ],
         "distinguishedName":"$${ent1Value.entitlement_value?.replace('\\', '\\\\')?.replace('/', '\\/')}",
         "addGroup":"$${ent2Value.entitlement_value?.replace('\\', '\\\\')?.replace('/', '\\/')}"
      }
   ],
   "requestConfiguration":{
      "grpMemExistenceChk":{
         "enable":true
      }
   }
}
EOF
  removeaccessentitlementjson = <<EOF
{
   "objects":[
      {
         "objectClasses":[
            "group"
         ],
         "distinguishedName":"$${ent1Value.entitlement_value?.replace('\\', '\\\\')?.replace('/', '\\/')}",
         "removeGroup":"$${ent2Value.entitlement_value?.replace('\\', '\\\\')?.replace('/', '\\/')}"
      }
   ],
   "requestConfiguration":{
      "grpMemExistenceChk":{
         "enable":true
      }
   }
}
EOF
  createserviceaccountjson    = <<EOF
{
  "objects": [
    {
      "objectClasses": [
        "user",
        "top",
        "Person",
        "OrganizationalPerson"
      ],
      "baseDn": "$${baseDN}",
      "password": "$${password}",
      "attributes": {
        "sAMAccountName": "$${task.accountName}",
        "cn": "$${task.accountName}",
        "displayname": "testDP",
        "userAccountControl": 512
      }
    }
  ]
}
EOF
  updateserviceaccountjson    = <<EOF
{
  "objects": [
    {
      "objectClasses": [
        "user",
        "top",
        "Person",
        "OrganizationalPerson"
      ],
      "distinguishedName": "$${account.accountID?.replace('\\', '\\\\')?.replace('/', '\\/')}",
      "attributes": {
        "pwdLastSet": 0,
        "displayName": "testDPUpdated"
      }
    }
  ]
}
EOF
  removeserviceaccountjson    = <<EOF
{
  "objects": [
    {
      "distinguishedName": "$${account.accountID?.replace('\\', '\\\\')?.replace('/', '\\/')}",
      "removeAction": "DELETE",
      "deleteChildObjects": false
    }
  ]
}
EOF
}