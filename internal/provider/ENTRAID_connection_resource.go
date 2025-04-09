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

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	s "github.com/saviynt/saviynt-api-go-client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ENTRAIDConnectorResourceModel struct {
	BaseConnector
	ID                              types.String `tfsdk:"id"`
	ClientId                        types.String `tfsdk:"client_id"`
	ClientSecret                    types.String `tfsdk:"client_secret"`
	AccessToken                     types.String `tfsdk:"access_token"`
	AadTenantId                     types.String `tfsdk:"aad_tenant_id"`
	AzureMgmtAccessToken            types.String `tfsdk:"azure_mgmt_access_token"`
	AuthenticationEndpoint          types.String `tfsdk:"authentication_endpoint"`
	MicrosoftGraphEndpoint          types.String `tfsdk:"microsoft_graph_endpoint"`
	AzureManagementEndpoint         types.String `tfsdk:"azure_management_endpoint"`
	ImportUserJson                  types.String `tfsdk:"import_user_json"`
	CreateUsers                     types.String `tfsdk:"create_users"`
	WindowsConnectorJson            types.String `tfsdk:"windows_connector_json"`
	CreateNewEndpoints              types.String `tfsdk:"create_new_endpoints"`
	ManagedAccountType              types.String `tfsdk:"managed_account_type"`
	AccountAttributes               types.String `tfsdk:"account_attributes"`
	ServiceAccountAttributes        types.String `tfsdk:"service_account_attributes"`
	DeltaTokensJson                 types.String `tfsdk:"delta_tokens_json"`
	AccountImportFields             types.String `tfsdk:"account_import_fields"`
	ImportDepth                     types.String `tfsdk:"import_depth"`
	EntitlementAttribute            types.String `tfsdk:"entitlement_attribute"`
	CreateAccountJson               types.String `tfsdk:"create_account_json"`
	UpdateAccountJson               types.String `tfsdk:"update_account_json"`
	EnableAccountJson               types.String `tfsdk:"enable_account_json"`
	DisableAccountJson              types.String `tfsdk:"disable_account_json"`
	AddAccessJson                   types.String `tfsdk:"add_access_json"`
	RemoveAccessJson                types.String `tfsdk:"remove_access_json"`
	UpdateUserJson                  types.String `tfsdk:"update_user_json"`
	ChangePassJson                  types.String `tfsdk:"change_pass_json"`
	RemoveAccountJson               types.String `tfsdk:"remove_account_json"`
	ConnectionJson                  types.String `tfsdk:"connection_json"`
	CreateGroupJson                 types.String `tfsdk:"create_group_json"`
	UpdateGroupJson                 types.String `tfsdk:"update_group_json"`
	AddAccessToEntitlementJson      types.String `tfsdk:"add_access_to_entitlement_json"`
	RemoveAccessFromEntitlementJson types.String `tfsdk:"remove_access_from_entitlement_json"`
	DeleteGroupJson                 types.String `tfsdk:"delete_group_json"`
	CreateServicePrincipalJson      types.String `tfsdk:"create_service_principal_json"`
	UpdateServicePrincipalJson      types.String `tfsdk:"update_service_principal_json"`
	RemoveServicePrincipalJson      types.String `tfsdk:"remove_service_principal_json"`
	EntitlementFilterJson           types.String `tfsdk:"entitlement_filter_json"`
	CreateTeamJson                  types.String `tfsdk:"create_team_json"`
	CreateChannelJson               types.String `tfsdk:"create_channel_json"`
	StatusThresholdConfig           types.String `tfsdk:"status_threshold_config"`
	AccountsFilter                  types.String `tfsdk:"accounts_filter"`
	PamConfig                       types.String `tfsdk:"pam_config"`
	EndpointsFilter                 types.String `tfsdk:"endpoints_filter"`
	ConfigJson                      types.String `tfsdk:"config_json"`
	ModifyUserdataJson              types.String `tfsdk:"modify_user_data_json"`
	EnhancedDirectoryRoles          types.String `tfsdk:"enhanced_directory_roles"`
}

