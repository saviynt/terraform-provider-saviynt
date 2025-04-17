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
  resource "saviynt_security_system_resource" "example" {
	systemname = "nameforthesecuritysystem"
  	display_name="displaynameforthesecuritysystem"
	hostname                           = "jira.example.com"
	port                               = "443"
	access_add_workflow                = "AccessAdd_Default_Workflow"
	access_remove_workflow             = "AccessRemove_Approval_Workflow"
	add_service_account_workflow       = "ServiceAccountAdd_Approval"
	remove_service_account_workflow    = "ServiceAccountRemove_Immediate"
	connection_parameters = <<EOT
	{
	"authType": "basic",
	"apiKey": "abc123",
	"timeout": 30
	}
	EOT
	automated_provisioning             = "true"
	use_open_connector                 = "false"
	recon_application                  = "JIRA_RECON_APP"
	instant_provision                  = "true"
	provisioning_tries                 = "3"
	provisioning_comments              = "Auto-provisioned by Terraform"
	proposed_account_owners_workflow  = "AccountOwnerApprovalWorkflow"
	firefighterid_workflow            = "FirefighterDefaultWorkflow"
	firefighterid_request_access_workflow = "FirefighterAccessRequest"
	policy_rule                        = "Default_Policy_Rule"
	policy_rule_service_account       = "ServiceAccountPolicy"
	connectionname                     = "jira_connection"
	provisioning_connection            = "jira_provisioning_conn"
	service_desk_connection            = "servicenow_sd_conn"
	external_risk_connection_json = <<EOT
	{
	"url": "https://external-risk-api.example.com",
	"apiKey": "external-api-key",
	"threshold": 0.7
	}
	EOT
	inherent_sod_report_fields = <<EOT
	[
	"Application Name",
	"Risk Level",
	"Control Owner",
	"Violation Count"
	]
	EOT
  }