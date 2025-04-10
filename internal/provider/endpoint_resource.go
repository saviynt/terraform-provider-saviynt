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

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
	openapi "github.com/saviynt/saviynt-api-go-client/endpoints"
)

type endpointResourceModel struct {
	ID                                      types.String          `tfsdk:"id"`
	EndpointName                            types.String          `tfsdk:"endpointname"`
	DisplayName                             types.String          `tfsdk:"display_name"`
	SecuritySystem                          types.String          `tfsdk:"security_system"`
	Description                             types.String          `tfsdk:"description"`
	OwnerType                               types.String          `tfsdk:"owner_type"`
	Owner                                   types.String          `tfsdk:"owner"`
	ResourceOwnerType                       types.String          `tfsdk:"resource_owner_type"`
	ResourceOwner                           types.String          `tfsdk:"resource_owner"`
	AccessQuery                             types.String          `tfsdk:"access_query"`
	EnableCopyAccess                        types.String          `tfsdk:"enable_copy_access"`
	DisableNewAccountRequestIfAccountExists types.String          `tfsdk:"disable_new_account_request_if_account_exists"`
	DisableRemoveAccount                    types.String          `tfsdk:"disable_remove_account"`
	DisableModifyAccount                    types.String          `tfsdk:"disable_modify_account"`
	UserAccountCorrelationRule              types.String          `tfsdk:"user_account_correlation_rule"`
	CreateEntTaskforRemoveAcc               types.String          `tfsdk:"create_ent_task_for_remove_acc"`
	OutOfBandAction                         types.String          `tfsdk:"out_of_band_action"`
	ConnectionConfig                        types.String          `tfsdk:"connection_config"`
	Requestable                             types.String          `tfsdk:"requestable"`
	ParentAccountPattern                    types.String          `tfsdk:"parent_account_pattern"`
	ServiceAccountNameRule                  types.String          `tfsdk:"service_account_name_rule"`
	ServiceAccountAccessQuery               types.String          `tfsdk:"service_account_access_query"`
	BlockInflightRequest                    types.String          `tfsdk:"block_inflight_request"`
	AccountNameRule                         types.String          `tfsdk:"account_name_rule"`
	AllowChangePasswordSQLQuery             types.String          `tfsdk:"allow_change_password_sql_query"`
	AccountNameValidatorRegex               types.String          `tfsdk:"account_name_validator_regex"`
	StatusConfig                            types.String          `tfsdk:"status_config"`
	PluginConfigs                           types.String          `tfsdk:"plugin_configs"`
	EndpointConfig                          types.String          `tfsdk:"endpoint_config"`
	CustomProperty1                         types.String          `tfsdk:"custom_property1"`
	CustomProperty2                         types.String          `tfsdk:"custom_property2"`
	CustomProperty3                         types.String          `tfsdk:"custom_property3"`
	CustomProperty4                         types.String          `tfsdk:"custom_property4"`
	CustomProperty5                         types.String          `tfsdk:"custom_property5"`
	CustomProperty6                         types.String          `tfsdk:"custom_property6"`
	CustomProperty7                         types.String          `tfsdk:"custom_property7"`
	CustomProperty8                         types.String          `tfsdk:"custom_property8"`
	CustomProperty9                         types.String          `tfsdk:"custom_property9"`
	CustomProperty10                        types.String          `tfsdk:"custom_property10"`
	CustomProperty11                        types.String          `tfsdk:"custom_property11"`
	CustomProperty12                        types.String          `tfsdk:"custom_property12"`
	CustomProperty13                        types.String          `tfsdk:"custom_property13"`
	CustomProperty14                        types.String          `tfsdk:"custom_property14"`
	CustomProperty15                        types.String          `tfsdk:"custom_property15"`
	CustomProperty16                        types.String          `tfsdk:"custom_property16"`
	CustomProperty17                        types.String          `tfsdk:"custom_property17"`
	CustomProperty18                        types.String          `tfsdk:"custom_property18"`
	CustomProperty19                        types.String          `tfsdk:"custom_property19"`
	CustomProperty20                        types.String          `tfsdk:"custom_property20"`
	CustomProperty21                        types.String          `tfsdk:"custom_property21"`
	CustomProperty22                        types.String          `tfsdk:"custom_property22"`
	CustomProperty23                        types.String          `tfsdk:"custom_property23"`
	CustomProperty24                        types.String          `tfsdk:"custom_property24"`
	CustomProperty25                        types.String          `tfsdk:"custom_property25"`
	CustomProperty26                        types.String          `tfsdk:"custom_property26"`
	CustomProperty27                        types.String          `tfsdk:"custom_property27"`
	CustomProperty28                        types.String          `tfsdk:"custom_property28"`
	CustomProperty29                        types.String          `tfsdk:"custom_property29"`
	CustomProperty30                        types.String          `tfsdk:"custom_property30"`
	CustomProperty31                        types.String          `tfsdk:"custom_property31"`
	CustomProperty32                        types.String          `tfsdk:"custom_property32"`
	CustomProperty33                        types.String          `tfsdk:"custom_property33"`
	CustomProperty34                        types.String          `tfsdk:"custom_property34"`
	CustomProperty35                        types.String          `tfsdk:"custom_property35"`
	CustomProperty36                        types.String          `tfsdk:"custom_property36"`
	CustomProperty37                        types.String          `tfsdk:"custom_property37"`
	CustomProperty38                        types.String          `tfsdk:"custom_property38"`
	CustomProperty39                        types.String          `tfsdk:"custom_property39"`
	CustomProperty40                        types.String          `tfsdk:"custom_property40"`
	CustomProperty41                        types.String          `tfsdk:"custom_property41"`
	CustomProperty42                        types.String          `tfsdk:"custom_property42"`
	CustomProperty43                        types.String          `tfsdk:"custom_property43"`
	CustomProperty44                        types.String          `tfsdk:"custom_property44"`
	CustomProperty45                        types.String          `tfsdk:"custom_property45"`
	CustomProperty1Label                    types.String          `tfsdk:"custom_property1_label"`
	CustomProperty2Label                    types.String          `tfsdk:"custom_property2_label"`
	CustomProperty3Label                    types.String          `tfsdk:"custom_property3_label"`
	CustomProperty4Label                    types.String          `tfsdk:"custom_property4_label"`
	CustomProperty5Label                    types.String          `tfsdk:"custom_property5_label"`
	CustomProperty6Label                    types.String          `tfsdk:"custom_property6_label"`
	CustomProperty7Label                    types.String          `tfsdk:"custom_property7_label"`
	CustomProperty8Label                    types.String          `tfsdk:"custom_property8_label"`
	CustomProperty9Label                    types.String          `tfsdk:"custom_property9_label"`
	CustomProperty10Label                   types.String          `tfsdk:"custom_property10_label"`
	CustomProperty11Label                   types.String          `tfsdk:"custom_property11_label"`
	CustomProperty12Label                   types.String          `tfsdk:"custom_property12_label"`
	CustomProperty13Label                   types.String          `tfsdk:"custom_property13_label"`
	CustomProperty14Label                   types.String          `tfsdk:"custom_property14_label"`
	CustomProperty15Label                   types.String          `tfsdk:"custom_property15_label"`
	CustomProperty16Label                   types.String          `tfsdk:"custom_property16_label"`
	CustomProperty17Label                   types.String          `tfsdk:"custom_property17_label"`
	CustomProperty18Label                   types.String          `tfsdk:"custom_property18_label"`
	CustomProperty19Label                   types.String          `tfsdk:"custom_property19_label"`
	CustomProperty20Label                   types.String          `tfsdk:"custom_property20_label"`
	CustomProperty21Label                   types.String          `tfsdk:"custom_property21_label"`
	CustomProperty22Label                   types.String          `tfsdk:"custom_property22_label"`
	CustomProperty23Label                   types.String          `tfsdk:"custom_property23_label"`
	CustomProperty24Label                   types.String          `tfsdk:"custom_property24_label"`
	CustomProperty25Label                   types.String          `tfsdk:"custom_property25_label"`
	CustomProperty26Label                   types.String          `tfsdk:"custom_property26_label"`
	CustomProperty27Label                   types.String          `tfsdk:"custom_property27_label"`
	CustomProperty28Label                   types.String          `tfsdk:"custom_property28_label"`
	CustomProperty29Label                   types.String          `tfsdk:"custom_property29_label"`
	CustomProperty30Label                   types.String          `tfsdk:"custom_property30_label"`
	CustomProperty31Label                   types.String          `tfsdk:"custom_property31_label"`
	CustomProperty32Label                   types.String          `tfsdk:"custom_property32_label"`
	CustomProperty33Label                   types.String          `tfsdk:"custom_property33_label"`
	CustomProperty34Label                   types.String          `tfsdk:"custom_property34_label"`
	CustomProperty35Label                   types.String          `tfsdk:"custom_property35_label"`
	CustomProperty36Label                   types.String          `tfsdk:"custom_property36_label"`
	CustomProperty37Label                   types.String          `tfsdk:"custom_property37_label"`
	CustomProperty38Label                   types.String          `tfsdk:"custom_property38_label"`
	CustomProperty39Label                   types.String          `tfsdk:"custom_property39_label"`
	CustomProperty40Label                   types.String          `tfsdk:"custom_property40_label"`
	CustomProperty41Label                   types.String          `tfsdk:"custom_property41_label"`
	CustomProperty42Label                   types.String          `tfsdk:"custom_property42_label"`
	CustomProperty43Label                   types.String          `tfsdk:"custom_property43_label"`
	CustomProperty44Label                   types.String          `tfsdk:"custom_property44_label"`
	CustomProperty45Label                   types.String          `tfsdk:"custom_property45_label"`
	CustomProperty46Label                   types.String          `tfsdk:"custom_property46_label"`
	CustomProperty47Label                   types.String          `tfsdk:"custom_property47_label"`
	CustomProperty48Label                   types.String          `tfsdk:"custom_property48_label"`
	CustomProperty49Label                   types.String          `tfsdk:"custom_property49_label"`
	CustomProperty50Label                   types.String          `tfsdk:"custom_property50_label"`
	CustomProperty51Label                   types.String          `tfsdk:"custom_property51_label"`
	CustomProperty52Label                   types.String          `tfsdk:"custom_property52_label"`
	CustomProperty53Label                   types.String          `tfsdk:"custom_property53_label"`
	CustomProperty54Label                   types.String          `tfsdk:"custom_property54_label"`
	CustomProperty55Label                   types.String          `tfsdk:"custom_property55_label"`
	CustomProperty56Label                   types.String          `tfsdk:"custom_property56_label"`
	CustomProperty57Label                   types.String          `tfsdk:"custom_property57_label"`
	CustomProperty58Label                   types.String          `tfsdk:"custom_property58_label"`
	CustomProperty59Label                   types.String          `tfsdk:"custom_property59_label"`
	CustomProperty60Label                   types.String          `tfsdk:"custom_property60_label"`
	AllowRemoveAllRoleOnRequest             types.String          `tfsdk:"allow_remove_all_role_on_request"`
	ChangePasswordAccessQuery               types.String          `tfsdk:"change_password_access_query"`
	RequestableRoleType                     []RequestableRoleType `tfsdk:"requestable_role_type"`
	EmailTemplate                           []EmailTemplate       `tfsdk:"email_template"`
	MappedEndpoints                         []MappedEndpoint      `tfsdk:"mapped_endpoints"`

	Result    types.String `tfsdk:"result"`
	Msg       types.String `tfsdk:"msg"`
	ErrorCode types.String `tfsdk:"error_code"`
}

