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

resource "saviynt_rest_connection_resource" "rest" {
  connection_type = "REST"
  connection_name = "Terraform_Rest_Connector20"
  connection_json = jsonencode({
    url      = "https://api.example.com"
    authType = "Bearer"
    token    = "$${access_token}"
  })
  import_user_json = jsonencode({
    method = "GET"
    url    = "https://api.example.com/users"
    headers = {
      Authorization = "Bearer $${access_token}"
    }
    keyField = "id"
    colsToPropsMap = {
      userID   = "id~#~char"
      username = "username~#~char"
      email    = "email~#~char"
    }
  })
  import_account_ent_json = jsonencode({
    method   = "GET"
    url      = "https://api.example.com/userEntitlements"
    keyField = "accountID"
    colsToPropsMap = {
      accountID = "account.id~#~char"
      entName   = "entitlement.name~#~char"
    }
  })
  status_threshold_config = jsonencode({
    threshold = 5
    unit      = "days"
  })
  create_account_json = jsonencode({
    method = "POST"
    url    = "https://api.example.com/accounts"
    body = {
      username = "$${user.username}"
      email    = "$${user.email}"
      role     = "$${user.defaultRole}"
    }
  })
  update_account_json = jsonencode({
    method = "PUT"
    url    = "https://api.example.com/accounts/$${user.accountID}"
    body = {
      displayName = "$${user.fullname}"
    }
  })
  enable_account_json = jsonencode({
    method = "PATCH"
    url    = "https://api.example.com/accounts/$${user.accountID}/enable"
  })
  disable_account_json = jsonencode({
    method = "PATCH"
    url    = "https://api.example.com/accounts/$${user.accountID}/disable"
  })
  add_access_json = jsonencode({
    method = "POST"
    url    = "https://api.example.com/accounts/$${user.accountID}/access"
    body = {
      entitlementID = "$${entitlement.id}"
    }
  })
  remove_access_json = jsonencode({
    method = "DELETE"
    url    = "https://api.example.com/accounts/$${user.accountID}/access/$${entitlement.id}"
  })
  update_user_json = jsonencode({
    method = "PATCH"
    url    = "https://api.example.com/users/$${user.userID}"
    body = {
      phone = "$${user.phone}"
    }
  })
  change_pass_json = jsonencode({
    method = "POST"
    url    = "https://api.example.com/accounts/$${user.accountID}/password"
    body = {
      newPassword = "$${user.new_password}"
    }
  })
  remove_account_json = jsonencode({
    method = "DELETE"
    url    = "https://api.example.com/accounts/$${user.accountID}"
  })
  ticket_status_json = jsonencode({
    method = "GET"
    url    = "https://servicedesk.example.com/tickets/$${ticketID}/status"
  })
  create_ticket_json = jsonencode({
    method = "POST"
    url    = "https://servicedesk.example.com/tickets"
    body = {
      subject     = "Access Request for $${user.username}"
      description = "Please grant access to $${entitlement.name}"
      priority    = "Medium"
    }
  })
  endpoints_filter = "endpoint_type = 'REST'"
  passwd_policy_json = jsonencode({
    minLength      = 8
    maxLength      = 16
    numUpperCase   = 1
    numDigits      = 2
    numSpecialChar = 1
  })
  config_json = jsonencode({
    showLogs="true"
  })
  add_ffid_access_json = jsonencode({
    method = "POST"
    url    = "https://api.example.com/ffid/acces111"
    body = {
      userID = "$${user.userID}"
      ffid   = "$${entitlement.ffid}"
    }
  })
  remove_ffid_access_json = jsonencode({
    method = "DELETE"
    url    = "https://api.example.com/ffid/access/$${entitlement.ffid}?user=$${user.userID}"
  })
  modify_user_data_json = jsonencode({
    method = "PATCH"
    url    = "https://api.example.com/users/$${user.userID}"
    body = {
      department = "$${user.department}"
    }
  })
  send_otp_json = jsonencode({
    method = "POST"
    url    = "https://api.example.com/otp/send"
    body = {
      user = "$${user.email}"
    }
  })
  validate_otp_json = jsonencode({
    method = "POST"
    url    = "https://api.example.com/otp/validate"
    body = {
      otp   = "$${otp}"
      email = "$${user.email}"
    }
  })
  pam_config = jsonencode({
    pam_enabled  = true
    vault_system = "CyberArk"
    rotation     = "on_login"
  })
}
