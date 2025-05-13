variable "CLIENT_ID" {
  type        = string
  description = "Saviynt AzureAD CLIENT_ID"
  sensitive   = true
}

variable "CLIENT_SECRET" {
  type        = string
  description = "Saviynt AzureAD CLIENT_SECRET"
  sensitive   = true
}

variable "TENANT_ID" {
  type        = string
  description = "Saviynt AzureAD TENANT_ID"
  sensitive   = true
}
resource "saviynt_entraid_connection_resource" "example" {
  connection_type           = "AzureAD"
  connection_name           = "Terraform_EntraId_Connector"
  client_id                 = var.CLIENT_ID
  client_secret             = var.CLIENT_SECRET
  aad_tenant_id             = var.TENANT_ID
  authentication_endpoint   = "https://login.microsoftonline.com"
  microsoft_graph_endpoint  = "https://graph.microsoft.com"
  azure_management_endpoint = "https://management.azure.com"
  create_users              = "YES"
  create_new_endpoints      = "YES"
  managed_account_type      = "ACCOUNTS"
  import_depth              = "FINE GRAINED"
  import_user_json = jsonencode({
    connection = "userAuth"
    method     = "GET"
    url        = "https://graph.microsoft.com/v1.0/users?$select=Id,userPrincipalName,accountEnabled,mail,userType,createdDateTime,country,preferredLanguage,displayName,surname,givenName,mobilePhone,businessPhones,mailNickname,mail",
    headers = {
      Authorization = "Bearer $${access_token}"
      Accept        = "application/json"
    }
    statusConfig = {
      active   = "true"
      inactive = "false"
    }
    colsToPropsMap = {
      username        = "userPrincipalName~#~char"
      displayname     = "displayName~#~char"
      firstname       = "givenName~#~char"
      lastname        = "surname~#~char"
      country         = "country~#~char"
      phonenumber     = "mobilePhone~#~char"
      statuskey       = "accountEnabled~#~char"
      email           = "mail~#~char"
      employeetype    = "userType~#~char"
      customproperty1 = "preferredLanguage~#~char"
      customproperty2 = "businessPhones~#~char"
      customproperty3 = "mailNickname~#~char"
      customproperty4 = "Id~#~char"
      customproperty5 = "userPrincipalName~#~char"
      customproperty6 = "createdDateTime~#~char"
    }
    userResponsePath = "value"
    pagination = {
      nextUrl = {
        nextUrlPath = "$${(response?.completeResponseMap?.get('@odata.nextLink')==null)? null : response?.completeResponseMap?.get('@odata.nextLink')}"
      }
    }
  })

  account_attributes = jsonencode({
    acctLabels = {
      customproperty1  = "FirstName"
      customproperty2  = "LastName"
      customproperty3  = "OfficePhone"
      customproperty10 = "AccountStatus"
    },
    colsToPropsMap = {
      accountID        = "id~#~char"
      name             = "userPrincipalName~#~char"
      customproperty1  = "givenName~#~char"
      customproperty2  = "surname~#~char"
      customproperty3  = "businessPhones~#~char"
      customproperty10 = "accountEnabled~#~char"
    }
  })

  account_import_fields = "accountEnabled,mail,businessPhone,surname,givenName,displayName,userPrincipalName,id"
  entitlement_attribute = jsonencode(
    {
      entitlementAttribute = {
        AADGroup = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "displayName~#~char",
            customproperty1   = "deletedDateTime~#~char",
            customproperty2   = "description~#~char",
            customproperty5   = "onPremisesSyncEnabled~#~char",
            customproperty6   = "onPremisesLastSyncDateTime~#~char",
            customproperty7   = "mail~#~char",
            customproperty8   = "mailEnabled~#~char",
            customproperty9   = "onPremisesSecurityIdentifier~#~char",
            customproperty10  = "securityEnabled~#~char",
            customproperty11  = "groupTypes~#~listAsString",
            customproperty12  = "membershipRule~#~char",
            customproperty13  = "membershipRuleProcessingState~#~char",
            customproperty16  = "resourceProvisioningOptions~#~char"
          }
        },
        AADGroupOwners = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "displayName~#~char"
          }
        },
        Team = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "displayName~#~char",
            description       = "description~#~char",
            customproperty1   = "internalId~#~char",
            customproperty2   = "webUrl~#~char",
            customproperty3   = "discoverySettings~#~char",
            customproperty6   = "isArchived~#~char"
          }
        },
        Channel = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "displayName~#~char",
            description       = "description~#~char",
            customproperty1   = "email~#~char",
            customproperty2   = "webUrl~#~char"
          }
        },
        DirectoryRole = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "displayName~#~char",
            customproperty4   = "description~#~char",
            customproperty6   = "deletedDateTime~#~char",
            customproperty8   = "roleTemplateId~#~char"
          }
        },
        Subscription = {
          colsToPropsMap = {
            entitlementID     = "subscriptionId~#~char",
            entitlement_value = "displayName~#~char",
            displayname       = "displayName~#~char",
            customproperty1   = "state~#~char",
            customproperty2   = "subscriptionPolicies.locationPlacementId~#~char",
            customproperty4   = "subscriptionPolicies.quotaId~#~char",
            customproperty6   = "subscriptionPolicies.spendingLimit~#~char",
            customproperty7   = "authorizationSource~#~char"
          }
        },
        Application = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "displayName~#~char",
            customproperty1   = "id~#~bool",
            customproperty2   = "resourceAppId~#~bool",
            customproperty4   = "orgRestrictions~#~boolListInverse",
            customproperty5   = "oauth2AllowImplicitFlow~#~bool",
            customproperty6   = "allowPublicClient~#~bool",
            customproperty7   = "createdDateTime~#~char"
          }
        },
        ApplicationInstance = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "displayName~#~char",
            displayname       = "appDisplayName~#~char",
            customproperty1   = "appId~#~char",
            customproperty4   = "appOwnerOrganizationId~#~char",
            customproperty5   = "appRoleAssignmentRequired~#~char",
            customproperty6   = "servicePrincipalNames~#~char",
            customproperty7   = "accountEnabled~#~bool",
            customproperty9   = "publisherName~#~char"
          }
        },
        SKU = {
          colsToPropsMap = {
            entitlementID     = "skuId~#~char",
            entitlement_value = "skuPartNumber~#~char",
            customproperty1   = "appliesTo~#~char",
            customproperty2   = "capabilityStatus~#~char",
            customproperty5   = "consumedUnits~#~char",
            customproperty7   = "prepaidUnits~#~listAsString"
          }
        },
        updateApplications = {
          colsToPropsMap = {
            customproperty1  = "id~#~char",
            customproperty11 = "createdDateTime~#~char"
          }
        },
        AppRole = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "displayName~#~char",
            customproperty1   = "isEnabled~#~char",
            customproperty2   = "value~#~char",
            customproperty4   = "id~#~char",
            customproperty5   = "allowedMemberTypes~#~char"
          }
        },
        Oauth2Permission = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "userConsentDisplayName~#~char",
            customproperty1   = "isEnabled~#~char",
            customproperty2   = "adminConsentDisplayName~#~char",
            customproperty4   = "id~#~char",
            customproperty5   = "type~#~char",
            customproperty6   = "userConsentDescription~#~char",
            customproperty7   = "adminConsentDescription~#~char",
            customproperty8   = "value~#~char"
          }
        },
        ApplicationInstanceAppRole = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "displayName~#~char",
            customproperty1   = "isEnabled~#~char",
            customproperty2   = "value~#~char",
            customproperty4   = "id~#~char",
            customproperty5   = "allowedMemberTypes~#~char"
          }
        },
        ApplicationInstanceOauth2Permission = {
          colsToPropsMap = {
            entitlementID     = "id~#~char",
            entitlement_value = "userConsentDisplayName~#~char",
            customproperty1   = "isEnabled~#~char",
            customproperty2   = "adminConsentDescription~#~char",
            customproperty4   = "id~#~char",
            customproperty6   = "userConsentDescription~#~char",
            customproperty7   = "adminConsentDisplayName~#~char",
            customproperty8   = "value~#~char"
          }
        },
        SKUServicePlans = {
          colsToPropsMap = {
            entitlementID     = "servicePlanId~#~char",
            entitlement_value = "servicePlanName~#~char",
            customproperty1   = "provisioningStatus~#~char",
            customproperty2   = "appliesTo~#~char",
            customproperty4   = "servicePlanId~#~char"
          }
        }
      }
    }
  )

  create_account_json = jsonencode({
    accountIdPath = "call1.message.id",
    dateFormat    = "yyyy-MM-dd'T'HH:mm:ssXXX",
    responseColsToPropsMap = {
      comments    = "call1.message.displayName~#~char",
      displayName = "call1.message.displayName~#~char",
      name        = "call1.message.userPrincipalName~#~char"
    },
    call = [
      {
        name       = "call1",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users",
        httpMethod = "POST",
        httpParams = "{\"accountEnabled\":\"true\",\"displayName\":\"$${user.displayname}\",\"passwordProfile\":\r\n{\"password\":\"Passw0rd\",\"forceChangePasswordNextSignIn\":\"true\"},\"UsageLocation\":\"US\",\"userPrincipalName\":\"$${user.email}\",\"mailNickname\":\"$${user.firstname}\",\"givenName\":\"$${user.firstname}\",\"surname\":\"$${user.lastname}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  update_account_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
        httpMethod = "PATCH",
        httpParams = "{\"userprincipalname\": \"$${user.email}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  enable_account_json = jsonencode({
    call = [{
      name       = "call1",
      connection = "userAuth",
      url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
      httpMethod = "PATCH",
      httpParams = "{\"accountEnabled\": true}",
      httpHeaders = {
        Authorization = "$${access_token}"
      },
      httpContentType = "application/json",
      successResponses = {
        statusCode = [200, 201, 204, 205]
      }
    }]
  })

  disable_account_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
        httpMethod = "PATCH",
        httpParams = "{\"accountEnabled\": false}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  add_access_json = jsonencode({
    "call" : [
      {
        name       = "SKU",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}/assignLicense",
        httpMethod = "POST",
        httpParams = "{\"addLicenses\": [{\"disabledPlans\": [],\"skuId\": \"$${entitlementValue.entitlementID}\"}],\"removeLicenses\": []}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      },
      {
        name       = "DirectoryRole",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/directoryRoles/$${entitlementValue.entitlementID}/members/\\$ref",
        httpMethod = "POST",
        httpParams = "{\"@odata.id\":\"https://graph.microsoft.com/v1.0/directoryObjects/$${account.accountID}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        },
        "unsuccessResponses" : {
          "odata~dot#error.code" : [
            "Request_BadRequest",
            "Authentication_MissingOrMalformed",
            "Request_ResourceNotFound",
            "Authorization_RequestDenied",
            "Authentication_Unauthorized"
          ]
        }
      },
      {
        name       = "AADGroup",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/groups/$${entitlementValue.entitlementID}/members/\\$ref",
        httpMethod = "POST",
        httpParams = "{\"@odata.id\":\"https://graph.microsoft.com/v1.0/directoryObjects/$${account.accountID}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      },
      {
        name       = "ApplicationInstance",
        connection = "entAuth",
        url        = "https://graph.windows.net/myorganization/users/$${account.accountID}/appRoleAssignedTo?api-version=1.6",
        httpMethod = "POST",
        httpParams = "{\"principalId\": \"$${account.accountID}\", \"id\": \"$${}\", \"resourceId\": \"$${entitlementValue.entitlementID}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      },
      {
        name       = "Team",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/groups/$${entitlementValue.entitlementID}/members/\\$ref",
        httpMethod = "POST",
        httpParams = "{\"@odata.id\":\"https://graph.microsoft.com/v1.0/directoryObjects/$$ {account.accountID}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })


  remove_access_json = jsonencode({
    call = [
      {
        name       = "SKU",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}/assignLicense",
        httpMethod = "POST",
        httpParams = "{\"addLicenses\": [],\"removeLicenses\": [\"$${entitlementValue.entitlementID}\"]}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      },
      {
        name       = "DirectoryRole",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/directoryRoles/$${entitlementValue.entitlementID}/members/$${account.accountID}/\\$ref",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      },
      {
        name       = "AADGroup",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/groups/$${entitlementValue.entitlementID}/members/$${account.accountID}/\\$ref",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      },
      {
        name       = "ApplicationInstance",
        connection = "entAuth",
        url        = "https://graph.windows.net/myOrganization/servicePrincipals/$${entitlementValue.entitlementID}/appRoleAssignedTo?api-version=1.6&\\$top=999",
        httpMethod = "GET",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201
          ]
        }
      },
      {
        name       = "ApplicationInstance",
        connection = "entAuth",
        url        = "https://graph.windows.net/myOrganization/servicePrincipals/$${entitlementValue.entitlementID}/appRoleAssignedTo/$${for (Map map : response.ApplicationInstance1.message.value){if (map.principalId.toString().equals(account.accountID)){return map.objectId;}}}?api-version=1.6",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  update_user_json = jsonencode({
    actions = {
      "Disable User" = {
        call = [
          {
            callOrder       = 0,
            connection      = "$${connectionName}",
            httpContentType = "application/json",
            httpHeaders = {
              Authorization = "$${access_token}"
            },
            httpMethod = "PATCH",
            httpParams = "{\"accountEnabled\": false}",
            name       = "Disable User",
            successResponses = {
              statusCode = [
                200,
                204
              ]
            },
            url = "https://graph.microsoft.com/v1.0/users/$${user.customproperty4}"
          }
        ]
      },
      "Enable User" = {
        call = [
          {
            callOrder       = 0,
            connection      = "$${connectionName}",
            httpContentType = "application/json",
            httpHeaders = {
              Authorization = "$${access_token}"
            },
            httpMethod = "PATCH",
            httpParams = "{\"accountEnabled\": true}",
            name       = "Enable User",
            successResponses = {
              statusCode = [
                200,
                204
              ]
            },
            url = "https://graph.microsoft.com/v1.0/users/$${user.customproperty4}"
          }
        ]
      },
      "Update Login" = {
        call = [
          {
            callOrder       = 0,
            connection      = "$${connectionName}",
            httpContentType = "application/json",
            httpHeaders = {
              Authorization = "$${access_token}"
            },
            httpMethod = "PATCH",
            httpParams = "{\"mobilePhone\": \"$${user.phonenumber}\"}",
            name       = "Update Login",
            successResponses = {
              statusCode = [
                200,
                204
              ]
            },
            url = "https://graph.microsoft.com/v1.0/users/$${user.customproperty4}"
          }
        ]
      }
    }
  })

  change_pass_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "$${connectionName}",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
        httpMethod = "PATCH",
        httpParams = "{\"passwordPolicies\" :\"DisableStrongPassword\",\"passwordProfile\" : {\"password\":\"$${password}\",\"forceChangePasswordNextSignIn\": false}}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  remove_account_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  connection_json = jsonencode({
    authentications = {
      userAuth = {
        authType   = "oauth2",
        url        = "https://login.microsoftonline.com/<tenantid>/oauth2/token",
        httpMethod = "POST",
        httpParams = {
          grant_type    = "client_credentials",
          client_secret = "<client_secret>",
          client_id     = "<client_id>",
          resource      = "https://graph.microsoft.com/"
        },
        httpHeaders = {
          contentType = "application/x-www-form-urlencoded"
        },
        httpContentType = "application/x-www-form-urlencoded",
        expiryError     = "ExpiredAuthenticationToken",
        authError = [
          "InvalidAuthenticationToken"
        ],
        retryFailureStatusCode = [
          401
        ],
        timeOutError       = "Read timed out",
        errorPath          = "error.code",
        maxRefreshTryCount = 5,
        tokenResponsePath  = "access_token",
        tokenType          = "Bearer",
        accessToken        = "Bearer abcd"
      },
      entAuth = {
        authType   = "oauth2",
        url        = "https://login.microsoftonline.com/<tenantid>/oauth2/token",
        httpMethod = "POST",
        httpParams = {
          grant_type    = "client_credentials",
          client_secret = "<client_secret>",
          client_id     = "<client_id>",
          resource      = "https://graph.windows.net/"
        },
        httpHeaders = {
          contentType = "application/x-www-form-urlencoded"
        },
        httpContentType = "application/x-www-form-urlencoded",
        expiryError     = "ExpiredAuthenticationToken",
        authError = [
          "InvalidAuthenticationToken",
          "Authentication_MissingOrMalformed"
        ],
        retryFailureStatusCode = [
          401
        ],
        timeOutError       = "Read timed out",
        errorPath          = "odata~dot#error.code",
        maxRefreshTryCount = 3,
        tokenResponsePath  = "access_token",
        tokenType          = "Bearer",
        accessToken        = "Bearer abcde"
      }
    }
  })

  create_group_json = jsonencode({
    connection = "userAuth",
    url        = "https://graph.microsoft.com/v1.0/groups",
    httpMethod = "Post",
    httpParams = "{\"description\": \"$${roles.description==null || roles.description==''? roles.displayname : roles.description}\", \"displayName\": \"$${roles.displayname==null || roles.displayname==''? roles.role_name : roles.displayname}\", \"groupTypes\": [\"$${roles.customproperty21=='Office365'? 'Unified' : ''}\"], \"mailEnabled\": \"$${roles.customproperty22 == '1' ? true : false}\", \"mailNickname\": \"$${roles.displayname==null || roles.displayname==''? roles.role_name : roles.displayname}\", \"securityEnabled\": \"$${roles.customproperty23 == '1' ? true : false}\",\"owners@odata.bind\": [\"$${allOwner}\"]}",
    httpHeaders = {
      Authorization = "$${access_token}",
      Content-Type  = "application/json"
    },
    httpContentType = "application/json"
  })

  update_group_json = jsonencode({
    connection = "userAuth",
    url        = "https://graph.microsoft.com/v1.0/groups/$${entitlementValue.entitlementID}",
    httpMethod = "PATCH",
    httpParams = "{\"description\": \"$${roles.description==null || roles.description==''? roles.displayname : roles.description}\", \"displayName\": \"$${roles.displayname==null || roles.displayname==''? roles.role_name : roles.displayname}\", \"groupTypes\": [\"$${roles.customproperty21=='Office365'? 'Unified' : ''}\"], \"mailEnabled\": \"$${roles.customproperty22 == '1' ? true : false}\", \"mailNickname\": \"$${roles.displayname==null || roles.displayname==''? roles.role_name : roles.displayname}\", \"securityEnabled\": \"$${roles.customproperty23 == '1' ? true : false}\",\"owners@odata.bind\": [\"$${allOwner}\"]}",
    httpHeaders = {
      Authorization = "$${access_token}",
      Content-Type  = "application/json"
    },
    httpContentType = "application/json"
  })

  add_access_to_entitlement_json = jsonencode({
    connection = "AzureADGroupProvisioning",
    url        = "https://graph.microsoft.com/v1.0/$${parentGroupType}/$${parentEntitlementValuesObj.entitlementID}/members/\\$ref",
    httpMethod = "POST",
    httpParams = "{\"@odata.id\": \"https://graph.microsoft.com/v1.0/groups/$${childEntitlementValuesObj.entitlementID}\"}",
    httpHeaders = {
      Authorization = "$${access_token}",
      Content-Type  = "application/json"
    },
    httpContentType = "application/json"
  })

  remove_access_from_entitlement_json = jsonencode({
    connection = "AzureADGroupProvisioning",
    url        = "https://graph.microsoft.com/v1.0/$${parentGroupType}/$${parentEntitlementValuesObj.entitlementID}/members/$${childEntitlementValuesObj.entitlementID}/\\$ref",
    httpMethod = "DELETE",
    httpParams = "",
    httpHeaders = {
      Authorization = "$${access_token}",
      Content-Type  = "application/json"
    },
    httpContentType = "application/json"
  })

  delete_group_json = jsonencode({
    connection = "userAuth",
    url        = "https://graph.microsoft.com/v1.0/groups/$${entitlementValue.entitlementID}",
    httpMethod = "DELETE",
    httpParams = "",
    httpHeaders = {
      Authorization = "$${access_token}",
      Content-Type  = "application/json"
    },
    httpContentType = "application/json"
  })

  create_service_principal_json = jsonencode({
    "accountIdPath" : "ServicePrincipal.message.id",
    "dateFormat" : "yyyy-MM-dd'T'HH:mm:ssXXX",
    "responseColsToPropsMap" : {
      "displayName" : "CreateApplication.message.displayName~#~char",
      "customproperty2" : "CreateApplication.message.appId~#~char"
    },
    "call" : [
      {
        name       = "CreateApplication",
        callOrder  = 0,
        connection = "$${connectionName}",
        url        = "https://graph.microsoft.com/v1.0/applications",
        httpMethod = "POST",
        httpParams = "{\"displayName\":\"$${accountName}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [200, 201]
        }
      },
      {
        name       = "ServicePrincipal",
        callOrder  = 1,
        connection = "$${connectionName}",
        url        = "https://graph.microsoft.com/v1.0/servicePrincipals",
        httpMethod = "POST",
        httpParams = "{\"appId\":\"$${response.CreateApplication.message.appId}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [200, 201]
        }
      }
    ]
    }
  )

  update_service_principal_json = jsonencode({
    call = [
      {
        name       = "UpdateAccount",
        connection = "$${connectionName}",
        url        = "https://graph.microsoft.com/v1.0/servicePrincipals/$${account.accountID}",
        httpMethod = "PATCH",
        httpParams = "{\"appRoleAssignmentRequired\": \"$${requestAccessAttributes.AssignmentRequired}\"}",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            204
          ]
        }
      }
    ]
  })

  remove_service_principal_json = jsonencode({
    call = [
      {
        name       = "DeleteSPN",
        callOrder  = "0",
        connection = "$${connectionName}",
        url        = "https://graph.microsoft.com/v1.0/servicePrincipals/$${account.accountID}",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [204]
        }
      },
      {
        name       = "GetAppId",
        callOrder  = "1",
        connection = "$${connectionName}",
        url        = "https://graph.microsoft.com/v1.0/applications?%24filter=appId%20eq%20%27$${account.customproperty2}%27",
        httpMethod = "GET",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [204, 200]
        }
      },
      {
        name       = "DeleteApp",
        callOrder  = "2",
        connection = "$${connectionName}",
        url        = "https://graph.microsoft.com/v1.0/applications/$${response.GetAppId.message.value[0].id}",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [204]
        }
      }
    ]
  })

  entitlement_filter_json = jsonencode({
    filter = "(displayName startsWith '$${filterValue}')"
  })

  create_team_json = jsonencode({
    connection = "userAuth",
    url        = "https://graph.microsoft.com/v1.0/groups/$${groupId}/team/",
    httpMethod = "PUT",
    httpParams = "{\"memberSettings\":{\"allowCreateUpdateChannels\": false,\"allowDeleteChannels\": false,\"allowAddRemoveApps\": true,\"allowCreateUpdateRemoveTabs\": true,\"allowCreateUpdateRemoveConnectors\": true},\"guestSettings\":{\"allowCreateUpdateChannels\": false,\"allowDeleteChannels\": true},\"messagingSettings\":{\"allowUserEditMessages\": true,\"allowUserDeleteMessages\": true,\"allowOwnerDeleteMessages\": true,\"allowTeamMentions\": true,\"allowChannelMentions\": true},\"funSettings\":{\"allowGiphy\": true,\"giphyContentRating\":\"strict\",\"allowStickersAndMemes\":true,\"allowCustomMemes\":true}}",
    httpHeaders = {
      Authorization = "$${access_token}",
      Content-Type  = "application/json"
    },
    httpContentType = "application/json"
  })

  create_channel_json = jsonencode({
    connection = "userAuth",
    url        = "https://graph.microsoft.com/v1.0/teams/$${groupId}/channels",
    httpMethod = "POST",
    httpParams = "{\"description\": \"$${rolesObj.customproperty27}\", \"displayName\": \"$${rolesObj.customproperty26}\"}",
    httpHeaders = {
      Authorization = "$${access_token}",
      Content-Type  = "application/json"
    },
    httpContentType = "application/json"
  })

  status_threshold_config = jsonencode({
    statusAndThresholdConfig = {
      accountThresholdValue     = 1000,
      appAccountThresholdValue  = 50,
      correlateInactiveAccounts = true,
      statusColumn              = "customproperty10",
      activeStatus              = ["true"],
      deleteLinks               = true
    }
  })

  endpoints_filter = jsonencode({
    APPLICATION_DEV = [
      {
        AADGROUP = [
          "GROUP_IN_ENGG",
          "GROUP_IN_FINANCE",
          "GROUP_IN_MARKETTING"
        ]
      }
    ]
  })
  accounts_filter = "(userType%20eq%20%27Member%27%20and%20(employeeType%20eq%20%27Employee%27%20or%20employeeType%20eq%20%27External%27%20or%20employeeType%20eq%20%27AdminAccount%27%20or%20employeeType%20eq%20%27Frontline%27)"
  config_json = jsonencode({
    connectionTimeoutConfig = {
      connectionTimeout = 10,
      readTimeout       = 60,
      writeTimeout      = 60,
      retryWait         = 5,
      retryCount        = 3
    }
  })

  modify_user_data_json = jsonencode({
    ADDITIONALTABLES = {
      USERS = "SELECT USERKEY,customproperty5,location,startdate,statuskey,employeeclass,customproperty7,departmentNumber,customproperty10,customproperty14,customproperty16,endDate,termDate,customproperty13,customproperty26,FIRSTNAME,LASTNAME,employeetype FROM USERS"
    },
    COMPUTEDCOLUMNS = [
      "customproperty2",
      "customproperty5",
      "siteid",
      "customproperty8",
      "statuskey",
      "customproperty30",
      "employeetype",
      "departmentNumber",
      "customproperty10",
      "customproperty14",
      "customproperty16",
      "termDate",
      "customproperty13",
      "customproperty45",
      "customproperty34"
    ],
    PREPROCESSQUERIES = [
      "UPDATE NEWUSERDATA SET customproperty2 = 'Transport'",
      "UPDATE NEWUSERDATA SET customproperty34 = 'AT'",
      "UPDATE NEWUSERDATA SET customproperty5 =FROM_UNIXTIME((CAST('0'+SUBSTRING(customproperty5 , 7, 13) AS UNSIGNED))/1000,'%Y-%m-%d')",
      "UPDATE NEWUSERDATA SET siteid = CASE WHEN customproperty7 = 'true' THEN customproperty23 WHEN customproperty7 = 'false' THEN siteid END",
      "UPDATE NEWUSERDATA SET customproperty8 = CASE WHEN customproperty7 = 'true' THEN customproperty25 WHEN customproperty7 = 'false' THEN SUBSTRING(SUBSTRING_INDEX(SUBSTRING_INDEX(location , '(', -1),')',1), 1, 2) END",
      "UPDATE NEWUSERDATA SET statuskey = CASE WHEN (FROM_UNIXTIME((CAST(SUBSTRING(startdate, locate('(' , startdate)+1,(locate('\\)',startdate)-locate('(' , startdate)-1)) AS UNSIGNED))/1000) between sysdate() and DATE_ADD(NOW(),INTERVAL 20 DAY)) THEN 't' ELSE statuskey END",
      "UPDATE NEWUSERDATA SET customproperty30 = CASE WHEN customproperty8 in ('AD','AE','SW','AT','AZ','BE','CH','CZ','DE','DK','DZ','EG','ES','FI','FR','GB','GR','HR','HU','IE','IL','IT','KZ','KA','LV','MA','MY','NG','NL','NO','PL','PT','PF','QA','RO','RU','SA','SE','SN','TN','TR','UZ','ZA','ZW','ZM','YE','EH','UG','TM','TG','TZ','TJ','SY','SZ','SD','SO','SL','SC','ST','RW','OM','NE','NA','MZ','MU','MR','ML','MW','MG','LY','LR','LS','LB','KG','KW','KE','IQ','GW','GN','GH','GE','GA','ET','ER','GQ','DJ','CI','CD','CG','KM','TD','CF','CV','CM','BI','BF','BW','BJ','BH','AM','AO','VG','SJ','GS','SI','SK','RS','SM','SH','AN','MS','ME','MD','MT','MK','LU','LT','LI','XK','JE','IS','VA','GL','GI','GM','FO','FK','EE','CY','KY','IO','BV','BA','BY','AW','AI','AL','BG','UA','PM','RE','NC','MC','YT','MQ','GP','TF','GF','AX','GG','SS') THEN 'EMEA' WHEN customproperty8 in ('AU','CN','HK','IN','ID','JO','JP','MM','PH','SG','KR','TW','TH','VN','WF','VU','TV','TC','TO','TK','TL','LK','SB','WS','PN','PG','PW','PK','MP','NF','NU','NZ','NP','NR','MN','FM','MH','MV','MO','LA','KI','HM','GU','FJ','CK','CC','CX','KH','BN','BT','BD','AQ','AS','AF','KP') THEN 'APAC' WHEN customproperty8 in ('AR','BR','CA','CL','CO','EC','IR','MX','PA','PE','SDQ','US','VE','UY','TT','SR','VC','LC','KN','PY','NI','JM','HN','HT','GY','GT','GD','SV','DM','CU','CR','BO','BM','BZ','BB','BS','AG','DO','VI','UM','PR','BL','BQ','CW','MF','SX') THEN 'AMER' END",
      "UPDATE NEWUSERDATA SET employeetype = CASE WHEN employeeClass in ('Apprentice','Employee','Trainee','VIE France','Hired Staff') THEN 'Internal' WHEN customproperty7 ='true' then 'External' END",
      "UPDATE NEWUSERDATA SET departmentNumber = CASE WHEN customproperty7 = 'true' THEN customproperty24 WHEN customproperty7 = 'false' THEN departmentNumber END",
      "UPDATE NEWUSERDATA SET customproperty14 = CASE WHEN UPPER(customproperty14) = 'abc' THEN '@abcgroup.com' WHEN UPPER(customproperty14) = 'abc1' THEN '@abc1.com' END",
      "UPDATE NEWUSERDATA SET customproperty16 = LOWER(customproperty16)",
      "UPDATE NEWUSERDATA SET termDate = CASE WHEN customproperty7 = 'true' THEN enddate WHEN customproperty7 = 'false' THEN termDate END",
      "UPDATE NEWUSERDATA SET customproperty13 = CASE WHEN customproperty7 = 'true' THEN customproperty13 WHEN customproperty7 = 'false' THEN customproperty26 END",
      "UPDATE NEWUSERDATA SET departmentName = SUBSTR(departmentName,1,CHAR_LENGTH(departmentName) - LOCATE('(', REVERSE(departmentName)))"
    ]
  })
  windows_connector_json = jsonencode({
    http = {
      url = "http://<domain-name>/FIMAzure/PS/ExecutePSCommand",
      httpHeaders = {
        Content-Type  = "application/json",
        Authorization = "DFLzDrU2qd71FfckmDW7AqDE2qh9XX6wpo8LdowRNowg8/sIpcfH+j+7mKzihKin",
        Username      = "G55L8KBHMjed9ahyXvByXsOpXbhm+CWaF02UCYSHdAqZm48imlKLTyDZ4yTHHKIg",
        Password      = "<encrypted password>",
        Command       = "CWHoJVaZjbi2IM0xq6vMiUUecMEPmZfyZda2Tqk2SH1aVROmE29pgxiP8/mNrK/tkctEh7ZMkC6C3HmiMaqSdAk1erZNv4LcNxAwF2Sun1FfvTYQ7+AOxSwK8ZjGRaU6vU/PfYtHVBVsAK7UZ2iQ3JCqkPWvuwp9SjWv6hxMiYXyrXRK9eif52pq+ebFh/kLqk0/mQJEQn/UGfiphEFmkqzpxho0SGtByMnJFcYKYRmyXShA8gzh0uDkcCbx3rluk2QQaryDDBjYM0CLL3eC9ApUClXUc/f3O5kADO85tbV8pLln1NazBt5bX/BY3sUYOBBN9jaSz8n/Jp4sa3vvKLAnqoc3r1ExIF71G2F6tnoVekT/6j/+ewVxokP/6pMmIW6lBU6m9iTTpTjub7DD8MDOoGGcq7wX62FIUCrstW8gsSS4TyTHpg4wcNAYNJrjNowKxZBybTt+pCCAZ69kkj/g8S98hlPKJne4e43BzEsW7WkG/mUwX+9iY5+POUffxPscGAWoDht8Smd8W0mD3tamfpMLlBW19dst76GOgu/Y/7L0HSlSiRRL9w4gPoFB7ECWwp5hX1HhQbeE0xe2hCI45uuSqU3kuPU8A8LLUdRtgCHLhht/n5rqhnh+U2ScPd2ruWap7n2s8z8aHX85BY2MIY52GqroULQssPM6al3XuIi1rE2KvKmkUG0Pq8MH5Ar2bvDwj4SZzHjH0IQR5ZCBXT9/lQSMR30r5pVBQ3U2KRyk1Qgn6WEgFRslpbWFI367vsbdTUMIbpwedvdG1n8AKpgY6FA1G2Yf2Y4kiRm/k+1AXZLr1h/zNB4tpNTdZLtLuIoNT+M7WeqI/9nPpr4EhioDpp/WZgd/SEtJn0hwUXTTgazeXndt6mzW9BV1aG9/2NaslA8c1Xyt8KK/n6Zzp6sMHnOFvWjM/104/9UfzrXN3zODBwRbBPeSwd2+Mdied04hMrNB4/M2/CZgmy9by3K7ygmc7pE23O8selXTQq+UnazopgQ6TQe4YDhKqcnf5tom8p33v45mVKGzTGmwlqh2GMtGYb5/zc3JwMzU5PS9CKKg5duaxopii/l1++alBgUKP5PB1R0Rc4Dcvk1xJVRUb4ogBAtqBoiZFvDalEgadb7JV8xz4/qO4x6UKc7aPTiemWj9RLf0yZpx6LAC+PQqFFt2f6XYNyNHfW2FAcBjx2wUIhs2mOne00mcs+SraNb8xbce5BbWTDokHQ=="
      },
      httpContentType = "application/json",
      httpMethod      = "GET"
    },
    listField = "value",
    keyField  = "accountID",
    colsToPropsMap = {
      accountID        = "ObjectId~#~char",
      customproperty17 = "WhenCreated~#~char",
      customproperty15 = "MFAState~#~char",
      customproperty16 = "MFADateTime~#~char"
    }
  })
  service_account_attributes = jsonencode({
    colsToPropsMap = {
      accountID        = "id~#~char",
      name             = "displayName~#~char",
      displayName      = "displayName~#~char",
      customproperty1  = "servicePrincipalNames~#~char",
      customproperty2  = "appId~#~char",
      status           = "accountEnabled~#~char",
      customproperty10 = "accountEnabled~#~char",
      customproperty3  = "appOwnerOrganizationId~#~char",
      customproperty4  = "appDescription~#~char",
      customproperty5  = "appDisplayName~#~char",
      customproperty6  = "accountEnabled~#~char",
      customproperty7  = "homepage~#~char",
      accountType      = "#CONST#Service Principal~#~char",
      accountClass     = "servicePrincipalType~#~char"
    }
  })
}