type endpointResource struct {
	client *s.Client
	token  string
}

type RequestableRoleType struct {
	RoleType       types.String `tfsdk:"role_type"`
	RequestOption  types.String `tfsdk:"request_option"`
	Required       types.Bool   `tfsdk:"required"`
	RequestedQuery types.String `tfsdk:"requested_query"`
	SelectedQuery  types.String `tfsdk:"selected_query"`
	ShowOn         types.String `tfsdk:"show_on"`
}

type EmailTemplate struct {
	EmailTemplateType types.String `tfsdk:"email_template_type"`
	TaskType          types.String `tfsdk:"task_type"`
	EmailTemplate     types.String `tfsdk:"email_template"`
}

type MappedEndpoint struct {
	SecuritySystem types.String `tfsdk:"security_system"`
	Endpoint       types.String `tfsdk:"endpoint"`
	Requestable    types.String `tfsdk:"requestable"`
	Operation      types.String `tfsdk:"operation"`
}

func NewEndpointResource() resource.Resource {
	return &endpointResource{}
}

func (r *endpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_endpoint_resource"
}

func (r *endpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and Manage endpoints",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the resource.",
			},
			"endpointname": schema.StringAttribute{
				Required:    true,
				Description: "Specify a name for the endpoint. Provide a logical name that will help you easily identify it.",
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "Enter a user-friendly display name for the endpoint that will be displayed in the user interface. Display Name can be different from Endpoint Name.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"security_system": schema.StringAttribute{
				Required:    true,
				Description: "Specify the Security system for which you want to create an endpoint.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "Specify a description for the endpoint.",
			},
			"owner_type": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the owner type for the endpoint. An endpoint can be owned by a User or Usergroup.",
			},
			"owner": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the owner of the endpoint. If the ownerType is User, then specify the username of the owner, and If it is is Usergroup then specify the name of the user group.",
			},
			"resource_owner_type": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the resource owner type for the endpoint. An endpoint can be owned by a User or Usergroup.",
			},
			"resource_owner": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the resource owner of the endpoint. If the resourceOwnerType is User, then specify the username of the owner and If it is Usergroup, specify the name of the user group.",
			},
			"access_query": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users.",
			},
			"enable_copy_access": schema.StringAttribute{
				Optional:    true,
				Description: "Specify true to display the Copy Access from User option in the Request pages.",
			},
			"disable_new_account_request_if_account_exists": schema.StringAttribute{
				Optional:    true,
				Description: "Specify true to disable users from requesting additional accounts on applications where they already have active accounts.",
			},
			"disable_remove_account": schema.StringAttribute{
				Optional:    true,
				Description: "Specify true to disable users from removing their existing application accounts.",
			},
			"disable_modify_account": schema.StringAttribute{
				Optional:    true,
				Description: "Specify true to disable users from modifying their application accounts.",
			},
			"user_account_correlation_rule": schema.StringAttribute{
				Optional:    true,
				Description: "Specify rule to map users in EIC with the accounts during account import.",
			},
			"create_ent_task_for_remove_acc": schema.StringAttribute{
				Optional:    true,
				Description: "If this is set to true, remove Access tasks will be created for entitlements (account entitlements and their dependent entitlements) when a user requests for removing an account.",
			},
			"out_of_band_action": schema.StringAttribute{
				Optional:    true,
				Description: "Use this parameter to determine if you need to remove the accesses which were granted outside Saviynt.",
			},
			"connection_config": schema.StringAttribute{
				Optional:    true,
				Description: "Use this configuration for processing the add access tasks and remove access tasks for AD and LDAP Connectors.",
			},
			"requestable": schema.StringAttribute{
				Optional:    true,
				Description: "Is this endpoint requestable.",
			},
			"parent_account_pattern": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the parent and child relationship for the Active Directory endpoint. The specified value is used to filter the parent and child objects in the Request Access tile.",
			},
			"service_account_name_rule": schema.StringAttribute{
				Optional:    true,
				Description: "Rule to generate a name for this endpoint while creating a new service account.",
			},
			"service_account_access_query": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the query to filter the access and display of the endpoint for specific users while managing service accounts.",
			},
			"block_inflight_request": schema.StringAttribute{
				Optional:    true,
				Description: "Specify true to prevent users from raising duplicate requests for the same applications.",
			},
			"account_name_rule": schema.StringAttribute{
				Optional:    true,
				Description: "Specify rule to generate an account name for this endpoint while creating a new account.",
			},
			"allow_change_password_sql_query": schema.StringAttribute{
				Optional:    true,
				Description: "SQL query to configure the accounts for which you can change passwords.",
			},
			"account_name_validator_regex": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the regular expression which will be used to validate the account name either generated by the rule or provided manually.",
			},
			"status_config": schema.StringAttribute{
				Optional:    true,
				Description: "Enable the State and Status options (Enable, Disable, Lock, Unlock) that would be available to request for a user and service accounts.",
			},
			"plugin_configs": schema.StringAttribute{
				Optional:    true,
				Description: "The Plugin Configuration drives the functionality of the Saviynt SmartAssist (Browserplugin).",
			},
			"endpoint_config": schema.StringAttribute{
				Optional:    true,
				Description: "Option to copy data in Step 3 of the service account request will be enabled.",
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

	for i := 1; i <= 45; i++ {
		key := fmt.Sprintf("custom_property%d", i)
		resp.Schema.Attributes[key] = schema.StringAttribute{
			Optional:    true,
			Description: fmt.Sprintf("Custom Property %d.", i),
		}
	}

	for i := 1; i <= 60; i++ {
		key := fmt.Sprintf("custom_property%d_label", i)
		resp.Schema.Attributes[key] = schema.StringAttribute{
			Optional:    true,
			Description: fmt.Sprintf("Label for the custom property %d of accounts of this endpoint.", i),
		}
	}

	resp.Schema.Attributes["allow_remove_all_role_on_request"] = schema.StringAttribute{
		Optional:    true,
		Description: "Specify true to displays the Remove All Roles option in the Request page that can be used to remove all the roles.",
	}

	resp.Schema.Attributes["change_password_access_query"] = schema.StringAttribute{
		Optional:    true,
		Description: "Specify query to restrict the access for changing the account password of the endpoint.",
	}

	resp.Schema.Attributes["requestable_role_type"] = schema.ListNestedAttribute{
		Description: "A list of requestable role types associated with the endpoint.",
		Optional:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"role_type": schema.StringAttribute{
					Description: "Type of role that can be requested.",
					Optional:    true,
				},
				"request_option": schema.StringAttribute{
					Description: "Option for requesting the role.",
					Optional:    true,
				},
				"required": schema.BoolAttribute{
					Description: "Indicates whether the role is required.",
					Optional:    true,
				},
				"requested_query": schema.StringAttribute{
					Description: "Query for requested role selection.",
					Optional:    true,
				},
				"selected_query": schema.StringAttribute{
					Description: "Query for selected role display.",
					Optional:    true,
				},
				"show_on": schema.StringAttribute{
					Description: "Specifies where the role should be shown.",
					Optional:    true,
				},
			},
		},
	}
	resp.Schema.Attributes["email_template"] = schema.ListNestedAttribute{
		Description: "A list of email templates associated with the endpoint.",
		Optional:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"email_template_type": schema.StringAttribute{
					Description: "Type of email template (e.g., Approval, Rejection).",
					Optional:    true,
				},
				"task_type": schema.StringAttribute{
					Description: "Task type associated with the email template (e.g., Create, Delete).",
					Optional:    true,
				},
				"email_template": schema.StringAttribute{
					Description: "The email template name to be used.",
					Optional:    true,
				},
			},
		},
	}

	resp.Schema.Attributes["mapped_endpoints"] = schema.ListNestedAttribute{
		Description: "List of mapped endpoints with individual security systems.",
		Optional:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"security_system": schema.StringAttribute{
					Description: "The security system specific to this mapped endpoint.",
					Required:    true,
				},
				"endpoint": schema.StringAttribute{
					Description: "Logical name of the endpoint.",
					Required:    true,
				},
				"requestable": schema.StringAttribute{
					Description: "Indicates whether the endpoint is requestable.",
					Optional:    true,
				},
				"operation": schema.StringAttribute{
					Description: "Specifies the operation associated with the endpoint.",
					Optional:    true,
				},
			},
		},
	}
}

