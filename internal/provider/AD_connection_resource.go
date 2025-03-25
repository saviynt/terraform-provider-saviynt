package provider

import (
	"context"
	"fmt"
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
	ConnectionName     *string `tfsdk:"connection_name"`
	ConnectionType     string  `tfsdk:"connection_type"`
	Description        *string `tfsdk:"description"`
	DefaultSavRoles    *string `tfsdk:"defaultsavroles"`
	EmailTemplate      *string `tfsdk:"email_template"`
	SSLCertificate     *string `tfsdk:"ssl_certificate"`
	VaultConnection    *string `tfsdk:"vault_connection"`
	VaultConfiguration *string `tfsdk:"vault_configuration"`
	SaveInVault        *string `tfsdk:"save_in_vault"`
}
type ADConnectorResourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnector

	// Optional additional fields
	URL                       *string `tfsdk:"url"`                           // LDAP or target system URL (example: "ldap://uscentral.com:8972/")
	Username                  *string `tfsdk:"username"`                      // System admin username
	Password                  string  `tfsdk:"password"`                      // Set the password
	LdapOrAd                  *string `tfsdk:"ldap_or_ad"`                    // e.g., default "AD"
	EntitlementAttribute      *string `tfsdk:"entitlement_attribute"`         // e.g., "memberOf"
	CheckForUnique            *string `tfsdk:"check_for_unique"`              // Uniqueness validation rule JSON (example: "{\"sAMAccountName\":\"${task.accountName}\"}")
	GroupSearchBaseDN         *string `tfsdk:"group_search_base_dn"`          // Base DN for group search
	CreateUpdateMappings      *string `tfsdk:"create_update_mappings"`        // Mapping for group creation/updation (JSON string)
	IncrementalConfig         *string `tfsdk:"incremental_config"`            // Incremental import configuration
	MaxChangeNumber           *string `tfsdk:"max_changenumber"`              // Maximum change number
	ReadOperationalAttributes *string `tfsdk:"read_operational_attributes"`   // Flag for reading operational attributes
	Base                      *string `tfsdk:"base"`                          // LDAP base DN
	DcLocator                 *string `tfsdk:"dc_locator"`                    // Domain controller locator
	StatusThresholdConfig     *string `tfsdk:"status_threshold_config"`       // JSON configuration for status thresholds
	RemoveAccountAction       *string `tfsdk:"remove_account_action"`         // Action on account removal
	AccountAttribute          *string `tfsdk:"account_attribute"`             // Mapping for LDAP user to EIC account attribute
	AccountNameRule           *string `tfsdk:"account_name_rule"`             // Rule to generate account name
	Advsearch                 *string `tfsdk:"advsearch"`                     // Advanced search settings
	Setdefaultpagesize        *string `tfsdk:"setdefaultpagesize"`            // Default page size setting
	ResetAndChangePasswrdJson *string `tfsdk:"reset_and_change_passwrd_json"` // JSON for reset/change password actions
	ReuseInactiveAccount      *string `tfsdk:"reuse_inactive_account"`        // Reuse inactive account flag
	ImportJson                *string `tfsdk:"import_json"`                   // JSON import configuration
	SupportEmptyString        *string `tfsdk:"support_empty_string"`          // Flag for sending empty strings
	EnableAccountJson         *string `tfsdk:"enable_account_json"`           // JSON configuration to enable account
	PageSize                  *string `tfsdk:"page_size"`                     // LDAP page size
	UserAttribute             *string `tfsdk:"user_attribute"`                // Mapping for LDAP user to EIC user attribute
	DefaultUserRole           *string `tfsdk:"default_user_role"`             // Default user role for imported users
	Searchfilter              *string `tfsdk:"searchfilter"`                  // LDAP search filter for users
	EndpointsFilter           *string `tfsdk:"endpoints_filter"`              // Configuration for child endpoints
	CreateAccountJson         *string `tfsdk:"create_account_json"`           // JSON to create an account
	UpdateAccountJson         *string `tfsdk:"update_account_json"`           // JSON to update an account
	ReuseAccountJson          *string `tfsdk:"reuse_account_json"`            // JSON to reuse an account
	EnforceTreeDeletion       *string `tfsdk:"enforce_tree_deletion"`         // Enforce tree deletion flag
	AdvanceFilterJson         *string `tfsdk:"advance_filter_json"`           // Advanced filter JSON configuration
	Filter                    *string `tfsdk:"filter"`                        // Simple filter string
	Objectfilter              *string `tfsdk:"objectfilter"`                  // LDAP object filter (example: "(objectClass=inetorgperson)")
	UpdateUserJson            *string `tfsdk:"update_user_json"`              // JSON to update a user
	SaveConnection            *string `tfsdk:"save_connection"`               // Flag to permanently save connection
	SystemName                *string `tfsdk:"system_name"`                   // Associated system name
	Setrandompassword         *string `tfsdk:"set_random_password"`           // Option to set a random password
	PasswordMinLength         *string `tfsdk:"password_min_length"`           // Minimum password length (example: "8")
	PasswordMaxLength         *string `tfsdk:"password_max_length"`           // Maximum password length (example: "12")
	PasswordNoofcapsalpha     *string `tfsdk:"password_noofcapsalpha"`        // Number of capital letters required
	PasswordNoofsplchars      *string `tfsdk:"password_noofsplchars"`         // Number of special characters required
	PasswordNoofdigits        *string `tfsdk:"password_noofdigits"`           // Number of digits required
	GroupImportMapping        *string `tfsdk:"group_import_mapping"`          // JSON mapping for LDAP groups
	UnlockAccountJson         *string `tfsdk:"unlock_account_json"`           // JSON to unlock accounts
	StatusKeyJson             *string `tfsdk:"status_key_json"`               // JSON for account status keys
	Enablegroupmanagement     *string `tfsdk:"enable_group_management"`       // Flag to enable group management
	DisableAccountJson        *string `tfsdk:"disable_account_json"`          // JSON to disable an account
	ModifyUserdataJson        *string `tfsdk:"modify_user_data_json"`         // JSON for inline user data transformation
	OrgBase                   *string `tfsdk:"org_base"`                      // Organization BASE for provision job
	OrganizationAttribute     *string `tfsdk:"organization_attribute"`        // Organization attributes
	Orgimportjson             *string `tfsdk:"org_import_json"`               // JSON for organization import
	Createorgjson             *string `tfsdk:"create_org_json"`               // JSON for organization creation
	Updateorgjson             *string `tfsdk:"update_org_json"`               // JSON for organization update
	ConfigJson                *string `tfsdk:"config_json"`                   // JSON for connection timeout configuration
	LastImportTime            *string `tfsdk:"last_import_time"`              // Last import timestamp
	PamConfig                 *string `tfsdk:"pam_config"`                    // JSON for PAM bootstrap configuration
	// Result                    string  `tfsdk:"result"`
	// Msg                       *string `tfsdk:"msg"`
	// ErrorCode                 *string `tfsdk:"error_code"`
}

