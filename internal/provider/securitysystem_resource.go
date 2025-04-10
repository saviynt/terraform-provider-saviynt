// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
	openapi "github.com/saviynt/saviynt-api-go-client/securitysystems"
)

// securitySystemResourceModel defines the state for our security system resource.
type securitySystemResourceModel struct {
	ID                                 types.String   `tfsdk:"id"`
	Systemname                         types.String   `tfsdk:"systemname"`
	DisplayName                        types.String   `tfsdk:"display_name"`
	Hostname                           types.String   `tfsdk:"hostname"`
	Port                               types.String   `tfsdk:"port"`
	AccessAddWorkflow                  types.String   `tfsdk:"access_add_workflow"`
	AccessRemoveWorkflow               types.String   `tfsdk:"access_remove_workflow"`
	AddServiceAccountWorkflow          types.String   `tfsdk:"add_service_account_workflow"`
	RemoveServiceAccountWorkflow       types.String   `tfsdk:"remove_service_account_workflow"`
	Connectionparameters               types.String   `tfsdk:"connection_parameters"`
	AutomatedProvisioning              types.String   `tfsdk:"automated_provisioning"`
	UseOpenConnector                   types.String   `tfsdk:"use_open_connector"`
	ReconApplication                   types.String   `tfsdk:"recon_application"`
	InstantProvision                   types.String   `tfsdk:"instant_provision"`
	ProvisioningTries                  types.String   `tfsdk:"provisioning_tries"`
	Provisioningcomments               types.String   `tfsdk:"provisioning_comments"`
	ProposedAccountOwnersWorkflow      types.String   `tfsdk:"proposed_account_owners_workflow"`
	FirefighterIDWorkflow              types.String   `tfsdk:"firefighterid_workflow"`
	FirefighterIDRequestAccessWorkflow types.String   `tfsdk:"firefighterid_request_access_workflow"`
	PolicyRule                         types.String   `tfsdk:"policy_rule"`
	PolicyRuleServiceAccount           types.String   `tfsdk:"policy_rule_service_account"`
	Connectionname                     types.String   `tfsdk:"connectionname"`
	ProvisioningConnection             types.String   `tfsdk:"provisioning_connection"`
	ServiceDeskConnection              types.String   `tfsdk:"service_desk_connection"`
	ExternalRiskConnectionJson         types.String   `tfsdk:"external_risk_connection_json"`
	// InherentSODReportFields            []types.String `tfsdk:"inherent_sod_report_fields"`
	InherentSODReportFields			types.List      `tfsdk:"inherent_sod_report_fields"`

	Msg       types.String `tfsdk:"msg"`
	ErrorCode types.String `tfsdk:"error_code"`
}

type SecuritySystemResource struct {
	client *s.Client
	token  string
}

func NewSecuritySystemResource() resource.Resource {
	return &SecuritySystemResource{}
}

func (r *SecuritySystemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_security_system_resource"
}

func (r *SecuritySystemResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and manage Security Systems in Saviynt",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the resource.",
			},
			"systemname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the security system.",
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "Specify a user-friendly display name that is shown on the the user interface.",
			},
			"hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Security system for which you want to create an endpoint.",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Description for the endpoint.",
			},
			"access_add_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the workflow to be used for approvals for an access request, which can be for an account, entitlements, role, and so on.",
			},
			"access_remove_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the workflow to be used when access has to be revoked, which can be for an account, entitlement, or any other de-provisioning task.",
			},
			"add_service_account_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Workflow for adding a service account.",
			},
			"remove_service_account_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Workflow for removing a service account.",
			},
			"proposed_account_owners_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users",
			},
			"firefighterid_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Firefighter ID Workflow.",
			},
			"firefighterid_request_access_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Firefighter ID Request Access Workflow.",
			},
			"connection_parameters": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users",
			},
			"automated_provisioning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to enable automated provisioning.",
			},
			"provisioning_tries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the number of tries to be used for performing provisioning / de-provisioning to the third-party application. You can specify provisioningTries between 1 to 20 based on your requirement.",
			},
			"connectionname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Select the connection name for performing reconciliation of identity objects from third-party application.",
			},
			"provisioning_connection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "You can use a separate connection to an endpoint where you are performing provisioning or deprovisioning. Based on your requirement, you can specify a separate connection where you want to perform provisioning and de-provisioning.",
			},
			"service_desk_connection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the Service Desk Connection used for integration with a ticketing system, which can be a disconnected system too.",
			},
			"provisioning_comments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify relevant comments for performing provisioning.",
			},
			"policy_rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this setting to assign the password policy for the security system.",
			},
			"policy_rule_service_account": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this setting to assign the password policy which will be used to set the service account passwords for the security system.",
			},
			"use_open_connector": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to enable the connectivity with any system over the open-source connectors such as REST.",
			},
			"recon_application": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to import data from the endpoint associated to the security system.",
			},
			"instant_provision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this flag to prevent users from raising duplicate requests for the same applications.",
			},
			"external_risk_connection_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Contains JSON configuration for external risk connections and is applicable only for a few connections like SAP.",
			},
			"inherent_sod_report_fields": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Optional:    true,
				Description: "You can use this option used to filter out columns in SOD.",
			},
			"msg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A message indicating the outcome of the operation.",
			},
			"error_code": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
			},
		},
	}
}

