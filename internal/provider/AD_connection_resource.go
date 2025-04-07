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

	s "github.com/saviynt/saviynt-api-go-client"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BaseConnector struct {
	ConnectionKey      types.Int64  `tfsdk:"connection_key"`
	ConnectionName     types.String `tfsdk:"connection_name"`
	ConnectionType     types.String `tfsdk:"connection_type"`
	Description        types.String `tfsdk:"description"`
	DefaultSavRoles    types.String `tfsdk:"defaultsavroles"`
	EmailTemplate      types.String `tfsdk:"email_template"`
	VaultConnection    types.String `tfsdk:"vault_connection"`
	VaultConfiguration types.String `tfsdk:"vault_configuration"`
	SaveInVault        types.String `tfsdk:"save_in_vault"`
	Msg                types.String `tfsdk:"msg"`
	ErrorCode          types.String `tfsdk:"error_code"`
}
type ADConnectorResourceModel struct {
	BaseConnector
	ID                        types.String `tfsdk:"id"`
	URL                       types.String `tfsdk:"url"`
	Username                  types.String `tfsdk:"username"`
	Password                  types.String `tfsdk:"password"`
	LdapOrAd                  types.String `tfsdk:"ldap_or_ad"`
	EntitlementAttribute      types.String `tfsdk:"entitlement_attribute"`
	CheckForUnique            types.String `tfsdk:"check_for_unique"`
	GroupSearchBaseDN         types.String `tfsdk:"group_search_base_dn"`
	CreateUpdateMappings      types.String `tfsdk:"create_update_mappings"`
	IncrementalConfig         types.String `tfsdk:"incremental_config"`
	MaxChangeNumber           types.String `tfsdk:"max_changenumber"`
	ReadOperationalAttributes types.String `tfsdk:"read_operational_attributes"`
	Base                      types.String `tfsdk:"base"`
	DcLocator                 types.String `tfsdk:"dc_locator"`
	StatusThresholdConfig     types.String `tfsdk:"status_threshold_config"`
	RemoveAccountAction       types.String `tfsdk:"remove_account_action"`
	AccountAttribute          types.String `tfsdk:"account_attribute"`
	AccountNameRule           types.String `tfsdk:"account_name_rule"`
	Advsearch                 types.String `tfsdk:"advsearch"`
	Setdefaultpagesize        types.String `tfsdk:"setdefaultpagesize"`
	ResetAndChangePasswrdJson types.String `tfsdk:"reset_and_change_passwrd_json"`
	ReuseInactiveAccount      types.String `tfsdk:"reuse_inactive_account"`
	ImportJson                types.String `tfsdk:"import_json"`
	SupportEmptyString        types.String `tfsdk:"support_empty_string"`
	EnableAccountJson         types.String `tfsdk:"enable_account_json"`
	PageSize                  types.String `tfsdk:"page_size"`
	UserAttribute             types.String `tfsdk:"user_attribute"`
	DefaultUserRole           types.String `tfsdk:"default_user_role"`
	Searchfilter              types.String `tfsdk:"searchfilter"`
	EndpointsFilter           types.String `tfsdk:"endpoints_filter"`
	CreateAccountJson         types.String `tfsdk:"create_account_json"`
	UpdateAccountJson         types.String `tfsdk:"update_account_json"`
	ReuseAccountJson          types.String `tfsdk:"reuse_account_json"`
	EnforceTreeDeletion       types.String `tfsdk:"enforce_tree_deletion"`
	AdvanceFilterJson         types.String `tfsdk:"advance_filter_json"`
	Filter                    types.String `tfsdk:"filter"`
	Objectfilter              types.String `tfsdk:"objectfilter"`
	UpdateUserJson            types.String `tfsdk:"update_user_json"`
	SaveConnection            types.String `tfsdk:"save_connection"`
	SystemName                types.String `tfsdk:"system_name"`
	Setrandompassword         types.String `tfsdk:"set_random_password"`
	PasswordMinLength         types.String `tfsdk:"password_min_length"`
	PasswordMaxLength         types.String `tfsdk:"password_max_length"`
	PasswordNoofcapsalpha     types.String `tfsdk:"password_noofcapsalpha"`
	PasswordNoofsplchars      types.String `tfsdk:"password_noofsplchars"`
	PasswordNoofdigits        types.String `tfsdk:"password_noofdigits"`
	GroupImportMapping        types.String `tfsdk:"group_import_mapping"`
	UnlockAccountJson         types.String `tfsdk:"unlock_account_json"`
	StatusKeyJson             types.String `tfsdk:"status_key_json"`
	DisableAccountJson        types.String `tfsdk:"disable_account_json"`
	ModifyUserdataJson        types.String `tfsdk:"modify_user_data_json"`
	OrgBase                   types.String `tfsdk:"org_base"`
	OrganizationAttribute     types.String `tfsdk:"organization_attribute"`
	Createorgjson             types.String `tfsdk:"create_org_json"`
	Updateorgjson             types.String `tfsdk:"update_org_json"`
	ConfigJson                types.String `tfsdk:"config_json"`
	PamConfig                 types.String `tfsdk:"pam_config"`
}

