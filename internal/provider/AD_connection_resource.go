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
	ConnectionName     types.String `tfsdk:"connection_name"`
	ConnectionType     types.String `tfsdk:"connection_type"`
	Description        types.String `tfsdk:"description"`
	DefaultSavRoles    types.String `tfsdk:"defaultsavroles"`
	EmailTemplate      types.String `tfsdk:"email_template"`
	VaultConnection    types.String `tfsdk:"vault_connection"`
	VaultConfiguration types.String `tfsdk:"vault_configuration"`
	SaveInVault        types.String `tfsdk:"save_in_vault"`
	Result             types.String `tfsdk:"result"`
	Msg                types.String `tfsdk:"msg"`
	ErrorCode          types.String `tfsdk:"error_code"`
}
type ADConnectorResourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnector
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

// testConnectionResource implements the resource.Resource interface.
type adConnectionResource struct {
	// client *openapi.APIClient
	client *s.Client
	token  string
}

// NewTestConnectionResource returns a new instance of testConnectionResource.
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
				Optional:    true,
				Description: "LDAP or target system URL. Example: \"ldap://uscentral.com:8972/\"",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Description: "System admin username.",
			},
			"password": schema.StringAttribute{
				Required:    true,
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
			"result": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The result of the API call.",
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
	adConn := openapi.ADConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "AD",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.SafeStringConnector(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnector(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnector(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		URL:                         util.SafeStringConnector(plan.URL.ValueString()),
		USERNAME:                    util.SafeStringConnector(plan.Username.ValueString()),
		PASSWORD:                    plan.Password.ValueString(),
		LDAP_OR_AD:                  util.SafeStringConnector(plan.LdapOrAd.ValueString()),
		ENTITLEMENT_ATTRIBUTE:       util.SafeStringConnector(plan.EntitlementAttribute.ValueString()),
		CHECKFORUNIQUE:              util.SafeStringConnector(plan.CheckForUnique.ValueString()),
		GroupSearchBaseDN:           util.SafeStringConnector(plan.GroupSearchBaseDN.ValueString()),
		CreateUpdateMappings:        util.SafeStringConnector(plan.CreateUpdateMappings.ValueString()),
		INCREMENTAL_CONFIG:          util.SafeStringConnector(plan.IncrementalConfig.ValueString()),
		MAX_CHANGENUMBER:            util.SafeStringConnector(plan.MaxChangeNumber.ValueString()),
		READ_OPERATIONAL_ATTRIBUTES: util.SafeStringConnector(plan.ReadOperationalAttributes.ValueString()),
		BASE:                        util.SafeStringConnector(plan.Base.ValueString()),
		DC_LOCATOR:                  util.SafeStringConnector(plan.DcLocator.ValueString()),
		STATUS_THRESHOLD_CONFIG:     util.SafeStringConnector(plan.StatusThresholdConfig.ValueString()),
		REMOVEACCOUNTACTION:         util.SafeStringConnector(plan.RemoveAccountAction.ValueString()),
		ACCOUNT_ATTRIBUTE:           util.SafeStringConnector(plan.AccountAttribute.ValueString()),
		ACCOUNTNAMERULE:             util.SafeStringConnector(plan.AccountNameRule.ValueString()),
		ADVSEARCH:                   util.SafeStringConnector(plan.Advsearch.ValueString()),
		SETDEFAULTPAGESIZE:          util.SafeStringConnector(plan.Setdefaultpagesize.ValueString()),
		RESETANDCHANGEPASSWRDJSON:   util.SafeStringConnector(plan.ResetAndChangePasswrdJson.ValueString()),
		REUSEINACTIVEACCOUNT:        util.SafeStringConnector(plan.ReuseInactiveAccount.ValueString()),
		IMPORTJSON:                  util.SafeStringConnector(plan.ImportJson.ValueString()),
		SUPPORTEMPTYSTRING:          util.SafeStringConnector(plan.SupportEmptyString.ValueString()),
		ENABLEACCOUNTJSON:           util.SafeStringConnector(plan.EnableAccountJson.ValueString()),
		PAGE_SIZE:                   util.SafeStringConnector(plan.PageSize.ValueString()),
		USER_ATTRIBUTE:              util.SafeStringConnector(plan.UserAttribute.ValueString()),
		DEFAULT_USER_ROLE:           util.SafeStringConnector(plan.DefaultUserRole.ValueString()),
		SEARCHFILTER:                util.SafeStringConnector(plan.Searchfilter.ValueString()),
		ENDPOINTS_FILTER:            util.SafeStringConnector(plan.EndpointsFilter.ValueString()),
		CREATEACCOUNTJSON:           util.SafeStringConnector(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:           util.SafeStringConnector(plan.UpdateAccountJson.ValueString()),
		REUSEACCOUNTJSON:            util.SafeStringConnector(plan.ReuseAccountJson.ValueString()),
		ENFORCE_TREE_DELETION:       util.SafeStringConnector(plan.EnforceTreeDeletion.ValueString()),
		ADVANCE_FILTER_JSON:         util.SafeStringConnector(plan.AdvanceFilterJson.ValueString()),
		FILTER:                      util.SafeStringConnector(plan.Filter.ValueString()),
		OBJECTFILTER:                util.SafeStringConnector(plan.Objectfilter.ValueString()),
		UPDATEUSERJSON:              util.SafeStringConnector(plan.UpdateUserJson.ValueString()),
		Saveconnection:              util.SafeStringConnector(plan.SaveConnection.ValueString()),
		Systemname:                  util.SafeStringConnector(plan.SystemName.ValueString()),
		SETRANDOMPASSWORD:           util.SafeStringConnector(plan.Setrandompassword.ValueString()),
		PASSWORD_MIN_LENGTH:         util.SafeStringConnector(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:         util.SafeStringConnector(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:      util.SafeStringConnector(plan.PasswordNoofcapsalpha.ValueString()),
		PASSWORD_NOOFSPLCHARS:       util.SafeStringConnector(plan.PasswordNoofsplchars.ValueString()),
		PASSWORD_NOOFDIGITS:         util.SafeStringConnector(plan.PasswordNoofdigits.ValueString()),
		GroupImportMapping:          util.SafeStringConnector(plan.GroupImportMapping.ValueString()),
		UNLOCKACCOUNTJSON:           util.SafeStringConnector(plan.UnlockAccountJson.ValueString()),
		STATUSKEYJSON:               util.SafeStringConnector(plan.StatusKeyJson.ValueString()),
		DISABLEACCOUNTJSON:          util.SafeStringConnector(plan.DisableAccountJson.ValueString()),
		MODIFYUSERDATAJSON:          util.SafeStringConnector(plan.ModifyUserdataJson.ValueString()),
		ORG_BASE:                    util.SafeStringConnector(plan.OrgBase.ValueString()),
		ORGANIZATION_ATTRIBUTE:      util.SafeStringConnector(plan.OrganizationAttribute.ValueString()),
		CREATEORGJSON:               util.SafeStringConnector(plan.Createorgjson.ValueString()),
		UPDATEORGJSON:               util.SafeStringConnector(plan.Updateorgjson.ValueString()),
		ConfigJSON:                  util.SafeStringConnector(plan.ConfigJson.ValueString()),
		PAM_CONFIG:                  util.SafeStringConnector(plan.PamConfig.ValueString()),
	}

	adConnRequest := openapi.CreateOrUpdateRequest{
		ADConnector: &adConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, httpResp, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adConnRequest).Execute()
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
	resultObj := map[string]string{
		"msg":        msgValue,
		"error_code": errorCodeValue,
	}
	resultJSON, err := util.MarshalDeterministic(resultObj)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Marshaling Result",
			fmt.Sprintf("Could not marshal API response: %v", err),
		)
		log.Printf("JSON Marshalling failed: ", err)
		return
	}
	plan.Result = types.StringValue(string(resultJSON))
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *adConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// If the API does not support a separate read operation, you can pass through the state.
}

func (r *adConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ADConnectorResourceModel

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

	adConn := openapi.ADConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "AD",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.SafeStringConnector(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnector(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnector(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		URL:                         util.SafeStringConnector(plan.URL.ValueString()),
		USERNAME:                    util.SafeStringConnector(plan.Username.ValueString()),
		PASSWORD:                    plan.Password.ValueString(),
		LDAP_OR_AD:                  util.SafeStringConnector(plan.LdapOrAd.ValueString()),
		ENTITLEMENT_ATTRIBUTE:       util.SafeStringConnector(plan.EntitlementAttribute.ValueString()),
		CHECKFORUNIQUE:              util.SafeStringConnector(plan.CheckForUnique.ValueString()),
		GroupSearchBaseDN:           util.SafeStringConnector(plan.GroupSearchBaseDN.ValueString()),
		CreateUpdateMappings:        util.SafeStringConnector(plan.CreateUpdateMappings.ValueString()),
		INCREMENTAL_CONFIG:          util.SafeStringConnector(plan.IncrementalConfig.ValueString()),
		MAX_CHANGENUMBER:            util.SafeStringConnector(plan.MaxChangeNumber.ValueString()),
		READ_OPERATIONAL_ATTRIBUTES: util.SafeStringConnector(plan.ReadOperationalAttributes.ValueString()),
		BASE:                        util.SafeStringConnector(plan.Base.ValueString()),
		DC_LOCATOR:                  util.SafeStringConnector(plan.DcLocator.ValueString()),
		STATUS_THRESHOLD_CONFIG:     util.SafeStringConnector(plan.StatusThresholdConfig.ValueString()),
		REMOVEACCOUNTACTION:         util.SafeStringConnector(plan.RemoveAccountAction.ValueString()),
		ACCOUNT_ATTRIBUTE:           util.SafeStringConnector(plan.AccountAttribute.ValueString()),
		ACCOUNTNAMERULE:             util.SafeStringConnector(plan.AccountNameRule.ValueString()),
		ADVSEARCH:                   util.SafeStringConnector(plan.Advsearch.ValueString()),
		SETDEFAULTPAGESIZE:          util.SafeStringConnector(plan.Setdefaultpagesize.ValueString()),
		RESETANDCHANGEPASSWRDJSON:   util.SafeStringConnector(plan.ResetAndChangePasswrdJson.ValueString()),
		REUSEINACTIVEACCOUNT:        util.SafeStringConnector(plan.ReuseInactiveAccount.ValueString()),
		IMPORTJSON:                  util.SafeStringConnector(plan.ImportJson.ValueString()),
		SUPPORTEMPTYSTRING:          util.SafeStringConnector(plan.SupportEmptyString.ValueString()),
		ENABLEACCOUNTJSON:           util.SafeStringConnector(plan.EnableAccountJson.ValueString()),
		PAGE_SIZE:                   util.SafeStringConnector(plan.PageSize.ValueString()),
		USER_ATTRIBUTE:              util.SafeStringConnector(plan.UserAttribute.ValueString()),
		DEFAULT_USER_ROLE:           util.SafeStringConnector(plan.DefaultUserRole.ValueString()),
		SEARCHFILTER:                util.SafeStringConnector(plan.Searchfilter.ValueString()),
		ENDPOINTS_FILTER:            util.SafeStringConnector(plan.EndpointsFilter.ValueString()),
		CREATEACCOUNTJSON:           util.SafeStringConnector(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:           util.SafeStringConnector(plan.UpdateAccountJson.ValueString()),
		REUSEACCOUNTJSON:            util.SafeStringConnector(plan.ReuseAccountJson.ValueString()),
		ENFORCE_TREE_DELETION:       util.SafeStringConnector(plan.EnforceTreeDeletion.ValueString()),
		ADVANCE_FILTER_JSON:         util.SafeStringConnector(plan.AdvanceFilterJson.ValueString()),
		FILTER:                      util.SafeStringConnector(plan.Filter.ValueString()),
		OBJECTFILTER:                util.SafeStringConnector(plan.Objectfilter.ValueString()),
		UPDATEUSERJSON:              util.SafeStringConnector(plan.UpdateUserJson.ValueString()),
		Saveconnection:              util.SafeStringConnector(plan.SaveConnection.ValueString()),
		Systemname:                  util.SafeStringConnector(plan.SystemName.ValueString()),
		SETRANDOMPASSWORD:           util.SafeStringConnector(plan.Setrandompassword.ValueString()),
		PASSWORD_MIN_LENGTH:         util.SafeStringConnector(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:         util.SafeStringConnector(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:      util.SafeStringConnector(plan.PasswordNoofcapsalpha.ValueString()),
		PASSWORD_NOOFSPLCHARS:       util.SafeStringConnector(plan.PasswordNoofsplchars.ValueString()),
		PASSWORD_NOOFDIGITS:         util.SafeStringConnector(plan.PasswordNoofdigits.ValueString()),
		GroupImportMapping:          util.SafeStringConnector(plan.GroupImportMapping.ValueString()),
		UNLOCKACCOUNTJSON:           util.SafeStringConnector(plan.UnlockAccountJson.ValueString()),
		STATUSKEYJSON:               util.SafeStringConnector(plan.StatusKeyJson.ValueString()),
		DISABLEACCOUNTJSON:          util.SafeStringConnector(plan.DisableAccountJson.ValueString()),
		MODIFYUSERDATAJSON:          util.SafeStringConnector(plan.ModifyUserdataJson.ValueString()),
		ORG_BASE:                    util.SafeStringConnector(plan.OrgBase.ValueString()),
		ORGANIZATION_ATTRIBUTE:      util.SafeStringConnector(plan.OrganizationAttribute.ValueString()),
		CREATEORGJSON:               util.SafeStringConnector(plan.Createorgjson.ValueString()),
		UPDATEORGJSON:               util.SafeStringConnector(plan.Updateorgjson.ValueString()),
		ConfigJSON:                  util.SafeStringConnector(plan.ConfigJson.ValueString()),
		PAM_CONFIG:                  util.SafeStringConnector(plan.PamConfig.ValueString()),
	}

	adConnRequest := openapi.CreateOrUpdateRequest{
		ADConnector: &adConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, httpResp, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adConnRequest).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating AD Connector",
			fmt.Sprintf("Error: %v\nHTTP Response: %v", err, httpResp),
		)
		log.Printf("[ERROR] API Called Failed: ", err)
		return
	}
	// Assign ID and result to the plan
	plan.ID = types.StringValue("test-connection-" + plan.ConnectionName.ValueString())

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)

	// Set the individual fields
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)
	resultObj := map[string]string{
		"msg":        msgValue,
		"error_code": errorCodeValue,
	}
	resultJSON, err := util.MarshalDeterministic(resultObj)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Marshaling Result",
			fmt.Sprintf("Could not marshal API response: %v", err),
		)
		log.Printf("JSON Marshalling failed: ", err)
		return
	}
	plan.Result = types.StringValue(string(resultJSON))

	// Store state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *adConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
