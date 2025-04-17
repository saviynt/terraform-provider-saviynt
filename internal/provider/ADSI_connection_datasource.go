// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// ADSIConnectionsDataSource defines the data source
type ADSIConnectionsDataSource struct {
	client *s.Client
	token  string
}

type ADSIConnectionDataSourceModel struct {
	BaseConnectionDataSourceModel
	ConnectionAttributes *ADSIConnectionAttributes `tfsdk:"connection_attributes"`
}

type ADSIConnectionAttributes struct {
	ImportNestedMembership      types.String            `tfsdk:"import_nested_membership"`
	PASSWDPOLICYJSON            types.String            `tfsdk:"password_policy_json"`
	CREATEACCOUNTJSON           types.String            `tfsdk:"create_account_json"`
	ENDPOINTS_FILTER            types.String            `tfsdk:"endpoints_filter"`
	DISABLEACCOUNTJSON          types.String            `tfsdk:"disable_account_json"`
	REMOVEACCESSENTITLEMENTJSON types.String            `tfsdk:"remove_access_entitlement_json"`
	GroupSearchBaseDN           types.String            `tfsdk:"group_search_base_dn"`
	ConnectionType              types.String            `tfsdk:"connection_type"`
	STATUSKEYJSON               types.String            `tfsdk:"status_key_json"`
	DEFAULT_USER_ROLE           types.String            `tfsdk:"default_user_role"`
	FOREST_DETAILS              types.String            `tfsdk:"forest_details"`
	USERNAME                    types.String            `tfsdk:"username"`
	UPDATESERVICEACCOUNTJSON    types.String            `tfsdk:"update_service_account_json"`
	ADDACCESSJSON               types.String            `tfsdk:"add_access_json"`
	CREATESERVICEACCOUNTJSON    types.String            `tfsdk:"create_service_account_json"`
	ACCOUNTNAMERULE             types.String            `tfsdk:"account_name_rule"`
	CONNECTION_URL              types.String            `tfsdk:"connection_url"`
	IsTimeoutSupported          types.Bool              `tfsdk:"is_timeout_supported"`
	CreateUpdateMappings        types.String            `tfsdk:"create_update_mappings"`
	ACCOUNT_ATTRIBUTE           types.String            `tfsdk:"account_attribute"`
	PASSWORD                    types.String            `tfsdk:"password"`
	PAM_CONFIG                  types.String            `tfsdk:"pam_config"`
	PAGE_SIZE                   types.String            `tfsdk:"page_size"`
	SEARCHFILTER                types.String            `tfsdk:"search_filter"`
	UPDATEGROUPJSON             types.String            `tfsdk:"update_group_json"`
	CREATEGROUPJSON             types.String            `tfsdk:"create_group_json"`
	ENTITLEMENT_ATTRIBUTE       types.String            `tfsdk:"entitlement_attribute"`
	CHECKFORUNIQUE              types.String            `tfsdk:"check_for_unique"`
	REMOVESERVICEACCOUNTJSON    types.String            `tfsdk:"remove_service_account_json"`
	ConnectionTimeoutConfig     ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	UPDATEUSERJSON              types.String            `tfsdk:"update_user_json"`
	URL                         types.String            `tfsdk:"url"`
	MOVEACCOUNTJSON             types.String            `tfsdk:"move_account_json"`
	CUSTOMCONFIGJSON            types.String            `tfsdk:"custom_config_json"`
	STATUS_THRESHOLD_CONFIG     types.String            `tfsdk:"status_threshold_config"`
	GroupImportMapping          types.String            `tfsdk:"group_import_mapping"`
	PROVISIONING_URL            types.String            `tfsdk:"provisioning_url"`
	REMOVEGROUPJSON             types.String            `tfsdk:"remove_group_json"`
	REMOVEACCESSJSON            types.String            `tfsdk:"remove_access_json"`
	IMPORTDATACOOKIES           types.String            `tfsdk:"import_data_cookies"`
	RESETANDCHANGEPASSWRDJSON   types.String            `tfsdk:"reset_and_change_password_json"`
	USER_ATTRIBUTE              types.String            `tfsdk:"user_attribute"`
	ADDACCESSENTITLEMENTJSON    types.String            `tfsdk:"add_access_entitlement_json"`
	MODIFYUSERDATAJSON          types.String            `tfsdk:"modify_user_data_json"`
	IsTimeoutConfigValidated    types.Bool              `tfsdk:"is_timeout_config_validated"`
	ENABLEGROUPMANAGEMENT       types.String            `tfsdk:"enable_group_management"`
	ENABLEACCOUNTJSON           types.String            `tfsdk:"enable_account_json"`
	FORESTLIST                  types.String            `tfsdk:"forest_list"`
	OBJECTFILTER                types.String            `tfsdk:"object_filter"`
	UPDATEACCOUNTJSON           types.String            `tfsdk:"update_account_json"`
	REMOVEACCOUNTJSON           types.String            `tfsdk:"remove_account_json"`
}

// Ensure the implementation satisfies Terraform framework interface
var _ datasource.DataSource = &ADSIConnectionsDataSource{}