type adConnectionResource struct {
	// client *openapi.APIClient
	client *s.Client
	token  string
}

func ADNewTestConnectionResource() resource.Resource {
	return &adConnectionResource{}
}

func (r *adConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_ad_connection_resource"
}

func (r *adConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and Manage Connections",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
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
			"url": schema.StringAttribute{
				Optional:    true,
				Description: "LDAP or target system URL. Example: \"ldap://uscentral.com:8972/\"",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Description: "System admin username.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Description: "Set the Password.",
			},
			"ldap_or_ad": schema.StringAttribute{
				Optional:    true,
				Description: "Type of Endpoint - LDAP or AD. Default is 'AD'. Example: \"AD\"",
			},
			"entitlement_attribute": schema.StringAttribute{
				Optional:    true,
				Description: "Attribute used for entitlements. Example: \"memberOf\"",
			},
			"check_for_unique": schema.StringAttribute{
				Optional:    true,
				Description: "Uniqueness validation rule JSON. Example: '{\"sAMAccountName\":\"${task.accountName}\"}'",
			},
			"group_search_base_dn": schema.StringAttribute{
				Optional:    true,
				Description: "Base DN for group search. Example: \"CN=Users,DC=Saviynt,DC=ABC,DC=Com\"",
			},
			"create_update_mappings": schema.StringAttribute{
				Optional:    true,
				Description: "Mapping for group creation/updation (JSON string). Example: '{\"cn\":\"${role?.customproperty27}\",\"objectCategory\":\"CN=Group,CN=Schema,CN=Configuration,...}'",
			},
			"incremental_config": schema.StringAttribute{
				Optional:    true,
				Description: "Incremental import configuration.",
			},
			"max_changenumber": schema.StringAttribute{
				Optional:    true,
				Description: "Maximum change number. Example: \"4\"",
			},
			"read_operational_attributes": schema.StringAttribute{
				Optional:    true,
				Description: "Flag for reading operational attributes. Example: \"FALSE\"",
			},
			"base": schema.StringAttribute{
				Optional:    true,
				Description: "LDAP base DN. Example: \"CN=Users,DC=Saviynt,DC=ABC,DC=Com\"",
			},
			"dc_locator": schema.StringAttribute{
				Optional:    true,
				Description: "Domain controller locator.",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Description: "JSON configuration for status thresholds. Example: '{\"statusAndThresholdConfig\":{...}}'",
			},
			"remove_account_action": schema.StringAttribute{
				Optional:    true,
				Description: "Action on account removal. Example: '{\"removeAction\":\"DELETE\"}'",
			},
			"account_attribute": schema.StringAttribute{
				Optional:    true,
				Description: "Mapping for LDAP user to EIC account attribute. Example: '[\"ACCOUNTID::objectGUID#Binary\", \"NAME::sAMAccountName#String\", ...]'",
			},
			"account_name_rule": schema.StringAttribute{
				Optional:    true,
				Description: "Rule to generate account name. Example: \"uid=${task.accountName.toString().toLowerCase()},ou=People,dc=racf,dc=com\"",
			},
			"advsearch": schema.StringAttribute{
				Optional:    true,
				Description: "Advanced search settings.",
			},
			"setdefaultpagesize": schema.StringAttribute{
				Optional:    true,
				Description: "Default page size setting. Example: \"FALSE\"",
			},
			"reset_and_change_passwrd_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for reset/change password actions. Example: '{\"RESET\":{\"pwdLastSet\":\"0\",\"title\":\"password reset\"},\"CHANGE\":{\"pwdLastSet\":\"-1\",\"title\":\"password changed\"}}'",
			},
			"reuse_inactive_account": schema.StringAttribute{
				Optional:    true,
				Description: "Reuse inactive account flag. Example: \"TRUE\"",
			},
			"import_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON import configuration. Example: '{\"envproperties\":{\"com.sun.jndi.ldap.connect.timeout\":\"10000\",...}}'",
			},
			"support_empty_string": schema.StringAttribute{
				Optional:    true,
				Description: "Flag for sending empty values. Example: \"FALSE\"",
			},
			"enable_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON configuration to enable account actions. Example: '{\"USEDNFROMACCOUNT\":\"NO\", ...}'",
			},
			"page_size": schema.StringAttribute{
				Optional:    true,
				Description: "LDAP page size. Example: \"1000\"",
			},
			"user_attribute": schema.StringAttribute{
				Optional:    true,
				Description: "Mapping for LDAP user to EIC user attribute. Example: '[\"USERNAME::sAMAccountName#String\", ...]'",
			},
			"default_user_role": schema.StringAttribute{
				Optional:    true,
				Description: "Default SAV Role for imported users. Example: \"ROLE_TASK_ADMIN\"",
			},
			"searchfilter": schema.StringAttribute{
				Optional:    true,
				Description: "LDAP search filter for users. Example: \"OU=Users,DC=domainname,DC=com\"",
			},
			"endpoints_filter": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration for child endpoints.",
			},
			"create_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to create an account. Example: '{\"cn\":\"${cn}\",\"displayname\":\"${user.displayname}\", ...}'",
			},
			"update_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to update an account. Example: '{\"uid\":\"${task.accountName.toString().toLowerCase()}\", ...}'",
			},
			"reuse_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to reuse an account. Example: '{\"ATTRIBUTESTOCHECK\":{\"userAccountControl\":\"514\",...}}'",
			},
			"enforce_tree_deletion": schema.StringAttribute{
				Optional:    true,
				Description: "Enforce tree deletion flag. Example: \"TRUE\"",
			},
			"advance_filter_json": schema.StringAttribute{
				Optional:    true,
				Description: "Advanced filter JSON configuration.",
			},
			"filter": schema.StringAttribute{
				Optional:    true,
				Description: "Simple filter string.",
			},
			"objectfilter": schema.StringAttribute{
				Optional:    true,
				Description: "LDAP object filter. Example: \"(objectClass=inetorgperson)\"",
			},
			"update_user_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to update a user. Example: '{\"mail\":\"${user.email}\", ...}'",
			},
			"save_connection": schema.StringAttribute{
				Optional:    true,
				Description: "Flag to permanently save connection. Example: \"N\"",
			},
			"system_name": schema.StringAttribute{
				Optional:    true,
				Description: "Associated system name. Example: \"Dummyapplication\"",
			},
			"set_random_password": schema.StringAttribute{
				Optional:    true,
				Description: "Option to set a random password.",
			},
			"password_min_length": schema.StringAttribute{
				Optional:    true,
				Description: "Minimum password length. Example: \"8\"",
			},
			"password_max_length": schema.StringAttribute{
				Optional:    true,
				Description: "Maximum password length. Example: \"12\"",
			},
			"password_noofcapsalpha": schema.StringAttribute{
				Optional:    true,
				Description: "Number of capital letters required. Example: \"2\"",
			},
			"password_noofsplchars": schema.StringAttribute{
				Optional:    true,
				Description: "Number of special characters required. Example: \"1\"",
			},
			"password_noofdigits": schema.StringAttribute{
				Optional:    true,
				Description: "Number of digits required. Example: \"5\"",
			},
			"group_import_mapping": schema.StringAttribute{
				Optional:    true,
				Description: "JSON mapping for LDAP groups. Example: '{\"entitlementTypeName\":\"memberOf\", ...}'",
			},
			"unlock_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to unlock accounts. Example: '{\"lockoutTime\":\"0\"}'",
			},
			"status_key_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for account status keys. Example: '{\"STATUS_ACTIVE\":[\"512\",\"544\"], ...}'",
			},
			"disable_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to disable an account. Example: '{\"userAccountControl\":\"546\", ...}'",
			},
			"modify_user_data_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for inline user data transformation.",
			},
			"org_base": schema.StringAttribute{
				Optional:    true,
				Description: "Organization BASE for provisioning.",
			},
			"organization_attribute": schema.StringAttribute{
				Optional:    true,
				Description: "Organization attributes.",
			},
			"create_org_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for organization creation.",
			},
			"update_org_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for organization update.",
			},
			"config_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for connection timeout configuration. Example: '{\"connectionTimeoutConfig\":{\"connectionTimeout\":10,\"readTimeout\":50,\"retryWait\":2,\"retryCount\":3}}'",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for PAM bootstrap configuration. Example: '{\"Connection\":\"AD\",...}'",
			},
			"msg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "A message indicating the outcome of the operation.",
			},
			"error_code": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
			},
		},
	}
}

