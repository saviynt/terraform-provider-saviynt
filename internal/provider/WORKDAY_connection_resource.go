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

type WORKDAYConnectorResourceModel struct {
	BaseConnector
	ID                     types.String `tfsdk:"id"`
	UsersLastImportTime    types.String `tfsdk:"users_last_import_time"`
	AccountsLastImportTime types.String `tfsdk:"accounts_last_import_time"`
	AccessLastImportTime   types.String `tfsdk:"access_last_import_time"`
	BaseURL                types.String `tfsdk:"base_url"`
	APIVersion             types.String `tfsdk:"api_version"`
	TenantName             types.String `tfsdk:"tenant_name"`
	ReportOwner            types.String `tfsdk:"report_owner"`
	UseOAuth               types.String `tfsdk:"use_oauth"`
	IncludeReferenceDesc   types.String `tfsdk:"include_reference_descriptors"`
	UseEnhancedOrgRole     types.String `tfsdk:"use_enhanced_orgrole"`
	UseX509AuthForSOAP     types.String `tfsdk:"use_x509auth_for_soap"`
	X509Key                types.String `tfsdk:"x509_key"`
	X509Cert               types.String `tfsdk:"x509_cert"`
	Username               types.String `tfsdk:"username"`
	Password               types.String `tfsdk:"password"`
	ClientID               types.String `tfsdk:"client_id"`
	ClientSecret           types.String `tfsdk:"client_secret"`
	RefreshToken           types.String `tfsdk:"refresh_token"`
	PageSize               types.String `tfsdk:"page_size"`
	UserImportPayload      types.String `tfsdk:"user_import_payload"`
	UserImportMapping      types.String `tfsdk:"user_import_mapping"`
	AccountImportPayload   types.String `tfsdk:"account_import_payload"`
	AccountImportMapping   types.String `tfsdk:"account_import_mapping"`
	AccessImportList       types.String `tfsdk:"access_import_list"`
	RAASMappingJSON        types.String `tfsdk:"raas_mapping_json"`
	AccessImportMapping    types.String `tfsdk:"access_import_mapping"`
	OrgRoleImportPayload   types.String `tfsdk:"orgrole_import_payload"`
	StatusKeyJSON          types.String `tfsdk:"status_key_json"`
	UserAttributeJSON      types.String `tfsdk:"userattributejson"`
	CustomConfig           types.String `tfsdk:"custom_config"`
	PAMConfig              types.String `tfsdk:"pam_config"`
	ModifyUserDataJSON     types.String `tfsdk:"modify_user_data_json"`
	StatusThresholdConfig  types.String `tfsdk:"status_threshold_config"`
	CreateAccountPayload   types.String `tfsdk:"create_account_payload"`
	UpdateAccountPayload   types.String `tfsdk:"update_account_payload"`
	UpdateUserPayload      types.String `tfsdk:"update_user_payload"`
	AssignOrgRolePayload   types.String `tfsdk:"assign_orgrole_payload"`
	RemoveOrgRolePayload   types.String `tfsdk:"remove_orgrole_payload"`
}

// restConnectionResource implements the resource.Resource interface.
type workdayConnectionResource struct {
	client *s.Client
	token  string
}

// NewTestConnectionResource returns a new instance of testConnectionResource.
func WorkdayNewTestConnectionResource() resource.Resource {
	return &workdayConnectionResource{}
}

func (r *workdayConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_workday_connection_resource"
}