// NewSecuritySystemsDataSource returns a new instance
func NewADSIConnectionsDataSource() datasource.DataSource {
	return &ADSIConnectionsDataSource{}
}

// Metadata defines the data source name
func (d *ADSIConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_adsi_connection_datasource"
}

// Schema defines the attributes for the data source
func (d *ADSIConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Retrieve the details for a given ADSI connector by its name or key",
		Attributes: map[string]schema.Attribute{
			"connection_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the connection.",
			},
			"connection_key": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The key of the connection.",
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"default_sav_roles": schema.StringAttribute{
				Computed: true,
			},
			"msg": schema.StringAttribute{
				Computed: true,
			},
			"email_template": schema.StringAttribute{
				Computed: true,
			},
			"connection_type": schema.StringAttribute{
				Computed: true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"created_by": schema.StringAttribute{
				Computed: true,
			},
			"updated_by": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.Int64Attribute{
				Computed: true,
			},
			"error_code": schema.Int64Attribute{
				Computed: true,
			},
			"connection_attributes": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"import_nested_membership":       schema.StringAttribute{Computed: true},
					"password_policy_json":           schema.StringAttribute{Computed: true},
					"create_account_json":            schema.StringAttribute{Computed: true},
					"endpoints_filter":               schema.StringAttribute{Computed: true},
					"disable_account_json":           schema.StringAttribute{Computed: true},
					"remove_access_entitlement_json": schema.StringAttribute{Computed: true},
					"group_search_base_dn":           schema.StringAttribute{Computed: true},
					"connection_type":                schema.StringAttribute{Computed: true},
					"status_key_json":                schema.StringAttribute{Computed: true},
					"default_user_role":              schema.StringAttribute{Computed: true},
					"forest_details":                 schema.StringAttribute{Computed: true},
					"username":                       schema.StringAttribute{Computed: true},
					"update_service_account_json":    schema.StringAttribute{Computed: true},
					"add_access_json":                schema.StringAttribute{Computed: true},
					"create_service_account_json":    schema.StringAttribute{Computed: true},
					"account_name_rule":              schema.StringAttribute{Computed: true},
					"connection_url":                 schema.StringAttribute{Computed: true},
					"is_timeout_supported":           schema.BoolAttribute{Computed: true},
					"create_update_mappings":         schema.StringAttribute{Computed: true},
					"account_attribute":              schema.StringAttribute{Computed: true},
					"password":                       schema.StringAttribute{Computed: true},
					"pam_config":                     schema.StringAttribute{Computed: true},
					"page_size":                      schema.StringAttribute{Computed: true},
					"search_filter":                  schema.StringAttribute{Computed: true},
					"update_group_json":              schema.StringAttribute{Computed: true},
					"create_group_json":              schema.StringAttribute{Computed: true},
					"entitlement_attribute":          schema.StringAttribute{Computed: true},
					"check_for_unique":               schema.StringAttribute{Computed: true},
					"remove_service_account_json":    schema.StringAttribute{Computed: true},
					"update_user_json":               schema.StringAttribute{Computed: true},
					"url":                            schema.StringAttribute{Computed: true},
					"move_account_json":              schema.StringAttribute{Computed: true},
					"custom_config_json":             schema.StringAttribute{Computed: true},
					"status_threshold_config":        schema.StringAttribute{Computed: true},
					"group_import_mapping":           schema.StringAttribute{Computed: true},
					"provisioning_url":               schema.StringAttribute{Computed: true},
					"remove_group_json":              schema.StringAttribute{Computed: true},
					"remove_access_json":             schema.StringAttribute{Computed: true},
					"import_data_cookies":            schema.StringAttribute{Computed: true},
					"reset_and_change_password_json": schema.StringAttribute{Computed: true},
					"user_attribute":                 schema.StringAttribute{Computed: true},
					"add_access_entitlement_json":    schema.StringAttribute{Computed: true},
					"modify_user_data_json":          schema.StringAttribute{Computed: true},
					"is_timeout_config_validated":    schema.BoolAttribute{Computed: true},
					"enable_group_management":        schema.StringAttribute{Computed: true},
					"enable_account_json":            schema.StringAttribute{Computed: true},
					"forest_list":                    schema.StringAttribute{Computed: true},
					"object_filter":                  schema.StringAttribute{Computed: true},
					"update_account_json":            schema.StringAttribute{Computed: true},
					"remove_account_json":            schema.StringAttribute{Computed: true},
					"connection_timeout_config": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"retry_wait":                  schema.Int64Attribute{Computed: true},
							"token_refresh_max_try_count": schema.Int64Attribute{Computed: true},
							"retry_wait_max_value":        schema.Int64Attribute{Computed: true},
							"retry_count":                 schema.Int64Attribute{Computed: true},
							"read_timeout":                schema.Int64Attribute{Computed: true},
							"connection_timeout":          schema.Int64Attribute{Computed: true},
							"retry_failure_status_code":   schema.Float64Attribute{Computed: true},
						},
					},
				},
			},
		},
	}
}

func (d *ADSIConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = prov.client
	d.token = prov.accessToken
}