func (r *adConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *adConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ADConnectorResourceModel
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
	adConn := openapi.ADConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "AD",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.SafeStringConnectorForNullHandling(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnectorForNullHandling(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnectorForNullHandling(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		URL:                         util.SafeStringConnectorForNullHandling(plan.URL.ValueString()),
		USERNAME:                    util.SafeStringConnectorForNullHandling(plan.Username.ValueString()),
		PASSWORD:                    plan.Password.ValueString(),
		LDAP_OR_AD:                  util.SafeStringConnectorForNullHandling(plan.LdapOrAd.ValueString()),
		ENTITLEMENT_ATTRIBUTE:       util.SafeStringConnectorForNullHandling(plan.EntitlementAttribute.ValueString()),
		CHECKFORUNIQUE:              util.SafeStringConnectorForNullHandling(plan.CheckForUnique.ValueString()),
		GroupSearchBaseDN:           util.SafeStringConnectorForNullHandling(plan.GroupSearchBaseDN.ValueString()),
		CreateUpdateMappings:        util.SafeStringConnectorForNullHandling(plan.CreateUpdateMappings.ValueString()),
		INCREMENTAL_CONFIG:          util.SafeStringConnectorForNullHandling(plan.IncrementalConfig.ValueString()),
		MAX_CHANGENUMBER:            util.SafeStringConnectorForNullHandling(plan.MaxChangeNumber.ValueString()),
		READ_OPERATIONAL_ATTRIBUTES: util.SafeStringConnectorForNullHandling(plan.ReadOperationalAttributes.ValueString()),
		BASE:                        util.SafeStringConnectorForNullHandling(plan.Base.ValueString()),
		DC_LOCATOR:                  util.SafeStringConnectorForNullHandling(plan.DcLocator.ValueString()),
		STATUS_THRESHOLD_CONFIG:     util.SafeStringConnectorForNullHandling(plan.StatusThresholdConfig.ValueString()),
		REMOVEACCOUNTACTION:         util.SafeStringConnectorForNullHandling(plan.RemoveAccountAction.ValueString()),
		ACCOUNT_ATTRIBUTE:           util.SafeStringConnectorForNullHandling(plan.AccountAttribute.ValueString()),
		ACCOUNTNAMERULE:             util.SafeStringConnectorForNullHandling(plan.AccountNameRule.ValueString()),
		ADVSEARCH:                   util.SafeStringConnectorForNullHandling(plan.Advsearch.ValueString()),
		SETDEFAULTPAGESIZE:          util.SafeStringConnectorForNullHandling(plan.Setdefaultpagesize.ValueString()),
		RESETANDCHANGEPASSWRDJSON:   util.SafeStringConnectorForNullHandling(plan.ResetAndChangePasswrdJson.ValueString()),
		REUSEINACTIVEACCOUNT:        util.SafeStringConnectorForNullHandling(plan.ReuseInactiveAccount.ValueString()),
		IMPORTJSON:                  util.SafeStringConnectorForNullHandling(plan.ImportJson.ValueString()),
		SUPPORTEMPTYSTRING:          util.SafeStringConnectorForNullHandling(plan.SupportEmptyString.ValueString()),
		ENABLEACCOUNTJSON:           util.SafeStringConnectorForNullHandling(plan.EnableAccountJson.ValueString()),
		PAGE_SIZE:                   util.SafeStringConnectorForNullHandling(plan.PageSize.ValueString()),
		USER_ATTRIBUTE:              util.SafeStringConnectorForNullHandling(plan.UserAttribute.ValueString()),
		DEFAULT_USER_ROLE:           util.SafeStringConnectorForNullHandling(plan.DefaultUserRole.ValueString()),
		SEARCHFILTER:                util.SafeStringConnectorForNullHandling(plan.Searchfilter.ValueString()),
		ENDPOINTS_FILTER:            util.SafeStringConnectorForNullHandling(plan.EndpointsFilter.ValueString()),
		CREATEACCOUNTJSON:           util.SafeStringConnectorForNullHandling(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:           util.SafeStringConnectorForNullHandling(plan.UpdateAccountJson.ValueString()),
		REUSEACCOUNTJSON:            util.SafeStringConnectorForNullHandling(plan.ReuseAccountJson.ValueString()),
		ENFORCE_TREE_DELETION:       util.SafeStringConnectorForNullHandling(plan.EnforceTreeDeletion.ValueString()),
		ADVANCE_FILTER_JSON:         util.SafeStringConnectorForNullHandling(plan.AdvanceFilterJson.ValueString()),
		FILTER:                      util.SafeStringConnectorForNullHandling(plan.Filter.ValueString()),
		OBJECTFILTER:                util.SafeStringConnectorForNullHandling(plan.Objectfilter.ValueString()),
		UPDATEUSERJSON:              util.SafeStringConnectorForNullHandling(plan.UpdateUserJson.ValueString()),
		Saveconnection:              util.SafeStringConnectorForNullHandling(plan.SaveConnection.ValueString()),
		Systemname:                  util.SafeStringConnectorForNullHandling(plan.SystemName.ValueString()),
		SETRANDOMPASSWORD:           util.SafeStringConnectorForNullHandling(plan.Setrandompassword.ValueString()),
		PASSWORD_MIN_LENGTH:         util.SafeStringConnectorForNullHandling(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:         util.SafeStringConnectorForNullHandling(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:      util.SafeStringConnectorForNullHandling(plan.PasswordNoofcapsalpha.ValueString()),
		PASSWORD_NOOFSPLCHARS:       util.SafeStringConnectorForNullHandling(plan.PasswordNoofsplchars.ValueString()),
		PASSWORD_NOOFDIGITS:         util.SafeStringConnectorForNullHandling(plan.PasswordNoofdigits.ValueString()),
		GroupImportMapping:          util.SafeStringConnectorForNullHandling(plan.GroupImportMapping.ValueString()),
		UNLOCKACCOUNTJSON:           util.SafeStringConnectorForNullHandling(plan.UnlockAccountJson.ValueString()),
		STATUSKEYJSON:               util.SafeStringConnectorForNullHandling(plan.StatusKeyJson.ValueString()),
		DISABLEACCOUNTJSON:          util.SafeStringConnectorForNullHandling(plan.DisableAccountJson.ValueString()),
		MODIFYUSERDATAJSON:          util.SafeStringConnectorForNullHandling(plan.ModifyUserdataJson.ValueString()),
		ORG_BASE:                    util.SafeStringConnectorForNullHandling(plan.OrgBase.ValueString()),
		ORGANIZATION_ATTRIBUTE:      util.SafeStringConnectorForNullHandling(plan.OrganizationAttribute.ValueString()),
		CREATEORGJSON:               util.SafeStringConnectorForNullHandling(plan.Createorgjson.ValueString()),
		UPDATEORGJSON:               util.SafeStringConnectorForNullHandling(plan.Updateorgjson.ValueString()),
		ConfigJSON:                  util.SafeStringConnectorForNullHandling(plan.ConfigJson.ValueString()),
		PAM_CONFIG:                  util.SafeStringConnectorForNullHandling(plan.PamConfig.ValueString()),
	}

	adConnRequest := openapi.CreateOrUpdateRequest{
		ADConnector: &adConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adConnRequest).Execute()
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

func (r *adConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ADConnectorResourceModel

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
		log.Printf("Unable to read properly baby gorl")
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)
	state.ConnectionKey = types.Int64Value(int64(*apiResp.ADConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ADConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.ADConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.ADConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectiontype)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.ADConnectionResponse.Emailtemplate)
	state.URL = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.URL)
	state.ConnectionType = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectiontype)
	state.Advsearch = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ADVSEARCH)
	state.CreateAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	state.DisableAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON)
	state.GroupSearchBaseDN = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.GroupSearchBaseDN)
	state.PasswordNoofsplchars = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS)
	state.PasswordNoofdigits = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS)
	state.StatusKeyJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.STATUSKEYJSON)
	state.Searchfilter = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SEARCHFILTER)
	state.ConfigJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ConfigJSON)
	state.RemoveAccountAction = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.REMOVEACCOUNTACTION)
	state.AccountAttribute = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTE)
	state.AccountNameRule = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ACCOUNTNAMERULE)
	state.Username = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.USERNAME)
	state.LdapOrAd = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.LDAP_OR_AD)
	state.EntitlementAttribute = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	state.Setrandompassword = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SETRANDOMPASSWORD)
	state.PasswordMinLength = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH)
	state.PasswordMaxLength = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH)
	state.PasswordNoofcapsalpha = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA)
	state.Setdefaultpagesize = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SETDEFAULTPAGESIZE)
	state.ReuseInactiveAccount = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.REUSEINACTIVEACCOUNT)
	state.ImportJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.IMPORTJSON)
	state.CreateUpdateMappings = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CreateUpdateMappings)
	state.AdvanceFilterJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ADVANCE_FILTER_JSON)
	state.PamConfig = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.PageSize = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PAGE_SIZE)
	state.Base = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.BASE)
	state.DcLocator = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.DC_LOCATOR)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.ResetAndChangePasswrdJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.RESETANDCHANGEPASSWRDJSON)
	state.SupportEmptyString = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SUPPORTEMPTYSTRING)
	state.ReadOperationalAttributes = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.READ_OPERATIONAL_ATTRIBUTES)
	state.EnableAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	state.UserAttribute = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.USER_ATTRIBUTE)
	state.DefaultUserRole = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.DEFAULT_USER_ROLE)
	state.EndpointsFilter = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	state.UpdateAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	state.ReuseAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.REUSEACCOUNTJSON)
	state.EnforceTreeDeletion = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENFORCE_TREE_DELETION)
	state.Filter = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.FILTER)
	state.Objectfilter = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.OBJECTFILTER)
	state.UpdateUserJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UPDATEUSERJSON)
	state.SaveConnection = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.Saveconnection)
	state.SystemName = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.Systemname)
	state.GroupImportMapping = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.GroupImportMapping)
	state.UnlockAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UNLOCKACCOUNTJSON)
	state.ModifyUserdataJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.OrgBase = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ORG_BASE)
	state.OrganizationAttribute = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ORGANIZATION_ATTRIBUTE)
	state.Createorgjson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CREATEORGJSON)
	state.Updateorgjson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UPDATEORGJSON)
	state.MaxChangeNumber = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.MAX_CHANGENUMBER)
	state.IncrementalConfig = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.INCREMENTAL_CONFIG)
	state.CheckForUnique = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CHECKFORUNIQUE)
	apiMessage := util.SafeDeref(apiResp.ADConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.ADConnectionResponse.Errorcode)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}

