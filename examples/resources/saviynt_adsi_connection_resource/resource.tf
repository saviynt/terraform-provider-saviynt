resource "saviynt_adsi_connection_resource" "example" {
    connection_type = "ADSI"
    connection_name = "Shaleen_testing_ADSI_terraform_3490"
    url="ldap://example-ldap.local:389"
    status_threshold_config = jsonencode({
    statusAndThresholdConfig = {
    accountStatusAttribute = "userAccountControl",
    activeStatusValue      = ["512", "544"],
    suspendedStatusValue   = ["546"],
    disableStatusValue     = ["514"],
    thresholdKey           = "lockoutThreshold",
    lockoutValue           = "3"
    }
    })
    username= "CN=admin,CN=Users,DC=example,DC=com"
    connection_url="http://ad.example.com:389"
    forestlist="saviyntlabs.org"
    password="XXXXXX"
    searchfilter="DC=saviyntlabs,DC=org"
    objectfilter="(&(objectCategory=person)(objectClass=user))"
    user_attribute="USERNAME::sAMAccountName#String"
    statuskeyjson=jsonencode({
        "STATUS_ACTIVE": [
          "512",
          "544",
          "66048"
        ],
        "STATUS_INACTIVE": [
          "546",
          "514"
        ]
    })
    account_attribute= "USERNAME::sAMAccountName#String"
    entitlement_attribute="memberof"
    accountnamerule="USERNAME::sAMAccountName#String"
    checkforunique = jsonencode({
      CheckForUnique = {
        Attributes = [
          {
            samaccountname = "customproperty1"
            RuleCheck      = "$${user.lastname}###$${user.lastname}1###$${user.lastname}2###$${user.lastname}3###$${user.lastname}4###$${user.lastname}5###$${user.lastname}6###$${user.lastname}7###$${user.lastname}8"
          },
        ]
      }
    })  
    group_search_base_dn="DC=saviyntlabs,DC=org"
    group_import_mapping=jsonencode({
    "importGroupHierarchy": "true",
    })
    import_nested_membership="false"
    page_size="1000"
    provisioning_url="http://ad.example.com:389"
    createaccountjson=jsonencode({
      "objects":
          "objectClasses": [
            "user",
            "top",
            "Person",
            "OrganizationalPerson"
          ],
    })
    updateaccountjson=jsonencode({
      "objects":
          "objectClasses": [
            "user"
          ],
    })
    addaccessjson=jsonencode({
      "objects": 
        {
          "objectClasses": [
            "user"
          ],
        }
      }
    )
    removeaccessjson=jsonencode()
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
    enableaccountjson=<<EOF
    {"objects":[{"objectClasses":["user"],"distinguishedName":"$${account.accountID}","deleteAllGroups":false,"attributes":{"userAccountControl":512}}]}
    EOF
    disableaccountjson=<<EOF
    {"objects":[{"objectClasses":["user"],"distinguishedName":"$${account.accountID}","deleteAllGroups":false,"attributes":{"userAccountControl":514}}]}
    EOF
    removeaccountjson=<<EOF
    {"objects": [{"distinguishedName": "$${account.accountID}","removeAction": "DELETE","deleteChildObjects": false}]}
    EOF
    resetandchangepasswrdjson=<<EOF
    { "objects":[ { "objectClasses":[ "user" ], "password":"$${password}", "distinguishedName":"", "attributes":{ "pwdLastSet":"$${pwdLastSet}" } } ]}
    EOF
    creategroupjson=jsonencode(
        {
      "objects": [
        {
          "objectClasses": [
            "group"
          ],
          "baseDn": "$${role.customproperty24}",
          "attributes": {
            "cn": "$${role.displayname}",
            "name": "$${role.displayname}",
            "description": "$${role.description}",
            "displayName": "$${role.displayname}",
            "groupType": "$${role?.customproperty21 == 'Security' && role?.customproperty22 == 'Global'?'-2147483646' : role?.customproperty21=='Security'&&role?.customproperty22=='Universal'?'-2147483640' : role?.customproperty21== 'Security'&&role?.customproperty22=='Domain Local' ? '-2147483644':role?.customproperty21=='Distribution'&&role?.customproperty22=='Global' ? '2':role?.customproperty21== 'Distribution'&&role?.customproperty22=='Universal'?'8':role?.customproperty21=='Distribution'&& role?.customproperty22=='Domain Local'?'4':''}"
          }
        }
      ]
    })
    updategroupjson=jsonencode({
      "objects": [
        {
          "objectClasses": [
            "group"
          ],
          "distinguishedName": "$${role.role_name}",
          "attributes": {
            "description": "$${role.description}",
            "proxyAddresses": "$${role.customproperty20}"
          }
        }
      ]
    })
    removegroupjson=<<EOF
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
    addaccessentitlementjson=<<EOF
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
    removeaccessentitlementjson=<<EOF
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
    createserviceaccountjson=<<EOF
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
    updateserviceaccountjson=<<EOF
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
    removeserviceaccountjson=<<EOF
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