func (d *ADSIConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ADSIConnectionDataSourceModel

	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configure API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(d.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+d.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)
	reqParams := openapi.GetConnectionDetailsRequest{}

	// Set filters based on provided parameters
	if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
		reqParams.SetConnectionname(state.ConnectionName.ValueString())
	}
	if !state.ConnectionKey.IsNull() {
		connectionKeyInt := state.ConnectionKey.ValueInt64()
		reqParams.SetConnectionkey(strconv.FormatInt(connectionKeyInt, 10))
	}
	apiReq := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams)

	// Execute API request
	apiResp, httpResp, err := apiReq.Execute()
	if err != nil {
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.ADSIConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Updatedby)
	state.Msg = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Msg)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Emailtemplate)

	if apiResp.ADSIConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &ADSIConnectionAttributes{
			ImportNestedMembership:      util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ImportNestedMembership),
			PASSWDPOLICYJSON:            util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PASSWDPOLICYJSON),
			CREATEACCOUNTJSON:           util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CREATEACCOUNTJSON),
			ENDPOINTS_FILTER:            util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENDPOINTS_FILTER),
			DISABLEACCOUNTJSON:          util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON),
			REMOVEACCESSENTITLEMENTJSON: util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCESSENTITLEMENTJSON),
			GroupSearchBaseDN:           util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.GroupSearchBaseDN),
			ConnectionType:              util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ConnectionType),
			STATUSKEYJSON:               util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.STATUSKEYJSON),
			DEFAULT_USER_ROLE:           util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.DEFAULT_USER_ROLE),
			FOREST_DETAILS:              util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.FOREST_DETAILS),
			USERNAME:                    util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.USERNAME),
			UPDATESERVICEACCOUNTJSON:    util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATESERVICEACCOUNTJSON),
			ADDACCESSJSON:               util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ADDACCESSJSON),
			CREATESERVICEACCOUNTJSON:    util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CREATESERVICEACCOUNTJSON),
			ACCOUNTNAMERULE:             util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ACCOUNTNAMERULE),
			CONNECTION_URL:              util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CONNECTION_URL),
			IsTimeoutSupported:          util.SafeBoolDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.IsTimeoutSupported),
			CreateUpdateMappings:        util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CreateUpdateMappings),
			ACCOUNT_ATTRIBUTE:           util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTE),
			PASSWORD:                    util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PASSWORD),
			PAM_CONFIG:                  util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PAM_CONFIG),
			PAGE_SIZE:                   util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PAGE_SIZE),
			SEARCHFILTER:                util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.SEARCHFILTER),
			UPDATEGROUPJSON:             util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATEGROUPJSON),
			CREATEGROUPJSON:             util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CREATEGROUPJSON),
			ENTITLEMENT_ATTRIBUTE:       util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE),
			CHECKFORUNIQUE:              util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CHECKFORUNIQUE),
			REMOVESERVICEACCOUNTJSON:    util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVESERVICEACCOUNTJSON),
			UPDATEUSERJSON:              util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATEUSERJSON),
			URL:                         util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.URL),
			MOVEACCOUNTJSON:             util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.MOVEACCOUNTJSON),
			CUSTOMCONFIGJSON:            util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CUSTOMCONFIGJSON),
			STATUS_THRESHOLD_CONFIG:     util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			GroupImportMapping:          util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.GroupImportMapping),
			PROVISIONING_URL:            util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PROVISIONING_URL),
			REMOVEGROUPJSON:             util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEGROUPJSON),
			REMOVEACCESSJSON:            util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCESSJSON),
			IMPORTDATACOOKIES:           util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.IMPORTDATACOOKIES),
			RESETANDCHANGEPASSWRDJSON:   util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.RESETANDCHANGEPASSWRDJSON),
			USER_ATTRIBUTE:              util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.USER_ATTRIBUTE),
			ADDACCESSENTITLEMENTJSON:    util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ADDACCESSENTITLEMENTJSON),
			MODIFYUSERDATAJSON:          util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON),
			IsTimeoutConfigValidated:    util.SafeBoolDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			ENABLEGROUPMANAGEMENT:       util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENABLEGROUPMANAGEMENT),
			ENABLEACCOUNTJSON:           util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON),
			FORESTLIST:                  util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.FORESTLIST),
			OBJECTFILTER:                util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.OBJECTFILTER),
			UPDATEACCOUNTJSON:           util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON),
			REMOVEACCOUNTJSON:           util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCOUNTJSON),
			ConnectionTimeoutConfig: ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				RetryWaitMaxValue:       util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:              util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:             util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:       util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
				// RetryFailureStatusCode:  SafeInt64FromStringPointer(apiResp.ADSIConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
				RetryFailureStatusCode: util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
			},
		}

	}
	if apiResp.ADSIConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
	}
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}

func SafeInt64FromStringPointer(s *string) types.Int64 {
	if s == nil || *s == "" {
		return types.Int64Value(0)
	}

	val, err := strconv.ParseInt(*s, 10, 64)
	if err != nil {
		return types.Int64Value(0)
	}

	return types.Int64Value(val)
}
