resource "saviynt_adsi_connection_resource" "example" {
    connection_type = "ADSI"
    connection_name = "namefortheconnection"
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
    })
    removeaccessjson=jsonencode(
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
    })
    enableaccountjson=jsonencode(
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
    })
    disableaccountjson=jsonencode(
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
    })
    removeaccountjson=jsonencode(
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
    })
    resetandchangepasswrdjson=jsonencode(
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
    })
    creategroupjson=jsonencode(
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
    })
    updategroupjson=jsonencode(
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
    })
    removegroupjson=jsonencode(
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
    })
    addaccessentitlementjson=jsonencode(
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
    })
    removeaccessentitlementjson=jsonencode(
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
    })
    createserviceaccountjson=jsonencode(
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
    })
    updateserviceaccountjson=jsonencode(
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
    })
    removeserviceaccountjson=jsonencode(
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
    })
}