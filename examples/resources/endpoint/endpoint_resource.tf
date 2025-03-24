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


resource "saviynt_endpoint_resource" "example" {
 endpointname   = var.endpoint_name
  display_name    = var.display_name
  security_system = var.security_system

  description = var.description
  owner_type  = var.owner_type
  owner       = var.owner
  resource_owner_type = var.resource_owner_type
  resource_owner      = var.resource_owner
  access_query        = var.access_query
  enable_copy_access  = var.enable_copy_access
  disable_new_account_request_if_account_exists = var.disable_new_account_request_if_account_exists
  disable_remove_account = var.disable_remove_account
  disable_modify_account = var.disable_modify_account
  user_account_correlation_rule = var.user_account_correlation_rule
  create_ent_taskfor_remove_acc = var.create_ent_taskfor_remove_acc
  out_of_band_action          = var.out_of_band_action
  connection_config           = var.connection_config
  requestable                 = var.requestable
  parent_account_pattern      = var.parent_account_pattern
  service_account_name_rule   = var.service_account_name_rule
  service_account_access_query = var.service_account_access_query
  block_inflight_request      = var.block_inflight_request
  account_name_rule           = var.account_name_rule
  allow_change_password_sql_query = var.allow_change_password_sql_query
  account_name_validator_regex    = var.account_name_validator_regex
  status_config               = var.status_config
  plugin_configs              = var.plugin_configs
  endpoint_config             = var.endpoint_config

  custom_property1  = var.custom_property1
  custom_property2  = var.custom_property2
  custom_property3  = var.custom_property3
  custom_property4  = var.custom_property4
  custom_property5  = var.custom_property5
  custom_property6  = var.custom_property6
  custom_property7  = var.custom_property7
  custom_property8  = var.custom_property8
  custom_property9  = var.custom_property9
  custom_property10 = var.custom_property10
  custom_property11 = var.custom_property11
  custom_property12 = var.custom_property12
  custom_property13 = var.custom_property13
  custom_property14 = var.custom_property14
  custom_property15 = var.custom_property15
  custom_property16 = var.custom_property16
  custom_property17 = var.custom_property17
  custom_property18 = var.custom_property18
  custom_property19 = var.custom_property19
  custom_property20 = var.custom_property20
  custom_property21 = var.custom_property21
  custom_property22 = var.custom_property22
  custom_property23 = var.custom_property23
  custom_property24 = var.custom_property24
  custom_property25 = var.custom_property25
  custom_property26 = var.custom_property26
  custom_property27 = var.custom_property27
  custom_property28 = var.custom_property28
  custom_property29 = var.custom_property29
  custom_property30 = var.custom_property30
  custom_property31 = var.custom_property31
  custom_property32 = var.custom_property32
  custom_property33 = var.custom_property33
  custom_property34 = var.custom_property34
  custom_property35 = var.custom_property35
  custom_property36 = var.custom_property36
  custom_property37 = var.custom_property37
  custom_property38 = var.custom_property38
  custom_property39 = var.custom_property39
  custom_property40 = var.custom_property40
  custom_property41 = var.custom_property41
  custom_property42 = var.custom_property42
  custom_property43 = var.custom_property43
  custom_property44 = var.custom_property44
  custom_property45 = var.custom_property45
  custom_property1_label  = var.custom_property1_label
  custom_property2_label  = var.custom_property2_label
  custom_property3_label  = var.custom_property3_label
  custom_property4_label  = var.custom_property4_label
  custom_property5_label  = var.custom_property5_label
  custom_property6_label  = var.custom_property6_label
  custom_property7_label  = var.custom_property7_label
  custom_property8_label  = var.custom_property8_label
  custom_property9_label  = var.custom_property9_label
  custom_property10_label = var.custom_property10_label
  custom_property11_label = var.custom_property11_label
  custom_property12_label = var.custom_property12_label
  custom_property13_label = var.custom_property13_label
  custom_property14_label = var.custom_property14_label
  custom_property15_label = var.custom_property15_label
  custom_property16_label = var.custom_property16_label
  custom_property17_label = var.custom_property17_label
  custom_property18_label = var.custom_property18_label
  custom_property19_label = var.custom_property19_label
  custom_property20_label = var.custom_property20_label
  custom_property21_label = var.custom_property21_label
  custom_property22_label = var.custom_property22_label
  custom_property23_label = var.custom_property23_label
  custom_property24_label = var.custom_property24_label
  custom_property25_label = var.custom_property25_label
  custom_property26_label = var.custom_property26_label
  custom_property27_label = var.custom_property27_label
  custom_property28_label = var.custom_property28_label
  custom_property29_label = var.custom_property29_label
  custom_property30_label = var.custom_property30_label
  custom_property31_label = var.custom_property31_label
  custom_property32_label = var.custom_property32_label
  custom_property33_label = var.custom_property33_label
  custom_property34_label = var.custom_property34_label
  custom_property35_label = var.custom_property35_label
  custom_property36_label = var.custom_property36_label
  custom_property37_label = var.custom_property37_label
  custom_property38_label = var.custom_property38_label
  custom_property39_label = var.custom_property39_label
  custom_property40_label = var.custom_property40_label
  custom_property41_label = var.custom_property41_label
  custom_property42_label = var.custom_property42_label
  custom_property43_label = var.custom_property43_label
  custom_property44_label = var.custom_property44_label
  custom_property45_label = var.custom_property45_label
  custom_property46_label = var.custom_property46_label
  custom_property47_label = var.custom_property47_label
  custom_property48_label = var.custom_property48_label
  custom_property49_label = var.custom_property49_label
  custom_property50_label = var.custom_property50_label
  custom_property60_label = var.custom_property60_label
}
