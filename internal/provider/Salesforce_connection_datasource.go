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

// SalesforceConnectionDataSource defines the data source
type SalesforceConnectionDataSource struct {
	client *s.Client
	token  string
}

type SalesforceConnectionDataSourceModel struct {
	BaseConnectionDataSourceModel
	ConnectionAttributes *SalesforceConnectionAttributes `tfsdk:"connection_attributes"`
}

type SalesforceConnectionAttributes struct {
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	ClientSecret             types.String             `tfsdk:"client_secret"`
	ObjectToBeImported       types.String             `tfsdk:"object_to_be_imported"`
	FeatureLicenseJson       types.String             `tfsdk:"feature_license_json"`
	CreateAccountJson        types.String             `tfsdk:"createaccountjson"`
	RedirectUri              types.String             `tfsdk:"redirect_uri"`
	RefreshToken             types.String             `tfsdk:"refresh_token"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	ModifyAccountJson        types.String             `tfsdk:"modifyaccountjson"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	ClientId                 types.String             `tfsdk:"client_id"`
	PamConfig                types.String             `tfsdk:"pam_config"`
	CustomConfigJson         types.String             `tfsdk:"customconfigjson"`
	FieldMappingJson         types.String             `tfsdk:"field_mapping_json"`
	StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	AccountFieldQuery        types.String             `tfsdk:"account_field_query"`
	CustomCreateAccountUrl   types.String             `tfsdk:"custom_createaccount_url"`
	AccountFilterQuery       types.String             `tfsdk:"account_filter_query"`
	InstanceUrl              types.String             `tfsdk:"instance_url"`
}

var _ datasource.DataSource = &SalesforceConnectionDataSource{}

func NewSalesforceConnectionsDataSource() datasource.DataSource {
	return &SalesforceConnectionDataSource{}
}

func (d *SalesforceConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_salesforce_connection_datasource"
}

func (d *SalesforceConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
					"is_timeout_supported": schema.BoolAttribute{
						Computed: true,
					},
					"client_secret": schema.StringAttribute{
						Computed: true,
					},
					"object_to_be_imported": schema.StringAttribute{
						Computed: true,
					},
					"feature_license_json": schema.StringAttribute{
						Computed: true,
					},
					"createaccountjson": schema.StringAttribute{
						Computed: true,
					},
					"redirect_uri": schema.StringAttribute{
						Computed: true,
					},
					"refresh_token": schema.StringAttribute{
						Computed: true,
					},
					"modifyaccountjson": schema.StringAttribute{
						Computed: true,
					},
					"connection_type": schema.StringAttribute{
						Computed: true,
					},
					"is_timeout_config_validated": schema.BoolAttribute{
						Computed: true,
					},
					"client_id": schema.StringAttribute{
						Computed: true,
					},
					"pam_config": schema.StringAttribute{
						Computed: true,
					},
					"customconfigjson": schema.StringAttribute{
						Computed: true,
					},
					"field_mapping_json": schema.StringAttribute{
						Computed: true,
					},
					"status_threshold_config": schema.StringAttribute{
						Computed: true,
					},
					"account_field_query": schema.StringAttribute{
						Computed: true,
					},
					"custom_createaccount_url": schema.StringAttribute{
						Computed: true,
					},
					"account_filter_query": schema.StringAttribute{
						Computed: true,
					},
					"instance_url": schema.StringAttribute{
						Computed: true,
					},
					"connection_timeout_config": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"retry_wait":                  schema.Int64Attribute{Computed: true},
							"token_refresh_max_try_count": schema.Int64Attribute{Computed: true},
							"retry_failure_status_code":   schema.Int64Attribute{Computed: true},
							"retry_wait_max_value":        schema.Int64Attribute{Computed: true},
							"retry_count":                 schema.Int64Attribute{Computed: true},
							"read_timeout":                schema.Int64Attribute{Computed: true},
							"connection_timeout":          schema.Int64Attribute{Computed: true},
						},
					},
				},
			},
		},
	}
}

func (d *SalesforceConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SalesforceConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SalesforceConnectionDataSourceModel

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

	state.Msg = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.SalesforceConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Updatedby)
	state.Msg = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Msg)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Emailtemplate)

	if apiResp.SalesforceConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &SalesforceConnectionAttributes{
			IsTimeoutSupported:       util.SafeBoolDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.IsTimeoutSupported),
			ClientSecret:             util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CLIENT_SECRET),
			ObjectToBeImported:       util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.OBJECT_TO_BE_IMPORTED),
			FeatureLicenseJson:       util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.FEATURE_LICENSE_JSON),
			CreateAccountJson:        util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CREATEACCOUNTJSON),
			RedirectUri:              util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.REDIRECT_URI),
			RefreshToken:             util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.REFRESH_TOKEN),
			ConnectionType:           util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionType),
			ModifyAccountJson:        util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.MODIFYACCOUNTJSON),
			IsTimeoutConfigValidated: util.SafeBoolDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			ClientId:                 util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CLIENT_ID),
			PamConfig:                util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.PAM_CONFIG),
			CustomConfigJson:         util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CUSTOMCONFIGJSON),
			FieldMappingJson:         util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.FIELD_MAPPING_JSON),
			StatusThresholdConfig:    util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			AccountFieldQuery:        util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FIELD_QUERY),
			CustomCreateAccountUrl:   util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CUSTOM_CREATEACCOUNT_URL),
			AccountFilterQuery:       util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FILTER_QUERY),
			InstanceUrl:              util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.INSTANCE_URL),
		}
		
		if apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
			state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				// RetryFailureStatusCode: util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
				RetryFailureStatusCode: SafeInt64FromStringPointer(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
				RetryWaitMaxValue:      util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:             util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:            util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:      util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
			}
		}
	}

	if apiResp.SalesforceConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
	}
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}
