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

type ADSIConnectorResourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnector
	URL                         types.String `tfsdk:"url"`
	Username                    types.String `tfsdk:"username"`
	Password                    types.String `tfsdk:"password"`
	ConnectionUrl               types.String `tfsdk:"connection_url"`
	ProvisioningUrl             types.String `tfsdk:"provisioning_url"`
	ForestList                  types.String `tfsdk:"forestlist"`
	DefaultUserRole             types.String `tfsdk:"default_user_role"`
	UpdateUserJson              types.String `tfsdk:"updateuserjson"`
	EndpointsFilter             types.String `tfsdk:"endpoints_filter"`
	SearchFilter                types.String `tfsdk:"searchfilter"`
	ObjectFilter                types.String `tfsdk:"objectfilter"`
	AccountAttribute            types.String `tfsdk:"account_attribute"`
	StatusThresholdConfig       types.String `tfsdk:"status_threshold_config"`
	EntitlementAttribute        types.String `tfsdk:"entitlement_attribute"`
	UserAttribute               types.String `tfsdk:"user_attribute"`
	GroupSearchBaseDN           types.String `tfsdk:"group_search_base_dn"`
	CheckForUnique              types.String `tfsdk:"checkforunique"`
	StatusKeyJson               types.String `tfsdk:"statuskeyjson"`
	GroupImportMapping          types.String `tfsdk:"group_import_mapping"`
	ImportNestedMembership      types.String `tfsdk:"import_nested_membership"`
	PageSize                    types.String `tfsdk:"page_size"`
	AccountNameRule             types.String `tfsdk:"accountnamerule"`
	CreateAccountJson           types.String `tfsdk:"createaccountjson"`
	UpdateAccountJson           types.String `tfsdk:"updateaccountjson"`
	EnableAccountJson           types.String `tfsdk:"enableaccountjson"`
	DisableAccountJson          types.String `tfsdk:"disableaccountjson"`
	RemoveAccountJson           types.String `tfsdk:"removeaccountjson"`
	AddAccessJson               types.String `tfsdk:"addaccessjson"`
	RemoveAccessJson            types.String `tfsdk:"removeaccessjson"`
	ResetAndChangePasswrdJson   types.String `tfsdk:"resetandchangepasswrdjson"`
	CreateGroupJson             types.String `tfsdk:"creategroupjson"`
	UpdateGroupJson             types.String `tfsdk:"updategroupjson"`
	RemoveGroupJson             types.String `tfsdk:"removegroupjson"`
	AddAccessEntitlementJson    types.String `tfsdk:"addaccessentitlementjson"`
	CustomConfigJson            types.String `tfsdk:"customconfigjson"`
	RemoveAccessEntitlementJson types.String `tfsdk:"removeaccessentitlementjson"`
	CreateServiceAccountJson    types.String `tfsdk:"createserviceaccountjson"`
	UpdateServiceAccountJson    types.String `tfsdk:"updateserviceaccountjson"`
	RemoveServiceAccountJson    types.String `tfsdk:"removeserviceaccountjson"`
	PamConfig                   types.String `tfsdk:"pam_config"`
	ModifyUserDataJson          types.String `tfsdk:"modifyuserdatajson"`
}

// testConnectionResource implements the resource.Resource interface.
type adsiConnectionResource struct {
	// client *openapi.APIClient
	client *s.Client
	token  string
}

// NewTestConnectionResource returns a new instance of testConnectionResource.
func ADSINewTestConnectionResource() resource.Resource {
	return &adsiConnectionResource{}
}

func (r *adsiConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_adsi_connection_resource"
}

