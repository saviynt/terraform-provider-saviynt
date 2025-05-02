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

// GithubRestConnectionDataSource defines the data source
type GithubRestConnectionDataSource struct {
	client *s.Client
	token  string
}

type GithubRestConnectionDataSourceModel struct {
	BaseConnectionDataSourceModel
	ConnectionAttributes *GithubRestConnectionAttributes `tfsdk:"connection_attributes"`
}

type GithubRestConnectionAttributes struct {
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	ConnectionJSON           types.String             `tfsdk:"connection_json"`
	OrganizationList         types.String             `tfsdk:"organization_list"`
	ImportAccountEntJSON     types.String             `tfsdk:"import_account_ent_json"`
	StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	AccessTokens             types.String             `tfsdk:"access_tokens"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
}

var _ datasource.DataSource = &GithubRestConnectionDataSource{}

func NewGithubRestConnectionsDataSource() datasource.DataSource {
	return &GithubRestConnectionDataSource{}
}

func (d *GithubRestConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_github_rest_connection_datasource"
}

func (d *GithubRestConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.GithubRestConnDataSourceDescription,
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
					"connection_json": schema.StringAttribute{
						Computed: true,
					},
					"organization_list": schema.StringAttribute{
						Computed: true,
					},
					"import_account_ent_json": schema.StringAttribute{
						Computed: true,
					},
					"status_threshold_config": schema.StringAttribute{
						Computed: true,
					},
					"access_tokens": schema.StringAttribute{
						Computed: true,
					},
					"connection_type": schema.StringAttribute{
						Computed: true,
					},
					"is_timeout_config_validated": schema.BoolAttribute{
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

func (d *GithubRestConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *GithubRestConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state GithubRestConnectionDataSourceModel

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

	state.Msg = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.GithubRESTConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Emailtemplate)

	if apiResp.GithubRESTConnectionResponse.Connectionattributes != nil {

		state.ConnectionAttributes = &GithubRestConnectionAttributes{
			IsTimeoutSupported:       util.SafeBoolDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.IsTimeoutSupported),
			ConnectionJSON:           util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionJSON),
			OrganizationList:         util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ORGANIZATION_LIST),
			ImportAccountEntJSON:     util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ImportAccountEntJSON),
			StatusThresholdConfig:    util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			AccessTokens:             util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ACCESS_TOKENS),
			ConnectionType:           util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionType),
			IsTimeoutConfigValidated: util.SafeBoolDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
		}
		if apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
			state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				RetryFailureStatusCode:  util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
				RetryWaitMaxValue:       util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:              util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:             util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:       util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
			}
		}
	}

	if apiResp.GithubRESTConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
	}
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}