func (r *endpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *endpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan endpointResourceModel

	planGetDiagnostics := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(planGetDiagnostics...)
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

	createReq := openapi.NewCreateEndpointRequest(
		plan.EndpointName.ValueString(),
		plan.DisplayName.ValueString(),
		plan.SecuritySystem.ValueString(),
	)

	if !plan.Description.IsNull() && plan.Description.ValueString() != "" {
		createReq.SetDescription(plan.Description.ValueString())
	}
	if !plan.OwnerType.IsNull() && plan.OwnerType.ValueString() != "" {
		createReq.SetOwnerType(plan.OwnerType.ValueString())
	}
	if !plan.Owner.IsNull() && plan.Owner.ValueString() != "" {
		createReq.SetOwner(plan.Owner.ValueString())
	}
	if !plan.ResourceOwnerType.IsNull() && plan.ResourceOwnerType.ValueString() != "" {
		createReq.SetResourceOwnerType(plan.ResourceOwnerType.ValueString())
	}
	if !plan.ResourceOwner.IsNull() && plan.ResourceOwner.ValueString() != "" {
		createReq.SetResourceOwner(plan.ResourceOwner.ValueString())
	}
	if !plan.AccessQuery.IsNull() && plan.AccessQuery.ValueString() != "" {
		createReq.SetAccessquery(plan.AccessQuery.ValueString())
	}
	if !plan.EnableCopyAccess.IsNull() && plan.EnableCopyAccess.ValueString() != "" {
		createReq.SetEnableCopyAccess(plan.EnableCopyAccess.ValueString())
	}
	if !plan.DisableNewAccountRequestIfAccountExists.IsNull() && plan.DisableNewAccountRequestIfAccountExists.ValueString() != "" {
		createReq.SetDisableNewAccountRequestIfAccountExists(plan.DisableNewAccountRequestIfAccountExists.ValueString())
	}
	if !plan.DisableRemoveAccount.IsNull() && plan.DisableRemoveAccount.ValueString() != "" {
		createReq.SetDisableRemoveAccount(plan.DisableRemoveAccount.ValueString())
	}
	if !plan.DisableModifyAccount.IsNull() && plan.DisableModifyAccount.ValueString() != "" {
		createReq.SetDisableModifyAccount(plan.DisableModifyAccount.ValueString())
	}
	if !plan.UserAccountCorrelationRule.IsNull() && plan.UserAccountCorrelationRule.ValueString() != "" {
		createReq.SetUserAccountCorrelationRule(plan.UserAccountCorrelationRule.ValueString())
	}
	if !plan.CreateEntTaskforRemoveAcc.IsNull() && plan.CreateEntTaskforRemoveAcc.ValueString() != "" {
		createReq.SetCreateEntTaskforRemoveAcc(plan.CreateEntTaskforRemoveAcc.ValueString())
	}
	if !plan.OutOfBandAction.IsNull() && plan.OutOfBandAction.ValueString() != "" {
		createReq.SetOutofbandaction(plan.OutOfBandAction.ValueString())
	}
	if !plan.ConnectionConfig.IsNull() && plan.ConnectionConfig.ValueString() != "" {
		createReq.SetConnectionconfig(plan.ConnectionConfig.ValueString())
	}
	if !plan.Requestable.IsNull() && plan.Requestable.ValueString() != "" {
		createReq.SetRequestable(plan.Requestable.ValueString())
	}
	if !plan.ParentAccountPattern.IsNull() && plan.ParentAccountPattern.ValueString() != "" {
		createReq.SetParentAccountPattern(plan.ParentAccountPattern.ValueString())
	}
	if !plan.ServiceAccountNameRule.IsNull() && plan.ServiceAccountNameRule.ValueString() != "" {
		createReq.SetServiceAccountNameRule(plan.ServiceAccountNameRule.ValueString())
	}
	if !plan.ServiceAccountAccessQuery.IsNull() && plan.ServiceAccountAccessQuery.ValueString() != "" {
		createReq.SetServiceAccountAccessQuery(plan.ServiceAccountAccessQuery.ValueString())
	}
	if !plan.BlockInflightRequest.IsNull() && plan.BlockInflightRequest.ValueString() != "" {
		createReq.SetBlockInflightRequest(plan.BlockInflightRequest.ValueString())
	}
	if !plan.AccountNameRule.IsNull() && plan.AccountNameRule.ValueString() != "" {
		createReq.SetAccountNameRule(plan.AccountNameRule.ValueString())
	}
	if !plan.AllowChangePasswordSQLQuery.IsNull() && plan.AllowChangePasswordSQLQuery.ValueString() != "" {
		createReq.SetAllowChangePasswordSqlquery(plan.AllowChangePasswordSQLQuery.ValueString())
	}
	if !plan.AccountNameValidatorRegex.IsNull() && plan.AccountNameValidatorRegex.ValueString() != "" {
		createReq.SetAccountNameValidatorRegex(plan.AccountNameValidatorRegex.ValueString())
	}
	if !plan.StatusConfig.IsNull() && plan.StatusConfig.ValueString() != "" {
		createReq.SetStatusConfig(plan.StatusConfig.ValueString())
	}
	if !plan.PluginConfigs.IsNull() && plan.PluginConfigs.ValueString() != "" {
		createReq.SetPluginConfigs(plan.PluginConfigs.ValueString())
	}
	if !plan.EndpointConfig.IsNull() && plan.EndpointConfig.ValueString() != "" {
		createReq.SetEndpointConfig(plan.EndpointConfig.ValueString())
	}
	if !plan.CustomProperty1.IsNull() && plan.CustomProperty1.ValueString() != "" {
		createReq.SetCustomproperty1(plan.CustomProperty1.ValueString())
	}
	if !plan.CustomProperty2.IsNull() && plan.CustomProperty2.ValueString() != "" {
		createReq.SetCustomproperty2(plan.CustomProperty2.ValueString())
	}
	if !plan.CustomProperty3.IsNull() && plan.CustomProperty3.ValueString() != "" {
		createReq.SetCustomproperty3(plan.CustomProperty3.ValueString())
	}
	if !plan.CustomProperty4.IsNull() && plan.CustomProperty4.ValueString() != "" {
		createReq.SetCustomproperty4(plan.CustomProperty4.ValueString())
	}
	if !plan.CustomProperty5.IsNull() && plan.CustomProperty5.ValueString() != "" {
		createReq.SetCustomproperty5(plan.CustomProperty5.ValueString())
	}
	if !plan.CustomProperty6.IsNull() && plan.CustomProperty6.ValueString() != "" {
		createReq.SetCustomproperty6(plan.CustomProperty6.ValueString())
	}
	if !plan.CustomProperty7.IsNull() && plan.CustomProperty7.ValueString() != "" {
		createReq.SetCustomproperty7(plan.CustomProperty7.ValueString())
	}
	if !plan.CustomProperty8.IsNull() && plan.CustomProperty8.ValueString() != "" {
		createReq.SetCustomproperty8(plan.CustomProperty8.ValueString())
	}
	if !plan.CustomProperty9.IsNull() && plan.CustomProperty9.ValueString() != "" {
		createReq.SetCustomproperty9(plan.CustomProperty9.ValueString())
	}
	if !plan.CustomProperty10.IsNull() && plan.CustomProperty10.ValueString() != "" {
		createReq.SetCustomproperty10(plan.CustomProperty10.ValueString())
	}
	if !plan.CustomProperty11.IsNull() && plan.CustomProperty11.ValueString() != "" {
		createReq.SetCustomproperty11(plan.CustomProperty11.ValueString())
	}
	if !plan.CustomProperty12.IsNull() && plan.CustomProperty12.ValueString() != "" {
		createReq.SetCustomproperty12(plan.CustomProperty12.ValueString())
	}
	if !plan.CustomProperty13.IsNull() && plan.CustomProperty13.ValueString() != "" {
		createReq.SetCustomproperty13(plan.CustomProperty13.ValueString())
	}
	if !plan.CustomProperty14.IsNull() && plan.CustomProperty14.ValueString() != "" {
		createReq.SetCustomproperty14(plan.CustomProperty14.ValueString())
	}
	if !plan.CustomProperty15.IsNull() && plan.CustomProperty15.ValueString() != "" {
		createReq.SetCustomproperty15(plan.CustomProperty15.ValueString())
	}
	if !plan.CustomProperty16.IsNull() && plan.CustomProperty16.ValueString() != "" {
		createReq.SetCustomproperty16(plan.CustomProperty16.ValueString())
	}
	if !plan.CustomProperty17.IsNull() && plan.CustomProperty17.ValueString() != "" {
		createReq.SetCustomproperty17(plan.CustomProperty17.ValueString())
	}
	if !plan.CustomProperty18.IsNull() && plan.CustomProperty18.ValueString() != "" {
		createReq.SetCustomproperty18(plan.CustomProperty18.ValueString())
	}
	if !plan.CustomProperty19.IsNull() && plan.CustomProperty19.ValueString() != "" {
		createReq.SetCustomproperty19(plan.CustomProperty19.ValueString())
	}
	if !plan.CustomProperty20.IsNull() && plan.CustomProperty20.ValueString() != "" {
		createReq.SetCustomproperty20(plan.CustomProperty20.ValueString())
	}
	if !plan.CustomProperty21.IsNull() && plan.CustomProperty21.ValueString() != "" {
		createReq.SetCustomproperty21(plan.CustomProperty21.ValueString())
	}
	if !plan.CustomProperty22.IsNull() && plan.CustomProperty22.ValueString() != "" {
		createReq.SetCustomproperty22(plan.CustomProperty22.ValueString())
	}
	if !plan.CustomProperty23.IsNull() && plan.CustomProperty23.ValueString() != "" {
		createReq.SetCustomproperty23(plan.CustomProperty23.ValueString())
	}
	if !plan.CustomProperty24.IsNull() && plan.CustomProperty24.ValueString() != "" {
		createReq.SetCustomproperty24(plan.CustomProperty24.ValueString())
	}
	if !plan.CustomProperty25.IsNull() && plan.CustomProperty25.ValueString() != "" {
		createReq.SetCustomproperty25(plan.CustomProperty25.ValueString())
	}
	if !plan.CustomProperty26.IsNull() && plan.CustomProperty26.ValueString() != "" {
		createReq.SetCustomproperty26(plan.CustomProperty26.ValueString())
	}
	if !plan.CustomProperty27.IsNull() && plan.CustomProperty27.ValueString() != "" {
		createReq.SetCustomproperty27(plan.CustomProperty27.ValueString())
	}
	if !plan.CustomProperty28.IsNull() && plan.CustomProperty28.ValueString() != "" {
		createReq.SetCustomproperty28(plan.CustomProperty28.ValueString())
	}
	if !plan.CustomProperty29.IsNull() && plan.CustomProperty29.ValueString() != "" {
		createReq.SetCustomproperty29(plan.CustomProperty29.ValueString())
	}
	if !plan.CustomProperty30.IsNull() && plan.CustomProperty30.ValueString() != "" {
		createReq.SetCustomproperty30(plan.CustomProperty30.ValueString())
	}
	if !plan.CustomProperty31.IsNull() && plan.CustomProperty31.ValueString() != "" {
		createReq.SetCustomproperty31(plan.CustomProperty31.ValueString())
	}
	if !plan.CustomProperty32.IsNull() && plan.CustomProperty32.ValueString() != "" {
		createReq.SetCustomproperty32(plan.CustomProperty32.ValueString())
	}
	if !plan.CustomProperty33.IsNull() && plan.CustomProperty33.ValueString() != "" {
		createReq.SetCustomproperty33(plan.CustomProperty33.ValueString())
	}
	if !plan.CustomProperty34.IsNull() && plan.CustomProperty34.ValueString() != "" {
		createReq.SetCustomproperty34(plan.CustomProperty34.ValueString())
	}
	if !plan.CustomProperty35.IsNull() && plan.CustomProperty35.ValueString() != "" {
		createReq.SetCustomproperty35(plan.CustomProperty35.ValueString())
	}
	if !plan.CustomProperty36.IsNull() && plan.CustomProperty36.ValueString() != "" {
		createReq.SetCustomproperty36(plan.CustomProperty36.ValueString())
	}
	if !plan.CustomProperty37.IsNull() && plan.CustomProperty37.ValueString() != "" {
		createReq.SetCustomproperty37(plan.CustomProperty37.ValueString())
	}
	if !plan.CustomProperty38.IsNull() && plan.CustomProperty38.ValueString() != "" {
		createReq.SetCustomproperty38(plan.CustomProperty38.ValueString())
	}
	if !plan.CustomProperty39.IsNull() && plan.CustomProperty39.ValueString() != "" {
		createReq.SetCustomproperty39(plan.CustomProperty39.ValueString())
	}
	if !plan.CustomProperty40.IsNull() && plan.CustomProperty40.ValueString() != "" {
		createReq.SetCustomproperty40(plan.CustomProperty40.ValueString())
	}
	if !plan.CustomProperty41.IsNull() && plan.CustomProperty41.ValueString() != "" {
		createReq.SetCustomproperty41(plan.CustomProperty41.ValueString())
	}
	if !plan.CustomProperty42.IsNull() && plan.CustomProperty42.ValueString() != "" {
		createReq.SetCustomproperty42(plan.CustomProperty42.ValueString())
	}
	if !plan.CustomProperty43.IsNull() && plan.CustomProperty43.ValueString() != "" {
		createReq.SetCustomproperty43(plan.CustomProperty43.ValueString())
	}
	if !plan.CustomProperty44.IsNull() && plan.CustomProperty44.ValueString() != "" {
		createReq.SetCustomproperty44(plan.CustomProperty44.ValueString())
	}
	if !plan.CustomProperty45.IsNull() && plan.CustomProperty45.ValueString() != "" {
		createReq.SetCustomproperty45(plan.CustomProperty45.ValueString())
	}
	if !plan.CustomProperty1Label.IsNull() && plan.CustomProperty1Label.ValueString() != "" {
		createReq.SetCustomproperty1Label(plan.CustomProperty1Label.ValueString())
	}
	if !plan.CustomProperty2Label.IsNull() && plan.CustomProperty2Label.ValueString() != "" {
		createReq.SetCustomproperty2Label(plan.CustomProperty2Label.ValueString())
	}
	if !plan.CustomProperty3Label.IsNull() && plan.CustomProperty3Label.ValueString() != "" {
		createReq.SetCustomproperty3Label(plan.CustomProperty3Label.ValueString())
	}
	if !plan.CustomProperty4Label.IsNull() && plan.CustomProperty4Label.ValueString() != "" {
		createReq.SetCustomproperty4Label(plan.CustomProperty4Label.ValueString())
	}
	if !plan.CustomProperty5Label.IsNull() && plan.CustomProperty5Label.ValueString() != "" {
		createReq.SetCustomproperty5Label(plan.CustomProperty5Label.ValueString())
	}
	if !plan.CustomProperty6Label.IsNull() && plan.CustomProperty6Label.ValueString() != "" {
		createReq.SetCustomproperty6Label(plan.CustomProperty6Label.ValueString())
	}
	if !plan.CustomProperty7Label.IsNull() && plan.CustomProperty7Label.ValueString() != "" {
		createReq.SetCustomproperty7Label(plan.CustomProperty7Label.ValueString())
	}
	if !plan.CustomProperty8Label.IsNull() && plan.CustomProperty8Label.ValueString() != "" {
		createReq.SetCustomproperty8Label(plan.CustomProperty8Label.ValueString())
	}
	if !plan.CustomProperty9Label.IsNull() && plan.CustomProperty9Label.ValueString() != "" {
		createReq.SetCustomproperty9Label(plan.CustomProperty9Label.ValueString())
	}
	if !plan.CustomProperty10Label.IsNull() && plan.CustomProperty10Label.ValueString() != "" {
		createReq.SetCustomproperty10Label(plan.CustomProperty10Label.ValueString())
	}
	if !plan.CustomProperty11Label.IsNull() && plan.CustomProperty11Label.ValueString() != "" {
		createReq.SetCustomproperty11Label(plan.CustomProperty11Label.ValueString())
	}
	if !plan.CustomProperty12Label.IsNull() && plan.CustomProperty12Label.ValueString() != "" {
		createReq.SetCustomproperty12Label(plan.CustomProperty12Label.ValueString())
	}
	if !plan.CustomProperty13Label.IsNull() && plan.CustomProperty13Label.ValueString() != "" {
		createReq.SetCustomproperty13Label(plan.CustomProperty13Label.ValueString())
	}
	if !plan.CustomProperty14Label.IsNull() && plan.CustomProperty14Label.ValueString() != "" {
		createReq.SetCustomproperty14Label(plan.CustomProperty14Label.ValueString())
	}
	if !plan.CustomProperty15Label.IsNull() && plan.CustomProperty15Label.ValueString() != "" {
		createReq.SetCustomproperty15Label(plan.CustomProperty15Label.ValueString())
	}
	if !plan.CustomProperty16Label.IsNull() && plan.CustomProperty16Label.ValueString() != "" {
		createReq.SetCustomproperty16Label(plan.CustomProperty16Label.ValueString())
	}
	if !plan.CustomProperty17Label.IsNull() && plan.CustomProperty17Label.ValueString() != "" {
		createReq.SetCustomproperty17Label(plan.CustomProperty17Label.ValueString())
	}
	if !plan.CustomProperty18Label.IsNull() && plan.CustomProperty18Label.ValueString() != "" {
		createReq.SetCustomproperty18Label(plan.CustomProperty18Label.ValueString())
	}
	if !plan.CustomProperty19Label.IsNull() && plan.CustomProperty19Label.ValueString() != "" {
		createReq.SetCustomproperty19Label(plan.CustomProperty19Label.ValueString())
	}
	if !plan.CustomProperty20Label.IsNull() && plan.CustomProperty20Label.ValueString() != "" {
		createReq.SetCustomproperty20Label(plan.CustomProperty20Label.ValueString())
	}
	if !plan.CustomProperty21Label.IsNull() && plan.CustomProperty21Label.ValueString() != "" {
		createReq.SetCustomproperty21Label(plan.CustomProperty21Label.ValueString())
	}
	if !plan.CustomProperty22Label.IsNull() && plan.CustomProperty22Label.ValueString() != "" {
		createReq.SetCustomproperty22Label(plan.CustomProperty22Label.ValueString())
	}
	if !plan.CustomProperty23Label.IsNull() && plan.CustomProperty23Label.ValueString() != "" {
		createReq.SetCustomproperty23Label(plan.CustomProperty23Label.ValueString())
	}
	if !plan.CustomProperty24Label.IsNull() && plan.CustomProperty24Label.ValueString() != "" {
		createReq.SetCustomproperty24Label(plan.CustomProperty24Label.ValueString())
	}
	if !plan.CustomProperty25Label.IsNull() && plan.CustomProperty25Label.ValueString() != "" {
		createReq.SetCustomproperty25Label(plan.CustomProperty25Label.ValueString())
	}
	if !plan.CustomProperty26Label.IsNull() && plan.CustomProperty26Label.ValueString() != "" {
		createReq.SetCustomproperty26Label(plan.CustomProperty26Label.ValueString())
	}
	if !plan.CustomProperty27Label.IsNull() && plan.CustomProperty27Label.ValueString() != "" {
		createReq.SetCustomproperty27Label(plan.CustomProperty27Label.ValueString())
	}
	if !plan.CustomProperty28Label.IsNull() && plan.CustomProperty28Label.ValueString() != "" {
		createReq.SetCustomproperty28Label(plan.CustomProperty28Label.ValueString())
	}
	if !plan.CustomProperty29Label.IsNull() && plan.CustomProperty29Label.ValueString() != "" {
		createReq.SetCustomproperty29Label(plan.CustomProperty29Label.ValueString())
	}
	if !plan.CustomProperty30Label.IsNull() && plan.CustomProperty30Label.ValueString() != "" {
		createReq.SetCustomproperty30Label(plan.CustomProperty30Label.ValueString())
	}
	if !plan.CustomProperty31Label.IsNull() && plan.CustomProperty31Label.ValueString() != "" {
		createReq.SetCustomproperty31Label(plan.CustomProperty31Label.ValueString())
	}
	if !plan.CustomProperty32Label.IsNull() && plan.CustomProperty32Label.ValueString() != "" {
		createReq.SetCustomproperty32Label(plan.CustomProperty32Label.ValueString())
	}
	if !plan.CustomProperty33Label.IsNull() && plan.CustomProperty33Label.ValueString() != "" {
		createReq.SetCustomproperty33Label(plan.CustomProperty33Label.ValueString())
	}
	if !plan.CustomProperty34Label.IsNull() && plan.CustomProperty34Label.ValueString() != "" {
		createReq.SetCustomproperty34Label(plan.CustomProperty34Label.ValueString())
	}
	if !plan.CustomProperty35Label.IsNull() && plan.CustomProperty35Label.ValueString() != "" {
		createReq.SetCustomproperty35Label(plan.CustomProperty35Label.ValueString())
	}
	if !plan.CustomProperty36Label.IsNull() && plan.CustomProperty36Label.ValueString() != "" {
		createReq.SetCustomproperty36Label(plan.CustomProperty36Label.ValueString())
	}
	if !plan.CustomProperty37Label.IsNull() && plan.CustomProperty37Label.ValueString() != "" {
		createReq.SetCustomproperty37Label(plan.CustomProperty37Label.ValueString())
	}
	if !plan.CustomProperty38Label.IsNull() && plan.CustomProperty38Label.ValueString() != "" {
		createReq.SetCustomproperty38Label(plan.CustomProperty38Label.ValueString())
	}
	if !plan.CustomProperty39Label.IsNull() && plan.CustomProperty39Label.ValueString() != "" {
		createReq.SetCustomproperty39Label(plan.CustomProperty39Label.ValueString())
	}
	if !plan.CustomProperty40Label.IsNull() && plan.CustomProperty40Label.ValueString() != "" {
		createReq.SetCustomproperty40Label(plan.CustomProperty40Label.ValueString())
	}
	if !plan.CustomProperty41Label.IsNull() && plan.CustomProperty41Label.ValueString() != "" {
		createReq.SetCustomproperty41Label(plan.CustomProperty41Label.ValueString())
	}
	if !plan.CustomProperty42Label.IsNull() && plan.CustomProperty42Label.ValueString() != "" {
		createReq.SetCustomproperty42Label(plan.CustomProperty42Label.ValueString())
	}
	if !plan.CustomProperty43Label.IsNull() && plan.CustomProperty43Label.ValueString() != "" {
		createReq.SetCustomproperty43Label(plan.CustomProperty43Label.ValueString())
	}
	if !plan.CustomProperty44Label.IsNull() && plan.CustomProperty44Label.ValueString() != "" {
		createReq.SetCustomproperty44Label(plan.CustomProperty44Label.ValueString())
	}
	if !plan.CustomProperty45Label.IsNull() && plan.CustomProperty45Label.ValueString() != "" {
		createReq.SetCustomproperty45Label(plan.CustomProperty45Label.ValueString())
	}
	if !plan.CustomProperty46Label.IsNull() && plan.CustomProperty46Label.ValueString() != "" {
		createReq.SetCustomproperty46Label(plan.CustomProperty46Label.ValueString())
	}
	if !plan.CustomProperty47Label.IsNull() && plan.CustomProperty47Label.ValueString() != "" {
		createReq.SetCustomproperty47Label(plan.CustomProperty47Label.ValueString())
	}
	if !plan.CustomProperty48Label.IsNull() && plan.CustomProperty48Label.ValueString() != "" {
		createReq.SetCustomproperty48Label(plan.CustomProperty48Label.ValueString())
	}
	if !plan.CustomProperty49Label.IsNull() && plan.CustomProperty49Label.ValueString() != "" {
		createReq.SetCustomproperty49Label(plan.CustomProperty49Label.ValueString())
	}
	if !plan.CustomProperty50Label.IsNull() && plan.CustomProperty50Label.ValueString() != "" {
		createReq.SetCustomproperty50Label(plan.CustomProperty50Label.ValueString())
	}
	if !plan.CustomProperty60Label.IsNull() && plan.CustomProperty60Label.ValueString() != "" {
		createReq.SetCustomproperty60Label(plan.CustomProperty60Label.ValueString())
	}

	apiResp, httpResp, err := apiClient.EndpointsAPI.
		CreateEndpoint(ctx).
		CreateEndpointRequest(*createReq).
		Execute()
	if err != nil {
		log.Printf("Error Creating Endpoint: %v, HTTP Response: %v", err, httpResp)
		resp.Diagnostics.AddError(
			"Error Creating Endpoint",
			"Check logs for details.",
		)
		return
	}

	plan.ID = types.StringValue("endpoint-" + plan.EndpointName.ValueString())
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
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("JSON Marshall failed: ", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)
	plan.Result = types.StringValue(string(resultJSON))

	stateCreateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateCreateDiagnostics...)
}

