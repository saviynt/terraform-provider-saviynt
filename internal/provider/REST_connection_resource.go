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

// restConnectionResource implements the resource.Resource interface.
type restConnectionResource struct {
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
			"connection_key": schema.Int64Attribute{
				Computed:    true,
				Description: "Unique identifier of the connection returned by the API. Example: 1909",
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

func (r *restConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RESTConnectorResourceModel

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	//for connJson data conversion from string to map[string]interface{}
	var connJSON map[string]interface{}
	err := json.Unmarshal([]byte(plan.ConnectionJSON.ValueString()), &connJSON)
	if err != nil {
		log.Fatalf("Failed to unmarshal ConnectionJSON: %v", err)
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	restConn := openapi.RESTConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "REST",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.SafeStringConnectorForNullHandling(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnectorForNullHandling(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnectorForNullHandling(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		ConnectionJSON:          connJSON,
		ImportUserJSON:          util.SafeStringConnectorForNullHandling(plan.ImportUserJson.ValueString()),
		ImportAccountEntJSON:    util.SafeStringConnectorForNullHandling(plan.ImportAccountEntJson.ValueString()),
		STATUS_THRESHOLD_CONFIG: util.SafeStringConnectorForNullHandling(plan.StatusThresholdConfig.ValueString()),
		CreateAccountJSON:       util.SafeStringConnectorForNullHandling(plan.CreateAccountJson.ValueString()),
		UpdateAccountJSON:       util.SafeStringConnectorForNullHandling(plan.UpdateAccountJson.ValueString()),
		EnableAccountJSON:       util.SafeStringConnectorForNullHandling(plan.EnableAccountJson.ValueString()),
		DisableAccountJSON:      util.SafeStringConnectorForNullHandling(plan.DisableAccountJson.ValueString()),
		AddAccessJSON:           util.SafeStringConnectorForNullHandling(plan.AddAccessJson.ValueString()),
		RemoveAccessJSON:        util.SafeStringConnectorForNullHandling(plan.RemoveAccessJson.ValueString()),
		UpdateUserJSON:          util.SafeStringConnectorForNullHandling(plan.UpdateUserJson.ValueString()),
		ChangePassJSON:          util.SafeStringConnectorForNullHandling(plan.ChangePassJson.ValueString()),
		RemoveAccountJSON:       util.SafeStringConnectorForNullHandling(plan.RemoveAccountJson.ValueString()),
		TicketStatusJSON:        util.SafeStringConnectorForNullHandling(plan.TicketStatusJson.ValueString()),
		CreateTicketJSON:        util.SafeStringConnectorForNullHandling(plan.CreateTicketJson.ValueString()),
		ENDPOINTS_FILTER:        util.SafeStringConnectorForNullHandling(plan.EndpointsFilter.ValueString()),
		PasswdPolicyJSON:        util.SafeStringConnectorForNullHandling(plan.PasswdPolicyJson.ValueString()),
		ConfigJSON:              util.SafeStringConnectorForNullHandling(plan.ConfigJSON.ValueString()),
		AddFFIDAccessJSON:       util.SafeStringConnectorForNullHandling(plan.AddFFIDAccessJson.ValueString()),
		RemoveFFIDAccessJSON:    util.SafeStringConnectorForNullHandling(plan.RemoveFFIDAccessJson.ValueString()),
		MODIFYUSERDATAJSON:      util.SafeStringConnectorForNullHandling(plan.ModifyUserdataJson.ValueString()),
		SendOtpJSON:             util.SafeStringConnectorForNullHandling(plan.SendOtpJson.ValueString()),
		ValidateOtpJSON:         util.SafeStringConnectorForNullHandling(plan.ValidateOtpJson.ValueString()),
		PAM_CONFIG:              util.SafeStringConnectorForNullHandling(plan.PamConfig.ValueString()),
	}
	restConnRequest := openapi.CreateOrUpdateRequest{
		RESTConnector: &restConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(restConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	r.Read(ctx, resource.ReadRequest{State: resp.State}, &resource.ReadResponse{State: resp.State})
}

func (r *restConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RESTConnectorResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configure API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)
	reqParams := openapi.GetConnectionDetailsRequest{}

	reqParams.SetConnectionname(state.ConnectionName.ValueString())
	apiResp, httpResp, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function")
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)
	state.ConnectionKey = types.Int64Value(int64(*apiResp.RESTConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.RESTConnectionResponse.Connectionkey))
	state.ConnectionJSON = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ConnectionJSON)
	state.ImportUserJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ImportUserJSON)
	state.ImportAccountEntJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ImportAccountEntJSON)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.CreateAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.CreateAccountJSON)
	state.UpdateAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.UpdateAccountJSON)
	state.EnableAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.EnableAccountJSON)
	state.DisableAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.DisableAccountJSON)
	state.AddAccessJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.AddAccessJSON)
	state.RemoveAccessJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.RemoveAccessJSON)
	state.UpdateUserJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.UpdateUserJSON)
	state.ChangePassJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ChangePassJSON)
	state.RemoveAccountJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.RemoveAccountJSON)
	state.TicketStatusJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.TicketStatusJSON)
	state.CreateTicketJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.CreateTicketJSON)
	state.EndpointsFilter = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	state.PasswdPolicyJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.PasswdPolicyJSON)
	state.ConfigJSON = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ConfigJSON)
	state.AddFFIDAccessJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.AddFFIDAccessJSON)
	state.RemoveFFIDAccessJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.RemoveFFIDAccessJSON)
	state.ModifyUserdataJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.SendOtpJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.SendOtpJSON)
	state.ValidateOtpJson = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ValidateOtpJSON)
	state.PamConfig = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.PAM_CONFIG)
	apiMessage := util.SafeDeref(apiResp.RESTConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.RESTConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}
