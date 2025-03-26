// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"terraform-provider-Saviynt/util"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	s "github.com/saviynt/saviynt-api-go-client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RESTConnectorResourceModel struct {
	BaseConnector
	ID                    types.String `tfsdk:"id"`
	ConnectionJSON        types.String `tfsdk:"connection_json"`
	ImportUserJson        types.String `tfsdk:"import_user_json"`
	ImportAccountEntJson  types.String `tfsdk:"import_account_ent_json"`
	StatusThresholdConfig types.String `tfsdk:"status_threshold_config"`
	CreateAccountJson     types.String `tfsdk:"create_account_json"`
	UpdateAccountJson     types.String `tfsdk:"update_account_json"`
	EnableAccountJson     types.String `tfsdk:"enable_account_json"`
	DisableAccountJson    types.String `tfsdk:"disable_account_json"`
	AddAccessJson         types.String `tfsdk:"add_access_json"`
	RemoveAccessJson      types.String `tfsdk:"remove_access_json"`
	UpdateUserJson        types.String `tfsdk:"update_user_json"`
	ChangePassJson        types.String `tfsdk:"change_pass_json"`
	RemoveAccountJson     types.String `tfsdk:"remove_account_json"`
	TicketStatusJson      types.String `tfsdk:"ticket_status_json"`
	CreateTicketJson      types.String `tfsdk:"create_ticket_json"`
	EndpointsFilter       types.String `tfsdk:"endpoints_filter"`
	PasswdPolicyJson      types.String `tfsdk:"passwd_policy_json"`
	ConfigJSON            types.String `tfsdk:"config_json"`
	AddFFIDAccessJson     types.String `tfsdk:"add_ffid_access_json"`
	RemoveFFIDAccessJson  types.String `tfsdk:"remove_ffid_access_json"`
	ModifyUserdataJson    types.String `tfsdk:"modify_user_data_json"`
	SendOtpJson           types.String `tfsdk:"send_otp_json"`
	ValidateOtpJson       types.String `tfsdk:"validate_otp_json"`
	PamConfig             types.String `tfsdk:"pam_config"`
}

// testConnectionResource implements the resource.Resource interface.
type restConnectionResource struct {
	// client *openapi.APIClient
	client *s.Client
	token  string
}

// NewTestConnectionResource returns a new instance of testConnectionResource.
func RestNewTestConnectionResource() resource.Resource {
	return &restConnectionResource{}
}

func (r *restConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_rest_connection_resource"
}

func (r *restConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and Manage Connections",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Resource ID.",
			},
			"connection_name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the connection. Example: \"Active Directory_Doc\"",
			},
			"connection_type": schema.StringAttribute{
				Required:    true,
				Description: "Connection type (e.g., 'AD' for Active Directory). Example: \"AD\"",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "Description for the connection. Example: \"ORG_AD\"",
			},
			"defaultsavroles": schema.StringAttribute{
				Optional:    true,
				Description: "Default SAV roles for managing the connection. Example: \"ROLE_ORG\"",
			},
			"email_template": schema.StringAttribute{
				Optional:    true,
				Description: "Email template for notifications. Example: \"New Account Task Creation\"",
			},
			"ssl_certificate": schema.StringAttribute{
				Optional:    true,
				Description: "SSL certificates to secure the connection. Example: \"-----BEGIN CERTIFICATE----- ... -----END CERTIFICATE-----\"",
			},
			"vault_connection": schema.StringAttribute{
				Optional:    true,
				Description: "Specifies the type of vault connection being used (e.g., 'Hashicorp'). Example: \"Hashicorp\"",
			},
			"vault_configuration": schema.StringAttribute{
				Optional:    true,
				Description: "JSON string specifying vault configuration. Example: '{\"path\":\"/secrets/data/kv-dev-intgn1/-AD_Credential\",\"keyMapping\":{\"PASSWORD\":\"AD_PASSWORD~#~None\"}}'",
			},
			"save_in_vault": schema.StringAttribute{
				Optional:    true,
				Description: "Flag indicating whether the encrypted attribute should be saved in the configured vault. Example: \"false\"",
			},
			"connection_json": schema.StringAttribute{
				Optional:    true,
				Description: "Dynamic JSON configuration for the connection. Must be a valid JSON object string.",
			},
			"import_user_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for importing users.",
			},
			"import_account_ent_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for importing accounts and entitlements.",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Description: "JSON configuration for status thresholds.",
			},
			"create_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to create an account.",
			},
			"update_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to update an account.",
			},
			"enable_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON configuration to enable an account.",
			},
			"disable_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON configuration to disable an account.",
			},
			"add_access_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to add access.",
			},
			"remove_access_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to remove access.",
			},
			"update_user_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to update a user.",
			},
			"change_pass_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to change a user’s password.",
			},
			"remove_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to remove an account.",
			},
			"ticket_status_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to check ticket status.",
			},
			"create_ticket_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to create a ticket.",
			},
			"endpoints_filter": schema.StringAttribute{
				Optional:    true,
				Description: "Filter criteria for endpoints.",
			},
			"passwd_policy_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON defining the password policy.",
			},
			"config_json": schema.StringAttribute{
				Optional:    true,
				Description: "General configuration JSON for the REST connector.",
			},
			"add_ffid_access_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to add FFID access.",
			},
			"remove_ffid_access_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to remove FFID access.",
			},
			"modify_user_data_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for modifying user data.",
			},
			"send_otp_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to send OTP.",
			},
			"validate_otp_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to validate OTP.",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Description: "PAM configuration JSON.",
			},
			"result": schema.StringAttribute{
				Computed:    true,
				Description: "Result of the operation.",
			},
			"msg": schema.StringAttribute{
				Computed:    true,
				Description: "Message returned from the operation.",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code if the operation fails.",
			},
		},
	}
}

