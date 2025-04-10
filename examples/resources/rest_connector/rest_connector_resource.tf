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
  connection_name = "shaleen_test_rest_10100"
    vault_connection    = ""
  vault_configuration = ""
  save_in_vault       =""
  config_json     = ""

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
}