// testConnectionResource implements the resource.Resource interface.
type testConnectionResource struct {
	// client *openapi.APIClient
	client *s.Client
	token  string
}

// NewTestConnectionResource returns a new instance of testConnectionResource.
func NewTestConnectionResource() resource.Resource {
	return &testConnectionResource{}
}

func (r *testConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_test_connection"
}

func (r *testConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
			"ssl_certificate": schema.StringAttribute{
				Optional:    true,
				Description: "SSL certificates to secure the connection. Example: \"-----BEGIN CERTIFICATE----- ... -----END CERTIFICATE-----\"",
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
			"enable_group_management": schema.StringAttribute{
				Optional:    true,
				Description: "Flag to enable group management. Example: \"TRUE\"",
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
			"org_import_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for organization import.",
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
			"last_import_time": schema.StringAttribute{
				Optional:    true,
				Description: "Last import timestamp.",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Description: "JSON for PAM bootstrap configuration. Example: '{\"Connection\":\"AD\",...}'",
			},
			// "result": schema.StringAttribute{
			// 	Optional:    true,
			// 	Computed:    true,
			// 	Description: "The result of the API call.",
			// },
			// "msg": schema.StringAttribute{
			// 	Optional:    true,
			// 	Computed:    true,
			// 	Description: "A message indicating the outcome of the operation.",
			// },
			// "error_code": schema.StringAttribute{
			// 	Optional:    true,
			// 	Computed:    true,
			// 	Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
			// },
		},
	}
}

