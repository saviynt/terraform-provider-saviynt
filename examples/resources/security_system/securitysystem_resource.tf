terraform {
  required_providers {
    saviynt = {
      source  = "registry.terraform.io/local/saviynt"
      version = "1.0.0"
    }
  }
}


provider "saviynt" {
  server_url = var.saviynt_server_url
  username   = var.saviynt_username
  password   = var.saviynt_password
}


resource "saviynt_security_system_resource" "example" {
  systemname                         = var.systemname
  display_name                       = var.display_name
  hostname                           = var.hostname
  port                               = var.port
  access_add_workflow                = var.access_add_workflow
  access_remove_workflow             = var.access_remove_workflow
  add_service_account_workflow       = var.add_service_account_workflow
  remove_service_account_workflow    = var.remove_service_account_workflow
  firefighterid_workflow             = var.firefighterid_workflow
  firefighterid_request_access_workflow = var.firefighterid_request_access_workflow
  policy_rule                        = var.policy_rule
  policy_rule_service_account        = var.policy_rule_service_account
  recon_application                  = var.recon_application
  instant_provision                  = var.instant_provision
  provisioning_tries                 = var.provisioning_tries
  connection_parameters              = var.connection_parameters
  automated_provisioning             = var.automated_provisioning
  use_open_connector                 = var.use_open_connector
  provisioning_comments              = var.provisioning_comments
  proposed_account_owners_workflow   = var.proposed_account_owners_workflow
  connectionname                     = var.connectionname
  provisioning_connection            = var.provisioning_connection
  service_desk_connection            = var.service_desk_connection
  inherent_sod_report_fields         = var.inherent_sod_report_fields
  external_risk_connection_json=var.external_risk_connection_json
}

check "instance_health" {
  assert {
   condition = saviynt_security_system_resource.example.error_code != "1"
   error_message = "The error is: ${saviynt_security_system_resource.example.msg}"
  }
}
