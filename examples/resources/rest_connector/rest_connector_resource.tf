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

resource "saviynt_rest_connection_resource" "example" {
  connection_type = "REST"
  connection_name = "Merserk_Release_3490_Rajiv_2"
  config_json     = <<EOF
  {"showLogs":true}
   EOF
  create_account_json = jsonencode({
    "accountIdPath" : "call1.message.id",
    "dateFormat" : "yyyy-MM-dd'T'HH:mm:ssXXX",
    "responseColsToPropsMap" : {
      "displayName" : "call1.message.displayName~#~char",
      "name" : "call1.message.userPrincipalName~#~char"
    },
    "call" : [
      {
        "name" : "call1",
        "connection" : "EntraIDAuth",
        "url" : "https://graph.microsoft.com",
        "httpMethod" : "POST",
        "httpParams" : "{\"accountEnabled\":true,\"displayName\":\"$${user.firstname} $${user.lastname}\",\"mailNickname\":\"$${user.firstname}\",\"userPrincipalName\":\"$${user.firstname}.$${user.lastname}@saviyntlivedev.onmicrosoft.com\",\"passwordProfile\":{\"forceChangePasswordNextSignIn\":true,\"password\":\"$${password}\"}}",
        "httpHeaders" : {
          "Authorization" : "$${access_token}"
        },
        "httpContentType" : "application/json",
        "successResponses" : {
          "statusCode" : [
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
    "authentications" : {
      "EntraIDAuth" : {
        "authType" : "oauth2",
        "url" : "https://login.microsoftonline.com/oauth2/token",
        "httpMethod" : "POST",
        "httpParams" : {
          "grant_type" : "client_credentials",
          "client_secret" : "XXXXX",
          "client_id" : "XXXX",
          "resource" : "https://graph.microsoft.com"
        },
        "httpHeaders" : {
          "contentType" : "application/x-www-form-urlencoded"
        },
        "httpContentType" : "application/x-www-form-urlencoded",
        "errorPath" : "errors.type",
        "retryFailureStatusCode" : [
          401
        ],
        "maxRefreshTryCount" : 5,
        "tokenResponsePath" : "access_token",
        "tokenType" : "Bearer",
        "accessToken" : "Bearer access_token"
      }
    }
  })
  import_account_ent_json = jsonencode({
    "accountParams" : {
      "connection" : "EntraIDAuth",
      "processingType" : "SequentialAndIterative",
      "call" : {
        "call1" : {
          "callOrder" : 0,
          "stageNumber" : 0,
          "showJobHistory" : true,
          "http" : {
            "url" : "https://graph.microsoft.com/v1.0/",
            "httpContentType" : "application/json",
            "httpMethod" : "GET",
            "httpHeaders" : {
              "Authorization" : "$${access_token}",
              "Accept" : "application/json"
            }
          },
          "listField" : "value",
          "keyField" : "accountID",
          "colsToPropsMap" : {
            "accountID" : "id~#~char",
            "name" : "userPrincipalName~#~char",
            "displayName" : "displayName~#~char",
            "customproperty10" : "accountEnabled~#~char",
            "customproperty31" : "STORE#ACC#ENT#MAPPINGINFO~#~char"
          },
          "pagination" : {
            "nextUrl" : {
              "nextUrlPath" : "@odata.nextLink"
            }
          },
          "disableDeletedAccounts" : true
        }
      },
      "successResponses" : {
        "statusCode" : [
          200,
          201,
          202,
          203,
          204,
          205
        ]
      }
    },
    "entitlementParams" : {
      "connection" : "EntraIDAuth",
      "processingType" : "SequentialAndIterative",
      "unsuccessResponses" : null,
      "doNotChangeIfFailed" : true,
      "entTypes" : {
        "AccessPackages" : {
          "entTypeOrder" : 0,
          "entTypeLabels" : {
            "customproperty1" : "PolicyID"
          },
          "call" : {
            "call1" : {
              "connection" : "EntraIDAuth",
              "callOrder" : 0,
              "stageNumber" : 0,
              "showJobHistory" : true,
              "http" : {
                "url" : "https://graph.microsoft.com/v1.0",
                "httpContentType" : "application/json",
                "httpMethod" : "GET",
                "httpHeaders" : {
                  "Authorization" : "$${access_token}",
                  "Accept" : "application/json"
                }
              },
              "listField" : "value",
              "keyField" : "entitlementID",
              "colsToPropsMap" : {
                "entitlement_value" : "displayName~#~char",
                "entitlementID" : "id~#~char",
                "displayname" : "displayName~#~char",
                "entitlementMappingJson" : "STORE#ENT#MAPPINGINFO~#~char"
              }
            }
          },
          "entMappings" : {
            "AssignmentPolicy" : {
              "listPath" : "assignmentPolicies",
              "idPath" : "id",
              "idColumn" : "entitlementID",
              "mappingTypes" : [
                "ENT2"
              ]
            }
          }
        },
        "AssignmentPolicy" : {
          "entTypeOrder" : 1,
          "entTypeLabels" : {
            "customproperty1" : "PolicyID"
          },
          "call" : {
            "call1" : {
              "connection" : "EntraIDAuth",
              "callOrder" : 0,
              "stageNumber" : 0,
              "showJobHistory" : true,
              "http" : {
                "url" : "https://graph.microsoft.com/",
                "httpContentType" : "application/json",
                "httpMethod" : "GET",
                "httpHeaders" : {
                  "Authorization" : "$${access_token}",
                  "Accept" : "application/json"
                }
              },
              "listField" : "value",
              "keyField" : "entitlementID",
              "colsToPropsMap" : {
                "entitlement_value" : "displayName~#~char",
                "entitlementID" : "id~#~char",
                "displayname" : "displayName~#~char"
              },
              "pagination" : {
                "nextUrl" : {
                  "nextUrlPath" : "@odata.nextLink"
                }
              },
              "disableDeletedEntitlements" : true
            }
          }
        }
      },
      "successResponses" : {
        "statusCode" : [
          200,
          201,
          202,
          203,
          204,
          205
        ]
      }
    },
    "acctEntParams" : {
      "entTypes" : {
        "AccessPackages" : {
          "call" : {
            "call1" : {
              "connection" : "EntraIDAuth",
              "showJobHistory" : true,
              "callOrder" : 0,
              "stageNumber" : 0,
              "processingType" : "http",
              "http" : {
                "url" : "https://graph.microsoft.com/",

                "httpContentType" : "application/json",
                "httpMethod" : "GET",
                "httpHeaders" : {
                  "Authorization" : "$${access_token}",
                  "Accept" : "application/json"
                }
              },
              "listField" : "value",
              "acctIdPath" : "target.objectId",
              "acctKeyField" : "accountID",
              "entIdPath" : "accessPackage.id",
              "entKeyField" : "entitlementID",
              "pagination" : {
                "nextUrl" : {
                  "nextUrlPath" : "@odata.nextLink"
                }
              }
            }
          }
        }
      },
      "successResponses" : {
        "statusCode" : [
          200,
          201,
          202,
          203,
          204,
          205
        ]
      },
      "unsuccessResponses" : null
    }
  })
}

check "instance_health" {
  assert {
    condition     = saviynt_rest_connection_resource.example.error_code != "1"
    error_message = "The error is: ${saviynt_rest_connection_resource.example.msg}"
  }
}