func (r *testConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
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

func (r *testConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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
	connectionName := *plan.ConnectionName
	adConn := openapi.ADConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype: "AD",
			ConnectionName: util.StringPtr(connectionName),
		},
		URL:                     plan.URL,
		USERNAME:                plan.Username,
		PASSWORD:                plan.Password,
		LDAP_OR_AD:              plan.LdapOrAd,
		PAGE_SIZE:               plan.PageSize,
		BASE:                    plan.Base,
		STATUS_THRESHOLD_CONFIG: plan.StatusThresholdConfig,
		ENTITLEMENT_ATTRIBUTE:   plan.EntitlementAttribute,
		GroupSearchBaseDN:       plan.GroupSearchBaseDN,
	}
	testConnRequest := openapi.TestConnectionRequest{
		ADConnector: &adConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	reqAPI := apiClient.ConnectionsAPI.TestConnection(ctx).TestConnectionRequest(testConnRequest)
	_, httpResp, err := reqAPI.Execute()
	if err != nil {
		// Handle 404: resource no longer exists, remove from state
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("Failed to create connection: %s", err.Error()))
		return
	}
	// Assign ID and result to the plan
	plan.ID = types.StringValue("test-connection-" + connectionName)
	// msgValue := util.SafeDeref(apiResponse.Msg)
	// errorCodeValue := util.SafeDeref(apiResponse.ErrorCode)

	// Set the individual fields
	// plan.Msg = util.StringPtr(msgValue)
	// plan.ErrorCode = util.StringPtr(errorCodeValue)
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
	// 	return
	// }
	// plan.Result = types.StringValue(string(resultJSON))

	// Store state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *testConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// If the API does not support a separate read operation, you can pass through the state.
}

func (r *testConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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
	connectionName := *plan.ConnectionName
	adConn := openapi.ADConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype: "AD",
			ConnectionName: util.StringPtr(connectionName),
		},
		URL:                     plan.URL,
		USERNAME:                plan.Username,
		PASSWORD:                plan.Password,
		LDAP_OR_AD:              plan.LdapOrAd,
		PAGE_SIZE:               plan.PageSize,
		BASE:                    plan.Base,
		STATUS_THRESHOLD_CONFIG: plan.StatusThresholdConfig,
		ENTITLEMENT_ATTRIBUTE:   plan.EntitlementAttribute,
		GroupSearchBaseDN:       plan.GroupSearchBaseDN,
	}
	testConnRequest := openapi.TestConnectionRequest{
		ADConnector: &adConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	reqAPI := apiClient.ConnectionsAPI.TestConnection(ctx).TestConnectionRequest(testConnRequest)
	_, httpResp, err := reqAPI.Execute()
	if err != nil {
		// Handle 404: resource no longer exists, remove from state
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("Failed to create connection: %s", err.Error()))
		return
	}
	// Assign ID and result to the plan
	// Assign ID and result to the plan
	plan.ID = types.StringValue("test-connection-" + connectionName)
	// msgValue := util.SafeDeref(apiResponse.Msg)
	// errorCodeValue := util.SafeDeref(apiResponse.ErrorCode)

	// // Set the individual fields
	// plan.Msg = util.StringPtr(msgValue)
	// plan.ErrorCode = util.StringPtr(errorCodeValue)
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
	// 	return
	// }
	// plan.Result = types.StringValue(string(resultJSON))

	// Store state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *testConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