func (r *SecuritySystemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *saviyntProvider")
		return
	}

	// Set the client and token from the provider state.
	r.client = prov.client
	r.token = prov.accessToken
}

func (r *SecuritySystemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan securitySystemResourceModel

	// Extract the planned state from the request.
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Initialize OpenAPI Client Configuration.
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	// Build the create request using mandatory fields.
	createReq := openapi.CreateSecuritySystemRequest{
		Systemname:  plan.Systemname.ValueString(),
		DisplayName: plan.DisplayName.ValueString(),
	}

	// Add optional fields if they are not empty
	if !plan.Hostname.IsNull() && plan.Hostname.ValueString() != "" {
		createReq.SetHostname(plan.Hostname.ValueString())
	}
	if !plan.Port.IsNull() && plan.Port.ValueString() != "" {
		createReq.SetPort(plan.Port.ValueString())
	}
	if !plan.AccessAddWorkflow.IsNull() && plan.AccessAddWorkflow.ValueString() != "" {
		createReq.SetAccessAddWorkflow(plan.AccessAddWorkflow.ValueString())
	}
	if !plan.AccessRemoveWorkflow.IsNull() && plan.AccessRemoveWorkflow.ValueString() != "" {
		createReq.SetAccessRemoveWorkflow(plan.AccessRemoveWorkflow.ValueString())
	}
	if !plan.AddServiceAccountWorkflow.IsNull() && plan.AddServiceAccountWorkflow.ValueString() != "" {
		createReq.SetAddServiceAccountWorkflow(plan.AddServiceAccountWorkflow.ValueString())
	}
	if !plan.RemoveServiceAccountWorkflow.IsNull() && plan.RemoveServiceAccountWorkflow.ValueString() != "" {
		createReq.SetRemoveServiceAccountWorkflow(plan.RemoveServiceAccountWorkflow.ValueString())
	}
	if !plan.Connectionparameters.IsNull() && plan.Connectionparameters.ValueString() != "" {
		createReq.SetConnectionparameters(plan.Connectionparameters.ValueString())
	}
	if !plan.AutomatedProvisioning.IsNull() && plan.AutomatedProvisioning.ValueString() != "" {
		createReq.SetAutomatedProvisioning(plan.AutomatedProvisioning.ValueString())
	}
	if !plan.UseOpenConnector.IsNull() && plan.UseOpenConnector.ValueString() != "" {
		createReq.SetUseopenconnector(plan.UseOpenConnector.ValueString())
	}
	if !plan.ReconApplication.IsNull() && plan.ReconApplication.ValueString() != "" {
		createReq.SetReconApplication(plan.ReconApplication.ValueString())
	}
	if !plan.InstantProvision.IsNull() && plan.InstantProvision.ValueString() != "" {
		createReq.SetInstantprovision(plan.InstantProvision.ValueString())
	}
	if !plan.ProvisioningTries.IsNull() && plan.ProvisioningTries.ValueString() != "" {
		createReq.SetProvisioningTries(plan.ProvisioningTries.ValueString())
	}
	if !plan.Provisioningcomments.IsNull() && plan.Provisioningcomments.ValueString() != "" {
		createReq.SetProvisioningcomments(plan.Provisioningcomments.ValueString())
	}
	// Execute the API call.
	apiResp, httpResp, err := apiClient.SecuritySystemsAPI.CreateSecuritySystem(ctx).
		CreateSecuritySystemRequest(createReq).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Security System",
			fmt.Sprintf("Error: %v\nHTTP Response: %v", err, httpResp),
		)
		return
	}

	// Set the resource ID and store the API response in state.
	plan.ID = types.StringValue("security-system-" + plan.Systemname.ValueString())
	if plan.UseOpenConnector.IsNull() || plan.UseOpenConnector.IsUnknown() || plan.UseOpenConnector.ValueString() == "" {
		plan.UseOpenConnector = types.StringValue("false")
	}

	if plan.ReconApplication.IsNull() || plan.ReconApplication.IsUnknown() || plan.ReconApplication.ValueString() == "" {
		plan.ReconApplication = types.StringValue("true")
	}

	if plan.InstantProvision.IsNull() || plan.InstantProvision.IsUnknown() || plan.InstantProvision.ValueString() == "" {
		plan.InstantProvision = types.StringValue("false")
	}

	if plan.AutomatedProvisioning.IsNull() || plan.AutomatedProvisioning.IsUnknown() || plan.AutomatedProvisioning.ValueString() == "" {
		plan.AutomatedProvisioning = types.StringValue("false")
	}

	// plan.AccessAddWorkflow = util.SafeStringDatasource(plan.AccessAddWorkflow.ValueStringPointer())
	plan.Hostname = util.SafeStringDatasource(plan.Hostname.ValueStringPointer())
	plan.Port = util.SafeStringDatasource(plan.Port.ValueStringPointer())
	plan.ProvisioningTries = util.SafeStringDatasource(plan.ProvisioningTries.ValueStringPointer())
	plan.Connectionparameters = util.SafeStringDatasource(plan.Connectionparameters.ValueStringPointer())
	plan.AccessAddWorkflow = util.SafeStringDatasource(plan.AccessAddWorkflow.ValueStringPointer())
	plan.Provisioningcomments = util.SafeStringDatasource(plan.Provisioningcomments.ValueStringPointer())
	plan.AccessRemoveWorkflow = util.SafeStringDatasource(plan.AccessRemoveWorkflow.ValueStringPointer())
	plan.AddServiceAccountWorkflow = util.SafeStringDatasource(plan.AddServiceAccountWorkflow.ValueStringPointer())
	plan.RemoveServiceAccountWorkflow = util.SafeStringDatasource(plan.RemoveServiceAccountWorkflow.ValueStringPointer())
	plan.ProposedAccountOwnersWorkflow = util.SafeStringDatasource(plan.ProposedAccountOwnersWorkflow.ValueStringPointer())
	plan.FirefighterIDWorkflow = util.SafeStringDatasource(plan.FirefighterIDWorkflow.ValueStringPointer())
	plan.FirefighterIDRequestAccessWorkflow = util.SafeStringDatasource(plan.FirefighterIDRequestAccessWorkflow.ValueStringPointer())
	plan.Connectionname = util.SafeStringDatasource(plan.Connectionname.ValueStringPointer())
	plan.ProvisioningConnection = util.SafeStringDatasource(plan.ProvisioningConnection.ValueStringPointer())
	plan.ServiceDeskConnection = util.SafeStringDatasource(plan.ServiceDeskConnection.ValueStringPointer())
	plan.PolicyRule = util.SafeStringDatasource(plan.PolicyRule.ValueStringPointer())
	plan.PolicyRuleServiceAccount = util.SafeStringDatasource(plan.PolicyRuleServiceAccount.ValueStringPointer())
	plan.ExternalRiskConnectionJson = util.SafeStringDatasource(plan.ExternalRiskConnectionJson.ValueStringPointer())
	plan.InherentSODReportFields = util.NormalizeTFListString(plan.InherentSODReportFields)

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)

	// Set the individual fields
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SecuritySystemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state securitySystemResourceModel

	// Load current state from Terraform
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Initialize API client configuration
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)

	// Fetch current details of the security system using the unique identifier
	apiResp, httpResp, err := apiClient.SecuritySystemsAPI.GetSecuritySystems(ctx).Systemname(state.Systemname.ValueString()).Execute()
	if err != nil {
		// Handle 404: resource no longer exists, remove from state
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Security System", err.Error())
		return
	}
	var foundItem *openapi.GetSecuritySystems200ResponseSecuritySystemDetailsInner
	for _, item := range apiResp.SecuritySystemDetails {
		if item.Systemname != nil && *item.Systemname == state.Systemname.ValueString() {
			foundItem = &item
			break
		}
	}

	if foundItem == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	// Update state with the found item
	state.ID = types.StringValue("security-system-" + *foundItem.Systemname)
	log.Printf("[DEBUG] Recon in read application: ", util.SafeStringDatasource(foundItem.ReconApplication))
	state.DisplayName = types.StringValue(util.SafeDeref(foundItem.DisplayName))
	state.Hostname = util.SafeStringDatasource(foundItem.Hostname)
	state.Port = util.SafeStringDatasource(foundItem.Port)
	state.AccessAddWorkflow = util.SafeStringDatasource(foundItem.AccessAddWorkflow)
	log.Printf("[DEBUG] AccessAddWorkflow in read application: ", util.SafeStringDatasource(foundItem.AccessAddWorkflow))
	state.AccessRemoveWorkflow = util.SafeStringDatasource(foundItem.AccessRemoveWorkflow)
	state.AddServiceAccountWorkflow = util.SafeStringDatasource(foundItem.AddServiceAccountWorkflow)
	state.RemoveServiceAccountWorkflow = util.SafeStringDatasource(foundItem.RemoveServiceAccountWorkflow)
	state.Connectionparameters = util.SafeStringDatasource(foundItem.Connectionparameters)
	state.AutomatedProvisioning = util.SafeStringDatasource(foundItem.AutomatedProvisioning)
	state.UseOpenConnector = util.SafeStringDatasource(foundItem.Useopenconnector)
	state.ReconApplication = util.SafeStringDatasource(foundItem.ReconApplication)
	state.InstantProvision = util.SafeStringDatasource(foundItem.Instantprovision)
	state.ProvisioningTries = util.SafeStringDatasource(foundItem.ProvisioningTries)
	state.Provisioningcomments = util.SafeStringDatasource(foundItem.Provisioningcomments)
	state.ProposedAccountOwnersWorkflow = util.SafeStringDatasource(foundItem.ProposedAccountOwnersworkflow)
	state.FirefighterIDWorkflow = util.SafeStringDatasource(foundItem.FirefighteridWorkflow)
	state.FirefighterIDRequestAccessWorkflow = util.SafeStringDatasource(foundItem.FirefighteridRequestAccessWorkflow)
	state.PolicyRule = util.SafeStringDatasource(foundItem.PolicyRule)
	state.PolicyRuleServiceAccount = util.SafeStringDatasource(foundItem.PolicyRuleServiceAccount)
	state.Connectionname = util.SafeStringDatasource(foundItem.Connectionname)
	state.ProvisioningConnection = util.SafeStringDatasource(foundItem.ProvisioningConnection)
	state.ServiceDeskConnection = util.SafeStringDatasource(foundItem.ServiceDeskConnection)
	state.ExternalRiskConnectionJson = util.SafeStringDatasource(foundItem.ExternalRiskConnectionJson)

	state.InherentSODReportFields = util.ConvertStringsToTFListString(foundItem.InherentSODReportFields)

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)

	// Set the individual fields
	state.Msg = types.StringValue(msgValue)
	state.ErrorCode = types.StringValue(errorCodeValue)

	// Save updated state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
