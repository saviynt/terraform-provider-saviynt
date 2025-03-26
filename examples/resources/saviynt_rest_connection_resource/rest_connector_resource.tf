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
  connection_type         = "REST"
  connection_name         = "Maersk_AD_Connector_4"
  vault_connection        = var.VAULT_CONNECTION
  vault_configuration     = var.VAULT_CONFIG
  save_in_vault           = var.SAVE_IN_VAULT
  connection_json         = "{\"example_key\":\"connection_value\"}"
  import_user_json        = "{\"import_user\":\"data\"}"
  import_account_ent_json = "{\"import_account_ent\":\"data\"}"
  status_threshold_config = "{\"threshold\":100}"
  create_account_json     = "{\"create_account\":\"data\"}"
  update_account_json     = "{\"update_account\":\"data\"}"
  enable_account_json     = "{\"enable_account\":true}"
  disable_account_json    = "{\"disable_account\":false}"
  add_access_json         = "{\"add_access\":\"data\"}"
  remove_access_json      = "{\"remove_access\":\"data\"}"
  update_user_json        = "{\"update_user\":\"data\"}"
  change_pass_json        = "{\"change_pass\":\"data\"}"
  remove_account_json     = "{\"remove_account\":\"data\"}"
  ticket_status_json      = "{\"ticket_status\":\"data\"}"
  create_ticket_json      = "{\"create_ticket\":\"data\"}"
  endpoints_filter        = "{\"filter\":\"value\"}"
  passwd_policy_json      = "{\"min_length\":8,\"max_length\":16}"
  config_json             = "{\"timeout\":30}"
  add_ffid_access_json    = "{\"add_ffid_access\":\"data\"}"
  remove_ffid_access_json = "{\"remove_ffid_access\":\"data\"}"
  modify_user_data_json   = "{\"modify_user\":\"data\"}"
  send_otp_json           = "{\"send_otp\":\"data\"}"
  validate_otp_json       = "{\"validate_otp\":\"data\"}"
  pam_config              = "{\"pam\":\"config\"}"
}

check "instance_health" {
  assert {
    condition     = saviynt_rest_connection_resource.example.error_code != "1"
    error_message = "The error is: ${saviynt_rest_connection_resource.example.msg}"
  }
}
