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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SalesforceConnectorResourceModel struct {
	BaseConnector
	ID                     types.String `tfsdk:"id"`
	ClientId               types.String `tfsdk:"client_id"`
	ClientSecret           types.String `tfsdk:"client_secret"`
	RefreshToken           types.String `tfsdk:"refresh_token"`
	RedirectUri            types.String `tfsdk:"redirect_uri"`
	InstanceUrl            types.String `tfsdk:"instance_url"`
	ObjectToBeImported     types.String `tfsdk:"object_to_be_imported"`
	FeatureLicenseJson     types.String `tfsdk:"feature_license_json"`
	CustomCreateaccountUrl types.String `tfsdk:"custom_createaccount_url"`
	Createaccountjson      types.String `tfsdk:"createaccountjson"`
	AccountFilterQuery     types.String `tfsdk:"account_filter_query"`
	AccountFieldQuery      types.String `tfsdk:"account_field_query"`
	FieldMappingJson       types.String `tfsdk:"field_mapping_json"`
	Modifyaccountjson      types.String `tfsdk:"modifyaccountjson"`
	StatusThresholdConfig  types.String `tfsdk:"status_threshold_config"`
	Customconfigjson       types.String `tfsdk:"customconfigjson"`
	PamConfig              types.String `tfsdk:"pam_config"`
}

type salesforceConnectionResource struct {
	client *s.Client
	token  string
}

func NewSalesfoceTestConnectionResource() resource.Resource {
	return &salesforceConnectionResource{}
}

func (r *salesforceConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_salesforce_connection_resource"
}

func (r *salesforceConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SalesforceConnDescription,
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
			"client_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The OAuth client ID for Salesforce.",
			},
			"client_secret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "The OAuth client secret for Salesforce.",
			},
			"refresh_token": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "The OAuth refresh token used to get access tokens from Salesforce.",
			},
			"redirect_uri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The redirect URI used in OAuth flows. Example: https://@INSTANCE_NAME@.salesforce.com/services/oauth2/success",
			},
			"instance_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Salesforce instance base URL. Example: https://@INSTANCE_NAME@.salesforce.com",
			},
			"object_to_be_imported": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: `Comma-separated list of Salesforce objects to import. Example: "Profile,Role,Group,PermissionSet"`,
			},
			"feature_license_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON mapping of feature licenses to permission fields in Salesforce.",
			},
			"custom_createaccount_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom URL used when creating a Salesforce account.",
			},
			"createaccountjson": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON template used for account creation in Salesforce.",
			},
			"account_filter_query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query used to filter Salesforce accounts.",
			},
			"account_field_query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fields to retrieve for Salesforce accounts. Example: Id, Username, LastName, FirstName, etc.",
			},
			"field_mapping_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON mapping of local fields to Salesforce fields with data types.",
			},
			"modifyaccountjson": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON template used for modifying Salesforce accounts.",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON configuration to define active/inactive thresholds and lock statuses.",
			},
			"customconfigjson": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom configuration options for Salesforce connector.",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Privileged Access Management (PAM) configuration in JSON format.",
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

