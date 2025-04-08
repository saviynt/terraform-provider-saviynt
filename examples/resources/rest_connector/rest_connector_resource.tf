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
  connection_name = "shaleen_test_rest_24"
  config_json     = jsonencode({"showLogs":true})
  connection_json = jsonencode(
{
  "authentications": {
    "acctAuth": {
      "authType": "Basic",
      "url": "https://nit2669.zendesk.com/",
      "httpMethod": "POST",
      "httpParams": {},
      "httpHeaders": {},
      "httpContentType": "text/html",
      "properties": {
        "userName":"supriya.c@saviynt.com",
        "password":"Password@123"
      },
      "expiryError": "ExpiredAuthenticationToken",
      "authError": [
        "InvalidAuthenticationToken",
        "AuthenticationFailed"
      ],
      "timeOutError": "Read timed out",
      "errorPath": "error.code",
      "maxRefreshTryCount": 5,
      "tokenResponsePath": "access_token",
      "tokenType": "Basic",
      "accessToken": "Basic abcd",
          "testConnectionParams": {
        "http": {
          "url": "https://nit2669.zendesk.com/api/v2/users.json",
          "httpHeaders": {
            "Authorization": "$${access_token}"
          },
          "httpContentType": "application/json",
          "httpMethod": "GET"
        },
        "successResponse": [],
        "successResponsePath": "",
        "errors": [
          "Couldn't authenticate you"
        ],
        "errorPath": "error"
                
      }
         
    }
  }
})

  create_account_json = jsonencode({
  "accountIdPath": "call1.message.user.id",
  "dateFormat": "yyyy-MM-dd'T'HH:mm:ssXXX",
  "responseColsToPropsMap": {
    "displayname": "call1.message.user.name~#~char"
  },
  "call": [
    {
      "name": "call1",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/users",
      "httpMethod": "POST",
      "httpParams": "{\"user\": {\"name\": \"$${user.firstname} $${user.lastname}\", \"email\": \"$${user.email}\", \"role\":\"agent\"}}",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201
        ]
      }
    }
  ]
})
  import_account_ent_json = jsonencode({
  "accountParams": {
    "connection": "acctAuth",
    "processingType": "SequentialAndIterative",
    "statusAndThresholdConfig": {
      "statusColumn": "customproperty11",
      "activeStatus": [
        "false"
      ],
      "deleteLinks": true,
      "accountThresholdValue": 10,
      "correlateInactiveAccounts": false,
      "inactivateAccountsNotInFile": true,
      "deleteAccEntForActiveAccounts": true
    },
    "call": {
      "call1": {
        "callOrder": 0,
        "stageNumber": 0,
        "http": {
          "url": "https://saviynt6799.zendesk.com/api/v2/users.json?role[]=admin&role[]=agent",
          "httpHeaders": {
            "Authorization": "$${access_token}",
            "Accept": "application/json"
          },
          "httpContentType": "application/json",
          "httpMethod": "GET"
        },
        "listField": "users",
        "keyField": "accountID",
        "statusConfig": {
          "active": "true",
          "inactive": "false"
        },
        "colsToPropsMap": {
          "accountID": "id~#~char",
          "name": "email~#~char",
          "displayName": "name~#~char",
          "customproperty2": "email~#~char",
          "customproperty3": "created_at~#~char",
          "customproperty4": "updated_at~#~char",
          "customproperty5": "role~#~char",
          "status": "active~#~char",
          "customproperty6": "last_login_at~#~char",
          "customproperty7": "custom_role_id~#~char",
          "customproperty8": "default_group_id~#~char",
          "customproperty9": "created_at~#~char",
          "customproperty10": "updated_at~#~char",
          "customproperty11": "suspended~#~char",
          "customproperty31": "STORE#ACC#ENT#MAPPINGINFO~#~char"
        },
        "pagination": {
          "nextUrl": {
            "nextUrlPath": "$${response?.completeResponseMap?.next_page==null?null:response.completeResponseMap.next_page}"
          }
        }
      }
    },
    "acctEntMappings": {
      "Role": {
        "listPath": "",
        "idPath": "custom_role_id",
        "keyField": "entitlementID"
      }
    }
  },
  "entitlementParams": {
    "connection": "acctAuth",
    "processingType": "SequentialAndIterative",
    "entTypes": {
      "Group": {
        "entTypeOrder": 0,
        "entTypeLabels": {
          "customproperty1": "Deleted",
          "customproperty2": "CreatedAt",
          "customproperty3": "UpdatedAt"
        },
        "call": {
          "call1": {
            "callOrder": 0,
            "stageNumber": 0,
            "http": {
              "url": "https://saviynt6799.zendesk.com/api/v2/groups",
              "httpHeaders": {
                "Authorization": "$${access_token}",
                "Accept": "application/json"
              },
              "httpContentType": "application/json",
              "httpMethod": "GET"
            },
            "listField": "groups",
            "keyField": "entitlementID",
            "colsToPropsMap": {
              "entitlementID": "id~#~char",
              "entitlement_value": "name~#~char",
              "customproperty1": "deleted~#~char",
              "customproperty2": "created_at~#~char",
              "customproperty3": "updated_at~#~char"
            },
            "pagination": {
              "nextUrl": {
                "nextUrlPath": "$${response?.completeResponseMap?.next_page==null?null:response.completeResponseMap.next_page}"
              }
            },
            "disableDeletedEntitlements": true
          }
        }
      },
      "Role": {
        "entTypeOrder": 1,
        "entTypeLabels": {
          "customproperty1": "Description",
          "customproperty2": "CreatedAt",
          "customproperty3": "UpdatedAt"
        },
        "call": {
          "call1": {
            "callOrder": 0,
            "stageNumber": 0,
            "http": {
              "url": "https://saviynt6799.zendesk.com/api/v2/custom_roles.json",
              "httpHeaders": {
                "Authorization": "$${access_token}",
                "Accept": "application/json"
              },
              "httpContentType": "application/json",
              "httpMethod": "GET"
            },
            "listField": "custom_roles",
            "keyField": "entitlementID",
            "colsToPropsMap": {
              "entitlementID": "id~#~char",
              "entitlement_value": "name~#~char",
              "customproperty1": "description~#~char",
              "customproperty2": "created_at~#~char",
              "customproperty3": "updated_at~#~char"
            },
            "pagination": {
              "nextUrl": {
                "nextUrlPath": "$${response?.completeResponseMap?.next_page==null?null:response.completeResponseMap.next_page}"
              }
            },
            "disableDeletedEntitlements": true
          }
        }
      }
    }
  },
  "acctEntParams": {
    "connection": "acctAuth",
    "entTypes": {
      "Group": {
        "call": {
          "call1": {
            "callOrder": 0,
            "stageNumber": 0,
            "processingType": "httpEntToAcct",
            "http": {
              "httpHeaders": {
                "Authorization": "$${access_token}"
              },
              "url": "https://saviynt6799.zendesk.com/api/v2/groups/$${id}/memberships.json",
              "httpContentType": "application/x-www-form-urlencoded",
              "httpMethod": "GET"
            },
            "listField": "group_memberships",
            "entKeyField": "entitlementID",
            "acctIdPath": "user_id",
            "acctKeyField": "accountID",
            "pagination": {
              "nextUrl": {
                "nextUrlPath": "$${response?.completeResponseMap?.next_page==null?null:response.completeResponseMap.next_page}"
              }
            }
          }
        }
      },
      "Role": {
        "call": {
          "call1": {
            "callOrder": 0,
            "stageNumber": 0,
            "processingType": "acctToEntMapping"
          }
        }
      }
    }
  }
})
update_account_json=jsonencode(
  {
  "dateFormat": "yyyy-MM-dd'T'HH:mm:ssXXX",
  "responseColsToPropsMap": {
    "displayName": "call1.message.user.name~#~char"
  },
  "call": [
    {
      "name": "Role",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/users/$${account.accountID}",
      "httpMethod": "PUT",
      "httpParams": "{\"user\": {\"name\": \"$${user.firstname} $${user.lastname}\"}}",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201
        ]
      }
    }
  ]
}
)
enable_account_json=jsonencode(
  {
  "call": [
    {
      "name": "call1",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/users",
      "httpMethod": "PUT",
      "httpParams": "{\"user\":{\"suspended\": \"false\"}}",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201
        ]
      }
    }
  ]
}
)
disable_account_json=jsonencode(
  {
  "call": [
    {
      "name": "call1",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/users",
      "httpMethod": "PUT",
      "httpParams": "{\"user\":{\"suspended\": \"true\"}}",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201
        ]
      }
    }
  ]
}
)
add_access_json=jsonencode(
  {
  "call": [
    {
      "name": "Group",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/group_memberships",
      "httpMethod": "POST",
      "httpParams": "{\"group_membership\": {\"user_id\": \"$${account.accountID}\", \"group_id\": \"$${entitlementValue.entitlementID}\"}}",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201
        ]
      }
    },
    {
      "name": "Role",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/users/$${account.accountID}",
      "httpMethod": "PUT",
      "httpParams": "{\"user\": {\"custom_role_id\": $${entitlementValue.entitlementID}}}",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201
        ]
      }
    }
  ]
}
)
remove_access_json=jsonencode(
  {
  "call": [
    {
      "name": "Group",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/users/$${account.accountID}/group_memberships",
      "httpMethod": "GET",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201
        ]
      }
    },
    {
      "name": "Group",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/group_memberships/$${for (Map map : response.Group1.message.group_memberships){if (map.group_id.toString().equals(entitlementValue.entitlementID)){return map.id;}}}",
      "httpMethod": "DELETE",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
"successResponses": {
        "statusCode": [
          200,
          204
        ]
      }
    },
    {
      "name": "Role",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/users/$${account.accountID}",
      "httpMethod": "PUT",
      "httpParams": "{\"user\": {\"custom_role_id\": $${entitlementValue.entitlementID}}}",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201
        ]
      }
    }
  ]
}
)
remove_account_json=jsonencode(
  {
  "call": [
    {
      "name": "call1",
      "connection": "acctAuth",
      "url": "https://saviynt6799.zendesk.com/api/v2/users/$${account.accountID}",
      "httpMethod": "DELETE",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201,
204
        ]
      }
    }
  ]
}
)
change_pass_json=jsonencode(
  {
  "call": [
    {
      "name": "call1",
      "connection": "acctAuth",
      "url": "https://hostname/api/v2/users/$${account.accountID}/password.json",
      "httpMethod": "POST",
      "httpParams": "{\"password\": \"$${password}\"}",
      "httpHeaders": {
        "Authorization": "$${access_token}",
        "Accept": "application/json"
      },
      "httpContentType": "application/json",
      "successResponses": {
        "statusCode": [
          200,
          201,
          204
        ]
      }
    }
  ]
}
)
}

