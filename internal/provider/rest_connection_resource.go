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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RestConnectorResourceModel struct {
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

type restConnectionResource struct {
	client *s.Client
	token  string
}

func NewRestTestConnectionResource() resource.Resource {
	return &restConnectionResource{}
}

func (r *restConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_rest_connection_resource"
}

func (r *restConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.RestConnDescription,
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
				Optional:    true,
				Computed:    true,
				Description: "Connection type (e.g., 'AD' for Active Directory). Example: \"AD\"",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Description for the connection. Example: \"ORG_AD\"",
			},
			"defaultsavroles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default SAV roles for managing the connection. Example: \"ROLE_ORG\"",
			},
			"email_template": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "JSON for importing users.",
			},
			"import_account_ent_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON for importing accounts and entitlements.",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON configuration for status thresholds.",
			},
			"create_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to create an account.",
			},
			"update_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to update an account.",
			},
			"enable_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON configuration to enable an account.",
			},
			"disable_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON configuration to disable an account.",
			},
			"add_access_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to add access.",
			},
			"remove_access_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to remove access.",
			},
			"update_user_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to update a user.",
			},
			"change_pass_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to change a user’s password.",
			},
			"remove_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to remove an account.",
			},
			"ticket_status_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to check ticket status.",
			},
			"create_ticket_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to create a ticket.",
			},
			"endpoints_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Filter criteria for endpoints.",
			},
			"passwd_policy_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON defining the password policy.",
			},
			"config_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "General configuration JSON for the REST connector.",
			},
			"add_ffid_access_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to add FFID access.",
			},
			"remove_ffid_access_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to remove FFID access.",
			},
			"modify_user_data_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON for modifying user data.",
			},
			"send_otp_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to send OTP.",
			},
			"validate_otp_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to validate OTP.",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PAM configuration JSON.",
			},
			"msg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Message returned from the operation.",
			},
			"error_code": schema.StringAttribute{
				Optional:    true,
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
	var plan RestConnectorResourceModel
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
	apiClient := openapi.NewAPIClient(cfg)
	reqParams := openapi.GetConnectionDetailsRequest{}
	reqParams.SetConnectionname(plan.ConnectionName.ValueString())
	existingResource, _, _ := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if existingResource != nil &&
		existingResource.RESTConnectionResponse != nil &&
		existingResource.RESTConnectionResponse.Errorcode != nil &&
		*existingResource.RESTConnectionResponse.Errorcode == 0 {
		log.Printf("[ERROR] Connection name already exists. Please import or use a different name")
		resp.Diagnostics.AddError("API Create Failed", "Connection name already exists. Please import or use a different name")
		return
	}

	restConn := openapi.RESTConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "REST",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:        util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate),
			VaultConnection:    util.StringPointerOrEmpty(plan.VaultConnection),
			VaultConfiguration: util.StringPointerOrEmpty(plan.VaultConfiguration),
			Saveinvault:        util.StringPointerOrEmpty(plan.SaveInVault),
		},
		//optional fields
		ConnectionJSON:          connJSON,
		ImportUserJSON:          util.StringPointerOrEmpty(plan.ImportUserJson),
		ImportAccountEntJSON:    util.StringPointerOrEmpty(plan.ImportAccountEntJson),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CreateAccountJSON:       util.StringPointerOrEmpty(plan.CreateAccountJson),
		UpdateAccountJSON:       util.StringPointerOrEmpty(plan.UpdateAccountJson),
		EnableAccountJSON:       util.StringPointerOrEmpty(plan.EnableAccountJson),
		DisableAccountJSON:      util.StringPointerOrEmpty(plan.DisableAccountJson),
		AddAccessJSON:           util.StringPointerOrEmpty(plan.AddAccessJson),
		RemoveAccessJSON:        util.StringPointerOrEmpty(plan.RemoveAccessJson),
		UpdateUserJSON:          util.StringPointerOrEmpty(plan.UpdateUserJson),
		ChangePassJSON:          util.StringPointerOrEmpty(plan.ChangePassJson),
		RemoveAccountJSON:       util.StringPointerOrEmpty(plan.RemoveAccountJson),
		TicketStatusJSON:        util.StringPointerOrEmpty(plan.TicketStatusJson),
		CreateTicketJSON:        util.StringPointerOrEmpty(plan.CreateTicketJson),
		ENDPOINTS_FILTER:        util.StringPointerOrEmpty(plan.EndpointsFilter),
		PasswdPolicyJSON:        util.StringPointerOrEmpty(plan.PasswdPolicyJson),
		ConfigJSON:              util.StringPointerOrEmpty(plan.ConfigJSON),
		AddFFIDAccessJSON:       util.StringPointerOrEmpty(plan.AddFFIDAccessJson),
		RemoveFFIDAccessJSON:    util.StringPointerOrEmpty(plan.RemoveFFIDAccessJson),
		MODIFYUSERDATAJSON:      util.StringPointerOrEmpty(plan.ModifyUserdataJson),
		SendOtpJSON:             util.StringPointerOrEmpty(plan.SendOtpJson),
		ValidateOtpJSON:         util.StringPointerOrEmpty(plan.ValidateOtpJson),
		PAM_CONFIG:              util.StringPointerOrEmpty(plan.PamConfig),
	}
	restConnRequest := openapi.CreateOrUpdateRequest{
		RESTConnector: &restConn,
	}

	// Initialize API client
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(restConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", *apiResp.Msg)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", *apiResp.Msg))
		return
	}
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionType = types.StringValue("REST")
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ImportUserJson = util.SafeStringDatasource(plan.ImportUserJson.ValueStringPointer())
	plan.ImportAccountEntJson = util.SafeStringDatasource(plan.ImportAccountEntJson.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.CreateAccountJson = util.SafeStringDatasource(plan.CreateAccountJson.ValueStringPointer())
	plan.UpdateAccountJson = util.SafeStringDatasource(plan.UpdateAccountJson.ValueStringPointer())
	plan.EnableAccountJson = util.SafeStringDatasource(plan.EnableAccountJson.ValueStringPointer())
	plan.DisableAccountJson = util.SafeStringDatasource(plan.DisableAccountJson.ValueStringPointer())
	plan.AddAccessJson = util.SafeStringDatasource(plan.AddAccessJson.ValueStringPointer())
	plan.RemoveAccessJson = util.SafeStringDatasource(plan.RemoveAccessJson.ValueStringPointer())
	plan.UpdateUserJson = util.SafeStringDatasource(plan.UpdateUserJson.ValueStringPointer())
	plan.ChangePassJson = util.SafeStringDatasource(plan.ChangePassJson.ValueStringPointer())
	plan.RemoveAccountJson = util.SafeStringDatasource(plan.RemoveAccountJson.ValueStringPointer())
	plan.TicketStatusJson = util.SafeStringDatasource(plan.TicketStatusJson.ValueStringPointer())
	plan.CreateTicketJson = util.SafeStringDatasource(plan.CreateTicketJson.ValueStringPointer())
	plan.EndpointsFilter = util.SafeStringDatasource(plan.EndpointsFilter.ValueStringPointer())
	plan.PasswdPolicyJson = util.SafeStringDatasource(plan.PasswdPolicyJson.ValueStringPointer())
	plan.ConfigJSON = util.SafeStringDatasource(plan.ConfigJSON.ValueStringPointer())
	plan.AddFFIDAccessJson = util.SafeStringDatasource(plan.AddFFIDAccessJson.ValueStringPointer())
	plan.RemoveFFIDAccessJson = util.SafeStringDatasource(plan.RemoveFFIDAccessJson.ValueStringPointer())
	plan.ModifyUserdataJson = util.SafeStringDatasource(plan.ModifyUserdataJson.ValueStringPointer())
	plan.SendOtpJson = util.SafeStringDatasource(plan.SendOtpJson.ValueStringPointer())
	plan.ValidateOtpJson = util.SafeStringDatasource(plan.ValidateOtpJson.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	r.Read(ctx, resource.ReadRequest{State: resp.State}, &resource.ReadResponse{State: resp.State})
}

func (r *restConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RestConnectorResourceModel

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
	apiResp, _, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in read block")
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", *apiResp.RESTConnectionResponse.Msg))
		return
	}
	state.ConnectionKey = types.Int64Value(int64(*apiResp.RESTConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.RESTConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectiontype)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Emailtemplate)
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
	if resp.Diagnostics.HasError() {
		return
	}
}
func (r *restConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RestConnectorResourceModel
	var state RestConnectorResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.ConnectionName.ValueString() != state.ConnectionName.ValueString() {
		resp.Diagnostics.AddError("Error", "Connection name cannot be updated")
		log.Printf("[ERROR]: Connection name cannot be updated")
		return
	}
	if plan.ConnectionType.ValueString() != state.ConnectionType.ValueString() {
		resp.Diagnostics.AddError("Error", "Connection type cannot by updated")
		log.Printf("[ERROR]: Connection type cannot by updated")
		return
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

	//for connJson data conversion from string to map[string]interface{}
	var connJSON map[string]interface{}
	if !plan.ConnectionJSON.IsNull() && plan.ConnectionJSON.ValueString() != "" {
		err := json.Unmarshal([]byte(plan.ConnectionJSON.ValueString()), &connJSON)
		if err != nil {
			resp.Diagnostics.AddError("Invalid JSON", fmt.Sprintf("Failed to parse connection_json: %v", err))
			return
		}
	}

	restConn := openapi.RESTConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "REST",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:        util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate),
			VaultConnection:    util.StringPointerOrEmpty(plan.VaultConnection),
			VaultConfiguration: util.StringPointerOrEmpty(plan.VaultConfiguration),
			Saveinvault:        util.StringPointerOrEmpty(plan.SaveInVault),
		},
		//optional fields
		ConnectionJSON:          connJSON,
		ImportUserJSON:          util.StringPointerOrEmpty(plan.ImportUserJson),
		ImportAccountEntJSON:    util.StringPointerOrEmpty(plan.ImportAccountEntJson),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CreateAccountJSON:       util.StringPointerOrEmpty(plan.CreateAccountJson),
		UpdateAccountJSON:       util.StringPointerOrEmpty(plan.UpdateAccountJson),
		EnableAccountJSON:       util.StringPointerOrEmpty(plan.EnableAccountJson),
		DisableAccountJSON:      util.StringPointerOrEmpty(plan.DisableAccountJson),
		AddAccessJSON:           util.StringPointerOrEmpty(plan.AddAccessJson),
		RemoveAccessJSON:        util.StringPointerOrEmpty(plan.RemoveAccessJson),
		UpdateUserJSON:          util.StringPointerOrEmpty(plan.UpdateUserJson),
		ChangePassJSON:          util.StringPointerOrEmpty(plan.ChangePassJson),
		RemoveAccountJSON:       util.StringPointerOrEmpty(plan.RemoveAccountJson),
		TicketStatusJSON:        util.StringPointerOrEmpty(plan.TicketStatusJson),
		CreateTicketJSON:        util.StringPointerOrEmpty(plan.CreateTicketJson),
		ENDPOINTS_FILTER:        util.StringPointerOrEmpty(plan.EndpointsFilter),
		PasswdPolicyJSON:        util.StringPointerOrEmpty(plan.PasswdPolicyJson),
		ConfigJSON:              util.StringPointerOrEmpty(plan.ConfigJSON),
		AddFFIDAccessJSON:       util.StringPointerOrEmpty(plan.AddFFIDAccessJson),
		RemoveFFIDAccessJSON:    util.StringPointerOrEmpty(plan.RemoveFFIDAccessJson),
		MODIFYUSERDATAJSON:      util.StringPointerOrEmpty(plan.ModifyUserdataJson),
		SendOtpJSON:             util.StringPointerOrEmpty(plan.SendOtpJson),
		ValidateOtpJSON:         util.StringPointerOrEmpty(plan.ValidateOtpJson),
		PAM_CONFIG:              util.StringPointerOrEmpty(plan.PamConfig),
	}
	restConnRequest := openapi.CreateOrUpdateRequest{
		RESTConnector: &restConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(restConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("Problem with the update function")
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", *apiResp.Msg))
		return
	}
	reqParams := openapi.GetConnectionDetailsRequest{}

	reqParams.SetConnectionname(plan.ConnectionName.ValueString())
	getResp, _, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in update block")
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", *getResp.RESTConnectionResponse.Msg))
		return
	}
	plan.ConnectionKey = types.Int64Value(int64(*getResp.RESTConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.RESTConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionname)
	plan.Description = util.SafeStringDatasource(getResp.RESTConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.RESTConnectionResponse.Defaultsavroles)
	plan.ConnectionType = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectiontype)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.RESTConnectionResponse.Emailtemplate)
	plan.ImportUserJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.ImportUserJSON)
	plan.ImportAccountEntJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.ImportAccountEntJSON)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.CreateAccountJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.CreateAccountJSON)
	plan.UpdateAccountJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.UpdateAccountJSON)
	plan.EnableAccountJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.EnableAccountJSON)
	plan.DisableAccountJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.DisableAccountJSON)
	plan.AddAccessJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.AddAccessJSON)
	plan.RemoveAccessJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.RemoveAccessJSON)
	plan.UpdateUserJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.UpdateUserJSON)
	plan.ChangePassJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.ChangePassJSON)
	plan.RemoveAccountJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.RemoveAccountJSON)
	plan.TicketStatusJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.TicketStatusJSON)
	plan.CreateTicketJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.CreateTicketJSON)
	plan.EndpointsFilter = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	plan.PasswdPolicyJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.PasswdPolicyJSON)
	plan.ConfigJSON = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.ConfigJSON)
	plan.AddFFIDAccessJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.AddFFIDAccessJSON)
	plan.RemoveFFIDAccessJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.RemoveFFIDAccessJSON)
	plan.ModifyUserdataJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	plan.SendOtpJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.SendOtpJSON)
	plan.ValidateOtpJson = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.ValidateOtpJSON)
	plan.PamConfig = util.SafeStringDatasource(getResp.RESTConnectionResponse.Connectionattributes.PAM_CONFIG)
	apiMessage := util.SafeDeref(getResp.RESTConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.RESTConnectionResponse.Errorcode)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *restConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
func (r *restConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)
}
