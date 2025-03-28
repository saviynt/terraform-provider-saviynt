// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"terraform-provider-Saviynt/util"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	s "github.com/saviynt/saviynt-api-go-client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DBConnectorResourceModel struct {
	BaseConnector
	ID                     types.String `tfsdk:"id"`
	URL                    types.String `tfsdk:"url"`
	Username               types.String `tfsdk:"username"`
	Password               types.String `tfsdk:"password"`
	DriverName             types.String `tfsdk:"driver_name"`
	ConnectionProperties   types.String `tfsdk:"connection_properties"`
	PasswordMinLength      types.String `tfsdk:"password_min_length"`
	PasswordMaxLength      types.String `tfsdk:"password_max_length"`
	PasswordNoOfCapsAlpha  types.String `tfsdk:"password_no_of_caps_alpha"`
	PasswordNoOfDigits     types.String `tfsdk:"password_no_of_digits"`
	PasswordNoOfSplChars   types.String `tfsdk:"password_no_of_spl_chars"`
	CreateAccountJson      types.String `tfsdk:"create_account_json"`
	UpdateAccountJson      types.String `tfsdk:"update_account_json"`
	GrantAccessJson        types.String `tfsdk:"grant_access_json"`
	RevokeAccessJson       types.String `tfsdk:"revoke_access_json"`
	ChangePassJson         types.String `tfsdk:"change_pass_json"`
	DeleteAccountJson      types.String `tfsdk:"delete_account_json"`
	EnableAccountJson      types.String `tfsdk:"enable_account_json"`
	DisableAccountJson     types.String `tfsdk:"disable_account_json"`
	AccountExistsJson      types.String `tfsdk:"account_exists_json"`
	UpdateUserJson         types.String `tfsdk:"update_user_json"`
	AccountsImport         types.String `tfsdk:"accounts_import"`
	EntitlementValueImport types.String `tfsdk:"entitlement_value_import"`
	RoleOwnerImport        types.String `tfsdk:"role_owner_import"`
	RolesImport            types.String `tfsdk:"roles_import"`
	SystemImport           types.String `tfsdk:"system_import"`
	UserImport             types.String `tfsdk:"user_import"`
	ModifyUserDataJson     types.String `tfsdk:"modify_user_data_json"`
	StatusThresholdConfig  types.String `tfsdk:"status_threshold_config"`
	MaxPaginationSize      types.String `tfsdk:"max_pagination_size"`
	CliCommandJson         types.String `tfsdk:"cli_command_json"`
}

// testConnectionResource implements the resource.Resource interface.
type dbConnectionResource struct {
	// client *openapi.APIClient
	client *s.Client
	token  string
}

// NewTestConnectionResource returns a new instance of testConnectionResource.
func DBNewTestConnectionResource() resource.Resource {
	return &dbConnectionResource{}
}

func (r *dbConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_db_connection_resource"
}