func (r *adsiConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"url": schema.StringAttribute{
				Required:    true,
				Description: "Primary/root domain URL list (comma Separated)",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Service account username",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Description: "Service account password",
			},
			"connection_url": schema.StringAttribute{
				Required:    true,
				Description: "ADSI remote agent Connection URL",
			},
			"provisioning_url": schema.StringAttribute{
				Optional:    true,
				Description: "ADSI remote agent Provisioning URL",
			},
			"forestlist": schema.StringAttribute{
				Required:    true,
				Description: "Forest List (Comma Separated) which we need to manage",
			},
			"default_user_role": schema.StringAttribute{
				Optional:    true,
				Description: "Default SAV Role to be assigned to all the new users that gets imported via User Import",
			},
			"updateuserjson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the attribute Value which will be used to Update existing User",
			},
			"endpoints_filter": schema.StringAttribute{
				Optional:    true,
				Description: "Provide the configuration to create Child Endpoints and import associated accounts and entitlements",
			},
			"searchfilter": schema.StringAttribute{
				Optional:    true,
				Description: "Account Search Filter to specify the starting point of the directory from where the accounts needs to be imported. You can have multiple BaseDNs here separated by ###.",
			},
			"objectfilter": schema.StringAttribute{
				Optional:    true,
				Description: "Object Filter is used to filter the objects that will be returned.This filter will be same for all domains.",
			},
			"account_attribute": schema.StringAttribute{
				Optional:    true,
				Description: "Map EIC and AD attributes for account import (AD attributes must be in lower case)",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Description: "Account status and threshold related config",
			},
			"entitlement_attribute": schema.StringAttribute{
				Optional:    true,
				Description: "Account attribute that contains group membership",
			},
			"user_attribute": schema.StringAttribute{
				Optional:    true,
				Description: "Map EIC and AD attributes for user import (AD attributes must be in lower case)",
			},
			"group_search_base_dn": schema.StringAttribute{
				Optional:    true,
				Description: "Group Search Filter to specify the starting point of the directory from where the groups needs to be imported. You can have multiple BaseDNs here separated by ###.",
			},
			"checkforunique": schema.StringAttribute{
				Optional:    true,
				Description: "Evaluate the uniqueness of an attribute",
			},
			"statuskeyjson": schema.StringAttribute{
				Optional:    true,
				Description: "JSON configuration to specify Users status",
			},
			"group_import_mapping": schema.StringAttribute{
				Optional:    true,
				Description: "Map AD group attribute to EIC entitlement attribute for import",
			},
			"import_nested_membership": schema.StringAttribute{
				Optional:    true,
				Description: "Specify if you want the connector to import all indirect or nested membership of an account or a group during access import",
			},
			"page_size": schema.StringAttribute{
				Optional:    true,
				Description: "Page size defines the number of objects to be returned from each AD operation.",
			},
			"accountnamerule": schema.StringAttribute{
				Optional:    true,
				Description: "Rule to generate account name.",
			},
			"createaccountjson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the attributes values which will be used to Create the New Account.",
			},
			"updateaccountjson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the attributes values which will be used to Update existing Account.",
			},
			"enableaccountjson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the actions and attribute updates to be performed for enabling an account.",
			},
			"disableaccountjson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the actions and attributes updates to be performed for disabling an account.",
			},
			"removeaccountjson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the actions to be performed for deleting an account.",
			},
			"addaccessjson": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration to ADD Access (cross domain/forest group membership) to an account.",
			},
			"removeaccessjson": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration to REMOVE Access (cross domain/forest group membership) to an account.",
			},
			"resetandchangepasswrdjson": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration to Reset and Change Password.",
			},
			"creategroupjson": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration to Create a Group",
			},
			"updategroupjson": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration to Update a Group",
			},
			"removegroupjson": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration to Delete a Group",
			},
			"addaccessentitlementjson": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration to Add nested group hierarchy",
			},
			"customconfigjson": schema.StringAttribute{
				Optional:    true,
				Description: "Custom configuration JSON",
			},
			"removeaccessentitlementjson": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration to Remove nested group hierarchy",
			},
			"createserviceaccountjson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the Field Value which will be used to Create the New Service Account.",
			},
			"updateserviceaccountjson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the Field Value which will be used to update the existing Service Account.",
			},
			"removeserviceaccountjson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the actions to be performed while deleting a service account.",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify Bootstrap Config.",
			},
			"modifyuserdatajson": schema.StringAttribute{
				Optional:    true,
				Description: "Specify this parameter to transform the data during user import.",
			},
			// "result": schema.StringAttribute{
			// 	Optional:    true,
			// 	Computed:    true,
			// 	Description: "The result of the API call.",
			// },
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