func (r *restConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
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

func (r *restConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RESTConnectorResourceModel

	// Extract plan from request
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var connJSON map[string]interface{}
	err := json.Unmarshal([]byte(plan.ConnectionJSON.ValueString()), &connJSON)
	if err != nil {
		log.Fatalf("Failed to unmarshal ConnectionJSON: %v", err)
	}
	fmt.Print("shaleen is great %T\n", connJSON)
	cfg := openapi.NewConfiguration()
	apiBaseURL := r.client.APIBaseURL()
	if strings.HasPrefix(apiBaseURL, "https://") {
		apiBaseURL = strings.TrimPrefix(apiBaseURL, "https://")
	}
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	restConn := openapi.RESTConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype: "REST",
			ConnectionName: plan.ConnectionName.ValueString(),
			// Description:        util.StringPtr(plan.Description.ValueString()),
			// Defaultsavroles:    util.StringPtr(plan.DefaultSavRoles.ValueString()),
			// EmailTemplate:      util.StringPtr(plan.EmailTemplate.ValueString()),
			// SslCertificate:     util.StringPtr(plan.SSLCertificate.ValueString()),
			// VaultConnection:    util.StringPtr(plan.VaultConnection.ValueString()),
			// VaultConfiguration: util.StringPtr(plan.VaultConfiguration.ValueString()),
			// Saveinvault:        util.StringPtr(plan.SaveInVault.ValueString()),
		},
		ConnectionJSON:          connJSON,
		ImportUserJSON:          util.StringPtr(plan.ImportUserJson.ValueString()),
		ImportAccountEntJSON:    util.StringPtr(plan.ImportAccountEntJson.ValueString()),
		STATUS_THRESHOLD_CONFIG: util.StringPtr(plan.StatusThresholdConfig.ValueString()),
		CreateAccountJSON:       util.StringPtr(plan.CreateAccountJson.ValueString()),
		UpdateAccountJSON:       util.StringPtr(plan.UpdateAccountJson.ValueString()),
		EnableAccountJSON:       util.StringPtr(plan.EnableAccountJson.ValueString()),
		DisableAccountJSON:      util.StringPtr(plan.DisableAccountJson.ValueString()),
		AddAccessJSON:           util.StringPtr(plan.AddAccessJson.ValueString()),
		RemoveAccessJSON:        util.StringPtr(plan.RemoveAccessJson.ValueString()),
		UpdateUserJSON:          util.StringPtr(plan.UpdateUserJson.ValueString()),
		ChangePassJSON:          util.StringPtr(plan.ChangePassJson.ValueString()),
		RemoveAccountJSON:       util.StringPtr(plan.RemoveAccountJson.ValueString()),
		TicketStatusJSON:        util.StringPtr(plan.TicketStatusJson.ValueString()),
		CreateTicketJSON:        util.StringPtr(plan.CreateTicketJson.ValueString()),
		ENDPOINTS_FILTER:        util.StringPtr(plan.EndpointsFilter.ValueString()),
		PasswdPolicyJSON:        util.StringPtr(plan.PasswdPolicyJson.ValueString()),
		ConfigJSON:              util.StringPtr(plan.ConfigJSON.ValueString()),
		AddFFIDAccessJSON:       util.StringPtr(plan.AddFFIDAccessJson.ValueString()),
		RemoveFFIDAccessJSON:    util.StringPtr(plan.RemoveFFIDAccessJson.ValueString()),
		MODIFYUSERDATAJSON:      util.StringPtr(plan.ModifyUserdataJson.ValueString()),
		SendOtpJSON:             util.StringPtr(plan.SendOtpJson.ValueString()),
		ValidateOtpJSON:         util.StringPtr(plan.ValidateOtpJson.ValueString()),
		PAM_CONFIG:              util.StringPtr(plan.PamConfig.ValueString()),
	}
	resttConnRequest := openapi.TestConnectionRequest{
		RESTConnector: &restConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, httpResp, err := apiClient.ConnectionsAPI.TestConnection(ctx).TestConnectionRequest(resttConnRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating AD Connector",
			fmt.Sprintf("Error: %v\nHTTP Response: %v", err, httpResp),
		)
		return
	}
	// Assign ID and result to the plan
	plan.ID = types.StringValue("test-connection-" + plan.ConnectionName.ValueString())

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)

	// Set the individual fields
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)
	resultObj := map[string]string{
		"msg":        msgValue,
		"error_code": errorCodeValue,
	}
	resultJSON, err := util.MarshalDeterministic(resultObj)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Marshaling Result",
			fmt.Sprintf("Could not marshal API response: %v", err),
		)
		return
	}
	plan.Result = types.StringValue(string(resultJSON))
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *restConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// If the API does not support a separate read operation, you can pass through the state.
}