func (r *endpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state endpointResourceModel

	// Retrieve the current state
	stateRetrievalDiagnostics := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(stateRetrievalDiagnostics...)
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
	endpointName := state.EndpointName.ValueString()
	readReq := openapi.NewGetEndpointsRequest()
	readReq.Endpointname = &endpointName
	apiResp, httpResp, err := apiClient.EndpointsAPI.GetEndpoints(ctx).GetEndpointsRequest(*readReq).Execute()
	if err != nil {
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	var foundItem *openapi.GetEndpoints200ResponseEndpointsInner
	for _, item := range apiResp.Endpoints {
		if item.Endpointname != nil && *item.Endpointname == state.EndpointName.ValueString() {
			foundItem = &item
			break
		}
	}
	if foundItem == nil {
		resp.State.RemoveResource(ctx)
	}
	state.ID = types.StringValue("endpoint-" + state.EndpointName.ValueString())
	state.OwnerType = types.StringValue(util.SafeDeref(foundItem.OwnerType))
	state.SecuritySystem = types.StringValue(util.SafeDeref(foundItem.Securitysystem))
	state.EndpointName = types.StringValue(util.SafeDeref(foundItem.Endpointname))
	state.DisplayName = types.StringValue(util.SafeDeref(foundItem.DisplayName))
	state.AllowRemoveAllRoleOnRequest = types.StringValue(util.SafeDeref(foundItem.AllowRemoveAllRoleOnRequest))
	state.ChangePasswordAccessQuery = types.StringValue(util.SafeDeref(foundItem.ChangePasswordAccessQuery))
	state.PluginConfigs = types.StringValue(util.SafeDeref(foundItem.PluginConfigs))
	state.CreateEntTaskforRemoveAcc = types.StringValue(util.SafeDeref(foundItem.CreateEntTaskforRemoveAcc))
	state.EnableCopyAccess = types.StringValue(util.SafeDeref(foundItem.EnableCopyAccess))
	state.EndpointConfig = types.StringValue(util.SafeDeref(foundItem.EndpointConfig))
	state.ServiceAccountAccessQuery = types.StringValue(util.SafeDeref(foundItem.ServiceAccountAccessQuery))
	state.UserAccountCorrelationRule = types.StringValue(util.SafeDeref(foundItem.UserAccountCorrelationRule))
	state.StatusConfig = types.StringValue(util.SafeDeref(foundItem.StatusConfig))

	// Custom properties (Custom Property 1 to Custom Property 45)
	state.CustomProperty1 = types.StringValue(util.SafeDeref(foundItem.CustomProperty1))
	state.CustomProperty2 = types.StringValue(util.SafeDeref(foundItem.CustomProperty2))
	state.CustomProperty3 = types.StringValue(util.SafeDeref(foundItem.CustomProperty3))
	state.CustomProperty4 = types.StringValue(util.SafeDeref(foundItem.CustomProperty4))
	state.CustomProperty5 = types.StringValue(util.SafeDeref(foundItem.CustomProperty5))
	state.CustomProperty6 = types.StringValue(util.SafeDeref(foundItem.CustomProperty6))
	state.CustomProperty7 = types.StringValue(util.SafeDeref(foundItem.CustomProperty7))
	state.CustomProperty8 = types.StringValue(util.SafeDeref(foundItem.CustomProperty8))
	state.CustomProperty9 = types.StringValue(util.SafeDeref(foundItem.CustomProperty9))
	state.CustomProperty10 = types.StringValue(util.SafeDeref(foundItem.CustomProperty10))
	state.CustomProperty11 = types.StringValue(util.SafeDeref(foundItem.CustomProperty11))
	state.CustomProperty12 = types.StringValue(util.SafeDeref(foundItem.CustomProperty12))
	state.CustomProperty13 = types.StringValue(util.SafeDeref(foundItem.CustomProperty13))
	state.CustomProperty14 = types.StringValue(util.SafeDeref(foundItem.CustomProperty14))
	state.CustomProperty15 = types.StringValue(util.SafeDeref(foundItem.CustomProperty15))
	state.CustomProperty16 = types.StringValue(util.SafeDeref(foundItem.CustomProperty16))
	state.CustomProperty17 = types.StringValue(util.SafeDeref(foundItem.CustomProperty17))
	state.CustomProperty18 = types.StringValue(util.SafeDeref(foundItem.CustomProperty18))
	state.CustomProperty19 = types.StringValue(util.SafeDeref(foundItem.CustomProperty19))
	state.CustomProperty20 = types.StringValue(util.SafeDeref(foundItem.CustomProperty20))
	state.CustomProperty21 = types.StringValue(util.SafeDeref(foundItem.CustomProperty21))
	state.CustomProperty22 = types.StringValue(util.SafeDeref(foundItem.CustomProperty22))
	state.CustomProperty23 = types.StringValue(util.SafeDeref(foundItem.CustomProperty23))
	state.CustomProperty24 = types.StringValue(util.SafeDeref(foundItem.CustomProperty24))
	state.CustomProperty25 = types.StringValue(util.SafeDeref(foundItem.CustomProperty25))
	state.CustomProperty26 = types.StringValue(util.SafeDeref(foundItem.CustomProperty26))
	state.CustomProperty27 = types.StringValue(util.SafeDeref(foundItem.CustomProperty27))
	state.CustomProperty28 = types.StringValue(util.SafeDeref(foundItem.CustomProperty28))
	state.CustomProperty29 = types.StringValue(util.SafeDeref(foundItem.CustomProperty29))
	state.CustomProperty30 = types.StringValue(util.SafeDeref(foundItem.CustomProperty30))
	state.CustomProperty31 = types.StringValue(util.SafeDeref(foundItem.CustomProperty31))
	state.CustomProperty32 = types.StringValue(util.SafeDeref(foundItem.CustomProperty32))
	state.CustomProperty33 = types.StringValue(util.SafeDeref(foundItem.CustomProperty33))
	state.CustomProperty34 = types.StringValue(util.SafeDeref(foundItem.CustomProperty34))
	state.CustomProperty35 = types.StringValue(util.SafeDeref(foundItem.CustomProperty35))
	state.CustomProperty36 = types.StringValue(util.SafeDeref(foundItem.CustomProperty36))
	state.CustomProperty37 = types.StringValue(util.SafeDeref(foundItem.CustomProperty37))
	state.CustomProperty38 = types.StringValue(util.SafeDeref(foundItem.CustomProperty38))
	state.CustomProperty39 = types.StringValue(util.SafeDeref(foundItem.CustomProperty39))
	state.CustomProperty40 = types.StringValue(util.SafeDeref(foundItem.CustomProperty40))
	state.CustomProperty41 = types.StringValue(util.SafeDeref(foundItem.CustomProperty41))
	state.CustomProperty42 = types.StringValue(util.SafeDeref(foundItem.CustomProperty42))
	state.CustomProperty43 = types.StringValue(util.SafeDeref(foundItem.CustomProperty43))
	state.CustomProperty44 = types.StringValue(util.SafeDeref(foundItem.CustomProperty44))
	state.CustomProperty45 = types.StringValue(util.SafeDeref(foundItem.CustomProperty45))
	msgValue := util.SafeDeref(apiResp.Message)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)
	state.Msg = types.StringValue(msgValue)
	state.ErrorCode = types.StringValue(errorCodeValue)
	resultObj := map[string]string{
		"msg":        msgValue,
		"error_code": errorCodeValue,
	}
	resultJSON, err := util.MarshalDeterministic(resultObj)
	if err != nil {
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)
	state.Result = types.StringValue(string(resultJSON))

	stateSetDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateSetDiagnostics...)
}