type entraidConnectionResource struct {
	client *s.Client
	token  string
}

func ENTRAIDNewTestConnectionResource() resource.Resource {
	return &entraidConnectionResource{}
}

func (r *entraidConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_entraid_connection_resource"
}

func (r *entraidConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"client_id": schema.StringAttribute{
				Optional:    true,
				Description: "Client ID for authentication.",
			},
			"client_secret": schema.StringAttribute{
				Optional:    true,
				Description: "Client Secret for authentication.",
			},
			"access_token": schema.StringAttribute{
				Optional:    true,
				Description: "Access token used for API calls.",
			},
			"aad_tenant_id": schema.StringAttribute{
				Optional:    true,
				Description: "Azure Active Directory tenant ID.",
			},
			"azure_mgmt_access_token": schema.StringAttribute{
				Optional:    true,
				Description: "Access token for Azure management APIs.",
			},
			"authentication_endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "Authentication endpoint URL.",
			},
			"microsoft_graph_endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "Microsoft Graph API endpoint.",
			},
			"azure_management_endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "Azure management endpoint URL.",
			},
			"import_user_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON configuration for importing users.",
			},
			"create_users": schema.StringAttribute{
				Optional:    true,
				Description: "Flag or configuration for creating users.",
			},
			"windows_connector_json": schema.StringAttribute{
				Optional:    true,
				Description: "Windows connector JSON configuration.",
			},
			"create_new_endpoints": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration to create new endpoints.",
			},
			"managed_account_type": schema.StringAttribute{
				Optional:    true,
				Description: "Type of managed accounts.",
			},
			"account_attributes": schema.StringAttribute{
				Optional:    true,
				Description: "Attributes for account configuration.",
			},
			"service_account_attributes": schema.StringAttribute{
				Optional:    true,
				Description: "Attributes for service account configuration.",
			},
			"delta_tokens_json": schema.StringAttribute{
				Optional:    true,
				Description: "Delta tokens JSON data.",
			},
			"account_import_fields": schema.StringAttribute{
				Optional:    true,
				Description: "Fields to import for accounts.",
			},
			"import_depth": schema.StringAttribute{
				Optional:    true,
				Description: "Depth level for import.",
			},
			"entitlement_attribute": schema.StringAttribute{
				Optional:    true,
				Description: "Attribute used for entitlement.",
			},
			"create_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON template to create an account.",
			},
			"update_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON template to update an account.",
			},
			"enable_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON template to enable an account.",
			},
			"disable_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON template to disable an account.",
			},
			"add_access_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON template to add access.",
			},
			"remove_access_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON template to remove access.",
			},
			"update_user_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON template to update user.",
			},
			"change_pass_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON template to change password.",
			},
			"remove_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON template to remove account.",
			},
			"connection_json": schema.StringAttribute{
				Optional:    true,
				Description: "Connection JSON configuration.",
			},
			"create_group_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to create group.",
			},
			"update_group_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to update group.",
			},
			"add_access_to_entitlement_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to add access to entitlement.",
			},
			"remove_access_from_entitlement_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to remove access from entitlement.",
			},
			"delete_group_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to delete group.",
			},
			"create_service_principal_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to create service principal.",
			},
			"update_service_principal_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to update service principal.",
			},
			"remove_service_principal_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to remove service principal.",
			},
			"entitlement_filter_json": schema.StringAttribute{
				Optional:    true,
				Description: "Filter JSON for entitlements.",
			},
			"create_team_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to create team.",
			},
			"create_channel_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to create channel.",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration for status thresholds.",
			},
			"accounts_filter": schema.StringAttribute{
				Optional:    true,
				Description: "Filter for accounts.",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Description: "PAM configuration.",
			},
			"endpoints_filter": schema.StringAttribute{
				Optional:    true,
				Description: "Endpoints filter configuration.",
			},
			"config_json": schema.StringAttribute{
				Optional:    true,
				Description: "Main config JSON.",
			},
			"modify_user_data_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to modify user data.",
			},
			"enhanced_directory_roles": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration for enhanced directory roles.",
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

