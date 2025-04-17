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

type GITHUBRESTConnectorResourceModel struct {
	BaseConnector
	ID                      types.String `tfsdk:"id"`
	ConnectionJSON          types.String `tfsdk:"connection_json"`
	ImportAccountEntJSON    types.String `tfsdk:"import_account_ent_json"`
	Access_Tokens           types.String `tfsdk:"access_tokens"`
	Organization_List       types.String `tfsdk:"organization_list"`
	Status_Threshold_Config types.String `tfsdk:"status_threshold_config"`
	Pam_Config              types.String `tfsdk:"pam_config"`
}

type githubRestConnectionResource struct {
	client *s.Client
	token  string
}

func GITHUBRESTNewTestConnectionResource() resource.Resource {
	return &githubRestConnectionResource{}
}

func (r *githubRestConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_github_rest_connection_resource"
}

func (r *githubRestConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and Manage GITHUB_REST Connections",
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
				Description: "Property for ConnectionJSON",
			},
			"import_account_ent_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Property for ImportAccountEntJSON",
			},
			"access_tokens": schema.StringAttribute{
				Optional:    true,
				Description: "Property for ACCESS_TOKENS",
			},
			"organization_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Property for ORGANIZATION_LIST",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Property for STATUS_THRESHOLD_CONFIG",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Property for PAM_CONFIG",
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

func (r *githubRestConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("ProviderData is nil, returning early.")
		resp.Diagnostics.AddError(
			"Provider Data Not Found",
			"Provider data is not set. Please configure the provider.",
		)
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		log.Print("[ERROR] ProviderData is not of type *saviyntProvider")
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *saviyntProvider")
		return
	}

	// Set the client and token from the provider state.
	r.client = prov.client
	r.token = prov.accessToken
}

func (r *githubRestConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GITHUBRESTConnectorResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get plan from request. Error: %v", resp.Diagnostics)
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)

	cfg.HTTPClient = http.DefaultClient
	githubRestConn := openapi.GithubRESTConnector{
		BaseConnector: openapi.BaseConnector{
			//required values
			Connectiontype: "GithubRest",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional values
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//optional values
		ConnectionJSON:          util.StringPointerOrEmpty(plan.ConnectionJSON),
		ImportAccountEntJSON:    util.StringPointerOrEmpty(plan.ImportAccountEntJSON),
		ACCESS_TOKENS:           util.StringPointerOrEmpty(plan.Access_Tokens),
		ORGANIZATION_LIST:       util.StringPointerOrEmpty(plan.Organization_List),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.Status_Threshold_Config),
		PAM_CONFIG:              util.StringPointerOrEmpty(plan.Pam_Config),
	}
	if plan.VaultConnection.ValueString() != "" {
		githubRestConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		githubRestConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		githubRestConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}
	githubRestRequest := openapi.CreateOrUpdateRequest{
		GithubRESTConnector: &githubRestConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(githubRestRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ImportAccountEntJSON = util.SafeStringDatasource(plan.ImportAccountEntJSON.ValueStringPointer())
	plan.Organization_List = util.SafeStringDatasource(plan.Organization_List.ValueStringPointer())
	plan.Status_Threshold_Config = util.SafeStringDatasource(plan.Status_Threshold_Config.ValueStringPointer())
	plan.Pam_Config = util.SafeStringDatasource(plan.Pam_Config.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	r.Read(ctx, resource.ReadRequest{State: resp.State}, &resource.ReadResponse{State: resp.State})
}

func (r *githubRestConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GITHUBRESTConnectorResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get state from request. Error: %v", resp.Diagnostics)
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
	state.ConnectionKey = types.Int64Value(int64(*apiResp.GithubRESTConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.GithubRESTConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectiontype)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Emailtemplate)
	state.ImportAccountEntJSON = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ImportAccountEntJSON)
	state.Organization_List = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ORGANIZATION_LIST)
	state.Status_Threshold_Config = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.Pam_Config = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.PAM_CONFIG)
	apiMessage := util.SafeDeref(apiResp.GithubRESTConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.GithubRESTConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to set state. Error: %v", resp.Diagnostics)
		return
	}
}

func (r *githubRestConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan GITHUBRESTConnectorResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get plan from request. Error: %v", resp.Diagnostics)
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	githubRestConn := openapi.GithubRESTConnector{
		BaseConnector: openapi.BaseConnector{
			//required values
			Connectiontype: "GithubRest",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional values
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//optional values
		ConnectionJSON:          util.StringPointerOrEmpty(plan.ConnectionJSON),
		ImportAccountEntJSON:    util.StringPointerOrEmpty(plan.ImportAccountEntJSON),
		ACCESS_TOKENS:           util.StringPointerOrEmpty(plan.Access_Tokens),
		ORGANIZATION_LIST:       util.StringPointerOrEmpty(plan.Organization_List),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.Status_Threshold_Config),
		PAM_CONFIG:              util.StringPointerOrEmpty(plan.Pam_Config),
	}
	if plan.VaultConnection.ValueString() != "" {
		githubRestConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		githubRestConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		githubRestConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	} else {
		emptyStr := ""
		githubRestConn.BaseConnector.VaultConnection = &emptyStr
		githubRestConn.BaseConnector.VaultConfiguration = &emptyStr
		githubRestConn.BaseConnector.Saveinvault = &emptyStr
	}
	githubRestRequest := openapi.CreateOrUpdateRequest{
		GithubRESTConnector: &githubRestConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(githubRestRequest).Execute()
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
	plan.ConnectionKey = types.Int64Value(int64(*getResp.GithubRESTConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.GithubRESTConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.GithubRESTConnectionResponse.Connectionname)
	plan.Description = util.SafeStringDatasource(getResp.GithubRESTConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.GithubRESTConnectionResponse.Defaultsavroles)
	plan.ConnectionType = util.SafeStringDatasource(getResp.GithubRESTConnectionResponse.Connectiontype)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.GithubRESTConnectionResponse.Emailtemplate)
	plan.ImportAccountEntJSON = util.SafeStringDatasource(getResp.GithubRESTConnectionResponse.Connectionattributes.ImportAccountEntJSON)
	plan.Organization_List = util.SafeStringDatasource(getResp.GithubRESTConnectionResponse.Connectionattributes.ORGANIZATION_LIST)
	plan.Status_Threshold_Config = util.SafeStringDatasource(getResp.GithubRESTConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.Pam_Config = util.SafeStringDatasource(getResp.GithubRESTConnectionResponse.Connectionattributes.PAM_CONFIG)
	apiMessage := util.SafeDeref(getResp.GithubRESTConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.GithubRESTConnectionResponse.Errorcode)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *githubRestConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