func (r *restConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RESTConnectorResourceModel

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	//for connJson data conversion from string to map[string]interface{}
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
			Description:        util.SafeStringConnectorForNullHandling(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnectorForNullHandling(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnectorForNullHandling(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		ConnectionJSON:          connJSON,
		ImportUserJSON:          util.SafeStringConnectorForNullHandling(plan.ImportUserJson.ValueString()),
		ImportAccountEntJSON:    util.SafeStringConnectorForNullHandling(plan.ImportAccountEntJson.ValueString()),
		STATUS_THRESHOLD_CONFIG: util.SafeStringConnectorForNullHandling(plan.StatusThresholdConfig.ValueString()),
		CreateAccountJSON:       util.SafeStringConnectorForNullHandling(plan.CreateAccountJson.ValueString()),
		UpdateAccountJSON:       util.SafeStringConnectorForNullHandling(plan.UpdateAccountJson.ValueString()),
		EnableAccountJSON:       util.SafeStringConnectorForNullHandling(plan.EnableAccountJson.ValueString()),
		DisableAccountJSON:      util.SafeStringConnectorForNullHandling(plan.DisableAccountJson.ValueString()),
		AddAccessJSON:           util.SafeStringConnectorForNullHandling(plan.AddAccessJson.ValueString()),
		RemoveAccessJSON:        util.SafeStringConnectorForNullHandling(plan.RemoveAccessJson.ValueString()),
		UpdateUserJSON:          util.SafeStringConnectorForNullHandling(plan.UpdateUserJson.ValueString()),
		ChangePassJSON:          util.SafeStringConnectorForNullHandling(plan.ChangePassJson.ValueString()),
		RemoveAccountJSON:       util.SafeStringConnectorForNullHandling(plan.RemoveAccountJson.ValueString()),
		TicketStatusJSON:        util.SafeStringConnectorForNullHandling(plan.TicketStatusJson.ValueString()),
		CreateTicketJSON:        util.SafeStringConnectorForNullHandling(plan.CreateTicketJson.ValueString()),
		ENDPOINTS_FILTER:        util.SafeStringConnectorForNullHandling(plan.EndpointsFilter.ValueString()),
		PasswdPolicyJSON:        util.SafeStringConnectorForNullHandling(plan.PasswdPolicyJson.ValueString()),
		ConfigJSON:              util.SafeStringConnectorForNullHandling(plan.ConfigJSON.ValueString()),
		AddFFIDAccessJSON:       util.SafeStringConnectorForNullHandling(plan.AddFFIDAccessJson.ValueString()),
		RemoveFFIDAccessJSON:    util.SafeStringConnectorForNullHandling(plan.RemoveFFIDAccessJson.ValueString()),
		MODIFYUSERDATAJSON:      util.SafeStringConnectorForNullHandling(plan.ModifyUserdataJson.ValueString()),
		SendOtpJSON:             util.SafeStringConnectorForNullHandling(plan.SendOtpJson.ValueString()),
		ValidateOtpJSON:         util.SafeStringConnectorForNullHandling(plan.ValidateOtpJson.ValueString()),
		PAM_CONFIG:              util.SafeStringConnectorForNullHandling(plan.PamConfig.ValueString()),
	}
	restConnRequest := openapi.CreateOrUpdateRequest{
		RESTConnector: &restConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(restConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("Problem with the update function")
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *restConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