func (r *restConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RESTConnectorResourceModel

	// Extract plan from request
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var connJSON map[string]interface{}
	if !plan.ConnectionJSON.IsNull() && plan.ConnectionJSON.ValueString() != "" {
		err := json.Unmarshal([]byte(plan.ConnectionJSON.ValueString()), &connJSON)
		if err != nil {
			resp.Diagnostics.AddError("Invalid JSON", fmt.Sprintf("Failed to parse connection_json: %v", err))
			return
		}
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := r.client.APIBaseURL()
	if strings.HasPrefix(apiBaseURL, "https://") {
		apiBaseURL = strings.TrimPrefix(apiBaseURL, "https://")
	}
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	restConn := openapi.RESTConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "REST",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.StringPtr(plan.Description.ValueString()),
			Defaultsavroles:    util.StringPtr(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.StringPtr(plan.EmailTemplate.ValueString()),
			SslCertificate:     util.StringPtr(plan.SSLCertificate.ValueString()),
			VaultConnection:    util.StringPtr(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.StringPtr(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.StringPtr(plan.SaveInVault.ValueString()),
		},
		ConnectionJSON:          connJSON,
		ImportUserJSON:          util.StringPtr(plan.ImportUserJson.ValueString()),
		ImportAccountEntJSON:    util.StringPtr(plan.ImportAccountEntJson.ValueString()),
		STATUS_THRESHOLD_CONFIG: util.StringPtr(plan.StatusThresholdConfig.ValueString()),
		CreateAccountJSON:       util.StringPtr(plan.CreateAccountJson.ValueString()),
		UpdateAccountJSON:       util.StringPtr(plan.UpdateAccountJson.ValueString()),
		EnableAccountJSON:       util.StringPtr(plan.EnableAccountJson.ValueString()),
		DisableAccountJSON:      util.StringPtr(plan.DisableAccountJson.ValueString()),
		AddAccessJSON:           util.StringPtr(plan.AddAccessJson.ValueString()),
		RemoveAccessJSON:        util.StringPtr(plan.RemoveAccessJson.ValueString()),
		UpdateUserJSON:          util.StringPtr(plan.UpdateUserJson.ValueString()),
		ChangePassJSON:          util.StringPtr(plan.ChangePassJson.ValueString()),
		RemoveAccountJSON:       util.StringPtr(plan.RemoveAccountJson.ValueString()),
		TicketStatusJSON:        util.StringPtr(plan.TicketStatusJson.ValueString()),
		CreateTicketJSON:        util.StringPtr(plan.CreateTicketJson.ValueString()),
		ENDPOINTS_FILTER:        util.StringPtr(plan.EndpointsFilter.ValueString()),
		PasswdPolicyJSON:        util.StringPtr(plan.PasswdPolicyJson.ValueString()),
		ConfigJSON:              util.StringPtr(plan.ConfigJSON.ValueString()),
		AddFFIDAccessJSON:       util.StringPtr(plan.AddFFIDAccessJson.ValueString()),
		RemoveFFIDAccessJSON:    util.StringPtr(plan.RemoveFFIDAccessJson.ValueString()),
		MODIFYUSERDATAJSON:      util.StringPtr(plan.ModifyUserdataJson.ValueString()),
		SendOtpJSON:             util.StringPtr(plan.SendOtpJson.ValueString()),
		ValidateOtpJSON:         util.StringPtr(plan.ValidateOtpJson.ValueString()),
		PAM_CONFIG:              util.StringPtr(plan.PamConfig.ValueString()),
	}
	resttConnRequest := openapi.TestConnectionRequest{
		RESTConnector: &restConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, httpResp, err := apiClient.ConnectionsAPI.TestConnection(ctx).TestConnectionRequest(resttConnRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating AD Connector",
			fmt.Sprintf("Error: %v\nHTTP Response: %v", err, httpResp),
		)
		return
	}
	// Assign ID and result to the plan
	plan.ID = types.StringValue("test-connection-" + plan.ConnectionName.ValueString())

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)

	// Set the individual fields
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)
	resultObj := map[string]string{
		"msg":        msgValue,
		"error_code": errorCodeValue,
	}
	resultJSON, err := util.MarshalDeterministic(resultObj)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Marshaling Result",
			fmt.Sprintf("Could not marshal API response: %v", err),
		)
		return
	}
	plan.Result = types.StringValue(string(resultJSON))

	// Store state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *restConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
