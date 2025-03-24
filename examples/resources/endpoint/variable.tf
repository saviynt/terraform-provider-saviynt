variable "saviynt_server_url" {
  type        = string
  description = "Saviynt API Server URL (without https://)"
}

variable "saviynt_username" {
  type        = string
  description = "Saviynt API Username"
}

variable "saviynt_password" {
  type        = string
  description = "Saviynt API Password"
  sensitive   = true
}

variable "endpoint_name" {
  description = "The name of the endpoint (required)."
  type        = string
}

variable "display_name" {
  description = "User-friendly display name for the endpoint (required)."
  type        = string
}

variable "security_system" {
  description = "The security system identifier for the endpoint (required)."
  type        = string
}

variable "description" {
  description = "Optional endpoint description."
  type        = string
  default     = ""
}

variable "owner_type" {
  description = "Optional owner type."
  type        = string
  default     = ""
}

variable "owner" {
  description = "Optional owner."
  type        = string
  default     = ""
}

variable "resource_owner_type" {
  description = "Optional resource owner type."
  type        = string
  default     = ""
}

variable "resource_owner" {
  description = "Optional resource owner."
  type        = string
  default     = ""
}

variable "access_query" {
  description = "Optional access query."
  type        = string
  default     = ""
}

variable "enable_copy_access" {
  description = "Optional flag to enable copy access."
  type        = string
  default     = ""
}

variable "disable_new_account_request_if_account_exists" {
  description = "Optional flag to disable new account request if account exists."
  type        = string
  default     = ""
}

variable "disable_remove_account" {
  description = "Optional flag to disable remove account."
  type        = string
  default     = ""
}

variable "disable_modify_account" {
  description = "Optional flag to disable modify account."
  type        = string
  default     = ""
}

variable "user_account_correlation_rule" {
  description = "Optional user account correlation rule."
  type        = string
  default     = ""
}

variable "create_ent_taskfor_remove_acc" {
  description = "Optional flag to create enterprise task for account removal."
  type        = string
  default     = ""
}

variable "out_of_band_action" {
  description = "Optional out-of-band action."
  type        = string
  default     = ""
}

variable "connection_config" {
  description = "Optional connection configuration."
  type        = string
  default     = ""
}

variable "requestable" {
  description = "Optional requestable parameter."
  type        = string
  default     = ""
}

variable "parent_account_pattern" {
  description = "Optional parent account pattern."
  type        = string
  default     = ""
}

variable "service_account_name_rule" {
  description = "Optional service account name rule."
  type        = string
  default     = ""
}

variable "service_account_access_query" {
  description = "Optional service account access query."
  type        = string
  default     = ""
}

variable "block_inflight_request" {
  description = "Optional flag to block inflight requests."
  type        = string
  default     = ""
}

variable "account_name_rule" {
  description = "Optional account name rule."
  type        = string
  default     = ""
}

variable "allow_change_password_sql_query" {
  description = "Optional SQL query to allow password changes."
  type        = string
  default     = ""
}

variable "account_name_validator_regex" {
  description = "Optional account name validator regex."
  type        = string
  default     = ""
}

variable "status_config" {
  description = "Optional status configuration."
  type        = string
  default     = ""
}

variable "plugin_configs" {
  description = "Optional plugin configurations."
  type        = string
  default     = ""
}

variable "endpoint_config" {
  description = "Optional endpoint configuration."
  type        = string
  default     = ""
}

variable "custom_property1" {
  description = "Optional custom property 1."
  type        = string
  default     = ""
}

variable "custom_property2" {
  description = "Optional custom property 2."
  type        = string
  default     = ""
}

variable "custom_property3" {
  description = "Optional custom property 3."
  type        = string
  default     = ""
}

variable "custom_property4" {
  description = "Optional custom property 4."
  type        = string
  default     = ""
}

variable "custom_property5" {
  description = "Optional custom property 5."
  type        = string
  default     = ""
}

variable "custom_property6" {
  description = "Optional custom property 6."
  type        = string
  default     = ""
}

variable "custom_property7" {
  description = "Optional custom property 7."
  type        = string
  default     = ""
}

variable "custom_property8" {
  description = "Optional custom property 8."
  type        = string
  default     = ""
}

variable "custom_property9" {
  description = "Optional custom property 9."
  type        = string
  default     = ""
}

variable "custom_property10" {
  description = "Optional custom property 10."
  type        = string
  default     = ""
}

variable "custom_property11" {
  description = "Optional custom property 11."
  type        = string
  default     = ""
}