func (r *workdayConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

			"users_last_import_time": schema.StringAttribute{
				Optional:    true,
				Description: "Property for USERS_LAST_IMPORT_TIME.",
			},
			"accounts_last_import_time": schema.StringAttribute{
				Optional:    true,
				Description: "Property for ACCOUNTS_LAST_IMPORT_TIME.",
			},
			"access_last_import_time": schema.StringAttribute{
				Optional:    true,
				Description: "Property for ACCESS_LAST_IMPORT_TIME.",
			},
			"base_url": schema.StringAttribute{
				Optional:    true,
				Description: "Base URL of the Workday tenant instance.",
			},
			"api_version": schema.StringAttribute{
				Optional:    true,
				Description: "Version of the SOAP API used for the connection.",
			},
			"tenant_name": schema.StringAttribute{
				Optional:    true,
				Description: "The name of your tenant.",
			},
			"report_owner": schema.StringAttribute{
				Optional:    true,
				Description: "Account name of the report owner used to build default RaaS URLs.",
			},
			"use_oauth": schema.StringAttribute{
				Required:    true,
				Description: "Whether to use OAuth authentication.",
			},
			"include_reference_descriptors": schema.StringAttribute{
				Optional:    true,
				Description: "Include descriptor attribute in response if set to TRUE.",
			},
			"use_enhanced_orgrole": schema.StringAttribute{
				Optional:    true,
				Description: "Set TRUE to utilize enhanced Organizational Role setup.",
			},
			"use_x509auth_for_soap": schema.StringAttribute{
				Optional:    true,
				Description: "Set TRUE to use certificate-based authentication for SOAP.",
			},
			"x509_key": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Private key for x509-based SOAP authentication.",
			},
			"x509_cert": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Certificate for x509-based SOAP authentication.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Description: "Username for SOAP authentication.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password for SOAP authentication.",
			},
			"client_id": schema.StringAttribute{
				Optional:    true,
				Description: "OAuth client ID.",
			},
			"client_secret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "OAuth client secret.",
			},
			"refresh_token": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "OAuth refresh token.",
			},
			"page_size": schema.StringAttribute{
				Optional:    true,
				Description: "Number of objects to return per page during import.",
			},
			"user_import_payload": schema.StringAttribute{
				Optional:    true,
				Description: "Request payload for importing users.",
			},
			"user_import_mapping": schema.StringAttribute{
				Optional:    true,
				Description: "Mapping configuration for user import.",
			},
			"account_import_payload": schema.StringAttribute{
				Optional:    true,
				Description: "Request payload for importing accounts.",
			},
			"account_import_mapping": schema.StringAttribute{
				Optional:    true,
				Description: "Mapping configuration for account import.",
			},
			"access_import_list": schema.StringAttribute{
				Optional:    true,
				Description: "Comma-separated list of access types to import.",
			},
			"raas_mapping_json": schema.StringAttribute{
				Optional:    true,
				Description: "Overrides default report mapping for RaaS.",
			},
			"access_import_mapping": schema.StringAttribute{
				Optional:    true,
				Description: "Additional access attribute mapping for Workday access objects.",
			},
			"orgrole_import_payload": schema.StringAttribute{
				Optional:    true,
				Description: "Custom SOAP body for organization role import.",
			},
			"status_key_json": schema.StringAttribute{
				Optional:    true,
				Description: "Mapping of user status.",
			},
			"userattributejson": schema.StringAttribute{
				Optional:    true,
				Description: "Specifies which job-related attributes are stored as user attributes.",
			},
			"custom_config": schema.StringAttribute{
				Optional:    true,
				Description: "Custom configuration for Workday connector.",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Description: "Privileged Access Management configuration.",
			},
			"modify_user_data_json": schema.StringAttribute{
				Optional:    true,
				Description: "Payload for modifying user data.",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Description: "Config for reading and importing status of account and entitlement.",
			},
			"create_account_payload": schema.StringAttribute{
				Optional:    true,
				Description: "Payload for creating an account.",
			},
			"update_account_payload": schema.StringAttribute{
				Optional:    true,
				Description: "Payload for updating an account.",
			},
			"update_user_payload": schema.StringAttribute{
				Optional:    true,
				Description: "Payload for updating a user.",
			},
			"assign_orgrole_payload": schema.StringAttribute{
				Optional:    true,
				Description: "Payload for assigning org role.",
			},
			"remove_orgrole_payload": schema.StringAttribute{
				Optional:    true,
				Description: "Payload for removing org role.",
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

func (r *workdayConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *workdayConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WORKDAYConnectorResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	workdayConn := openapi.WorkdayConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "WORKDAY",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.StringPointerOrEmpty(plan.Description.ValueString()),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		USERS_LAST_IMPORT_TIME:        util.StringPointerOrEmpty(plan.UsersLastImportTime.ValueString()),
		ACCOUNTS_LAST_IMPORT_TIME:     util.StringPointerOrEmpty(plan.AccountsLastImportTime.ValueString()),
		ACCESS_LAST_IMPORT_TIME:       util.StringPointerOrEmpty(plan.AccessLastImportTime.ValueString()),
		BASE_URL:                      util.StringPointerOrEmpty(plan.BaseURL.ValueString()),
		API_VERSION:                   util.StringPointerOrEmpty(plan.APIVersion.ValueString()),
		TENANT_NAME:                   util.StringPointerOrEmpty(plan.TenantName.ValueString()),
		REPORT_OWNER:                  util.StringPointerOrEmpty(plan.ReportOwner.ValueString()),
		USE_OAUTH:                     plan.UseOAuth.ValueString(), //Required Field
		INCLUDE_REFERENCE_DESCRIPTORS: util.StringPointerOrEmpty(plan.IncludeReferenceDesc.ValueString()),
		USE_ENHANCED_ORGROLE:          util.StringPointerOrEmpty(plan.UseEnhancedOrgRole.ValueString()),
		USEX509AUTHFORSOAP:            util.StringPointerOrEmpty(plan.UseX509AuthForSOAP.ValueString()),
		X509KEY:                       util.StringPointerOrEmpty(plan.X509Key.ValueString()),
		X509CERT:                      util.StringPointerOrEmpty(plan.X509Cert.ValueString()),
		USERNAME:                      util.StringPointerOrEmpty(plan.Username.ValueString()),
		PASSWORD:                      util.StringPointerOrEmpty(plan.Password.ValueString()),
		CLIENT_ID:                     util.StringPointerOrEmpty(plan.ClientID.ValueString()),
		CLIENT_SECRET:                 util.StringPointerOrEmpty(plan.ClientSecret.ValueString()),
		REFRESH_TOKEN:                 util.StringPointerOrEmpty(plan.RefreshToken.ValueString()),
		PAGE_SIZE:                     util.StringPointerOrEmpty(plan.PageSize.ValueString()),
		USER_IMPORT_PAYLOAD:           util.StringPointerOrEmpty(plan.UserImportPayload.ValueString()),
		USER_IMPORT_MAPPING:           util.StringPointerOrEmpty(plan.UserImportMapping.ValueString()),
		ACCOUNT_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.AccountImportPayload.ValueString()),
		ACCOUNT_IMPORT_MAPPING:        util.StringPointerOrEmpty(plan.AccountImportMapping.ValueString()),
		ACCESS_IMPORT_LIST:            util.StringPointerOrEmpty(plan.AccessImportList.ValueString()),
		RAAS_MAPPING_JSON:             util.StringPointerOrEmpty(plan.RAASMappingJSON.ValueString()),
		ACCESS_IMPORT_MAPPING:         util.StringPointerOrEmpty(plan.AccessImportMapping.ValueString()),
		ORGROLE_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.OrgRoleImportPayload.ValueString()),
		STATUS_KEY_JSON:               util.StringPointerOrEmpty(plan.StatusKeyJSON.ValueString()),
		USERATTRIBUTEJSON:             util.StringPointerOrEmpty(plan.UserAttributeJSON.ValueString()),
		CUSTOM_CONFIG:                 util.StringPointerOrEmpty(plan.CustomConfig.ValueString()),
		PAM_CONFIG:                    util.StringPointerOrEmpty(plan.PAMConfig.ValueString()),
		MODIFYUSERDATAJSON:            util.StringPointerOrEmpty(plan.ModifyUserDataJSON.ValueString()),
		STATUS_THRESHOLD_CONFIG:       util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()),
		CREATE_ACCOUNT_PAYLOAD:        util.StringPointerOrEmpty(plan.CreateAccountPayload.ValueString()),
		UPDATE_ACCOUNT_PAYLOAD:        util.StringPointerOrEmpty(plan.UpdateAccountPayload.ValueString()),
		UPDATE_USER_PAYLOAD:           util.StringPointerOrEmpty(plan.UpdateUserPayload.ValueString()),
		ASSIGN_ORGROLE_PAYLOAD:        util.StringPointerOrEmpty(plan.AssignOrgRolePayload.ValueString()),
		REMOVE_ORGROLE_PAYLOAD:        util.StringPointerOrEmpty(plan.RemoveOrgRolePayload.ValueString()),
	}
	workdayConnRequest := openapi.CreateOrUpdateRequest{
		WorkdayConnector: &workdayConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(workdayConnRequest).Execute()
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

func (r *workdayConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}
func (r *workdayConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WORKDAYConnectorResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	workdayConn := openapi.WorkdayConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "WORKDAY",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.StringPointerOrEmpty(plan.Description.ValueString()),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		USERS_LAST_IMPORT_TIME:        util.StringPointerOrEmpty(plan.UsersLastImportTime.ValueString()),
		ACCOUNTS_LAST_IMPORT_TIME:     util.StringPointerOrEmpty(plan.AccountsLastImportTime.ValueString()),
		ACCESS_LAST_IMPORT_TIME:       util.StringPointerOrEmpty(plan.AccessLastImportTime.ValueString()),
		BASE_URL:                      util.StringPointerOrEmpty(plan.BaseURL.ValueString()),
		API_VERSION:                   util.StringPointerOrEmpty(plan.APIVersion.ValueString()),
		TENANT_NAME:                   util.StringPointerOrEmpty(plan.TenantName.ValueString()),
		REPORT_OWNER:                  util.StringPointerOrEmpty(plan.ReportOwner.ValueString()),
		USE_OAUTH:                     plan.UseOAuth.ValueString(), //Required Field
		INCLUDE_REFERENCE_DESCRIPTORS: util.StringPointerOrEmpty(plan.IncludeReferenceDesc.ValueString()),
		USE_ENHANCED_ORGROLE:          util.StringPointerOrEmpty(plan.UseEnhancedOrgRole.ValueString()),
		USEX509AUTHFORSOAP:            util.StringPointerOrEmpty(plan.UseX509AuthForSOAP.ValueString()),
		X509KEY:                       util.StringPointerOrEmpty(plan.X509Key.ValueString()),
		X509CERT:                      util.StringPointerOrEmpty(plan.X509Cert.ValueString()),
		USERNAME:                      util.StringPointerOrEmpty(plan.Username.ValueString()),
		PASSWORD:                      util.StringPointerOrEmpty(plan.Password.ValueString()),
		CLIENT_ID:                     util.StringPointerOrEmpty(plan.ClientID.ValueString()),
		CLIENT_SECRET:                 util.StringPointerOrEmpty(plan.ClientSecret.ValueString()),
		REFRESH_TOKEN:                 util.StringPointerOrEmpty(plan.RefreshToken.ValueString()),
		PAGE_SIZE:                     util.StringPointerOrEmpty(plan.PageSize.ValueString()),
		USER_IMPORT_PAYLOAD:           util.StringPointerOrEmpty(plan.UserImportPayload.ValueString()),
		USER_IMPORT_MAPPING:           util.StringPointerOrEmpty(plan.UserImportMapping.ValueString()),
		ACCOUNT_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.AccountImportPayload.ValueString()),
		ACCOUNT_IMPORT_MAPPING:        util.StringPointerOrEmpty(plan.AccountImportMapping.ValueString()),
		ACCESS_IMPORT_LIST:            util.StringPointerOrEmpty(plan.AccessImportList.ValueString()),
		RAAS_MAPPING_JSON:             util.StringPointerOrEmpty(plan.RAASMappingJSON.ValueString()),
		ACCESS_IMPORT_MAPPING:         util.StringPointerOrEmpty(plan.AccessImportMapping.ValueString()),
		ORGROLE_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.OrgRoleImportPayload.ValueString()),
		STATUS_KEY_JSON:               util.StringPointerOrEmpty(plan.StatusKeyJSON.ValueString()),
		USERATTRIBUTEJSON:             util.StringPointerOrEmpty(plan.UserAttributeJSON.ValueString()),
		CUSTOM_CONFIG:                 util.StringPointerOrEmpty(plan.CustomConfig.ValueString()),
		PAM_CONFIG:                    util.StringPointerOrEmpty(plan.PAMConfig.ValueString()),
		MODIFYUSERDATAJSON:            util.StringPointerOrEmpty(plan.ModifyUserDataJSON.ValueString()),
		STATUS_THRESHOLD_CONFIG:       util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()),
		CREATE_ACCOUNT_PAYLOAD:        util.StringPointerOrEmpty(plan.CreateAccountPayload.ValueString()),
		UPDATE_ACCOUNT_PAYLOAD:        util.StringPointerOrEmpty(plan.UpdateAccountPayload.ValueString()),
		UPDATE_USER_PAYLOAD:           util.StringPointerOrEmpty(plan.UpdateUserPayload.ValueString()),
		ASSIGN_ORGROLE_PAYLOAD:        util.StringPointerOrEmpty(plan.AssignOrgRolePayload.ValueString()),
		REMOVE_ORGROLE_PAYLOAD:        util.StringPointerOrEmpty(plan.RemoveOrgRolePayload.ValueString()),
	}
	workdayConnRequest := openapi.CreateOrUpdateRequest{
		WorkdayConnector: &workdayConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(workdayConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *workdayConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
