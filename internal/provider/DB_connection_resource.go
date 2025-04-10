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

type dbConnectionResource struct {
	client *s.Client
	token  string
}

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
				Computed:    true,
				Description: "Properties that need to be added when connecting to the database",
			},
			"password_min_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the minimum length for the random password",
			},
			"password_max_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum length for the random password",
			},
			"password_no_of_caps_alpha": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the number of uppercase alphabets required for the random password",
			},
			"password_no_of_digits": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the number of digits required for the random password",
			},
			"password_no_of_spl_chars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the number of special characters required for the random password",
			},
			"create_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify the queries/stored procedures used to create a new account (e.g., randomPassword, task, user, accountName, role, endpoint, etc.)",
			},
			"update_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify the queries/stored procedures used to update an existing account",
			},
			"grant_access_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify the queries/stored procedures used to provide access",
			},
			"revoke_access_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify the queries/stored procedures used to revoke access",
			},
			"change_pass_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON to specify the queries/stored procedures used to change a password",
			},
			"delete_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify the queries/stored procedures used to delete an account",
			},
			"enable_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify the queries/stored procedures used to enable an account",
			},
			"disable_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify the queries/stored procedures used to disable an account",
			},
			"account_exists_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify the query used to check whether an account exists",
			},
			"update_user_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify the queries/stored procedures used to update user information",
			},
			"accounts_import": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Accounts Import XML file content",
			},
			"entitlement_value_import": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Entitlement Value Import XML file content",
			},
			"role_owner_import": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Role Owner Import XML file content",
			},
			"roles_import": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Roles Import XML file content",
			},
			"system_import": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "System Import XML file content",
			},
			"user_import": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User Import XML file content",
			},
			"modify_user_data_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Property for MODIFYUSERDATAJSON",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configuration for status and threshold (e.g., statusColumn, activeStatus, accountThresholdValue, etc.)",
			},
			"max_pagination_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Defines the maximum number of records to be processed per page",
			},
			"cli_command_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON to specify commands executable on the target server",
			},
			"msg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Message returned from the operation.",
			},
			"error_code": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Error code if the operation fails.",
			},
		},
	}
}

