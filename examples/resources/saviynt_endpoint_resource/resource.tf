resource "saviynt_endpoint_resource" "example" {
  endpointname                                  = "sample"
  display_name                                  = "sample"
  security_system                               = "samplesystem"
  description                                   = "Endpoint for Jira Production Access"
  owner_type                                    = "USER"
  owner                                         = "john.doe"
  resource_owner_type                           = "ROLE"
  resource_owner                                = "ROLE_JIRA_OWNER"
  access_query                                  = "SELECT * FROM ACCESS WHERE endpoint='JIRA'"
  enable_copy_access                            = "true"
  disable_new_account_request_if_account_exists = "false"
  disable_remove_account                        = "false"
  disable_modify_account                        = "false"
  user_account_correlation_rule                 = "MATCH_ON_USERNAME"
  create_ent_task_for_remove_acc                = "true"
  out_of_band_action                            = "RECONCILE"
  connection_config                             = "{\"host\":\"jira.example.com\", \"port\":443}"
  requestable                                   = "true"
  parent_account_pattern                        = "parent-${accountId}"
  service_account_name_rule                     = "svc-${username}"
  service_account_access_query                  = "SELECT * FROM ACCESS WHERE account_type='SERVICE'"
  block_inflight_request                        = "false"
  account_name_rule                             = "acct-${user.email}"
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
  # ...
  custom_property60_label = "Custom Label 60"

  allow_remove_all_role_on_request = "false"
  change_password_access_query     = "SELECT * FROM USERS WHERE changepassword = 1"

  # Complex blocks
  requestable_role_type = [
    {
      role_type       = "TECH_ROLE"
      request_option  = "SELF"
      required        = true
      requested_query = "SELECT * FROM ROLE WHERE type='TECH_ROLE'"
      selected_query  = "SELECT * FROM ROLE WHERE selected='true'"
      show_on         = "ACCESS_REQUEST"
    },
    {
      role_type       = "ENT_ROLE"
      request_option  = "MANAGER"
      required        = false
      requested_query = "SELECT * FROM ROLE WHERE type='ENT_ROLE'"
      selected_query  = "SELECT * FROM ROLE WHERE selected='true'"
      show_on         = "BIRTHRIGHT"
    }
  ]

  email_template = [
    {
      email_template_type = "CREATE"
      task_type           = "ACCOUNT_CREATE"
      email_template      = "AccountCreationNotification"
    },
    {
      email_template_type = "REVOKE"
      task_type           = "ACCOUNT_REMOVE"
      email_template      = "AccessRevokedNotification"
    }
  ]

  mapped_endpoints = [
    {
      security_system = "JIRA_SECURITY_SYSTEM"
      endpoint        = "JIRA_DEV_ENDPOINT"
      requestable     = "true"
      operation       = "SYNC"
    },
    {
      security_system = "JIRA_SECURITY_SYSTEM"
      endpoint        = "JIRA_QA_ENDPOINT"
      requestable     = "false"
      operation       = "READONLY"
    }
  ]
}