func (r *salesforceConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *salesforceConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SalesforceConnectorResourceModel
	// Extract plan from request
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
	apiClient := openapi.NewAPIClient(cfg)

	reqParams := openapi.GetConnectionDetailsRequest{}
	reqParams.SetConnectionname(plan.ConnectionName.ValueString())
	// reqParams.SetConnectionkey(state.ConnectionKey.String())
	existingResource, _, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in read block")
	}
	if existingResource != nil && existingResource.SalesforceConnectionResponse != nil && existingResource.SalesforceConnectionResponse.Errorcode != nil && *existingResource.SalesforceConnectionResponse.Errorcode == 0 {
		log.Printf("[ERROR] Connection name already exists. Please import or use a different name")
		resp.Diagnostics.AddError("API Create Failed", "Connection name already exists. Please import or use a different name")
		return
	}
	salesforceConn := openapi.SalesforceConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "SalesForce",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:        util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		CLIENT_ID:                util.StringPointerOrEmpty(plan.ClientId),
		CLIENT_SECRET:            util.StringPointerOrEmpty(plan.ClientSecret),
		REFRESH_TOKEN:            util.StringPointerOrEmpty(plan.RefreshToken),
		REDIRECT_URI:             util.StringPointerOrEmpty(plan.RedirectUri),
		INSTANCE_URL:             util.StringPointerOrEmpty(plan.InstanceUrl),
		OBJECT_TO_BE_IMPORTED:    util.StringPointerOrEmpty(plan.ObjectToBeImported),
		FEATURE_LICENSE_JSON:     util.StringPointerOrEmpty(plan.FeatureLicenseJson),
		CUSTOM_CREATEACCOUNT_URL: util.StringPointerOrEmpty(plan.CustomCreateaccountUrl),
		CREATEACCOUNTJSON:        util.StringPointerOrEmpty(plan.Createaccountjson),
		ACCOUNT_FILTER_QUERY:     util.StringPointerOrEmpty(plan.AccountFilterQuery),
		ACCOUNT_FIELD_QUERY:      util.StringPointerOrEmpty(plan.AccountFieldQuery),
		FIELD_MAPPING_JSON:       util.StringPointerOrEmpty(plan.FieldMappingJson),
		MODIFYACCOUNTJSON:        util.StringPointerOrEmpty(plan.Modifyaccountjson),
		STATUS_THRESHOLD_CONFIG:  util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CUSTOMCONFIGJSON:         util.StringPointerOrEmpty(plan.Customconfigjson),
		PAM_CONFIG:               util.StringPointerOrEmpty(plan.PamConfig),
	}

	salesforceConnRequest := openapi.CreateOrUpdateRequest{
		SalesforceConnector: &salesforceConn,
	}

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(salesforceConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionType = types.StringValue("SalesForce")
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ClientId = util.SafeStringDatasource(plan.ClientId.ValueStringPointer())
	plan.RedirectUri = util.SafeStringDatasource(plan.RedirectUri.ValueStringPointer())
	plan.InstanceUrl = util.SafeStringDatasource(plan.InstanceUrl.ValueStringPointer())
	plan.ObjectToBeImported = util.SafeStringDatasource(plan.ObjectToBeImported.ValueStringPointer())
	plan.FeatureLicenseJson = util.SafeStringDatasource(plan.FeatureLicenseJson.ValueStringPointer())
	plan.CustomCreateaccountUrl = util.SafeStringDatasource(plan.CustomCreateaccountUrl.ValueStringPointer())
	plan.Createaccountjson = util.SafeStringDatasource(plan.Createaccountjson.ValueStringPointer())
	plan.AccountFilterQuery = util.SafeStringDatasource(plan.AccountFilterQuery.ValueStringPointer())
	plan.AccountFieldQuery = util.SafeStringDatasource(plan.AccountFieldQuery.ValueStringPointer())
	plan.FieldMappingJson = util.SafeStringDatasource(plan.FieldMappingJson.ValueStringPointer())
	plan.Modifyaccountjson = util.SafeStringDatasource(plan.Modifyaccountjson.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.Customconfigjson = util.SafeStringDatasource(plan.Customconfigjson.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *salesforceConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SalesforceConnectorResourceModel

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
	state.ConnectionKey = types.Int64Value(int64(*apiResp.SalesforceConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.SalesforceConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectiontype)
	state.Msg = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Msg)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Emailtemplate)
	state.ObjectToBeImported = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.OBJECT_TO_BE_IMPORTED)
	state.FeatureLicenseJson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.FEATURE_LICENSE_JSON)
	state.Createaccountjson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	state.RedirectUri = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.REDIRECT_URI)
	state.Modifyaccountjson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.MODIFYACCOUNTJSON)
	state.ClientId = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CLIENT_ID)
	state.PamConfig = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.Customconfigjson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CUSTOMCONFIGJSON)
	state.FieldMappingJson = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.FIELD_MAPPING_JSON)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.AccountFieldQuery = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FIELD_QUERY)
	state.CustomCreateaccountUrl = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CUSTOM_CREATEACCOUNT_URL)
	state.AccountFilterQuery = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FILTER_QUERY)
	state.InstanceUrl = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.INSTANCE_URL)
	apiMessage := util.SafeDeref(apiResp.SalesforceConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.SalesforceConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *salesforceConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SalesforceConnectorResourceModel
	var state SalesforceConnectorResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	if plan.ConnectionName.ValueString() != state.ConnectionName.ValueString() {
		resp.Diagnostics.AddError("Error", "Connection name cannot be updated")
		log.Printf("[ERROR]: Connection name cannot be updated")
		return
	}
	if plan.ConnectionType.ValueString() != state.ConnectionType.ValueString() {
		resp.Diagnostics.AddError("Error", "Connection type cannot be updated")
		log.Printf("[ERROR]: Connection type cannot be updated")
		return
	}

	cfg.HTTPClient = http.DefaultClient
	salesforceConn := openapi.SalesforceConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "SalesForce",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:        util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		CLIENT_ID:                util.StringPointerOrEmpty(plan.ClientId),
		CLIENT_SECRET:            util.StringPointerOrEmpty(plan.ClientSecret),
		REFRESH_TOKEN:            util.StringPointerOrEmpty(plan.RefreshToken),
		REDIRECT_URI:             util.StringPointerOrEmpty(plan.RedirectUri),
		INSTANCE_URL:             util.StringPointerOrEmpty(plan.InstanceUrl),
		OBJECT_TO_BE_IMPORTED:    util.StringPointerOrEmpty(plan.ObjectToBeImported),
		FEATURE_LICENSE_JSON:     util.StringPointerOrEmpty(plan.FeatureLicenseJson),
		CUSTOM_CREATEACCOUNT_URL: util.StringPointerOrEmpty(plan.CustomCreateaccountUrl),
		CREATEACCOUNTJSON:        util.StringPointerOrEmpty(plan.Createaccountjson),
		ACCOUNT_FILTER_QUERY:     util.StringPointerOrEmpty(plan.AccountFilterQuery),
		ACCOUNT_FIELD_QUERY:      util.StringPointerOrEmpty(plan.AccountFieldQuery),
		FIELD_MAPPING_JSON:       util.StringPointerOrEmpty(plan.FieldMappingJson),
		MODIFYACCOUNTJSON:        util.StringPointerOrEmpty(plan.Modifyaccountjson),
		STATUS_THRESHOLD_CONFIG:  util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CUSTOMCONFIGJSON:         util.StringPointerOrEmpty(plan.Customconfigjson),
		PAM_CONFIG:               util.StringPointerOrEmpty(plan.PamConfig),
	}
	salesforceConnRequest := openapi.CreateOrUpdateRequest{
		SalesforceConnector: &salesforceConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(salesforceConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
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
	plan.ConnectionKey = types.Int64Value(int64(*getResp.SalesforceConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.SalesforceConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionname)
	plan.Description = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Defaultsavroles)
	plan.ConnectionType = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectiontype)
	plan.Msg = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Msg)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Emailtemplate)
	plan.ObjectToBeImported = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.OBJECT_TO_BE_IMPORTED)
	plan.FeatureLicenseJson = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.FEATURE_LICENSE_JSON)
	plan.Createaccountjson = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	plan.RedirectUri = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.REDIRECT_URI)
	plan.Modifyaccountjson = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.MODIFYACCOUNTJSON)
	plan.ClientId = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.CLIENT_ID)
	plan.PamConfig = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.PAM_CONFIG)
	plan.Customconfigjson = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.CUSTOMCONFIGJSON)
	plan.FieldMappingJson = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.FIELD_MAPPING_JSON)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.AccountFieldQuery = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FIELD_QUERY)
	plan.CustomCreateaccountUrl = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.CUSTOM_CREATEACCOUNT_URL)
	plan.AccountFilterQuery = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FILTER_QUERY)
	plan.InstanceUrl = util.SafeStringDatasource(getResp.SalesforceConnectionResponse.Connectionattributes.INSTANCE_URL)
	apiMessage := util.SafeDeref(getResp.SalesforceConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.SalesforceConnectionResponse.Errorcode)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}
func (r *salesforceConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *salesforceConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)
}