func (r *dbConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *dbConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DBConnectorResourceModel
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

	dbConn := openapi.DBConnector{
		BaseConnector: openapi.BaseConnector{
			//required field
			Connectiontype: "DB",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional field
			Description:     util.StringPointerOrEmpty(plan.Description.ValueString()),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate.ValueString()),
		},
		//required field
		URL:        plan.URL.ValueString(),
		USERNAME:   plan.Username.ValueString(),
		PASSWORD:   plan.Password.ValueString(),
		DRIVERNAME: plan.DriverName.ValueString(),
		//optional field
		CONNECTIONPROPERTIES:    util.StringPointerOrEmpty(plan.ConnectionProperties.ValueString()),
		PASSWORD_MIN_LENGTH:     util.StringPointerOrEmpty(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:     util.StringPointerOrEmpty(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:  util.StringPointerOrEmpty(plan.PasswordNoOfCapsAlpha.ValueString()),
		PASSWORD_NOOFDIGITS:     util.StringPointerOrEmpty(plan.PasswordNoOfDigits.ValueString()),
		PASSWORD_NOOFSPLCHARS:   util.StringPointerOrEmpty(plan.PasswordNoOfSplChars.ValueString()),
		CREATEACCOUNTJSON:       util.StringPointerOrEmpty(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:       util.StringPointerOrEmpty(plan.UpdateAccountJson.ValueString()),
		GRANTACCESSJSON:         util.StringPointerOrEmpty(plan.GrantAccessJson.ValueString()),
		REVOKEACCESSJSON:        util.StringPointerOrEmpty(plan.RevokeAccessJson.ValueString()),
		CHANGEPASSJSON:          util.StringPointerOrEmpty(plan.ChangePassJson.ValueString()),
		DELETEACCOUNTJSON:       util.StringPointerOrEmpty(plan.DeleteAccountJson.ValueString()),
		ENABLEACCOUNTJSON:       util.StringPointerOrEmpty(plan.EnableAccountJson.ValueString()),
		DISABLEACCOUNTJSON:      util.StringPointerOrEmpty(plan.DisableAccountJson.ValueString()),
		ACCOUNTEXISTSJSON:       util.StringPointerOrEmpty(plan.AccountExistsJson.ValueString()),
		UPDATEUSERJSON:          util.StringPointerOrEmpty(plan.UpdateUserJson.ValueString()),
		ACCOUNTSIMPORT:          util.StringPointerOrEmpty(plan.AccountsImport.ValueString()),
		ENTITLEMENTVALUEIMPORT:  util.StringPointerOrEmpty(plan.EntitlementValueImport.ValueString()),
		ROLEOWNERIMPORT:         util.StringPointerOrEmpty(plan.RoleOwnerImport.ValueString()),
		ROLESIMPORT:             util.StringPointerOrEmpty(plan.RolesImport.ValueString()),
		SYSTEMIMPORT:            util.StringPointerOrEmpty(plan.SystemImport.ValueString()),
		USERIMPORT:              util.StringPointerOrEmpty(plan.UserImport.ValueString()),
		MODIFYUSERDATAJSON:      util.StringPointerOrEmpty(plan.ModifyUserDataJson.ValueString()),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()),
		MAX_PAGINATION_SIZE:     util.StringPointerOrEmpty(plan.MaxPaginationSize.ValueString()),
		CLI_COMMAND_JSON:        util.StringPointerOrEmpty(plan.CliCommandJson.ValueString()),
	}
	if plan.VaultConnection.ValueString() != "" {
		dbConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		dbConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		dbConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}
	dbConnRequest := openapi.CreateOrUpdateRequest{
		DBConnector: &dbConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(dbConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = types.StringValue(*util.StringPointerOrEmpty(plan.Description.ValueString()))
	plan.DefaultSavRoles = types.StringValue(*util.StringPointerOrEmpty(plan.DefaultSavRoles.ValueString()))
	plan.EmailTemplate = types.StringValue(*util.StringPointerOrEmpty(plan.EmailTemplate.ValueString()))
	plan.ConnectionProperties = types.StringValue(*util.StringPointerOrEmpty(plan.ConnectionProperties.ValueString()))
	plan.PasswordMinLength = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordMinLength.ValueString()))
	plan.PasswordMaxLength = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordMaxLength.ValueString()))
	plan.PasswordNoOfCapsAlpha = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordNoOfCapsAlpha.ValueString()))
	plan.PasswordNoOfDigits = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordNoOfDigits.ValueString()))
	plan.PasswordNoOfSplChars = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordNoOfSplChars.ValueString()))
	plan.CreateAccountJson = types.StringValue(*util.StringPointerOrEmpty(plan.CreateAccountJson.ValueString()))
	plan.UpdateAccountJson = types.StringValue(*util.StringPointerOrEmpty(plan.UpdateAccountJson.ValueString()))
	plan.GrantAccessJson = types.StringValue(*util.StringPointerOrEmpty(plan.GrantAccessJson.ValueString()))
	plan.RevokeAccessJson = types.StringValue(*util.StringPointerOrEmpty(plan.RevokeAccessJson.ValueString()))
	plan.DeleteAccountJson = types.StringValue(*util.StringPointerOrEmpty(plan.DeleteAccountJson.ValueString()))
	plan.EnableAccountJson = types.StringValue(*util.StringPointerOrEmpty(plan.EnableAccountJson.ValueString()))
	plan.DisableAccountJson = types.StringValue(*util.StringPointerOrEmpty(plan.DisableAccountJson.ValueString()))
	plan.AccountExistsJson = types.StringValue(*util.StringPointerOrEmpty(plan.AccountExistsJson.ValueString()))
	plan.UpdateUserJson = types.StringValue(*util.StringPointerOrEmpty(plan.UpdateUserJson.ValueString()))
	plan.AccountsImport = types.StringValue(*util.StringPointerOrEmpty(plan.AccountsImport.ValueString()))
	plan.EntitlementValueImport = types.StringValue(*util.StringPointerOrEmpty(plan.EntitlementValueImport.ValueString()))
	plan.RoleOwnerImport = types.StringValue(*util.StringPointerOrEmpty(plan.RoleOwnerImport.ValueString()))
	plan.RolesImport = types.StringValue(*util.StringPointerOrEmpty(plan.RolesImport.ValueString()))
	plan.SystemImport = types.StringValue(*util.StringPointerOrEmpty(plan.SystemImport.ValueString()))
	plan.UserImport = types.StringValue(*util.StringPointerOrEmpty(plan.UserImport.ValueString()))
	plan.ModifyUserDataJson = types.StringValue(*util.StringPointerOrEmpty(plan.ModifyUserDataJson.ValueString()))
	plan.StatusThresholdConfig = types.StringValue(*util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()))
	plan.MaxPaginationSize = types.StringValue(*util.StringPointerOrEmpty(plan.MaxPaginationSize.ValueString()))
	plan.CliCommandJson = types.StringValue(*util.StringPointerOrEmpty(plan.CliCommandJson.ValueString()))
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	r.Read(ctx, resource.ReadRequest{State: resp.State}, &resource.ReadResponse{State: resp.State})
}