func (r *dbConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Description: "Host Name for connection",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Username for connection",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Description: "Password for connection",
			},
			"driver_name": schema.StringAttribute{
				Required:    true,
				Description: "Driver name for the connection",
			},
			"connection_properties": schema.StringAttribute{
				Optional:    true,
				Description: "Properties that need to be added when connecting to the database",
			},
			"password_min_length": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the minimum length for the random password",
			},
			"password_max_length": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the maximum length for the random password",
			},
			"password_no_of_caps_alpha": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the number of uppercase alphabets required for the random password",
			},
			"password_no_of_digits": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the number of digits required for the random password",
			},
			"password_no_of_spl_chars": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the number of special characters required for the random password",
			},
			"create_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to create a new account (e.g., randomPassword, task, user, accountName, role, endpoint, etc.)",
			},
			"update_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to update an existing account",
			},
			"grant_access_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to provide access",
			},
			"revoke_access_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to revoke access",
			},
			"change_pass_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to change a password",
			},
			"delete_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to delete an account",
			},
			"enable_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to enable an account",
			},
			"disable_account_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to disable an account",
			},
			"account_exists_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the query used to check whether an account exists",
			},
			"update_user_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to update user information",
			},
			"accounts_import": schema.StringAttribute{
				Optional:    true,
				Description: "Accounts Import XML file content",
			},
			"entitlement_value_import": schema.StringAttribute{
				Optional:    true,
				Description: "Entitlement Value Import XML file content",
			},
			"role_owner_import": schema.StringAttribute{
				Optional:    true,
				Description: "Role Owner Import XML file content",
			},
			"roles_import": schema.StringAttribute{
				Optional:    true,
				Description: "Roles Import XML file content",
			},
			"system_import": schema.StringAttribute{
				Optional:    true,
				Description: "System Import XML file content",
			},
			"user_import": schema.StringAttribute{
				Optional:    true,
				Description: "User Import XML file content",
			},
			"modify_user_data_json": schema.StringAttribute{
				Optional:    true,
				Description: "Property for MODIFYUSERDATAJSON",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Description: "Configuration for status and threshold (e.g., statusColumn, activeStatus, accountThresholdValue, etc.)",
			},
			"max_pagination_size": schema.StringAttribute{
				Optional:    true,
				Description: "Defines the maximum number of records to be processed per page",
			},
			"cli_command_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify commands executable on the target server",
			},
			"result": schema.StringAttribute{
				Computed:    true,
				Description: "Result of the operation.",
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

func (r *dbConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *dbConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DBConnectorResourceModel

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

	dbConn := openapi.DBConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "DB",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.SafeStringConnector(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnector(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnector(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		URL:                     plan.URL.ValueString(),
		USERNAME:                plan.Username.ValueString(),
		PASSWORD:                plan.Password.ValueString(),
		DRIVERNAME:              plan.DriverName.ValueString(),
		CONNECTIONPROPERTIES:    util.SafeStringConnector(plan.ConnectionProperties.ValueString()),
		PASSWORD_MIN_LENGTH:     util.SafeStringConnector(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:     util.SafeStringConnector(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:  util.SafeStringConnector(plan.PasswordNoOfCapsAlpha.ValueString()),
		PASSWORD_NOOFDIGITS:     util.SafeStringConnector(plan.PasswordNoOfDigits.ValueString()),
		PASSWORD_NOOFSPLCHARS:   util.SafeStringConnector(plan.PasswordNoOfSplChars.ValueString()),
		CREATEACCOUNTJSON:       util.SafeStringConnector(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:       util.SafeStringConnector(plan.UpdateAccountJson.ValueString()),
		GRANTACCESSJSON:         util.SafeStringConnector(plan.GrantAccessJson.ValueString()),
		REVOKEACCESSJSON:        util.SafeStringConnector(plan.RevokeAccessJson.ValueString()),
		CHANGEPASSJSON:          util.SafeStringConnector(plan.ChangePassJson.ValueString()),
		DELETEACCOUNTJSON:       util.SafeStringConnector(plan.DeleteAccountJson.ValueString()),
		ENABLEACCOUNTJSON:       util.SafeStringConnector(plan.EnableAccountJson.ValueString()),
		DISABLEACCOUNTJSON:      util.SafeStringConnector(plan.DisableAccountJson.ValueString()),
		ACCOUNTEXISTSJSON:       util.SafeStringConnector(plan.AccountExistsJson.ValueString()),
		UPDATEUSERJSON:          util.SafeStringConnector(plan.UpdateUserJson.ValueString()),
		ACCOUNTSIMPORT:          util.SafeStringConnector(plan.AccountsImport.ValueString()),
		ENTITLEMENTVALUEIMPORT:  util.SafeStringConnector(plan.EntitlementValueImport.ValueString()),
		ROLEOWNERIMPORT:         util.SafeStringConnector(plan.RoleOwnerImport.ValueString()),
		ROLESIMPORT:             util.SafeStringConnector(plan.RolesImport.ValueString()),
		SYSTEMIMPORT:            util.SafeStringConnector(plan.SystemImport.ValueString()),
		USERIMPORT:              util.SafeStringConnector(plan.UserImport.ValueString()),
		MODIFYUSERDATAJSON:      util.SafeStringConnector(plan.ModifyUserDataJson.ValueString()),
		STATUS_THRESHOLD_CONFIG: util.SafeStringConnector(plan.StatusThresholdConfig.ValueString()),
		MAX_PAGINATION_SIZE:     util.SafeStringConnector(plan.MaxPaginationSize.ValueString()),
		CLI_COMMAND_JSON:        util.SafeStringConnector(plan.CliCommandJson.ValueString()),
	}
	dbConnRequest := openapi.CreateOrUpdateRequest{
		DBConnector: &dbConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, httpResp, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(dbConnRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating AD Connector",
			fmt.Sprintf("Error: %v\nHTTP Response: %v", err, httpResp),
		)
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
		return
	}
	plan.Result = types.StringValue(string(resultJSON))
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *dbConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// If the API does not support a separate read operation, you can pass through the state.
}

func (r *dbConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DBConnectorResourceModel

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

	dbConn := openapi.DBConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:     "DB",
			ConnectionName:     plan.ConnectionName.ValueString(),
			Description:        util.SafeStringConnector(plan.Description.ValueString()),
			Defaultsavroles:    util.SafeStringConnector(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.SafeStringConnector(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		URL:                     plan.URL.ValueString(),
		USERNAME:                plan.Username.ValueString(),
		PASSWORD:                plan.Password.ValueString(),
		DRIVERNAME:              plan.DriverName.ValueString(),
		CONNECTIONPROPERTIES:    util.SafeStringConnector(plan.ConnectionProperties.ValueString()),
		PASSWORD_MIN_LENGTH:     util.SafeStringConnector(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:     util.SafeStringConnector(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:  util.SafeStringConnector(plan.PasswordNoOfCapsAlpha.ValueString()),
		PASSWORD_NOOFDIGITS:     util.SafeStringConnector(plan.PasswordNoOfDigits.ValueString()),
		PASSWORD_NOOFSPLCHARS:   util.SafeStringConnector(plan.PasswordNoOfSplChars.ValueString()),
		CREATEACCOUNTJSON:       util.SafeStringConnector(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:       util.SafeStringConnector(plan.UpdateAccountJson.ValueString()),
		GRANTACCESSJSON:         util.SafeStringConnector(plan.GrantAccessJson.ValueString()),
		REVOKEACCESSJSON:        util.SafeStringConnector(plan.RevokeAccessJson.ValueString()),
		CHANGEPASSJSON:          util.SafeStringConnector(plan.ChangePassJson.ValueString()),
		DELETEACCOUNTJSON:       util.SafeStringConnector(plan.DeleteAccountJson.ValueString()),
		ENABLEACCOUNTJSON:       util.SafeStringConnector(plan.EnableAccountJson.ValueString()),
		DISABLEACCOUNTJSON:      util.SafeStringConnector(plan.DisableAccountJson.ValueString()),
		ACCOUNTEXISTSJSON:       util.SafeStringConnector(plan.AccountExistsJson.ValueString()),
		UPDATEUSERJSON:          util.SafeStringConnector(plan.UpdateUserJson.ValueString()),
		ACCOUNTSIMPORT:          util.SafeStringConnector(plan.AccountsImport.ValueString()),
		ENTITLEMENTVALUEIMPORT:  util.SafeStringConnector(plan.EntitlementValueImport.ValueString()),
		ROLEOWNERIMPORT:         util.SafeStringConnector(plan.RoleOwnerImport.ValueString()),
		ROLESIMPORT:             util.SafeStringConnector(plan.RolesImport.ValueString()),
		SYSTEMIMPORT:            util.SafeStringConnector(plan.SystemImport.ValueString()),
		USERIMPORT:              util.SafeStringConnector(plan.UserImport.ValueString()),
		MODIFYUSERDATAJSON:      util.SafeStringConnector(plan.ModifyUserDataJson.ValueString()),
		STATUS_THRESHOLD_CONFIG: util.SafeStringConnector(plan.StatusThresholdConfig.ValueString()),
		MAX_PAGINATION_SIZE:     util.SafeStringConnector(plan.MaxPaginationSize.ValueString()),
		CLI_COMMAND_JSON:        util.SafeStringConnector(plan.CliCommandJson.ValueString()),
	}
	dbConnRequest := openapi.CreateOrUpdateRequest{
		DBConnector: &dbConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, httpResp, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(dbConnRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating AD Connector",
			fmt.Sprintf("Error: %v\nHTTP Response: %v", err, httpResp),
		)
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
		return
	}
	plan.Result = types.StringValue(string(resultJSON))
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *dbConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