variable "custom_property12" {
  description = "Optional custom property 12."
  type        = string
  default     = ""
}

variable "custom_property13" {
  description = "Optional custom property 13."
  type        = string
  default     = ""
}

variable "custom_property14" {
  description = "Optional custom property 14."
  type        = string
  default     = ""
}

variable "custom_property15" {
  description = "Optional custom property 15."
  type        = string
  default     = ""
}

variable "custom_property16" {
  description = "Optional custom property 16."
  type        = string
  default     = ""
}

variable "custom_property17" {
  description = "Optional custom property 17."
  type        = string
  default     = ""
}

variable "custom_property18" {
  description = "Optional custom property 18."
  type        = string
  default     = ""
}

variable "custom_property19" {
  description = "Optional custom property 19."
  type        = string
  default     = ""
}

variable "custom_property20" {
  description = "Optional custom property 20."
  type        = string
  default     = ""
}

variable "custom_property21" {
  description = "Optional custom property 21."
  type        = string
  default     = ""
}

variable "custom_property22" {
  description = "Optional custom property 22."
  type        = string
  default     = ""
}

variable "custom_property23" {
  description = "Optional custom property 23."
  type        = string
  default     = ""
}

variable "custom_property24" {
  description = "Optional custom property 24."
  type        = string
  default     = ""
}

variable "custom_property25" {
  description = "Optional custom property 25."
  type        = string
  default     = ""
}

variable "custom_property26" {
  description = "Optional custom property 26."
  type        = string
  default     = ""
}

variable "custom_property27" {
  description = "Optional custom property 27."
  type        = string
  default     = ""
}

variable "custom_property28" {
  description = "Optional custom property 28."
  type        = string
  default     = ""
}

variable "custom_property29" {
  description = "Optional custom property 29."
  type        = string
  default     = ""
}

variable "custom_property30" {
  description = "Optional custom property 30."
  type        = string
  default     = ""
}

variable "custom_property31" {
  description = "Optional custom property 31."
  type        = string
  default     = ""
}

variable "custom_property32" {
  description = "Optional custom property 32."
  type        = string
  default     = ""
}

variable "custom_property33" {
  description = "Optional custom property 33."
  type        = string
  default     = ""
}

variable "custom_property34" {
  description = "Optional custom property 34."
  type        = string
  default     = ""
}

variable "custom_property35" {
  description = "Optional custom property 35."
  type        = string
  default     = ""
}

variable "custom_property36" {
  description = "Optional custom property 36."
  type        = string
  default     = ""
}

variable "custom_property37" {
  description = "Optional custom property 37."
  type        = string
  default     = ""
}

variable "custom_property38" {
  description = "Optional custom property 38."
  type        = string
  default     = ""
}

variable "custom_property39" {
  description = "Optional custom property 39."
  type        = string
  default     = ""
}

variable "custom_property40" {
  description = "Optional custom property 40."
  type        = string
  default     = ""
}

variable "custom_property41" {
  description = "Optional custom property 41."
  type        = string
  default     = ""
}

variable "custom_property42" {
  description = "Optional custom property 42."
  type        = string
  default     = ""
}

variable "custom_property43" {
  description = "Optional custom property 43."
  type        = string
  default     = ""
}

variable "custom_property44" {
  description = "Optional custom property 44."
  type        = string
  default     = ""
}

variable "custom_property45" {
  description = "Optional custom property 45."
  type        = string
  default     = ""
}

##################################
# Custom Property Label Variables#
##################################

variable "custom_property1_label" {
  description = "Optional label for custom property 1."
  type        = string
  default     = ""
}

variable "custom_property2_label" {
  description = "Optional label for custom property 2."
  type        = string
  default     = ""
}

variable "custom_property3_label" {
  description = "Optional label for custom property 3."
  type        = string
  default     = ""
}

variable "custom_property4_label" {
  description = "Optional label for custom property 4."
  type        = string
  default     = ""
}

variable "custom_property5_label" {
  description = "Optional label for custom property 5."
  type        = string
  default     = ""
}

variable "custom_property6_label" {
  description = "Optional label for custom property 6."
  type        = string
  default     = ""
}

variable "custom_property7_label" {
  description = "Optional label for custom property 7."
  type        = string
  default     = ""
}

variable "custom_property8_label" {
  description = "Optional label for custom property 8."
  type        = string
  default     = ""
}

variable "custom_property9_label" {
  description = "Optional label for custom property 9."
  type        = string
  default     = ""
}

variable "custom_property10_label" {
  description = "Optional label for custom property 10."
  type        = string
  default     = ""
}