func (r *dbConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DBConnectorResourceModel

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
	state.ConnectionKey = types.Int64Value(int64(*apiResp.DBConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.DBConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.DBConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.DBConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.DBConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectiontype)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.DBConnectionResponse.Emailtemplate)
	state.PasswordMinLength = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH)
	state.AccountExistsJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ACCOUNTEXISTSJSON)
	state.RolesImport = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ROLESIMPORT)
	state.RoleOwnerImport = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ROLEOWNERIMPORT)
	state.CreateAccountJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	state.UserImport = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.USERIMPORT)
	state.DisableAccountJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON)
	state.EntitlementValueImport = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ENTITLEMENTVALUEIMPORT)
	state.UpdateUserJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.UPDATEUSERJSON)
	state.PasswordNoOfSplChars = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS)
	state.RevokeAccessJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.REVOKEACCESSJSON)
	state.URL = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.URL)
	state.SystemImport = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.SYSTEMIMPORT)
	state.DriverName = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.DRIVERNAME)
	state.DeleteAccountJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.DELETEACCOUNTJSON)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.Username = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.USERNAME)
	state.PasswordNoOfCapsAlpha = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA)
	state.PasswordNoOfDigits = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS)
	state.ConnectionProperties = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.CONNECTIONPROPERTIES)
	state.ModifyUserDataJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.AccountsImport = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ACCOUNTSIMPORT)
	state.EnableAccountJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	state.PasswordMaxLength = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH)
	state.MaxPaginationSize = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.MAX_PAGINATION_SIZE)
	state.UpdateAccountJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	state.GrantAccessJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.GRANTACCESSJSON)
	state.CliCommandJson = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.CLI_COMMAND_JSON)
	apiMessage := util.SafeDeref(apiResp.DBConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.DBConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *dbConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DBConnectorResourceModel
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

	dbConn := openapi.DBConnector{
		BaseConnector: openapi.BaseConnector{
			//required field
			Connectiontype: "DB",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional field
			Description:        util.StringPointerOrEmpty(plan.Description.ValueString()),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		//required field
		URL:        plan.URL.ValueString(),
		USERNAME:   plan.Username.ValueString(),
		PASSWORD:   plan.Password.ValueString(),
		DRIVERNAME: plan.DriverName.ValueString(),
		//optional field
		CONNECTIONPROPERTIES:    util.StringPointerOrEmpty(plan.ConnectionProperties.ValueString()),
		PASSWORD_MIN_LENGTH:     util.StringPointerOrEmpty(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:     util.StringPointerOrEmpty(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:  util.StringPointerOrEmpty(plan.PasswordNoOfCapsAlpha.ValueString()),
		PASSWORD_NOOFDIGITS:     util.StringPointerOrEmpty(plan.PasswordNoOfDigits.ValueString()),
		PASSWORD_NOOFSPLCHARS:   util.StringPointerOrEmpty(plan.PasswordNoOfSplChars.ValueString()),
		CREATEACCOUNTJSON:       util.StringPointerOrEmpty(plan.CreateAccountJson.ValueString()),
		UPDATEACCOUNTJSON:       util.StringPointerOrEmpty(plan.UpdateAccountJson.ValueString()),
		GRANTACCESSJSON:         util.StringPointerOrEmpty(plan.GrantAccessJson.ValueString()),
		REVOKEACCESSJSON:        util.StringPointerOrEmpty(plan.RevokeAccessJson.ValueString()),
		CHANGEPASSJSON:          util.StringPointerOrEmpty(plan.ChangePassJson.ValueString()),
		DELETEACCOUNTJSON:       util.StringPointerOrEmpty(plan.DeleteAccountJson.ValueString()),
		ENABLEACCOUNTJSON:       util.StringPointerOrEmpty(plan.EnableAccountJson.ValueString()),
		DISABLEACCOUNTJSON:      util.StringPointerOrEmpty(plan.DisableAccountJson.ValueString()),
		ACCOUNTEXISTSJSON:       util.StringPointerOrEmpty(plan.AccountExistsJson.ValueString()),
		UPDATEUSERJSON:          util.StringPointerOrEmpty(plan.UpdateUserJson.ValueString()),
		ACCOUNTSIMPORT:          util.StringPointerOrEmpty(plan.AccountsImport.ValueString()),
		ENTITLEMENTVALUEIMPORT:  util.StringPointerOrEmpty(plan.EntitlementValueImport.ValueString()),
		ROLEOWNERIMPORT:         util.StringPointerOrEmpty(plan.RoleOwnerImport.ValueString()),
		ROLESIMPORT:             util.StringPointerOrEmpty(plan.RolesImport.ValueString()),
		SYSTEMIMPORT:            util.StringPointerOrEmpty(plan.SystemImport.ValueString()),
		USERIMPORT:              util.StringPointerOrEmpty(plan.UserImport.ValueString()),
		MODIFYUSERDATAJSON:      util.StringPointerOrEmpty(plan.ModifyUserDataJson.ValueString()),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()),
		MAX_PAGINATION_SIZE:     util.StringPointerOrEmpty(plan.MaxPaginationSize.ValueString()),
		CLI_COMMAND_JSON:        util.StringPointerOrEmpty(plan.CliCommandJson.ValueString()),
	}
	if plan.VaultConnection.ValueString() != "" {
		dbConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		dbConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		dbConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	} else {
		emptyStr := ""
		dbConn.BaseConnector.VaultConnection = &emptyStr
		dbConn.BaseConnector.VaultConfiguration = &emptyStr
		dbConn.BaseConnector.Saveinvault = &emptyStr
	}
	dbConnRequest := openapi.CreateOrUpdateRequest{
		DBConnector: &dbConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(dbConnRequest).Execute()
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
	plan.ConnectionKey = types.Int64Value(int64(*getResp.DBConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.DBConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionname)
	plan.ConnectionKey = util.SafeInt64(getResp.DBConnectionResponse.Connectionkey)
	plan.Description = util.SafeStringDatasource(getResp.DBConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.DBConnectionResponse.Defaultsavroles)
	plan.ConnectionType = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectiontype)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.DBConnectionResponse.Emailtemplate)
	plan.PasswordMinLength = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH)
	plan.AccountExistsJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.ACCOUNTEXISTSJSON)
	plan.RolesImport = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.ROLESIMPORT)
	plan.RoleOwnerImport = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.ROLEOWNERIMPORT)
	plan.CreateAccountJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	plan.UserImport = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.USERIMPORT)
	plan.DisableAccountJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON)
	plan.EntitlementValueImport = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.ENTITLEMENTVALUEIMPORT)
	plan.UpdateUserJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.UPDATEUSERJSON)
	plan.PasswordNoOfSplChars = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS)
	plan.RevokeAccessJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.REVOKEACCESSJSON)
	plan.URL = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.URL)
	plan.SystemImport = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.SYSTEMIMPORT)
	plan.DriverName = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.DRIVERNAME)
	plan.DeleteAccountJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.DELETEACCOUNTJSON)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.Username = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.USERNAME)
	plan.PasswordNoOfCapsAlpha = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA)
	plan.PasswordNoOfDigits = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS)
	plan.ConnectionProperties = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.CONNECTIONPROPERTIES)
	plan.ModifyUserDataJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	plan.AccountsImport = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.ACCOUNTSIMPORT)
	plan.EnableAccountJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	plan.PasswordMaxLength = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH)
	plan.MaxPaginationSize = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.MAX_PAGINATION_SIZE)
	plan.UpdateAccountJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	plan.GrantAccessJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.GRANTACCESSJSON)
	plan.CliCommandJson = util.SafeStringDatasource(getResp.DBConnectionResponse.Connectionattributes.CLI_COMMAND_JSON)
	apiMessage := util.SafeDeref(getResp.DBConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.DBConnectionResponse.Errorcode)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *dbConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
