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
  connection_name = "Terraform_Rest_Connector10"
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
    retryCount = 3
    timeout    = 30
  })
  add_ffid_access_json = jsonencode({
    method = "POST"
    url    = "https://api.example.com/ffid/access"
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
resource "saviynt_security_system_resource" "ss" {
  systemname                      = "Terraform_Security_System10"
  display_name                    = "Terraform_Security_System10"
  hostname                        = "EntitlementsOnly"
  port                            = "443"
  access_add_workflow             = "autoapprovalwf"
  access_remove_workflow          = "autoapprovalwf"
  add_service_account_workflow    = "autoapprovalwf"
  remove_service_account_workflow = "autoapprovalwf"
  automated_provisioning          = "true"
  use_open_connector              = "true"
  recon_application               = "true"
  instant_provision               = "true"
  provisioning_tries              = "3"
  provisioning_comments           = "Auto-provisioned by Terraform"
  connectionname="Terraform_Rest_Connector10"
}
resource "saviynt_endpoint_resource" "ep" {
  endpointname                                  = "Terraform_Endpoint10"
  display_name                                  = "Terraform_Endpoint10"
  security_system                               = "Terraform_Security_System10"
  description                                   = "Endpoint for Jira Production Access"
  owner_type                                    = "USER"
  owner                                         = "admin"
  resource_owner_type                           = "User"
  resource_owner                                = "admin"
  access_query                                  = "SELECT * FROM ACCESS WHERE endpoint='JIRA'"
  enable_copy_access                            = "true"
  disable_new_account_request_if_account_exists = "false"
  disable_remove_account                        = "false"
  disable_modify_account                        = "false"
  user_account_correlation_rule                 = "MATCH_ON_USERNAME"
  create_ent_task_for_remove_acc                = "true"
  out_of_band_action                            = "2"
  requestable                                   = "true"
  service_account_access_query                  = "SELECT * FROM ACCESS WHERE account_type='SERVICE'"
  block_inflight_request                        = "ON"
  account_name_rule                             = "acct-$${user.email}"
  allow_change_password_sql_query               = "SELECT 1 FROM dual"
  account_name_validator_regex                  = "^[a-zA-Z0-9_.-]{5,15}$"
  status_config   = "{\"enabled\":true, \"checkInterval\":10}"
  plugin_configs  = "{\"pluginVersion\":\"1.2.3\"}"
  endpoint_config = "{\"audit\":true}"
  # Sample custom properties (only showing 1–5 for brevity)
  custom_property1 = "BusinessUnit"
  custom_property2 = "ApplicationName"
  custom_property3 = "Region"
  custom_property4 = "Environment"
  custom_property5 = "IntegrationID"
  # Labels for custom properties
  custom_property1_label = "Business Unit"
  custom_property2_label = "App Name"
  custom_property3_label = "Region"
  custom_property4_label = "Environment"
  custom_property5_label = "Integration ID"
  # The rest can be filled similarly (up to 60)
  custom_property60_label = "Custom Label 60"
  allow_remove_all_role_on_request = "false"
  change_password_access_query     = "SELECT * FROM USERS WHERE changepassword = 1"
}
 