func (r *endpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan endpointResourceModel

	planGetDiagnostics := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(planGetDiagnostics...)
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

	updateReq := openapi.NewUpdateEndpointRequest(plan.EndpointName.ValueString())

	if !plan.DisplayName.IsNull() && plan.DisplayName.ValueString() != "" {
		updateReq.SetDisplayName(plan.DisplayName.ValueString())
	}
	if !plan.Description.IsNull() && plan.Description.ValueString() != "" {
		updateReq.SetDescription(plan.Description.ValueString())
	}
	if !plan.OwnerType.IsNull() && plan.OwnerType.ValueString() != "" {
		updateReq.SetOwnerType(plan.OwnerType.ValueString())
	}
	if !plan.Owner.IsNull() && plan.Owner.ValueString() != "" {
		updateReq.SetOwner(plan.Owner.ValueString())
	}
	if !plan.ResourceOwnerType.IsNull() && plan.ResourceOwnerType.ValueString() != "" {
		updateReq.SetResourceOwnerType(plan.ResourceOwnerType.ValueString())
	}
	if !plan.ResourceOwner.IsNull() && plan.ResourceOwner.ValueString() != "" {
		updateReq.SetResourceOwner(plan.ResourceOwner.ValueString())
	}
	if !plan.AccessQuery.IsNull() && plan.AccessQuery.ValueString() != "" {
		updateReq.SetAccessquery(plan.AccessQuery.ValueString())
	}
	if !plan.EnableCopyAccess.IsNull() && plan.EnableCopyAccess.ValueString() != "" {
		updateReq.SetEnableCopyAccess(plan.EnableCopyAccess.ValueString())
	}
	if !plan.DisableNewAccountRequestIfAccountExists.IsNull() && plan.DisableNewAccountRequestIfAccountExists.ValueString() != "" {
		updateReq.SetDisableNewAccountRequestIfAccountExists(plan.DisableNewAccountRequestIfAccountExists.ValueString())
	}
	if !plan.DisableRemoveAccount.IsNull() && plan.DisableRemoveAccount.ValueString() != "" {
		updateReq.SetDisableRemoveAccount(plan.DisableRemoveAccount.ValueString())
	}
	if !plan.DisableModifyAccount.IsNull() && plan.DisableModifyAccount.ValueString() != "" {
		updateReq.SetDisableModifyAccount(plan.DisableModifyAccount.ValueString())
	}
	if !plan.UserAccountCorrelationRule.IsNull() && plan.UserAccountCorrelationRule.ValueString() != "" {
		updateReq.SetUserAccountCorrelationRule(plan.UserAccountCorrelationRule.ValueString())
	}
	if !plan.CreateEntTaskforRemoveAcc.IsNull() && plan.CreateEntTaskforRemoveAcc.ValueString() != "" {
		updateReq.SetCreateEntTaskforRemoveAcc(plan.CreateEntTaskforRemoveAcc.ValueString())
	}
	if !plan.OutOfBandAction.IsNull() && plan.OutOfBandAction.ValueString() != "" {
		updateReq.SetOutofbandaction(plan.OutOfBandAction.ValueString())
	}
	if !plan.ConnectionConfig.IsNull() && plan.ConnectionConfig.ValueString() != "" {
		updateReq.SetConnectionconfig(plan.ConnectionConfig.ValueString())
	}
	if !plan.Requestable.IsNull() && plan.Requestable.ValueString() != "" {
		updateReq.SetRequestable(plan.Requestable.ValueString())
	}
	if !plan.ParentAccountPattern.IsNull() && plan.ParentAccountPattern.ValueString() != "" {
		updateReq.SetParentAccountPattern(plan.ParentAccountPattern.ValueString())
	}
	if !plan.ServiceAccountNameRule.IsNull() && plan.ServiceAccountNameRule.ValueString() != "" {
		updateReq.SetServiceAccountNameRule(plan.ServiceAccountNameRule.ValueString())
	}
	if !plan.ServiceAccountAccessQuery.IsNull() && plan.ServiceAccountAccessQuery.ValueString() != "" {
		updateReq.SetServiceAccountAccessQuery(plan.ServiceAccountAccessQuery.ValueString())
	}
	if !plan.BlockInflightRequest.IsNull() && plan.BlockInflightRequest.ValueString() != "" {
		updateReq.SetBlockInflightRequest(plan.BlockInflightRequest.ValueString())
	}
	if !plan.AccountNameRule.IsNull() && plan.AccountNameRule.ValueString() != "" {
		updateReq.SetAccountNameRule(plan.AccountNameRule.ValueString())
	}
	if !plan.AllowChangePasswordSQLQuery.IsNull() && plan.AllowChangePasswordSQLQuery.ValueString() != "" {
		updateReq.SetAllowChangePasswordSqlquery(plan.AllowChangePasswordSQLQuery.ValueString())
	}
	if !plan.AccountNameValidatorRegex.IsNull() && plan.AccountNameValidatorRegex.ValueString() != "" {
		updateReq.SetAccountNameValidatorRegex(plan.AccountNameValidatorRegex.ValueString())
	}
	if !plan.StatusConfig.IsNull() && plan.StatusConfig.ValueString() != "" {
		updateReq.SetStatusConfig(plan.StatusConfig.ValueString())
	}
	if !plan.PluginConfigs.IsNull() && plan.PluginConfigs.ValueString() != "" {
		updateReq.SetPluginConfigs(plan.PluginConfigs.ValueString())
	}
	if !plan.EndpointConfig.IsNull() && plan.EndpointConfig.ValueString() != "" {
		updateReq.SetEndpointConfig(plan.EndpointConfig.ValueString())
	}
	if !plan.CustomProperty1.IsNull() && plan.CustomProperty1.ValueString() != "" {
		updateReq.SetCustomproperty1(plan.CustomProperty1.ValueString())
	}
	if !plan.CustomProperty2.IsNull() && plan.CustomProperty2.ValueString() != "" {
		updateReq.SetCustomproperty2(plan.CustomProperty2.ValueString())
	}
	if !plan.CustomProperty3.IsNull() && plan.CustomProperty3.ValueString() != "" {
		updateReq.SetCustomproperty3(plan.CustomProperty3.ValueString())
	}
	if !plan.CustomProperty4.IsNull() && plan.CustomProperty4.ValueString() != "" {
		updateReq.SetCustomproperty4(plan.CustomProperty4.ValueString())
	}
	if !plan.CustomProperty5.IsNull() && plan.CustomProperty5.ValueString() != "" {
		updateReq.SetCustomproperty5(plan.CustomProperty5.ValueString())
	}
	if !plan.CustomProperty6.IsNull() && plan.CustomProperty6.ValueString() != "" {
		updateReq.SetCustomproperty6(plan.CustomProperty6.ValueString())
	}
	if !plan.CustomProperty7.IsNull() && plan.CustomProperty7.ValueString() != "" {
		updateReq.SetCustomproperty7(plan.CustomProperty7.ValueString())
	}
	if !plan.CustomProperty8.IsNull() && plan.CustomProperty8.ValueString() != "" {
		updateReq.SetCustomproperty8(plan.CustomProperty8.ValueString())
	}
	if !plan.CustomProperty9.IsNull() && plan.CustomProperty9.ValueString() != "" {
		updateReq.SetCustomproperty9(plan.CustomProperty9.ValueString())
	}
	if !plan.CustomProperty10.IsNull() && plan.CustomProperty10.ValueString() != "" {
		updateReq.SetCustomproperty10(plan.CustomProperty10.ValueString())
	}
	if !plan.CustomProperty11.IsNull() && plan.CustomProperty11.ValueString() != "" {
		updateReq.SetCustomproperty11(plan.CustomProperty11.ValueString())
	}
	if !plan.CustomProperty12.IsNull() && plan.CustomProperty12.ValueString() != "" {
		updateReq.SetCustomproperty12(plan.CustomProperty12.ValueString())
	}
	if !plan.CustomProperty13.IsNull() && plan.CustomProperty13.ValueString() != "" {
		updateReq.SetCustomproperty13(plan.CustomProperty13.ValueString())
	}
	if !plan.CustomProperty14.IsNull() && plan.CustomProperty14.ValueString() != "" {
		updateReq.SetCustomproperty14(plan.CustomProperty14.ValueString())
	}
	if !plan.CustomProperty15.IsNull() && plan.CustomProperty15.ValueString() != "" {
		updateReq.SetCustomproperty15(plan.CustomProperty15.ValueString())
	}
	if !plan.CustomProperty16.IsNull() && plan.CustomProperty16.ValueString() != "" {
		updateReq.SetCustomproperty16(plan.CustomProperty16.ValueString())
	}
	if !plan.CustomProperty17.IsNull() && plan.CustomProperty17.ValueString() != "" {
		updateReq.SetCustomproperty17(plan.CustomProperty17.ValueString())
	}
	if !plan.CustomProperty18.IsNull() && plan.CustomProperty18.ValueString() != "" {
		updateReq.SetCustomproperty18(plan.CustomProperty18.ValueString())
	}
	if !plan.CustomProperty19.IsNull() && plan.CustomProperty19.ValueString() != "" {
		updateReq.SetCustomproperty19(plan.CustomProperty19.ValueString())
	}
	if !plan.CustomProperty20.IsNull() && plan.CustomProperty20.ValueString() != "" {
		updateReq.SetCustomproperty20(plan.CustomProperty20.ValueString())
	}
	if !plan.CustomProperty21.IsNull() && plan.CustomProperty21.ValueString() != "" {
		updateReq.SetCustomproperty21(plan.CustomProperty21.ValueString())
	}
	if !plan.CustomProperty22.IsNull() && plan.CustomProperty22.ValueString() != "" {
		updateReq.SetCustomproperty22(plan.CustomProperty22.ValueString())
	}
	if !plan.CustomProperty23.IsNull() && plan.CustomProperty23.ValueString() != "" {
		updateReq.SetCustomproperty23(plan.CustomProperty23.ValueString())
	}
	if !plan.CustomProperty24.IsNull() && plan.CustomProperty24.ValueString() != "" {
		updateReq.SetCustomproperty24(plan.CustomProperty24.ValueString())
	}
	if !plan.CustomProperty25.IsNull() && plan.CustomProperty25.ValueString() != "" {
		updateReq.SetCustomproperty25(plan.CustomProperty25.ValueString())
	}
	if !plan.CustomProperty26.IsNull() && plan.CustomProperty26.ValueString() != "" {
		updateReq.SetCustomproperty26(plan.CustomProperty26.ValueString())
	}
	if !plan.CustomProperty27.IsNull() && plan.CustomProperty27.ValueString() != "" {
		updateReq.SetCustomproperty27(plan.CustomProperty27.ValueString())
	}
	if !plan.CustomProperty28.IsNull() && plan.CustomProperty28.ValueString() != "" {
		updateReq.SetCustomproperty28(plan.CustomProperty28.ValueString())
	}
	if !plan.CustomProperty29.IsNull() && plan.CustomProperty29.ValueString() != "" {
		updateReq.SetCustomproperty29(plan.CustomProperty29.ValueString())
	}
	if !plan.CustomProperty30.IsNull() && plan.CustomProperty30.ValueString() != "" {
		updateReq.SetCustomproperty30(plan.CustomProperty30.ValueString())
	}
	if !plan.CustomProperty31.IsNull() && plan.CustomProperty31.ValueString() != "" {
		updateReq.SetCustomproperty31(plan.CustomProperty31.ValueString())
	}
	if !plan.CustomProperty32.IsNull() && plan.CustomProperty32.ValueString() != "" {
		updateReq.SetCustomproperty32(plan.CustomProperty32.ValueString())
	}
	if !plan.CustomProperty33.IsNull() && plan.CustomProperty33.ValueString() != "" {
		updateReq.SetCustomproperty33(plan.CustomProperty33.ValueString())
	}
	if !plan.CustomProperty34.IsNull() && plan.CustomProperty34.ValueString() != "" {
		updateReq.SetCustomproperty34(plan.CustomProperty34.ValueString())
	}
	if !plan.CustomProperty35.IsNull() && plan.CustomProperty35.ValueString() != "" {
		updateReq.SetCustomproperty35(plan.CustomProperty35.ValueString())
	}
	if !plan.CustomProperty36.IsNull() && plan.CustomProperty36.ValueString() != "" {
		updateReq.SetCustomproperty36(plan.CustomProperty36.ValueString())
	}
	if !plan.CustomProperty37.IsNull() && plan.CustomProperty37.ValueString() != "" {
		updateReq.SetCustomproperty37(plan.CustomProperty37.ValueString())
	}
	if !plan.CustomProperty38.IsNull() && plan.CustomProperty38.ValueString() != "" {
		updateReq.SetCustomproperty38(plan.CustomProperty38.ValueString())
	}
	if !plan.CustomProperty39.IsNull() && plan.CustomProperty39.ValueString() != "" {
		updateReq.SetCustomproperty39(plan.CustomProperty39.ValueString())
	}
	if !plan.CustomProperty40.IsNull() && plan.CustomProperty40.ValueString() != "" {
		updateReq.SetCustomproperty40(plan.CustomProperty40.ValueString())
	}
	if !plan.CustomProperty41.IsNull() && plan.CustomProperty41.ValueString() != "" {
		updateReq.SetCustomproperty41(plan.CustomProperty41.ValueString())
	}
	if !plan.CustomProperty42.IsNull() && plan.CustomProperty42.ValueString() != "" {
		updateReq.SetCustomproperty42(plan.CustomProperty42.ValueString())
	}
	if !plan.CustomProperty43.IsNull() && plan.CustomProperty43.ValueString() != "" {
		updateReq.SetCustomproperty43(plan.CustomProperty43.ValueString())
	}
	if !plan.CustomProperty44.IsNull() && plan.CustomProperty44.ValueString() != "" {
		updateReq.SetCustomproperty44(plan.CustomProperty44.ValueString())
	}
	if !plan.CustomProperty45.IsNull() && plan.CustomProperty45.ValueString() != "" {
		updateReq.SetCustomproperty45(plan.CustomProperty45.ValueString())
	}
	if !plan.CustomProperty1Label.IsNull() && plan.CustomProperty1Label.ValueString() != "" {
		updateReq.SetCustomproperty1Label(plan.CustomProperty1Label.ValueString())
	}
	if !plan.CustomProperty2Label.IsNull() && plan.CustomProperty2Label.ValueString() != "" {
		updateReq.SetCustomproperty2Label(plan.CustomProperty2Label.ValueString())
	}
	if !plan.CustomProperty3Label.IsNull() && plan.CustomProperty3Label.ValueString() != "" {
		updateReq.SetCustomproperty3Label(plan.CustomProperty3Label.ValueString())
	}
	if !plan.CustomProperty4Label.IsNull() && plan.CustomProperty4Label.ValueString() != "" {
		updateReq.SetCustomproperty4Label(plan.CustomProperty4Label.ValueString())
	}
	if !plan.CustomProperty5Label.IsNull() && plan.CustomProperty5Label.ValueString() != "" {
		updateReq.SetCustomproperty5Label(plan.CustomProperty5Label.ValueString())
	}
	if !plan.CustomProperty6Label.IsNull() && plan.CustomProperty6Label.ValueString() != "" {
		updateReq.SetCustomproperty6Label(plan.CustomProperty6Label.ValueString())
	}
	if !plan.CustomProperty7Label.IsNull() && plan.CustomProperty7Label.ValueString() != "" {
		updateReq.SetCustomproperty7Label(plan.CustomProperty7Label.ValueString())
	}
	if !plan.CustomProperty8Label.IsNull() && plan.CustomProperty8Label.ValueString() != "" {
		updateReq.SetCustomproperty8Label(plan.CustomProperty8Label.ValueString())
	}
	if !plan.CustomProperty9Label.IsNull() && plan.CustomProperty9Label.ValueString() != "" {
		updateReq.SetCustomproperty9Label(plan.CustomProperty9Label.ValueString())
	}
	if !plan.CustomProperty10Label.IsNull() && plan.CustomProperty10Label.ValueString() != "" {
		updateReq.SetCustomproperty10Label(plan.CustomProperty10Label.ValueString())
	}
	if !plan.CustomProperty11Label.IsNull() && plan.CustomProperty11Label.ValueString() != "" {
		updateReq.SetCustomproperty11Label(plan.CustomProperty11Label.ValueString())
	}
	if !plan.CustomProperty12Label.IsNull() && plan.CustomProperty12Label.ValueString() != "" {
		updateReq.SetCustomproperty12Label(plan.CustomProperty12Label.ValueString())
	}
	if !plan.CustomProperty13Label.IsNull() && plan.CustomProperty13Label.ValueString() != "" {
		updateReq.SetCustomproperty13Label(plan.CustomProperty13Label.ValueString())
	}
	if !plan.CustomProperty14Label.IsNull() && plan.CustomProperty14Label.ValueString() != "" {
		updateReq.SetCustomproperty14Label(plan.CustomProperty14Label.ValueString())
	}
	if !plan.CustomProperty15Label.IsNull() && plan.CustomProperty15Label.ValueString() != "" {
		updateReq.SetCustomproperty15Label(plan.CustomProperty15Label.ValueString())
	}
	if !plan.CustomProperty16Label.IsNull() && plan.CustomProperty16Label.ValueString() != "" {
		updateReq.SetCustomproperty16Label(plan.CustomProperty16Label.ValueString())
	}
	if !plan.CustomProperty17Label.IsNull() && plan.CustomProperty17Label.ValueString() != "" {
		updateReq.SetCustomproperty17Label(plan.CustomProperty17Label.ValueString())
	}
	if !plan.CustomProperty18Label.IsNull() && plan.CustomProperty18Label.ValueString() != "" {
		updateReq.SetCustomproperty18Label(plan.CustomProperty18Label.ValueString())
	}
	if !plan.CustomProperty19Label.IsNull() && plan.CustomProperty19Label.ValueString() != "" {
		updateReq.SetCustomproperty19Label(plan.CustomProperty19Label.ValueString())
	}
	if !plan.CustomProperty20Label.IsNull() && plan.CustomProperty20Label.ValueString() != "" {
		updateReq.SetCustomproperty20Label(plan.CustomProperty20Label.ValueString())
	}
	if !plan.CustomProperty21Label.IsNull() && plan.CustomProperty21Label.ValueString() != "" {
		updateReq.SetCustomproperty21Label(plan.CustomProperty21Label.ValueString())
	}
	if !plan.CustomProperty22Label.IsNull() && plan.CustomProperty22Label.ValueString() != "" {
		updateReq.SetCustomproperty22Label(plan.CustomProperty22Label.ValueString())
	}
	if !plan.CustomProperty23Label.IsNull() && plan.CustomProperty23Label.ValueString() != "" {
		updateReq.SetCustomproperty23Label(plan.CustomProperty23Label.ValueString())
	}
	if !plan.CustomProperty24Label.IsNull() && plan.CustomProperty24Label.ValueString() != "" {
		updateReq.SetCustomproperty24Label(plan.CustomProperty24Label.ValueString())
	}
	if !plan.CustomProperty25Label.IsNull() && plan.CustomProperty25Label.ValueString() != "" {
		updateReq.SetCustomproperty25Label(plan.CustomProperty25Label.ValueString())
	}
	if !plan.CustomProperty26Label.IsNull() && plan.CustomProperty26Label.ValueString() != "" {
		updateReq.SetCustomproperty26Label(plan.CustomProperty26Label.ValueString())
	}
	if !plan.CustomProperty27Label.IsNull() && plan.CustomProperty27Label.ValueString() != "" {
		updateReq.SetCustomproperty27Label(plan.CustomProperty27Label.ValueString())
	}
	if !plan.CustomProperty28Label.IsNull() && plan.CustomProperty28Label.ValueString() != "" {
		updateReq.SetCustomproperty28Label(plan.CustomProperty28Label.ValueString())
	}
	if !plan.CustomProperty29Label.IsNull() && plan.CustomProperty29Label.ValueString() != "" {
		updateReq.SetCustomproperty29Label(plan.CustomProperty29Label.ValueString())
	}
	if !plan.CustomProperty30Label.IsNull() && plan.CustomProperty30Label.ValueString() != "" {
		updateReq.SetCustomproperty30Label(plan.CustomProperty30Label.ValueString())
	}
	if !plan.CustomProperty31Label.IsNull() && plan.CustomProperty31Label.ValueString() != "" {
		updateReq.SetCustomproperty31Label(plan.CustomProperty31Label.ValueString())
	}
	if !plan.CustomProperty32Label.IsNull() && plan.CustomProperty32Label.ValueString() != "" {
		updateReq.SetCustomproperty32Label(plan.CustomProperty32Label.ValueString())
	}
	if !plan.CustomProperty33Label.IsNull() && plan.CustomProperty33Label.ValueString() != "" {
		updateReq.SetCustomproperty33Label(plan.CustomProperty33Label.ValueString())
	}
	if !plan.CustomProperty34Label.IsNull() && plan.CustomProperty34Label.ValueString() != "" {
		updateReq.SetCustomproperty34Label(plan.CustomProperty34Label.ValueString())
	}
	if !plan.CustomProperty35Label.IsNull() && plan.CustomProperty35Label.ValueString() != "" {
		updateReq.SetCustomproperty35Label(plan.CustomProperty35Label.ValueString())
	}
	if !plan.CustomProperty36Label.IsNull() && plan.CustomProperty36Label.ValueString() != "" {
		updateReq.SetCustomproperty36Label(plan.CustomProperty36Label.ValueString())
	}
	if !plan.CustomProperty37Label.IsNull() && plan.CustomProperty37Label.ValueString() != "" {
		updateReq.SetCustomproperty37Label(plan.CustomProperty37Label.ValueString())
	}
	if !plan.CustomProperty38Label.IsNull() && plan.CustomProperty38Label.ValueString() != "" {
		updateReq.SetCustomproperty38Label(plan.CustomProperty38Label.ValueString())
	}
	if !plan.CustomProperty39Label.IsNull() && plan.CustomProperty39Label.ValueString() != "" {
		updateReq.SetCustomproperty39Label(plan.CustomProperty39Label.ValueString())
	}
	if !plan.CustomProperty40Label.IsNull() && plan.CustomProperty40Label.ValueString() != "" {
		updateReq.SetCustomproperty40Label(plan.CustomProperty40Label.ValueString())
	}
	if !plan.CustomProperty41Label.IsNull() && plan.CustomProperty41Label.ValueString() != "" {
		updateReq.SetCustomproperty41Label(plan.CustomProperty41Label.ValueString())
	}
	if !plan.CustomProperty42Label.IsNull() && plan.CustomProperty42Label.ValueString() != "" {
		updateReq.SetCustomproperty42Label(plan.CustomProperty42Label.ValueString())
	}
	if !plan.CustomProperty43Label.IsNull() && plan.CustomProperty43Label.ValueString() != "" {
		updateReq.SetCustomproperty43Label(plan.CustomProperty43Label.ValueString())
	}
	if !plan.CustomProperty44Label.IsNull() && plan.CustomProperty44Label.ValueString() != "" {
		updateReq.SetCustomproperty44Label(plan.CustomProperty44Label.ValueString())
	}
	if !plan.CustomProperty45Label.IsNull() && plan.CustomProperty45Label.ValueString() != "" {
		updateReq.SetCustomproperty45Label(plan.CustomProperty45Label.ValueString())
	}
	if !plan.CustomProperty46Label.IsNull() && plan.CustomProperty46Label.ValueString() != "" {
		updateReq.SetCustomproperty46Label(plan.CustomProperty46Label.ValueString())
	}
	if !plan.CustomProperty47Label.IsNull() && plan.CustomProperty47Label.ValueString() != "" {
		updateReq.SetCustomproperty47Label(plan.CustomProperty47Label.ValueString())
	}
	if !plan.CustomProperty48Label.IsNull() && plan.CustomProperty48Label.ValueString() != "" {
		updateReq.SetCustomproperty48Label(plan.CustomProperty48Label.ValueString())
	}
	if !plan.CustomProperty49Label.IsNull() && plan.CustomProperty49Label.ValueString() != "" {
		updateReq.SetCustomproperty49Label(plan.CustomProperty49Label.ValueString())
	}
	if !plan.CustomProperty50Label.IsNull() && plan.CustomProperty50Label.ValueString() != "" {
		updateReq.SetCustomproperty50Label(plan.CustomProperty50Label.ValueString())
	}
	if !plan.CustomProperty60Label.IsNull() && plan.CustomProperty60Label.ValueString() != "" {
		updateReq.SetCustomproperty60Label(plan.CustomProperty60Label.ValueString())
	}

	var requestableRoleTypes []openapi.UpdateEndpointRequestRequestableRoleTypeInner

	for _, role := range plan.RequestableRoleType {
		var requestableRole openapi.UpdateEndpointRequestRequestableRoleTypeInner

		if !role.RoleType.IsNull() && role.RoleType.ValueString() != "" {
			requestableRole.RoleType = role.RoleType.ValueStringPointer()
		}
		if !role.RequestOption.IsNull() && role.RequestOption.ValueString() != "" {
			requestableRole.RequestOption = role.RequestOption.ValueStringPointer()
		}
		if !role.Required.IsNull() {
			requestableRole.Required = role.Required.ValueBoolPointer()
		}
		if !role.RequestedQuery.IsNull() && role.RequestedQuery.ValueString() != "" {
			requestableRole.RequestedQuery = role.RequestedQuery.ValueStringPointer()
		}
		if !role.SelectedQuery.IsNull() && role.SelectedQuery.ValueString() != "" {
			requestableRole.SelectedQuery = role.SelectedQuery.ValueStringPointer()
		}
		if !role.ShowOn.IsNull() && role.ShowOn.ValueString() != "" {
			requestableRole.ShowOn = role.ShowOn.ValueStringPointer()
		}

		// Add the role type if any value is set
		if requestableRole.RoleType != nil || requestableRole.RequestOption != nil || requestableRole.Required != nil ||
			requestableRole.RequestedQuery != nil || requestableRole.SelectedQuery != nil || requestableRole.ShowOn != nil {
			requestableRoleTypes = append(requestableRoleTypes, requestableRole)
		}
	}

	// Assign to update request if populated
	if len(requestableRoleTypes) > 0 {
		updateReq.RequestableRoleType = requestableRoleTypes
	}

	var emailTemplates []openapi.UpdateEndpointRequestEmailTemplateInner

	// Iterate over email templates from Terraform plan
	for _, template := range plan.EmailTemplate {
		var emailTemplate openapi.UpdateEndpointRequestEmailTemplateInner

		// Check and set each field if it's not null or empty
		if !template.EmailTemplateType.IsNull() && template.EmailTemplateType.ValueString() != "" {
			emailTemplate.EmailTemplateType = template.EmailTemplateType.ValueStringPointer()
		}
		if !template.TaskType.IsNull() && template.TaskType.ValueString() != "" {
			emailTemplate.TaskType = template.TaskType.ValueStringPointer()
		}
		if !template.EmailTemplate.IsNull() && template.EmailTemplate.ValueString() != "" {
			emailTemplate.EmailTemplate = template.EmailTemplate.ValueStringPointer()
		}

		// Append only if at least one field is set
		if emailTemplate.EmailTemplateType != nil || emailTemplate.TaskType != nil || emailTemplate.EmailTemplate != nil {
			emailTemplates = append(emailTemplates, emailTemplate)
		}
	}

	// Assign to update request if email templates exist
	if len(emailTemplates) > 0 {
		updateReq.EmailTemplate = emailTemplates
	}

	var mappedEndpoints []openapi.UpdateEndpointRequestMappedEndpointsInner

	// Iterate over mapped endpoints from Terraform plan
	for _, endpoint := range plan.MappedEndpoints {
		var mappedEndpoint openapi.UpdateEndpointRequestMappedEndpointsInner

		// Check and set each field if it's not null or empty
		if !endpoint.SecuritySystem.IsNull() && endpoint.SecuritySystem.ValueString() != "" {
			mappedEndpoint.Securitysystem = endpoint.SecuritySystem.ValueStringPointer()
		}
		if !endpoint.Endpoint.IsNull() && endpoint.Endpoint.ValueString() != "" {
			mappedEndpoint.Endpoint = endpoint.Endpoint.ValueStringPointer()
		}
		if !endpoint.Requestable.IsNull() && endpoint.Requestable.ValueString() != "" {
			mappedEndpoint.Requestable = endpoint.Requestable.ValueStringPointer()
		}
		if !endpoint.Operation.IsNull() && endpoint.Operation.ValueString() != "" {
			mappedEndpoint.Operation = endpoint.Operation.ValueStringPointer()
		}

		// Append only if at least one field is set
		if mappedEndpoint.Securitysystem != nil || mappedEndpoint.Endpoint != nil ||
			mappedEndpoint.Requestable != nil || mappedEndpoint.Operation != nil {
			mappedEndpoints = append(mappedEndpoints, mappedEndpoint)
		}
	}

	// Assign to update request if mapped endpoints exist
	if len(mappedEndpoints) > 0 {
		updateReq.MappedEndpoints = mappedEndpoints
	}

	apiResp, httpResp, err := apiClient.EndpointsAPI.
		UpdateEndpoint(ctx).
		UpdateEndpointRequest(*updateReq).
		Execute()
	if err != nil {
		log.Printf("Error Updating Endpoint: %v, HTTP Response: %v", err, httpResp)
		resp.Diagnostics.AddError(
			"Error Updating Endpoint",
			"Check logs for details.",
		)
		return
	}

	if plan.ID.IsUnknown() || plan.ID.IsNull() {
		plan.ID = types.StringValue("endpoint-" + plan.EndpointName.ValueString())
	}
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
		log.Printf("Error marshaling result: %v", err)
		return
	}
	plan.Result = types.StringValue(string(resultJSON))

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *endpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