func (r *SecuritySystemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan securitySystemResourceModel

	// Extract the desired state from the request.
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Initialize OpenAPI Client Configuration.
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	// Build the update request using the required systemname.
	updateReq := openapi.UpdateSecuritySystemRequest{
		Systemname:  plan.Systemname.ValueString(),
		DisplayName: plan.DisplayName.ValueString(),
	}

	// Optional fields: set only if provided.
	if !plan.Hostname.IsNull() && plan.Hostname.ValueString() != "" {
		updateReq.SetHostname(plan.Hostname.ValueString())
	}
	if !plan.Port.IsNull() && plan.Port.ValueString() != "" {
		updateReq.SetPort(plan.Port.ValueString())
	}
	if !plan.AccessAddWorkflow.IsNull() && plan.AccessAddWorkflow.ValueString() != "" {
		updateReq.SetAccessAddWorkflow(plan.AccessAddWorkflow.ValueString())
	}
	if !plan.AccessRemoveWorkflow.IsNull() && plan.AccessRemoveWorkflow.ValueString() != "" {
		updateReq.SetAccessRemoveWorkflow(plan.AccessRemoveWorkflow.ValueString())
	}
	if !plan.AddServiceAccountWorkflow.IsNull() && plan.AddServiceAccountWorkflow.ValueString() != "" {
		updateReq.SetAddServiceAccountWorkflow(plan.AddServiceAccountWorkflow.ValueString())
	}
	if !plan.RemoveServiceAccountWorkflow.IsNull() && plan.RemoveServiceAccountWorkflow.ValueString() != "" {
		updateReq.SetRemoveServiceAccountWorkflow(plan.RemoveServiceAccountWorkflow.ValueString())
	}
	if !plan.Connectionparameters.IsNull() && plan.Connectionparameters.ValueString() != "" {
		updateReq.SetConnectionparameters(plan.Connectionparameters.ValueString())
	}
	if !plan.AutomatedProvisioning.IsNull() && plan.AutomatedProvisioning.ValueString() != "" {
		updateReq.SetAutomatedProvisioning(plan.AutomatedProvisioning.ValueString())
	}
	if !plan.UseOpenConnector.IsNull() && plan.UseOpenConnector.ValueString() != "" {
		updateReq.SetUseopenconnector(plan.UseOpenConnector.ValueString())
	}
	if !plan.ReconApplication.IsNull() && plan.ReconApplication.ValueString() != "" {
		updateReq.SetReconApplication(plan.ReconApplication.ValueString())
	}
	if !plan.InstantProvision.IsNull() && plan.InstantProvision.ValueString() != "" {
		updateReq.SetInstantprovision(plan.InstantProvision.ValueString())
	}
	if !plan.ProvisioningTries.IsNull() && plan.ProvisioningTries.ValueString() != "" {
		updateReq.SetProvisioningTries(plan.ProvisioningTries.ValueString())
	}
	if !plan.Provisioningcomments.IsNull() && plan.Provisioningcomments.ValueString() != "" {
		updateReq.SetProvisioningcomments(plan.Provisioningcomments.ValueString())
	}
	if !plan.ProposedAccountOwnersWorkflow.IsNull() && plan.ProposedAccountOwnersWorkflow.ValueString() != "" {
		updateReq.SetProposedAccountOwnersworkflow(plan.ProposedAccountOwnersWorkflow.ValueString())
	}
	if !plan.FirefighterIDWorkflow.IsNull() && plan.FirefighterIDWorkflow.ValueString() != "" {
		updateReq.SetFirefighteridWorkflow(plan.FirefighterIDWorkflow.ValueString())
	}
	if !plan.FirefighterIDRequestAccessWorkflow.IsNull() && plan.FirefighterIDRequestAccessWorkflow.ValueString() != "" {
		updateReq.SetFirefighteridRequestAccessWorkflow(plan.FirefighterIDRequestAccessWorkflow.ValueString())
	}
	if !plan.PolicyRule.IsNull() && plan.PolicyRule.ValueString() != "" {
		updateReq.SetPolicyRule(plan.PolicyRule.ValueString())
	}
	if !plan.PolicyRuleServiceAccount.IsNull() && plan.PolicyRuleServiceAccount.ValueString() != "" {
		updateReq.SetPolicyRuleServiceAccount(plan.PolicyRuleServiceAccount.ValueString())
	}
	if !plan.Connectionname.IsNull() && plan.Connectionname.ValueString() != "" {
		updateReq.SetConnectionname(plan.Connectionname.ValueString())
	}
	if !plan.ProvisioningConnection.IsNull() && plan.ProvisioningConnection.ValueString() != "" {
		updateReq.SetProvisioningConnection(plan.ProvisioningConnection.ValueString())
	}
	if !plan.ServiceDeskConnection.IsNull() && plan.ServiceDeskConnection.ValueString() != "" {
		updateReq.SetServiceDeskConnection(plan.ServiceDeskConnection.ValueString())
	}
	if !plan.ExternalRiskConnectionJson.IsNull() && plan.ExternalRiskConnectionJson.ValueString() != "" {
		updateReq.SetExternalRiskConnectionJson(plan.ExternalRiskConnectionJson.ValueString())
	}
	if !plan.InherentSODReportFields.IsNull() && !plan.InherentSODReportFields.IsUnknown() {
		elements := plan.InherentSODReportFields.Elements()
		if len(elements) > 0 {
			inherentFields := util.ConvertTFStringsToGoStrings(plan.InherentSODReportFields)
			updateReq.SetInherentSODReportFields(inherentFields)
		}
	}	
	// Execute the update API call.
	apiResp, httpResp, err := apiClient.SecuritySystemsAPI.
		UpdateSecuritySystem(ctx).
		UpdateSecuritySystemRequest(updateReq).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Security System",
			fmt.Sprintf("Error: %v\nHTTP Response: %v", err, httpResp),
		)
		log.Printf("API call error: ", err)
		return
	}

	getResp, httpResp, err := apiClient.SecuritySystemsAPI.GetSecuritySystems(ctx).Systemname(plan.Systemname.ValueString()).Execute()
	if err != nil {
		// Handle 404: resource no longer exists, remove from state
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Security System", err.Error())
		return
	}
	var foundItem *openapi.GetSecuritySystems200ResponseSecuritySystemDetailsInner
	for _, item := range getResp.SecuritySystemDetails {
		if item.Systemname != nil && *item.Systemname == plan.Systemname.ValueString() {
			foundItem = &item
			break
		}
	}

	if foundItem == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	log.Printf("[DEBUG] Recon in update application: ", util.SafeStringDatasource(foundItem.ReconApplication))

	// Update state with the found item
	plan.ID = types.StringValue("security-system-" + *foundItem.Systemname)
	plan.DisplayName = types.StringValue(util.SafeDeref(foundItem.DisplayName))
	plan.Hostname = util.SafeStringDatasource(foundItem.Hostname)
	plan.Port = util.SafeStringDatasource(foundItem.Port)
	plan.AccessAddWorkflow = util.SafeStringDatasource(foundItem.AccessAddWorkflow)
	plan.AccessRemoveWorkflow = util.SafeStringDatasource(foundItem.AccessRemoveWorkflow)
	log.Printf("[DEBUG] AccessAddWorkflow in update application: ", util.SafeStringDatasource(foundItem.AccessAddWorkflow))
	plan.AddServiceAccountWorkflow = util.SafeStringDatasource(foundItem.AddServiceAccountWorkflow)
	plan.RemoveServiceAccountWorkflow = util.SafeStringDatasource(foundItem.RemoveServiceAccountWorkflow)
	plan.Connectionparameters = util.SafeStringDatasource(foundItem.Connectionparameters)
	plan.AutomatedProvisioning = util.SafeStringDatasource(foundItem.AutomatedProvisioning)
	plan.UseOpenConnector = util.SafeStringDatasource(foundItem.Useopenconnector)
	plan.ReconApplication = util.SafeStringDatasource(foundItem.ReconApplication)
	plan.InstantProvision = util.SafeStringDatasource(foundItem.Instantprovision)
	plan.ProvisioningTries = util.SafeStringDatasource(foundItem.ProvisioningTries)
	plan.Provisioningcomments = util.SafeStringDatasource(foundItem.Provisioningcomments)
	plan.ProposedAccountOwnersWorkflow = util.SafeStringDatasource(foundItem.ProposedAccountOwnersworkflow)
	plan.FirefighterIDWorkflow = util.SafeStringDatasource(foundItem.FirefighteridWorkflow)
	plan.FirefighterIDRequestAccessWorkflow = util.SafeStringDatasource(foundItem.FirefighteridRequestAccessWorkflow)
	plan.PolicyRule = util.SafeStringDatasource(foundItem.PolicyRule)
	plan.PolicyRuleServiceAccount = util.SafeStringDatasource(foundItem.PolicyRuleServiceAccount)
	plan.Connectionname = util.SafeStringDatasource(foundItem.Connectionname)
	plan.ProvisioningConnection = util.SafeStringDatasource(foundItem.ProvisioningConnection)
	plan.ServiceDeskConnection = util.SafeStringDatasource(foundItem.ServiceDeskConnection)
	plan.ExternalRiskConnectionJson = util.SafeStringDatasource(foundItem.ExternalRiskConnectionJson)

	// Handle list attributes
	plan.InherentSODReportFields = util.ConvertStringsToTFListString(foundItem.InherentSODReportFields)

	// Optional: Save response as debug info
	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)

	// Set the individual fields
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SecuritySystemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// If deletion is supported by the API, call the corresponding API endpoint.
	// Otherwise, remove the resource from state.
	resp.State.RemoveResource(ctx)
}