func (r *adsiConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *adsiConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ADSIConnectorResourceModel

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

	adsiConn := openapi.ADSIConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "ADSI",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.SafeStringConnector(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnector(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnector(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		URL:                         plan.URL.ValueString(),
		USERNAME:                    plan.Username.ValueString(),
		PASSWORD:                    plan.Password.ValueString(),
		CONNECTION_URL:              plan.ConnectionUrl.ValueString(),
		FORESTLIST:                  plan.ForestList.ValueString(),
		PROVISIONING_URL:            util.SafeStringConnector(plan.ProvisioningUrl.ValueString()),
		DEFAULT_USER_ROLE:           util.SafeStringConnector(plan.DefaultUserRole.ValueString()),
		UPDATEUSERJSON:              util.SafeStringConnector(plan.UpdateUserJson.ValueString()),
		ENDPOINTS_FILTER:            util.SafeStringConnector(plan.EndpointsFilter.ValueString()),
		SEARCHFILTER:                util.SafeStringConnector(plan.SearchFilter.ValueString()),
		OBJECTFILTER:                util.SafeStringConnector(plan.ObjectFilter.ValueString()),
		ACCOUNT_ATTRIBUTE:           util.SafeStringConnector(plan.AccountAttribute.ValueString()),
		STATUS_THRESHOLD_CONFIG:     util.SafeStringConnector(plan.StatusThresholdConfig.ValueString()),
		ENTITLEMENT_ATTRIBUTE:       util.SafeStringConnector(plan.EntitlementAttribute.ValueString()),
		USER_ATTRIBUTE:              util.SafeStringConnector(plan.UserAttribute.ValueString()),
		GroupSearchBaseDN:           util.SafeStringConnector(plan.GroupSearchBaseDN.ValueString()),
		CHECKFORUNIQUE:              util.SafeStringConnector(plan.CheckForUnique.ValueString()),
		STATUSKEYJSON:               util.SafeStringConnector(plan.StatusKeyJson.ValueString()),
		GroupImportMapping:          util.SafeStringConnector(plan.GroupImportMapping.ValueString()),
		ImportNestedMembership:      util.SafeStringConnector(plan.ImportNestedMembership.ValueString()),
		PAGE_SIZE:                   util.SafeStringConnector(plan.PageSize.ValueString()),
		ACCOUNTNAMERULE:             util.SafeStringConnector(plan.AccountNameRule.ValueString()),
		CREATEACCOUNTJSON:           util.SafeStringConnector(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:           util.SafeStringConnector(plan.UpdateAccountJson.ValueString()),
		ENABLEACCOUNTJSON:           util.SafeStringConnector(plan.EnableAccountJson.ValueString()),
		DISABLEACCOUNTJSON:          util.SafeStringConnector(plan.DisableAccountJson.ValueString()),
		REMOVEACCOUNTJSON:           util.SafeStringConnector(plan.RemoveAccessJson.ValueString()),
		ADDACCESSJSON:               util.SafeStringConnector(plan.AddAccessJson.ValueString()),
		REMOVEACCESSJSON:            util.SafeStringConnector(plan.RemoveAccountJson.ValueString()),
		RESETANDCHANGEPASSWRDJSON:   util.SafeStringConnector(plan.ResetAndChangePasswrdJson.ValueString()),
		CREATEGROUPJSON:             util.SafeStringConnector(plan.CreateGroupJson.ValueString()),
		UPDATEGROUPJSON:             util.SafeStringConnector(plan.UpdateGroupJson.ValueString()),
		REMOVEGROUPJSON:             util.SafeStringConnector(plan.RemoveGroupJson.ValueString()),
		ADDACCESSENTITLEMENTJSON:    util.SafeStringConnector(plan.AddAccessEntitlementJson.ValueString()),
		CUSTOMCONFIGJSON:            util.SafeStringConnector(plan.CustomConfigJson.ValueString()),
		REMOVEACCESSENTITLEMENTJSON: util.SafeStringConnector(plan.RemoveAccessEntitlementJson.ValueString()),
		CREATESERVICEACCOUNTJSON:    util.SafeStringConnector(plan.CreateServiceAccountJson.ValueString()),
		UPDATESERVICEACCOUNTJSON:    util.SafeStringConnector(plan.UpdateServiceAccountJson.ValueString()),
		REMOVESERVICEACCOUNTJSON:    util.SafeStringConnector(plan.RemoveServiceAccountJson.ValueString()),
		PAM_CONFIG:                  util.SafeStringConnector(plan.PamConfig.ValueString()),
		MODIFYUSERDATAJSON:          util.SafeStringConnector(plan.ModifyUserDataJson.ValueString()),
	}

	adsiConnRequest := openapi.CreateOrUpdateRequest{
		ADSIConnector: &adsiConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, httpResp, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adsiConnRequest).Execute()
	if err != nil {
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)
	// Assign ID and result to the plan
	plan.ID = types.StringValue("test-connection-" + plan.ConnectionName.ValueString())

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)

	// Set the individual fields
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)
	// resultObj := map[string]string{
	// 	"msg":        msgValue,
	// 	"error_code": errorCodeValue,
	// }
	// resultJSON, err := util.MarshalDeterministic(resultObj)
	// if err != nil {
	// 	resp.Diagnostics.AddError(
	// 		"Error Marshaling Result",
	// 		fmt.Sprintf("Could not marshal API response: %v", err),
	// 	)
	// 	log.Printf("JSON Marshalling failed: ", err)
	// 	return
	// }
	// plan.Result = types.StringValue(string(resultJSON))
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *adsiConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// If the API does not support a separate read operation, you can pass through the state.
}

func (r *adsiConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ADSIConnectorResourceModel

	// Extract plan from request
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient

	adsiConn := openapi.ADSIConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "ADSI",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.SafeStringConnector(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnector(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnector(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		URL:                         plan.URL.ValueString(),
		USERNAME:                    plan.Username.ValueString(),
		PASSWORD:                    plan.Password.ValueString(),
		CONNECTION_URL:              plan.ConnectionUrl.ValueString(),
		FORESTLIST:                  plan.ForestList.ValueString(),
		PROVISIONING_URL:            util.SafeStringConnector(plan.ProvisioningUrl.ValueString()),
		DEFAULT_USER_ROLE:           util.SafeStringConnector(plan.DefaultUserRole.ValueString()),
		UPDATEUSERJSON:              util.SafeStringConnector(plan.UpdateUserJson.ValueString()),
		ENDPOINTS_FILTER:            util.SafeStringConnector(plan.EndpointsFilter.ValueString()),
		SEARCHFILTER:                util.SafeStringConnector(plan.SearchFilter.ValueString()),
		OBJECTFILTER:                util.SafeStringConnector(plan.ObjectFilter.ValueString()),
		ACCOUNT_ATTRIBUTE:           util.SafeStringConnector(plan.AccountAttribute.ValueString()),
		STATUS_THRESHOLD_CONFIG:     util.SafeStringConnector(plan.StatusThresholdConfig.ValueString()),
		ENTITLEMENT_ATTRIBUTE:       util.SafeStringConnector(plan.EntitlementAttribute.ValueString()),
		USER_ATTRIBUTE:              util.SafeStringConnector(plan.UserAttribute.ValueString()),
		GroupSearchBaseDN:           util.SafeStringConnector(plan.GroupSearchBaseDN.ValueString()),
		CHECKFORUNIQUE:              util.SafeStringConnector(plan.CheckForUnique.ValueString()),
		STATUSKEYJSON:               util.SafeStringConnector(plan.StatusKeyJson.ValueString()),
		GroupImportMapping:          util.SafeStringConnector(plan.GroupImportMapping.ValueString()),
		ImportNestedMembership:      util.SafeStringConnector(plan.ImportNestedMembership.ValueString()),
		PAGE_SIZE:                   util.SafeStringConnector(plan.PageSize.ValueString()),
		ACCOUNTNAMERULE:             util.SafeStringConnector(plan.AccountNameRule.ValueString()),
		CREATEACCOUNTJSON:           util.SafeStringConnector(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:           util.SafeStringConnector(plan.UpdateAccountJson.ValueString()),
		ENABLEACCOUNTJSON:           util.SafeStringConnector(plan.EnableAccountJson.ValueString()),
		DISABLEACCOUNTJSON:          util.SafeStringConnector(plan.DisableAccountJson.ValueString()),
		REMOVEACCOUNTJSON:           util.SafeStringConnector(plan.RemoveAccessJson.ValueString()),
		ADDACCESSJSON:               util.SafeStringConnector(plan.AddAccessJson.ValueString()),
		REMOVEACCESSJSON:            util.SafeStringConnector(plan.RemoveAccountJson.ValueString()),
		RESETANDCHANGEPASSWRDJSON:   util.SafeStringConnector(plan.ResetAndChangePasswrdJson.ValueString()),
		CREATEGROUPJSON:             util.SafeStringConnector(plan.CreateGroupJson.ValueString()),
		UPDATEGROUPJSON:             util.SafeStringConnector(plan.UpdateGroupJson.ValueString()),
		REMOVEGROUPJSON:             util.SafeStringConnector(plan.RemoveGroupJson.ValueString()),
		ADDACCESSENTITLEMENTJSON:    util.SafeStringConnector(plan.AddAccessEntitlementJson.ValueString()),
		CUSTOMCONFIGJSON:            util.SafeStringConnector(plan.CustomConfigJson.ValueString()),
		REMOVEACCESSENTITLEMENTJSON: util.SafeStringConnector(plan.RemoveAccessEntitlementJson.ValueString()),
		CREATESERVICEACCOUNTJSON:    util.SafeStringConnector(plan.CreateServiceAccountJson.ValueString()),
		UPDATESERVICEACCOUNTJSON:    util.SafeStringConnector(plan.UpdateServiceAccountJson.ValueString()),
		REMOVESERVICEACCOUNTJSON:    util.SafeStringConnector(plan.RemoveServiceAccountJson.ValueString()),
		PAM_CONFIG:                  util.SafeStringConnector(plan.PamConfig.ValueString()),
		MODIFYUSERDATAJSON:          util.SafeStringConnector(plan.ModifyUserDataJson.ValueString()),
	}

	adsiConnRequest := openapi.CreateOrUpdateRequest{
		ADSIConnector: &adsiConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, httpResp, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adsiConnRequest).Execute()
	if err != nil {
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	// Assign ID and result to the plan
	plan.ID = types.StringValue("test-connection-" + plan.ConnectionName.ValueString())

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)

	// Set the individual fields
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)
	// resultObj := map[string]string{
	// 	"msg":        msgValue,
	// 	"error_code": errorCodeValue,
	// }
	// resultJSON, err := util.MarshalDeterministic(resultObj)
	// if err != nil {
	// 	resp.Diagnostics.AddError(
	// 		"Error Marshaling Result",
	// 		fmt.Sprintf("Could not marshal API response: %v", err),
	// 	)
	// 	log.Printf("JSON Marshalling failed: ", err)
	// 	return
	// }
	// plan.Result = types.StringValue(string(resultJSON))

	// Store state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *adsiConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