variable "custom_property11_label" {
  description = "Optional label for custom property 11."
  type        = string
  default     = ""
}

variable "custom_property12_label" {
  description = "Optional label for custom property 12."
  type        = string
  default     = ""
}

variable "custom_property13_label" {
  description = "Optional label for custom property 13."
  type        = string
  default     = ""
}

variable "custom_property14_label" {
  description = "Optional label for custom property 14."
  type        = string
  default     = ""
}

variable "custom_property15_label" {
  description = "Optional label for custom property 15."
  type        = string
  default     = ""
}

variable "custom_property16_label" {
  description = "Optional label for custom property 16."
  type        = string
  default     = ""
}

variable "custom_property17_label" {
  description = "Optional label for custom property 17."
  type        = string
  default     = ""
}

variable "custom_property18_label" {
  description = "Optional label for custom property 18."
  type        = string
  default     = ""
}

variable "custom_property19_label" {
  description = "Optional label for custom property 19."
  type        = string
  default     = ""
}

variable "custom_property20_label" {
  description = "Optional label for custom property 20."
  type        = string
  default     = ""
}

variable "custom_property21_label" {
  description = "Optional label for custom property 21."
  type        = string
  default     = ""
}

variable "custom_property22_label" {
  description = "Optional label for custom property 22."
  type        = string
  default     = ""
}

variable "custom_property23_label" {
  description = "Optional label for custom property 23."
  type        = string
  default     = ""
}

variable "custom_property24_label" {
  description = "Optional label for custom property 24."
  type        = string
  default     = ""
}

variable "custom_property25_label" {
  description = "Optional label for custom property 25."
  type        = string
  default     = ""
}

variable "custom_property26_label" {
  description = "Optional label for custom property 26."
  type        = string
  default     = ""
}

variable "custom_property27_label" {
  description = "Optional label for custom property 27."
  type        = string
  default     = ""
}

variable "custom_property28_label" {
  description = "Optional label for custom property 28."
  type        = string
  default     = ""
}

variable "custom_property29_label" {
  description = "Optional label for custom property 29."
  type        = string
  default     = ""
}

variable "custom_property30_label" {
  description = "Optional label for custom property 30."
  type        = string
  default     = ""
}

variable "custom_property31_label" {
  description = "Optional label for custom property 31."
  type        = string
  default     = ""
}

variable "custom_property32_label" {
  description = "Optional label for custom property 32."
  type        = string
  default     = ""
}

variable "custom_property33_label" {
  description = "Optional label for custom property 33."
  type        = string
  default     = ""
}

variable "custom_property34_label" {
  description = "Optional label for custom property 34."
  type        = string
  default     = ""
}

variable "custom_property35_label" {
  description = "Optional label for custom property 35."
  type        = string
  default     = ""
}

variable "custom_property36_label" {
  description = "Optional label for custom property 36."
  type        = string
  default     = ""
}

variable "custom_property37_label" {
  description = "Optional label for custom property 37."
  type        = string
  default     = ""
}

variable "custom_property38_label" {
  description = "Optional label for custom property 38."
  type        = string
  default     = ""
}

variable "custom_property39_label" {
  description = "Optional label for custom property 39."
  type        = string
  default     = ""
}

variable "custom_property40_label" {
  description = "Optional label for custom property 40."
  type        = string
  default     = ""
}

variable "custom_property41_label" {
  description = "Optional label for custom property 41."
  type        = string
  default     = ""
}

variable "custom_property42_label" {
  description = "Optional label for custom property 42."
  type        = string
  default     = ""
}

variable "custom_property43_label" {
  description = "Optional label for custom property 43."
  type        = string
  default     = ""
}

variable "custom_property44_label" {
  description = "Optional label for custom property 44."
  type        = string
  default     = ""
}

variable "custom_property45_label" {
  description = "Optional label for custom property 45."
  type        = string
  default     = ""
}

variable "custom_property46_label" {
  description = "Optional label for custom property 46."
  type        = string
  default     = ""
}

variable "custom_property47_label" {
  description = "Optional label for custom property 47."
  type        = string
  default     = ""
}

variable "custom_property48_label" {
  description = "Optional label for custom property 48."
  type        = string
  default     = ""
}

variable "custom_property49_label" {
  description = "Optional label for custom property 49."
  type        = string
  default     = ""
}

variable "custom_property50_label" {
  description = "Optional label for custom property 50."
  type        = string
  default     = ""
}

variable "custom_property60_label" {
  description = "Optional label for custom property 60."
  type        = string
  default     = ""
}