func (r *entraidConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *entraidConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ENTRAIDConnectorResourceModel
	// Extract plan from request
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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
	entraidConn := openapi.EntraIDConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "AzureAD",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:        util.StringPointerOrEmpty(plan.Description.ValueString()),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		//required fields
		CLIENT_ID:     plan.ClientId.ValueString(),
		CLIENT_SECRET: plan.ClientSecret.ValueString(),
		AAD_TENANT_ID: plan.AadTenantId.ValueString(),
		//optional fields
		ACCESS_TOKEN:                    util.StringPointerOrEmpty(plan.AccessToken.ValueString()),
		AZURE_MGMT_ACCESS_TOKEN:         util.StringPointerOrEmpty(plan.AzureMgmtAccessToken.ValueString()),
		AUTHENTICATION_ENDPOINT:         util.StringPointerOrEmpty(plan.AuthenticationEndpoint.ValueString()),
		MICROSOFT_GRAPH_ENDPOINT:        util.StringPointerOrEmpty(plan.MicrosoftGraphEndpoint.ValueString()),
		AZURE_MANAGEMENT_ENDPOINT:       util.StringPointerOrEmpty(plan.AzureManagementEndpoint.ValueString()),
		ImportUserJSON:                  util.StringPointerOrEmpty(plan.ImportUserJson.ValueString()),
		CREATEUSERS:                     util.StringPointerOrEmpty(plan.CreateUsers.ValueString()),
		WINDOWS_CONNECTOR_JSON:          util.StringPointerOrEmpty(plan.WindowsConnectorJson.ValueString()),
		CREATE_NEW_ENDPOINTS:            util.StringPointerOrEmpty(plan.CreateNewEndpoints.ValueString()),
		MANAGED_ACCOUNT_TYPE:            util.StringPointerOrEmpty(plan.ManagedAccountType.ValueString()),
		ACCOUNT_ATTRIBUTES:              util.StringPointerOrEmpty(plan.AccountAttributes.ValueString()),
		SERVICE_ACCOUNT_ATTRIBUTES:      util.StringPointerOrEmpty(plan.ServiceAccountAttributes.ValueString()),
		DELTATOKENSJSON:                 util.StringPointerOrEmpty(plan.DeltaTokensJson.ValueString()),
		ACCOUNT_IMPORT_FIELDS:           util.StringPointerOrEmpty(plan.AccountImportFields.ValueString()),
		IMPORT_DEPTH:                    util.StringPointerOrEmpty(plan.ImportDepth.ValueString()),
		ENTITLEMENT_ATTRIBUTE:           util.StringPointerOrEmpty(plan.EntitlementAttribute.ValueString()),
		CreateAccountJSON:               util.StringPointerOrEmpty(plan.CreateAccountJson.ValueString()),
		UpdateAccountJSON:               util.StringPointerOrEmpty(plan.UpdateAccountJson.ValueString()),
		EnableAccountJSON:               util.StringPointerOrEmpty(plan.EnableAccountJson.ValueString()),
		DisableAccountJSON:              util.StringPointerOrEmpty(plan.DisableAccountJson.ValueString()),
		AddAccessJSON:                   util.StringPointerOrEmpty(plan.AddAccessJson.ValueString()),
		RemoveAccessJSON:                util.StringPointerOrEmpty(plan.RemoveAccessJson.ValueString()),
		UpdateUserJSON:                  util.StringPointerOrEmpty(plan.UpdateUserJson.ValueString()),
		ChangePassJSON:                  util.StringPointerOrEmpty(plan.ChangePassJson.ValueString()),
		RemoveAccountJSON:               util.StringPointerOrEmpty(plan.RemoveAccountJson.ValueString()),
		ConnectionJSON:                  util.StringPointerOrEmpty(plan.ConnectionJson.ValueString()),
		CreateGroupJSON:                 util.StringPointerOrEmpty(plan.CreateGroupJson.ValueString()),
		UpdateGroupJSON:                 util.StringPointerOrEmpty(plan.UpdateGroupJson.ValueString()),
		AddAccessToEntitlementJSON:      util.StringPointerOrEmpty(plan.AddAccessToEntitlementJson.ValueString()),
		RemoveAccessFromEntitlementJSON: util.StringPointerOrEmpty(plan.RemoveAccessFromEntitlementJson.ValueString()),
		DeleteGroupJSON:                 util.StringPointerOrEmpty(plan.DeleteGroupJson.ValueString()),
		CreateServicePrincipalJSON:      util.StringPointerOrEmpty(plan.CreateServicePrincipalJson.ValueString()),
		UpdateServicePrincipalJSON:      util.StringPointerOrEmpty(plan.UpdateServicePrincipalJson.ValueString()),
		RemoveServicePrincipalJSON:      util.StringPointerOrEmpty(plan.RemoveServicePrincipalJson.ValueString()),
		ENTITLEMENT_FILTER_JSON:         util.StringPointerOrEmpty(plan.EntitlementFilterJson.ValueString()),
		CreateTeamJSON:                  util.StringPointerOrEmpty(plan.CreateTeamJson.ValueString()),
		CreateChannelJSON:               util.StringPointerOrEmpty(plan.CreateChannelJson.ValueString()),
		STATUS_THRESHOLD_CONFIG:         util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()),
		ACCOUNTS_FILTER:                 util.StringPointerOrEmpty(plan.AccountsFilter.ValueString()),
		PAM_CONFIG:                      util.StringPointerOrEmpty(plan.PamConfig.ValueString()),
		ENDPOINTS_FILTER:                util.StringPointerOrEmpty(plan.EndpointsFilter.ValueString()),
		ConfigJSON:                      util.StringPointerOrEmpty(plan.ConfigJson.ValueString()),
		MODIFYUSERDATAJSON:              util.StringPointerOrEmpty(plan.ModifyUserdataJson.ValueString()),
		ENHANCEDDIRECTORYROLES:          util.StringPointerOrEmpty(plan.EnhancedDirectoryRoles.ValueString()),
	}

	entraidConnRequest := openapi.CreateOrUpdateRequest{
		EntraIDConnector: &entraidConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(entraidConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
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

func (r *entraidConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ENTRAIDConnectorResourceModel

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
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	state.ConnectionKey = types.Int64Value(int64(*apiResp.EntraIDConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.EntraIDConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectiontype)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Emailtemplate)
	state.UpdateUserJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateUserJSON)
	state.MicrosoftGraphEndpoint = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MICROSOFT_GRAPH_ENDPOINT)
	state.EndpointsFilter = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	state.ImportUserJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ImportUserJSON)
	state.EnableAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.EnableAccountJSON)
	state.ConnectionJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionJSON)
	state.ClientId = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CLIENT_ID)
	state.DeleteGroupJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DeleteGroupJSON)
	state.ConfigJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ConfigJSON)
	state.AccessToken = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCESS_TOKEN)
	state.AddAccessJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AddAccessJSON)
	state.CreateChannelJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateChannelJSON)
	state.UpdateAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateAccountJSON)
	state.RemoveServicePrincipalJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveServicePrincipalJSON)
	state.ImportDepth = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.IMPORT_DEPTH)
	state.CreateAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateAccountJSON)
	state.PamConfig = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.UpdateServicePrincipalJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateServicePrincipalJSON)
	state.AzureManagementEndpoint = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AZURE_MANAGEMENT_ENDPOINT)
	state.EntitlementAttribute = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	state.AccountsFilter = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNTS_FILTER)
	state.WindowsConnectorJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.WINDOWS_CONNECTOR_JSON)
	state.DeltaTokensJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DELTATOKENSJSON)
	state.AzureMgmtAccessToken = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AZURE_MGMT_ACCESS_TOKEN)
	state.CreateTeamJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateTeamJSON)
	state.EnhancedDirectoryRoles = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENHANCEDDIRECTORYROLES)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.AccountImportFields = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_FIELDS)
	state.RemoveAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccountJSON)
	state.ChangePassJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ChangePassJSON)
	state.ClientSecret = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CLIENT_SECRET)
	state.EntitlementFilterJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_FILTER_JSON)
	state.ServiceAccountAttributes = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.SERVICE_ACCOUNT_ATTRIBUTES)
	state.AddAccessToEntitlementJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AddAccessToEntitlementJSON)
	state.AuthenticationEndpoint = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AUTHENTICATION_ENDPOINT)
	state.CreateServicePrincipalJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateServicePrincipalJSON)
	state.ModifyUserdataJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.RemoveAccessJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessJSON)
	state.CreateUsers = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CREATEUSERS)
	state.RemoveAccessFromEntitlementJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessFromEntitlementJSON)
	state.DisableAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DisableAccountJSON)
	state.CreateNewEndpoints = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CREATE_NEW_ENDPOINTS)
	state.ManagedAccountType = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MANAGED_ACCOUNT_TYPE)
	state.AccountAttributes = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTES)
	state.AadTenantId = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AAD_TENANT_ID)
	state.UpdateGroupJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateGroupJSON)
	state.CreateGroupJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateGroupJSON)
	apiMessage := util.SafeDeref(apiResp.EntraIDConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.EntraIDConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *entraidConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ENTRAIDConnectorResourceModel
	// Extract plan from request
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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
	entraidConn := openapi.EntraIDConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "AzureAD",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:        util.StringPointerOrEmpty(plan.Description.ValueString()),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		//required fields
		CLIENT_ID:     plan.ClientId.ValueString(),
		CLIENT_SECRET: plan.ClientSecret.ValueString(),
		AAD_TENANT_ID: plan.AadTenantId.ValueString(),
		//optional fields
		ACCESS_TOKEN:                    util.StringPointerOrEmpty(plan.AccessToken.ValueString()),
		AZURE_MGMT_ACCESS_TOKEN:         util.StringPointerOrEmpty(plan.AzureMgmtAccessToken.ValueString()),
		AUTHENTICATION_ENDPOINT:         util.StringPointerOrEmpty(plan.AuthenticationEndpoint.ValueString()),
		MICROSOFT_GRAPH_ENDPOINT:        util.StringPointerOrEmpty(plan.MicrosoftGraphEndpoint.ValueString()),
		AZURE_MANAGEMENT_ENDPOINT:       util.StringPointerOrEmpty(plan.AzureManagementEndpoint.ValueString()),
		ImportUserJSON:                  util.StringPointerOrEmpty(plan.ImportUserJson.ValueString()),
		CREATEUSERS:                     util.StringPointerOrEmpty(plan.CreateUsers.ValueString()),
		WINDOWS_CONNECTOR_JSON:          util.StringPointerOrEmpty(plan.WindowsConnectorJson.ValueString()),
		CREATE_NEW_ENDPOINTS:            util.StringPointerOrEmpty(plan.CreateNewEndpoints.ValueString()),
		MANAGED_ACCOUNT_TYPE:            util.StringPointerOrEmpty(plan.ManagedAccountType.ValueString()),
		ACCOUNT_ATTRIBUTES:              util.StringPointerOrEmpty(plan.AccountAttributes.ValueString()),
		SERVICE_ACCOUNT_ATTRIBUTES:      util.StringPointerOrEmpty(plan.ServiceAccountAttributes.ValueString()),
		DELTATOKENSJSON:                 util.StringPointerOrEmpty(plan.DeltaTokensJson.ValueString()),
		ACCOUNT_IMPORT_FIELDS:           util.StringPointerOrEmpty(plan.AccountImportFields.ValueString()),
		IMPORT_DEPTH:                    util.StringPointerOrEmpty(plan.ImportDepth.ValueString()),
		ENTITLEMENT_ATTRIBUTE:           util.StringPointerOrEmpty(plan.EntitlementAttribute.ValueString()),
		CreateAccountJSON:               util.StringPointerOrEmpty(plan.CreateAccountJson.ValueString()),
		UpdateAccountJSON:               util.StringPointerOrEmpty(plan.UpdateAccountJson.ValueString()),
		EnableAccountJSON:               util.StringPointerOrEmpty(plan.EnableAccountJson.ValueString()),
		DisableAccountJSON:              util.StringPointerOrEmpty(plan.DisableAccountJson.ValueString()),
		AddAccessJSON:                   util.StringPointerOrEmpty(plan.AddAccessJson.ValueString()),
		RemoveAccessJSON:                util.StringPointerOrEmpty(plan.RemoveAccessJson.ValueString()),
		UpdateUserJSON:                  util.StringPointerOrEmpty(plan.UpdateUserJson.ValueString()),
		ChangePassJSON:                  util.StringPointerOrEmpty(plan.ChangePassJson.ValueString()),
		RemoveAccountJSON:               util.StringPointerOrEmpty(plan.RemoveAccountJson.ValueString()),
		ConnectionJSON:                  util.StringPointerOrEmpty(plan.ConnectionJson.ValueString()),
		CreateGroupJSON:                 util.StringPointerOrEmpty(plan.CreateGroupJson.ValueString()),
		UpdateGroupJSON:                 util.StringPointerOrEmpty(plan.UpdateGroupJson.ValueString()),
		AddAccessToEntitlementJSON:      util.StringPointerOrEmpty(plan.AddAccessToEntitlementJson.ValueString()),
		RemoveAccessFromEntitlementJSON: util.StringPointerOrEmpty(plan.RemoveAccessFromEntitlementJson.ValueString()),
		DeleteGroupJSON:                 util.StringPointerOrEmpty(plan.DeleteGroupJson.ValueString()),
		CreateServicePrincipalJSON:      util.StringPointerOrEmpty(plan.CreateServicePrincipalJson.ValueString()),
		UpdateServicePrincipalJSON:      util.StringPointerOrEmpty(plan.UpdateServicePrincipalJson.ValueString()),
		RemoveServicePrincipalJSON:      util.StringPointerOrEmpty(plan.RemoveServicePrincipalJson.ValueString()),
		ENTITLEMENT_FILTER_JSON:         util.StringPointerOrEmpty(plan.EntitlementFilterJson.ValueString()),
		CreateTeamJSON:                  util.StringPointerOrEmpty(plan.CreateTeamJson.ValueString()),
		CreateChannelJSON:               util.StringPointerOrEmpty(plan.CreateChannelJson.ValueString()),
		STATUS_THRESHOLD_CONFIG:         util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()),
		ACCOUNTS_FILTER:                 util.StringPointerOrEmpty(plan.AccountsFilter.ValueString()),
		PAM_CONFIG:                      util.StringPointerOrEmpty(plan.PamConfig.ValueString()),
		ENDPOINTS_FILTER:                util.StringPointerOrEmpty(plan.EndpointsFilter.ValueString()),
		ConfigJSON:                      util.StringPointerOrEmpty(plan.ConfigJson.ValueString()),
		MODIFYUSERDATAJSON:              util.StringPointerOrEmpty(plan.ModifyUserdataJson.ValueString()),
		ENHANCEDDIRECTORYROLES:          util.StringPointerOrEmpty(plan.EnhancedDirectoryRoles.ValueString()),
	}

	entraidConnRequest := openapi.CreateOrUpdateRequest{
		EntraIDConnector: &entraidConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(entraidConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("Problem with the update function")
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	reqParams := openapi.GetConnectionDetailsRequest{}

	reqParams.SetConnectionname(plan.ConnectionName.ValueString())
	getResp, _, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in update block")
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	plan.ConnectionKey = types.Int64Value(int64(*getResp.EntraIDConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.EntraIDConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionname)
	plan.ConnectionKey = util.SafeInt64(getResp.EntraIDConnectionResponse.Connectionkey)
	plan.Description = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Defaultsavroles)
	plan.ConnectionType = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectiontype)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Emailtemplate)
	plan.UpdateUserJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.UpdateUserJSON)
	plan.MicrosoftGraphEndpoint = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.MICROSOFT_GRAPH_ENDPOINT)
	plan.EndpointsFilter = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	plan.ImportUserJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ImportUserJSON)
	plan.EnableAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.EnableAccountJSON)
	plan.ConnectionJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ConnectionJSON)
	plan.ClientId = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CLIENT_ID)
	plan.DeleteGroupJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.DeleteGroupJSON)
	plan.ConfigJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ConfigJSON)
	plan.AccessToken = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ACCESS_TOKEN)
	plan.AddAccessJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AddAccessJSON)
	plan.CreateChannelJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateChannelJSON)
	plan.UpdateAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.UpdateAccountJSON)
	plan.RemoveServicePrincipalJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.RemoveServicePrincipalJSON)
	plan.ImportDepth = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.IMPORT_DEPTH)
	plan.CreateAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateAccountJSON)
	plan.PamConfig = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.PAM_CONFIG)
	plan.UpdateServicePrincipalJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.UpdateServicePrincipalJSON)
	plan.AzureManagementEndpoint = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AZURE_MANAGEMENT_ENDPOINT)
	plan.EntitlementAttribute = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	plan.AccountsFilter = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNTS_FILTER)
	plan.WindowsConnectorJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.WINDOWS_CONNECTOR_JSON)
	plan.DeltaTokensJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.DELTATOKENSJSON)
	plan.AzureMgmtAccessToken = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AZURE_MGMT_ACCESS_TOKEN)
	plan.CreateTeamJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateTeamJSON)
	plan.EnhancedDirectoryRoles = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ENHANCEDDIRECTORYROLES)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.AccountImportFields = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_FIELDS)
	plan.RemoveAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccountJSON)
	plan.ChangePassJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ChangePassJSON)
	plan.ClientSecret = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CLIENT_SECRET)
	plan.EntitlementFilterJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_FILTER_JSON)
	plan.ServiceAccountAttributes = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.SERVICE_ACCOUNT_ATTRIBUTES)
	plan.AddAccessToEntitlementJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AddAccessToEntitlementJSON)
	plan.AuthenticationEndpoint = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AUTHENTICATION_ENDPOINT)
	plan.CreateServicePrincipalJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateServicePrincipalJSON)
	plan.ModifyUserdataJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	plan.RemoveAccessJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessJSON)
	plan.CreateUsers = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CREATEUSERS)
	plan.RemoveAccessFromEntitlementJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessFromEntitlementJSON)
	plan.DisableAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.DisableAccountJSON)
	plan.CreateNewEndpoints = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CREATE_NEW_ENDPOINTS)
	plan.ManagedAccountType = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.MANAGED_ACCOUNT_TYPE)
	plan.AccountAttributes = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTES)
	plan.AadTenantId = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AAD_TENANT_ID)
	plan.UpdateGroupJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.UpdateGroupJSON)
	plan.CreateGroupJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateGroupJSON)
	apiMessage := util.SafeDeref(getResp.EntraIDConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.EntraIDConnectionResponse.Errorcode)
}
func (r *entraidConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
