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

variable "systemname" {
  type        = string
  description = "System name of the security system."
}

variable "display_name" {
  type        = string
  description = "Display name of the security system."
}

variable "hostname" {
  type        = string
  description = "Hostname of the security system."
}

variable "port" {
  type        = string
  description = "Port used by the security system."
}

variable "access_add_workflow" {
  type        = string
  description = "Workflow used to add access."
}

variable "access_remove_workflow" {
  type        = string
  description = "Workflow used to remove access."
}

variable "add_service_account_workflow" {
  type        = string
  description = "Workflow to add a service account."
}

variable "remove_service_account_workflow" {
  type        = string
  description = "Workflow to remove a service account."
}

variable "connection_parameters" {
  type        = string
  description = "Connection parameters in JSON format."
}

variable "automated_provisioning" {
  type        = string
  description = "Flag indicating if automated provisioning is enabled."
}

variable "use_open_connector" {
  type        = string
  description = "Flag indicating whether to use the open connector."
}

variable "recon_application" {
  type        = string
  description = "Indicates if reconciliation is enabled for the application."
}

variable "instant_provision" {
  type        = string
  description = "Flag indicating if instant provisioning is enabled."
}

variable "provisioning_tries" {
  type        = string
  description = "The number of provisioning attempts."
}

variable "provisioning_comments" {
  type        = string
  description = "Comments regarding the provisioning process."
}

variable "proposed_account_owners_workflow" {
  type        = string
  description = "Workflow for proposed account owners."
}

variable "firefighterid_workflow" {
  type        = string
  description = "Workflow for firefighter ID."
}

variable "firefighterid_request_access_workflow" {
  type        = string
  description = "Workflow for firefighter ID request access."
}

variable "policy_rule" {
  type        = string
  description = "Policy rule applied to the security system."
}

variable "policy_rule_service_account" {
  type        = string
  description = "Policy rule for the service account."
}

variable "connectionname" {
  type        = string
  description = "Name of the connection."
}

variable "provisioning_connection" {
  type        = string
  description = "Provisioning connection identifier."
}

variable "service_desk_connection" {
  type        = string
  description = "Connection used for the service desk."
}

variable "external_risk_connection_json" {
  type        = string
  description = "JSON string for external risk connection details."
}

variable "inherent_sod_report_fields" {
  type        = list(string)
  description = "List of inherent SOD report fields."
}