func (r *adConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ADConnectorResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
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
	adConn := openapi.ADConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "AD",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.SafeStringConnectorForNullHandling(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnectorForNullHandling(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnectorForNullHandling(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		URL:                         util.SafeStringConnectorForNullHandling(plan.URL.ValueString()),
		USERNAME:                    util.SafeStringConnectorForNullHandling(plan.Username.ValueString()),
		PASSWORD:                    plan.Password.ValueString(),
		LDAP_OR_AD:                  util.SafeStringConnectorForNullHandling(plan.LdapOrAd.ValueString()),
		ENTITLEMENT_ATTRIBUTE:       util.SafeStringConnectorForNullHandling(plan.EntitlementAttribute.ValueString()),
		CHECKFORUNIQUE:              util.SafeStringConnectorForNullHandling(plan.CheckForUnique.ValueString()),
		GroupSearchBaseDN:           util.SafeStringConnectorForNullHandling(plan.GroupSearchBaseDN.ValueString()),
		CreateUpdateMappings:        util.SafeStringConnectorForNullHandling(plan.CreateUpdateMappings.ValueString()),
		INCREMENTAL_CONFIG:          util.SafeStringConnectorForNullHandling(plan.IncrementalConfig.ValueString()),
		MAX_CHANGENUMBER:            util.SafeStringConnectorForNullHandling(plan.MaxChangeNumber.ValueString()),
		READ_OPERATIONAL_ATTRIBUTES: util.SafeStringConnectorForNullHandling(plan.ReadOperationalAttributes.ValueString()),
		BASE:                        util.SafeStringConnectorForNullHandling(plan.Base.ValueString()),
		DC_LOCATOR:                  util.SafeStringConnectorForNullHandling(plan.DcLocator.ValueString()),
		STATUS_THRESHOLD_CONFIG:     util.SafeStringConnectorForNullHandling(plan.StatusThresholdConfig.ValueString()),
		REMOVEACCOUNTACTION:         util.SafeStringConnectorForNullHandling(plan.RemoveAccountAction.ValueString()),
		ACCOUNT_ATTRIBUTE:           util.SafeStringConnectorForNullHandling(plan.AccountAttribute.ValueString()),
		ACCOUNTNAMERULE:             util.SafeStringConnectorForNullHandling(plan.AccountNameRule.ValueString()),
		ADVSEARCH:                   util.SafeStringConnectorForNullHandling(plan.Advsearch.ValueString()),
		SETDEFAULTPAGESIZE:          util.SafeStringConnectorForNullHandling(plan.Setdefaultpagesize.ValueString()),
		RESETANDCHANGEPASSWRDJSON:   util.SafeStringConnectorForNullHandling(plan.ResetAndChangePasswrdJson.ValueString()),
		REUSEINACTIVEACCOUNT:        util.SafeStringConnectorForNullHandling(plan.ReuseInactiveAccount.ValueString()),
		IMPORTJSON:                  util.SafeStringConnectorForNullHandling(plan.ImportJson.ValueString()),
		SUPPORTEMPTYSTRING:          util.SafeStringConnectorForNullHandling(plan.SupportEmptyString.ValueString()),
		ENABLEACCOUNTJSON:           util.SafeStringConnectorForNullHandling(plan.EnableAccountJson.ValueString()),
		PAGE_SIZE:                   util.SafeStringConnectorForNullHandling(plan.PageSize.ValueString()),
		USER_ATTRIBUTE:              util.SafeStringConnectorForNullHandling(plan.UserAttribute.ValueString()),
		DEFAULT_USER_ROLE:           util.SafeStringConnectorForNullHandling(plan.DefaultUserRole.ValueString()),
		SEARCHFILTER:                util.SafeStringConnectorForNullHandling(plan.Searchfilter.ValueString()),
		ENDPOINTS_FILTER:            util.SafeStringConnectorForNullHandling(plan.EndpointsFilter.ValueString()),
		CREATEACCOUNTJSON:           util.SafeStringConnectorForNullHandling(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:           util.SafeStringConnectorForNullHandling(plan.UpdateAccountJson.ValueString()),
		REUSEACCOUNTJSON:            util.SafeStringConnectorForNullHandling(plan.ReuseAccountJson.ValueString()),
		ENFORCE_TREE_DELETION:       util.SafeStringConnectorForNullHandling(plan.EnforceTreeDeletion.ValueString()),
		ADVANCE_FILTER_JSON:         util.SafeStringConnectorForNullHandling(plan.AdvanceFilterJson.ValueString()),
		FILTER:                      util.SafeStringConnectorForNullHandling(plan.Filter.ValueString()),
		OBJECTFILTER:                util.SafeStringConnectorForNullHandling(plan.Objectfilter.ValueString()),
		UPDATEUSERJSON:              util.SafeStringConnectorForNullHandling(plan.UpdateUserJson.ValueString()),
		Saveconnection:              util.SafeStringConnectorForNullHandling(plan.SaveConnection.ValueString()),
		Systemname:                  util.SafeStringConnectorForNullHandling(plan.SystemName.ValueString()),
		SETRANDOMPASSWORD:           util.SafeStringConnectorForNullHandling(plan.Setrandompassword.ValueString()),
		PASSWORD_MIN_LENGTH:         util.SafeStringConnectorForNullHandling(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:         util.SafeStringConnectorForNullHandling(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:      util.SafeStringConnectorForNullHandling(plan.PasswordNoofcapsalpha.ValueString()),
		PASSWORD_NOOFSPLCHARS:       util.SafeStringConnectorForNullHandling(plan.PasswordNoofsplchars.ValueString()),
		PASSWORD_NOOFDIGITS:         util.SafeStringConnectorForNullHandling(plan.PasswordNoofdigits.ValueString()),
		GroupImportMapping:          util.SafeStringConnectorForNullHandling(plan.GroupImportMapping.ValueString()),
		UNLOCKACCOUNTJSON:           util.SafeStringConnectorForNullHandling(plan.UnlockAccountJson.ValueString()),
		STATUSKEYJSON:               util.SafeStringConnectorForNullHandling(plan.StatusKeyJson.ValueString()),
		DISABLEACCOUNTJSON:          util.SafeStringConnectorForNullHandling(plan.DisableAccountJson.ValueString()),
		MODIFYUSERDATAJSON:          util.SafeStringConnectorForNullHandling(plan.ModifyUserdataJson.ValueString()),
		ORG_BASE:                    util.SafeStringConnectorForNullHandling(plan.OrgBase.ValueString()),
		ORGANIZATION_ATTRIBUTE:      util.SafeStringConnectorForNullHandling(plan.OrganizationAttribute.ValueString()),
		CREATEORGJSON:               util.SafeStringConnectorForNullHandling(plan.Createorgjson.ValueString()),
		UPDATEORGJSON:               util.SafeStringConnectorForNullHandling(plan.Updateorgjson.ValueString()),
		ConfigJSON:                  util.SafeStringConnectorForNullHandling(plan.ConfigJson.ValueString()),
		PAM_CONFIG:                  util.SafeStringConnectorForNullHandling(plan.PamConfig.ValueString()),
	}

	adConnRequest := openapi.CreateOrUpdateRequest{
		ADConnector: &adConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("Unable to update properly baby gorl")
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("sab sahi chal raha hai update mei")
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *adConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
