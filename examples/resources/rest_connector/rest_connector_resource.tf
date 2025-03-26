terraform {
  required_providers {
    saviynt = {
      source  = "registry.terraform.io/local/saviynt"
      version = "1.0.0"
    }
  }
}
provider "saviynt" {
  server_url      = var.SAVIYNT_SERVER_URL
  username = var.SAVIYNT_USERNAME
  password = var.SAVIYNT_PASSWORD
}

resource "saviynt_rest_connection_resource" "example" {
  connection_type = "REST"
  connection_name = "Merserk_Release_3490_Rajiv_1"
  remove_account_json=jsonencode({
  "call": [
    {
      "name": "call1",
      "connection": "EntraIDAuth",
      "url": "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
      "httpMethod": "DELETE",
      "httpHeaders": {
        "Authorization": "$${access_token}"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201,
          204,
          205
        ]
      }
    }
  ]
  })
  create_account_json=jsonencode({
  "accountIdPath": "call1.message.id",
  "dateFormat": "yyyy-MM-dd'T'HH:mm:ssXXX",
  "responseColsToPropsMap": {
    "displayName": "call1.message.displayName~#~char",
    "name": "call1.message.userPrincipalName~#~char"
  },
  "call": [
    {
      "name": "call1",
      "connection": "EntraIDAuth",
      "url": "https://graph.microsoft.com/v1.0/users",
      "httpMethod": "POST",
      "httpParams": "{\"accountEnabled\":true,\"displayName\":\"$${user.firstname} $${user.lastname}\",\"mailNickname\":\"$${user.firstname}\",\"userPrincipalName\":\"$${user.firstname}.$${user.lastname}@saviyntlivedev.onmicrosoft.com\",\"passwordProfile\":{\"forceChangePasswordNextSignIn\":true,\"password\":\"$${password}\"}}",
      "httpHeaders": {
        "Authorization": "$${access_token}"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
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
    "authentications": {
    "EntraIDAuth": {
      "authType": "oauth2",
      "url": "https://login.microsoftonline.com/18fe0463-2c69-49f9-af8e-fb04fe901808/oauth2/token",
      "httpMethod": "POST",
      "httpParams": {
        "grant_type": "client_credentials",
        "client_secret": "WqZ8Q~fhxq1wvDGod5VmfI73jWb24caX3Lz4naP2",
        "client_id": "8555b9ed-11f1-4c07-bfda-4af9373a8414",
        "resource": "https://graph.microsoft.com"
      },
      "httpHeaders": {
        "contentType": "application/x-www-form-urlencoded"
      },
      "httpContentType": "application/x-www-form-urlencoded",
      "errorPath": "errors.type",
      "retryFailureStatusCode": [
        401
      ],
      "maxRefreshTryCount": 5,
      "tokenResponsePath": "access_token",
      "tokenType": "Bearer",
      "accessToken": "Bearer eyJ0eXAiOiJKV1QiLCJub25jZSI6Ik5yZGtRajNQVGpQWFkyX3BKY1pjTGxRU1RxdnJjdVFnZS1XMk9NMERDcm8iLCJhbGciOiJSUzI1NiIsIng1dCI6IkpETmFfNGk0cjdGZ2lnTDNzSElsSTN4Vi1JVSIsImtpZCI6IkpETmFfNGk0cjdGZ2lnTDNzSElsSTN4Vi1JVSJ9.eyJhdWQiOiJodHRwczovL2dyYXBoLm1pY3Jvc29mdC5jb20iLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC8xOGZlMDQ2My0yYzY5LTQ5ZjktYWY4ZS1mYjA0ZmU5MDE4MDgvIiwiaWF0IjoxNzQyOTY5MDc5LCJuYmYiOjE3NDI5NjkwNzksImV4cCI6MTc0Mjk3Mjk3OSwiYWlvIjoiazJSZ1lPRCtZUnBadUtkQVZVZmc3SloySjlHOUFBPT0iLCJhcHBfZGlzcGxheW5hbWUiOiJTYXZpeW50SUdBIiwiYXBwaWQiOiI4NTU1YjllZC0xMWYxLTRjMDctYmZkYS00YWY5MzczYTg0MTQiLCJhcHBpZGFjciI6IjEiLCJpZHAiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC8xOGZlMDQ2My0yYzY5LTQ5ZjktYWY4ZS1mYjA0ZmU5MDE4MDgvIiwiaWR0eXAiOiJhcHAiLCJvaWQiOiIxOTJlNjUwZi1kOGZlLTQ0ZjAtYWM2MC0wODkwY2MxOGJhMjEiLCJyaCI6IjEuQVc4Qll3VC1HR2tzLVVtdmp2c0VfcEFZQ0FNQUFBQUFBQUFBd0FBQUFBQUFBQUJ2QVFCdkFRLiIsInJvbGVzIjpbIlRlYW1TZXR0aW5ncy5SZWFkV3JpdGUuQWxsIiwiVGVhbU1lbWJlci5SZWFkLkFsbCIsIlVzZXIuUmVhZEJhc2ljLkFsbCIsIlJvbGVNYW5hZ2VtZW50LlJlYWRXcml0ZS5EZWZlbmRlciIsIkRpcmVjdG9yeVJlY29tbWVuZGF0aW9ucy5SZWFkV3JpdGUuQWxsIiwiZURpc2NvdmVyeS5SZWFkLkFsbCIsIkdyb3VwLUNvbnZlcnNhdGlvbi5SZWFkV3JpdGUuQWxsIiwiU2VydmljZVByaW5jaXBhbEVuZHBvaW50LlJlYWQuQWxsIiwiVGVhbXNUYWIuUmVhZC5BbGwiLCJNYWlsYm94Rm9sZGVyLlJlYWQuQWxsIiwiUm9sZU1hbmFnZW1lbnQuUmVhZFdyaXRlLkV4Y2hhbmdlIiwiVXNlckF1dGhlbnRpY2F0aW9uTWV0aG9kLlJlYWQuQWxsIiwiVGVhbXNBY3Rpdml0eS5SZWFkLkFsbCIsIlRlYW1NZW1iZXIuUmVhZFdyaXRlTm9uT3duZXJSb2xlLkFsbCIsIk1haWwuUmVhZFdyaXRlIiwiVGVhbXNUYWIuUmVhZFdyaXRlRm9yQ2hhdC5BbGwiLCJVc2VyLlJldm9rZVNlc3Npb25zLkFsbCIsIk9yZ0NvbnRhY3QuUmVhZC5BbGwiLCJVc2VyLU1haWwuUmVhZFdyaXRlLkFsbCIsIkdyb3VwLUNvbnZlcnNhdGlvbi5SZWFkLkFsbCIsIkNoYW5uZWwuRGVsZXRlLkFsbCIsIlVzZXIuUmVhZFdyaXRlLkFsbCIsIkRvbWFpbi5SZWFkV3JpdGUuQWxsIiwiQXBwbGljYXRpb24uUmVhZFdyaXRlLk93bmVkQnkiLCJVc2VyLkRlbGV0ZVJlc3RvcmUuQWxsIiwiTWFpbGJveEl0ZW0uUmVhZC5BbGwiLCJDaGFubmVsU2V0dGluZ3MuUmVhZC5BbGwiLCJUZWFtc1RhYi5DcmVhdGUiLCJVc2VyQXV0aGVudGljYXRpb25NZXRob2QuUmVhZFdyaXRlLkFsbCIsIkRlbGVnYXRlZFBlcm1pc3Npb25HcmFudC5SZWFkV3JpdGUuQWxsIiwiTWFpbC5SZWFkQmFzaWMuQWxsIiwiRGVsZWdhdGVkUGVybWlzc2lvbkdyYW50LlJlYWQuQWxsIiwiUm9sZU1hbmFnZW1lbnQuUmVhZC5EaXJlY3RvcnkiLCJDaGFubmVsLlJlYWRCYXNpYy5BbGwiLCJBcHBsaWNhdGlvbi5SZWFkV3JpdGUuQWxsIiwiR3JvdXAuUmVhZC5BbGwiLCJUZWFtc1RhYi5SZWFkV3JpdGVGb3JVc2VyLkFsbCIsIkRpcmVjdG9yeS5SZWFkV3JpdGUuQWxsIiwiVGVhbXNBY3Rpdml0eS5TZW5kIiwiUm9sZU1hbmFnZW1lbnQuUmVhZC5FeGNoYW5nZSIsIlRlYW1zVGFiLlJlYWRXcml0ZS5BbGwiLCJNYWlsYm94U2V0dGluZ3MuUmVhZCIsIkNvbnRhY3RzLlJlYWRXcml0ZSIsIlRlYW0uQ3JlYXRlIiwiR3JvdXAuQ3JlYXRlIiwiR3JvdXAuUmVhZFdyaXRlLkFsbCIsIlVzZXIuRW5hYmxlRGlzYWJsZUFjY291bnQuQWxsIiwiRGlyZWN0b3J5UmVjb21tZW5kYXRpb25zLlJlYWQuQWxsIiwiRW50aXRsZW1lbnRNYW5hZ2VtZW50LlJlYWQuQWxsIiwiVXNlci5JbnZpdGUuQWxsIiwiVGVhbXNUYWIuUmVhZFdyaXRlU2VsZkZvclRlYW0uQWxsIiwiRGlyZWN0b3J5LlJlYWQuQWxsIiwiQ29uc2VudFJlcXVlc3QuUmVhZC5BbGwiLCJSb2xlTWFuYWdlbWVudC5SZWFkLkFsbCIsIk1haWxib3hJdGVtLkltcG9ydEV4cG9ydC5BbGwiLCJVc2VyLlJlYWQuQWxsIiwiRG9tYWluLlJlYWQuQWxsIiwiQ2hhbm5lbE1lbWJlci5SZWFkLkFsbCIsIlJvbGVNYW5hZ2VtZW50LlJlYWRXcml0ZS5DbG91ZFBDIiwiTWFpbGJveEZvbGRlci5SZWFkV3JpdGUuQWxsIiwiR3JvdXBNZW1iZXIuUmVhZC5BbGwiLCJUZWFtTWVtYmVyLlJlYWRXcml0ZS5BbGwiLCJJZGVudGl0eVByb3ZpZGVyLlJlYWRXcml0ZS5BbGwiLCJSb2xlTWFuYWdlbWVudC5SZWFkLkNsb3VkUEMiLCJTZXJ2aWNlUHJpbmNpcGFsRW5kcG9pbnQuUmVhZFdyaXRlLkFsbCIsIkVudGl0bGVtZW50TWFuYWdlbWVudC5SZWFkV3JpdGUuQWxsIiwiVGVhbXNUYWIuUmVhZFdyaXRlU2VsZkZvckNoYXQuQWxsIiwiVGVhbS5SZWFkQmFzaWMuQWxsIiwiQ29uc2VudFJlcXVlc3QuUmVhZFdyaXRlLkFsbCIsIk1haWwuUmVhZCIsIkNoYW5uZWxNZXNzYWdlLlJlYWQuQWxsIiwiZURpc2NvdmVyeS5SZWFkV3JpdGUuQWxsIiwiVXNlci5FeHBvcnQuQWxsIiwiSWRlbnRpdHlQcm92aWRlci5SZWFkLkFsbCIsIlRlYW1TZXR0aW5ncy5SZWFkLkFsbCIsIk1haWwuU2VuZCIsIlRlYW1zVGFiLlJlYWRXcml0ZVNlbGZGb3JVc2VyLkFsbCIsIlVzZXIuTWFuYWdlSWRlbnRpdGllcy5BbGwiLCJDaGFubmVsTWVzc2FnZS5VcGRhdGVQb2xpY3lWaW9sYXRpb24uQWxsIiwiTWFpbGJveFNldHRpbmdzLlJlYWRXcml0ZSIsIkNoYW5uZWxNZW1iZXIuUmVhZFdyaXRlLkFsbCIsIlJvbGVNYW5hZ2VtZW50LlJlYWRXcml0ZS5EaXJlY3RvcnkiLCJHcm91cE1lbWJlci5SZWFkV3JpdGUuQWxsIiwiQ29udGFjdHMuUmVhZCIsIlJvbGVNYW5hZ2VtZW50LlJlYWQuRGVmZW5kZXIiLCJNYWlsLlJlYWRCYXNpYyIsIkNoYW5uZWxTZXR0aW5ncy5SZWFkV3JpdGUuQWxsIiwiQXVkaXRMb2cuUmVhZC5BbGwiLCJDaGFubmVsLkNyZWF0ZSIsIk1lbWJlci5SZWFkLkhpZGRlbiIsIkFwcGxpY2F0aW9uLlJlYWQuQWxsIiwiVGVhbXNUYWIuUmVhZFdyaXRlRm9yVGVhbS5BbGwiXSwic3ViIjoiMTkyZTY1MGYtZDhmZS00NGYwLWFjNjAtMDg5MGNjMThiYTIxIiwidGVuYW50X3JlZ2lvbl9zY29wZSI6Ik5BIiwidGlkIjoiMThmZTA0NjMtMmM2OS00OWY5LWFmOGUtZmIwNGZlOTAxODA4IiwidXRpIjoiU0VSLVdXZ3Y1VUdWSEFweFM5Nm9BQSIsInZlciI6IjEuMCIsIndpZHMiOlsiMDk5N2ExZDAtMGQxZC00YWNiLWI0MDgtZDVjYTczMTIxZTkwIl0sInhtc19pZHJlbCI6IjIwIDciLCJ4bXNfdGNkdCI6MTcyNTUzODM1NH0.ECcVG7gAWNiOs5-5xhXZQzp2dGZcqtgcgQ-i4Jw0NgClCZMWK_dGci1gKDupC9IdR3lZCMaYY1JvtNP77DkpVgF3tckYvNyxS18yVw_3E4YkWl_X6oGXQBmcGYfbeh6W73_2QNHLPtT7GFTQnnPcdstR_YjJEnZN4Mq2iDYAO_ALlejLsRF_g6gR8L_EZedhsP-HVoCneAjgo9eKBykt41LUVN759SWOZ6mtp-UTx52m0XOo07VIRLlWg03kUoQ4e1CbkmZSJh4yWkzqajHp3tceIoCTcrWs9RmO2CDYi0ij17F-6GU4tUrcwlWaOddPPq53t8U2MjQ9FF4nWDgu3A"
    }
  }
  })
import_account_ent_json = jsonencode({
  "accountParams": {
    "connection": "EntraIDAuth",
    "processingType": "SequentialAndIterative",
    "call": {
      "call1": {
        "callOrder": 0,
        "stageNumber": 0,
        "showJobHistory": true,
        "http": {
          "url": "https://graph.microsoft.com/v1.0/users?%24select=Id,userPrincipalName,accountEnabled,mail,userType,createdDateTime,country,preferredLanguage,displayName,surname,givenName,mobilePhone,businessPhones,mailNickname,mail,state",
          "httpContentType": "application/json",
          "httpMethod": "GET",
          "httpHeaders": {
            "Authorization": "$${access_token}",
            "Accept": "application/json"
          }
        },
        "listField": "value",
        "keyField": "accountID",
        "colsToPropsMap": {
          "accountID": "id~#~char",
          "name": "userPrincipalName~#~char",
          "displayName": "displayName~#~char",
          "customproperty10": "accountEnabled~#~char",
          "customproperty31": "STORE#ACC#ENT#MAPPINGINFO~#~char"
        },
        "pagination": {
          "nextUrl": {
            "nextUrlPath": "@odata.nextLink"
          }
        },
        "disableDeletedAccounts": true
      }
    },
    "successResponses": {
      "statusCode": [
        200,
        201,
        202,
        203,
        204,
        205
      ]
    }
  },
  "entitlementParams": {
    "connection": "EntraIDAuth",
    "processingType": "SequentialAndIterative",
    "unsuccessResponses": null,
    "doNotChangeIfFailed": true,
    "entTypes": {
      "AccessPackages": {
        "entTypeOrder": 0,
        "entTypeLabels": {
          "customproperty1": "PolicyID"
        },
        "call": {
          "call1": {
            "connection": "EntraIDAuth",
            "callOrder": 0,
            "stageNumber": 0,
            "showJobHistory": true,
            "http": {
              "url": "https://graph.microsoft.com/v1.0/identityGovernance/entitlementManagement/accessPackages?%24expand=assignmentPolicies&$select=id,displayName,uniqueName,description,isHidden,createdDateTime,modifiedDateTime",
              "httpContentType": "application/json",
              "httpMethod": "GET",
              "httpHeaders": {
                "Authorization": "$${access_token}",
                "Accept": "application/json"
              }
            },
            "listField": "value",
            "keyField": "entitlementID",
            "colsToPropsMap": {
              "entitlement_value": "displayName~#~char",
              "entitlementID": "id~#~char",
              "displayname": "displayName~#~char",
              "entitlementMappingJson": "STORE#ENT#MAPPINGINFO~#~char"
            }
          }
        },
        "entMappings": {
          "AssignmentPolicy": {
            "listPath": "assignmentPolicies",
            "idPath": "id",
            "idColumn": "entitlementID",
            "mappingTypes": [
              "ENT2"
            ]
          }
        }
      },
      "AssignmentPolicy": {
        "entTypeOrder": 1,
        "entTypeLabels": {
          "customproperty1": "PolicyID"
        },
        "call": {
          "call1": {
            "connection": "EntraIDAuth",
            "callOrder": 0,
            "stageNumber": 0,
            "showJobHistory": true,
            "http": {
              "url": "https://graph.microsoft.com/v1.0/identityGovernance/entitlementManagement/assignmentPolicies",
              "httpContentType": "application/json",
              "httpMethod": "GET",
              "httpHeaders": {
                "Authorization": "$${access_token}",
                "Accept": "application/json"
              }
            },
            "listField": "value",
            "keyField": "entitlementID",
            "colsToPropsMap": {
              "entitlement_value": "displayName~#~char",
              "entitlementID": "id~#~char",
              "displayname": "displayName~#~char"
            },
            "pagination": {
              "nextUrl": {
                "nextUrlPath": "@odata.nextLink"
              }
            },
            "disableDeletedEntitlements": true
          }
        }
      }
    },
    "successResponses": {
      "statusCode": [
        200,
        201,
        202,
        203,
        204,
        205
      ]
    }
  },
  "acctEntParams": {
    "entTypes": {
      "AccessPackages": {
        "call": {
          "call1": {
            "connection": "EntraIDAuth",
            "showJobHistory": true,
            "callOrder": 0,
            "stageNumber": 0,
            "processingType": "http",
            "http": {
              "url": "https://graph.microsoft.com/v1.0/identityGovernance/entitlementManagement/assignments?%24filter=state%20eq%20'Delivered'&%24expand=target,accessPackage",
              "httpContentType": "application/json",
              "httpMethod": "GET",
              "httpHeaders": {
                "Authorization": "$${access_token}",
                "Accept": "application/json"
              }
            },
            "listField": "value",
            "acctIdPath": "target.objectId",
            "acctKeyField": "accountID",
            "entIdPath": "accessPackage.id",
            "entKeyField": "entitlementID",
            "pagination": {
              "nextUrl": {
                "nextUrlPath": "@odata.nextLink"
              }
            }
          }
        }
      }
    },
    "successResponses": {
      "statusCode": [
        200,
        201,
        202,
        203,
        204,
        205
      ]
    },
    "unsuccessResponses": null
  }
})
}

check "instance_health" {
  assert {
   condition = saviynt_rest_connection_resource.example.error_code != "1"
   error_message = "The error is: ${saviynt_rest_connection_resource.example.msg}"
  }
}
