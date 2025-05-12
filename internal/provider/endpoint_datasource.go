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

	openapi "github.com/saviynt/saviynt-api-go-client/endpoints"

	s "github.com/saviynt/saviynt-api-go-client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EndpointsDataSource struct {
	client *s.Client
	token  string
}

var _ datasource.DataSource = &EndpointsDataSource{}
var _ datasource.DataSourceWithConfigure = &EndpointsDataSource{}

func NewEndpointsDataSource() datasource.DataSource {
	return &EndpointsDataSource{}
}

type EndpointsDataSourceModel struct {
	Results        []Endpoint   `tfsdk:"results"`
	DisplayCount   types.Int64  `tfsdk:"display_count"`
	ErrorCode      types.String `tfsdk:"error_code"`
	TotalCount     types.Int64  `tfsdk:"total_count"`
	Message        types.String `tfsdk:"message"`
	EndpointName   types.String `tfsdk:"endpointname"`
	ConnectionType types.String `tfsdk:"connection_type"`
	Displayname    types.String `tfsdk:"displayname"`
	Owner          types.String `tfsdk:"owner"`
	FilterCriteria types.Map    `tfsdk:"filter_criteria"`
	Max            types.String `tfsdk:"max"`
}

type Endpoint struct {
	Id                              types.String `tfsdk:"id"`
	Description                     types.String `tfsdk:"description"`
	StatusForUniqueAccount          types.String `tfsdk:"status_for_unique_account"`
	Requestowner                    types.String `tfsdk:"requestowner"`
	Requestable                     types.String `tfsdk:"requestable"`
	PrimaryAccountType              types.String `tfsdk:"primary_account_type"`
	AccountTypeNoPasswordChange     types.String `tfsdk:"account_type_no_password_change"`
	ServiceAccountNameRule          types.String `tfsdk:"service_account_name_rule"`
	AccountNameValidatorRegex       types.String `tfsdk:"account_name_validator_regex"`
	AllowChangePasswordSqlquery     types.String `tfsdk:"allow_change_password_sqlquery"`
	ParentAccountPattern            types.String `tfsdk:"parent_account_pattern"`
	OwnerType                       types.String `tfsdk:"owner_type"`
	Securitysystem                  types.String `tfsdk:"securitysystem"`
	Endpointname                    types.String `tfsdk:"endpointname"`
	UpdatedBy                       types.String `tfsdk:"updated_by"`
	Accessquery                     types.String `tfsdk:"accessquery"`
	Status                          types.String `tfsdk:"status"`
	DisplayName                     types.String `tfsdk:"display_name"`
	UpdateDate                      types.String `tfsdk:"update_date"`
	AllowRemoveAllRoleOnRequest     types.String `tfsdk:"allow_remove_all_role_on_request"`
	RoleTypeAsJson                  types.String `tfsdk:"role_type_as_json"`
	EntsWithNewAccount              types.String `tfsdk:"ents_with_new_account"`
	ConnectionconfigAsJson          types.String `tfsdk:"connectionconfig_as_json"`
	Connectionconfig                types.String `tfsdk:"connectionconfig"`
	AccountNameRule                 types.String `tfsdk:"account_name_rule"`
	ChangePasswordAccessQuery       types.String `tfsdk:"change_password_access_query"`
	Disableaccountrequest           types.String `tfsdk:"disableaccountrequest"`
	PluginConfigs                   types.String `tfsdk:"plugin_configs"`
	DisableaccountrequestServiceAccount types.String `tfsdk:"disableaccountrequest_service_account"`
	Requestableapplication          types.String `tfsdk:"requestableapplication"`
	CreatedFrom                     types.String `tfsdk:"created_from"`
	CreatedBy                       types.String `tfsdk:"created_by"`
	CreateDate                      types.String `tfsdk:"create_date"`
	ParentEndpoint                  types.String `tfsdk:"parent_endpoint"`
	BaseLineConfig                  types.String `tfsdk:"base_line_config"`
	Requestownertype                types.String `tfsdk:"requestownertype"`
	CreateEntTaskforRemoveAcc       types.String `tfsdk:"create_ent_taskfor_remove_acc"`
	EnableCopyAccess                types.String `tfsdk:"enable_copy_access"`
	AccountTypeNoDeprovision        types.String `tfsdk:"account_type_no_deprovision"`
	EndpointConfig                  types.String `tfsdk:"endpoint_config"`
	Taskemailtemplates              types.String `tfsdk:"taskemailtemplates"`
	Ownerkey                        types.String `tfsdk:"ownerkey"`
	ServiceAccountAccessQuery       types.String `tfsdk:"service_account_access_query"`
	UserAccountCorrelationRule      types.String `tfsdk:"user_account_correlation_rule"`
	StatusConfig                    types.String `tfsdk:"status_config"`

	CustomProperty1  types.String `tfsdk:"custom_property_1"`
	CustomProperty2  types.String `tfsdk:"custom_property_2"`
	CustomProperty3  types.String `tfsdk:"custom_property_3"`
	CustomProperty4  types.String `tfsdk:"custom_property_4"`
	CustomProperty5  types.String `tfsdk:"custom_property_5"`
	CustomProperty6  types.String `tfsdk:"custom_property_6"`
	CustomProperty7  types.String `tfsdk:"custom_property_7"`
	CustomProperty8  types.String `tfsdk:"custom_property_8"`
	CustomProperty9  types.String `tfsdk:"custom_property_9"`
	CustomProperty10 types.String `tfsdk:"custom_property_10"`
	CustomProperty11 types.String `tfsdk:"custom_property_11"`
	CustomProperty12 types.String `tfsdk:"custom_property_12"`
	CustomProperty13 types.String `tfsdk:"custom_property_13"`
	CustomProperty14 types.String `tfsdk:"custom_property_14"`
	CustomProperty15 types.String `tfsdk:"custom_property_15"`
	CustomProperty16 types.String `tfsdk:"custom_property_16"`
	CustomProperty17 types.String `tfsdk:"custom_property_17"`
	CustomProperty18 types.String `tfsdk:"custom_property_18"`
	CustomProperty19 types.String `tfsdk:"custom_property_19"`
	CustomProperty20 types.String `tfsdk:"custom_property_20"`
	CustomProperty21 types.String `tfsdk:"custom_property_21"`
	CustomProperty22 types.String `tfsdk:"custom_property_22"`
	CustomProperty23 types.String `tfsdk:"custom_property_23"`
	CustomProperty24 types.String `tfsdk:"custom_property_24"`
	CustomProperty25 types.String `tfsdk:"custom_property_25"`
	CustomProperty26 types.String `tfsdk:"custom_property_26"`
	CustomProperty27 types.String `tfsdk:"custom_property_27"`
	CustomProperty28 types.String `tfsdk:"custom_property_28"`
	CustomProperty29 types.String `tfsdk:"custom_property_29"`
	CustomProperty30 types.String `tfsdk:"custom_property_30"`
	CustomProperty31 types.String `tfsdk:"custom_property_31"`
	CustomProperty32 types.String `tfsdk:"custom_property_32"`
	CustomProperty33 types.String `tfsdk:"custom_property_33"`
	CustomProperty34 types.String `tfsdk:"custom_property_34"`
	CustomProperty35 types.String `tfsdk:"custom_property_35"`
	CustomProperty36 types.String `tfsdk:"custom_property_36"`
	CustomProperty37 types.String `tfsdk:"custom_property_37"`
	CustomProperty38 types.String `tfsdk:"custom_property_38"`
	CustomProperty39 types.String `tfsdk:"custom_property_39"`
	CustomProperty40 types.String `tfsdk:"custom_property_40"`
	CustomProperty41 types.String `tfsdk:"custom_property_41"`
	CustomProperty42 types.String `tfsdk:"custom_property_42"`
	CustomProperty43 types.String `tfsdk:"custom_property_43"`
	CustomProperty44 types.String `tfsdk:"custom_property_44"`
	CustomProperty45 types.String `tfsdk:"custom_property_45"`

	AccountCustomProperty1Label  types.String `tfsdk:"account_custom_property_1_label"`
	AccountCustomProperty2Label  types.String `tfsdk:"account_custom_property_2_label"`
	AccountCustomProperty3Label  types.String `tfsdk:"account_custom_property_3_label"`
	AccountCustomProperty4Label  types.String `tfsdk:"account_custom_property_4_label"`
	AccountCustomProperty5Label  types.String `tfsdk:"account_custom_property_5_label"`
	AccountCustomProperty6Label  types.String `tfsdk:"account_custom_property_6_label"`
	AccountCustomProperty7Label  types.String `tfsdk:"account_custom_property_7_label"`
	AccountCustomProperty8Label  types.String `tfsdk:"account_custom_property_8_label"`
	AccountCustomProperty9Label  types.String `tfsdk:"account_custom_property_9_label"`
	AccountCustomProperty10Label types.String `tfsdk:"account_custom_property_10_label"`
	AccountCustomProperty11Label types.String `tfsdk:"account_custom_property_11_label"`
	AccountCustomProperty12Label types.String `tfsdk:"account_custom_property_12_label"`
	AccountCustomProperty13Label types.String `tfsdk:"account_custom_property_13_label"`
	AccountCustomProperty14Label types.String `tfsdk:"account_custom_property_14_label"`
	AccountCustomProperty15Label types.String `tfsdk:"account_custom_property_15_label"`
	AccountCustomProperty16Label types.String `tfsdk:"account_custom_property_16_label"`
	AccountCustomProperty17Label types.String `tfsdk:"account_custom_property_17_label"`
	AccountCustomProperty18Label types.String `tfsdk:"account_custom_property_18_label"`
	AccountCustomProperty19Label types.String `tfsdk:"account_custom_property_19_label"`
	AccountCustomProperty20Label types.String `tfsdk:"account_custom_property_20_label"`
	AccountCustomProperty21Label types.String `tfsdk:"account_custom_property_21_label"`
	AccountCustomProperty22Label types.String `tfsdk:"account_custom_property_22_label"`
	AccountCustomProperty23Label types.String `tfsdk:"account_custom_property_23_label"`
	AccountCustomProperty24Label types.String `tfsdk:"account_custom_property_24_label"`
	AccountCustomProperty25Label types.String `tfsdk:"account_custom_property_25_label"`
	AccountCustomProperty26Label types.String `tfsdk:"account_custom_property_26_label"`
	AccountCustomProperty27Label types.String `tfsdk:"account_custom_property_27_label"`
	AccountCustomProperty28Label types.String `tfsdk:"account_custom_property_28_label"`
	AccountCustomProperty29Label types.String `tfsdk:"account_custom_property_29_label"`
	AccountCustomProperty30Label types.String `tfsdk:"account_custom_property_30_label"`
	
	CustomProperty31Label                   types.String `tfsdk:"custom_property_31_label"`
	CustomProperty32Label                   types.String `tfsdk:"custom_property_32_label"`
	CustomProperty33Label                   types.String `tfsdk:"custom_property_33_label"`
	CustomProperty34Label                   types.String `tfsdk:"custom_property_34_label"`
	CustomProperty35Label                   types.String `tfsdk:"custom_property_35_label"`
	CustomProperty36Label                   types.String `tfsdk:"custom_property_36_label"`
	CustomProperty37Label                   types.String `tfsdk:"custom_property_37_label"`
	CustomProperty38Label                   types.String `tfsdk:"custom_property_38_label"`
	CustomProperty39Label                   types.String `tfsdk:"custom_property_39_label"`
	CustomProperty40Label                   types.String `tfsdk:"custom_property_40_label"`
	CustomProperty41Label                   types.String `tfsdk:"custom_property_41_label"`
	CustomProperty42Label                   types.String `tfsdk:"custom_property_42_label"`
	CustomProperty43Label                   types.String `tfsdk:"custom_property_43_label"`
	CustomProperty44Label                   types.String `tfsdk:"custom_property_44_label"`
	CustomProperty45Label                   types.String `tfsdk:"custom_property_45_label"`
	CustomProperty46Label                   types.String `tfsdk:"custom_property_46_label"`
	CustomProperty47Label                   types.String `tfsdk:"custom_property_47_label"`
	CustomProperty48Label                   types.String `tfsdk:"custom_property_48_label"`
	CustomProperty49Label                   types.String `tfsdk:"custom_property_49_label"`
	CustomProperty50Label                   types.String `tfsdk:"custom_property_50_label"`
	CustomProperty51Label                   types.String `tfsdk:"custom_property_51_label"`
	CustomProperty52Label                   types.String `tfsdk:"custom_property_52_label"`
	CustomProperty53Label                   types.String `tfsdk:"custom_property_53_label"`
	CustomProperty54Label                   types.String `tfsdk:"custom_property_54_label"`
	CustomProperty55Label                   types.String `tfsdk:"custom_property_55_label"`
	CustomProperty56Label                   types.String `tfsdk:"custom_property_56_label"`
	CustomProperty57Label                   types.String `tfsdk:"custom_property_57_label"`
	CustomProperty58Label                   types.String `tfsdk:"custom_property_58_label"`
	CustomProperty59Label                   types.String `tfsdk:"custom_property_59_label"`
	CustomProperty60Label                   types.String `tfsdk:"custom_property_60_label"`

}

func (d *EndpointsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_endpoints_datasource"
}
func (d *EndpointsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.EndpointDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"endpointname": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by endpoint name",
			},
			"connection_type": schema.StringAttribute{
				Optional:    true,                   
				Description: "Filter by connection type",
			},
			"displayname": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by display name",
			},
			"owner": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by owner",
			},
			"filter_criteria": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Filter criteria",
			},
			"max": schema.StringAttribute{
				Optional: true,
			},
			"message": schema.StringAttribute{
				Computed:    true,
				Description: "API response message",
			},
			"display_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of records returned in the response",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code from the API",
			},
			"total_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Total count of available records",
			},
			"results": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of endpoints retrieved",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "Unique ID of the endpoint",
						},
						"description": schema.StringAttribute{
							Computed:    true,
							Description: "Description for the endpoint",
						},
						"status_for_unique_account": schema.StringAttribute{
							Computed:    true,
							Description: "Status for unique account",
						},
						"requestowner": schema.StringAttribute{
							Computed:    true,
							Description: "Request owner",
						},
						"requestable": schema.StringAttribute{
							Computed:    true,
							Description: "Requestable flag",
						},
						"primary_account_type": schema.StringAttribute{
							Computed:    true,
							Description: "Primary account type",
						},
						"account_type_no_password_change": schema.StringAttribute{
							Computed:    true,
							Description: "Account types for which password change is not allowed",
						},
						"service_account_name_rule": schema.StringAttribute{
							Computed:    true,
							Description: "Rule for generating service account names",
						},
						"account_name_validator_regex": schema.StringAttribute{
							Computed:    true,
							Description: "Regex to validate account name",
						},
						"allow_change_password_sqlquery": schema.StringAttribute{
							Computed:    true,
							Description: "SQL query to allow change password",
						},
						"parent_account_pattern": schema.StringAttribute{
							Computed:    true,
							Description: "Pattern for parent account",
						},
						"owner_type": schema.StringAttribute{
							Computed:    true,
							Description: "Owner type of the endpoint (User/Usergroup)",
						},
						"securitysystem": schema.StringAttribute{
							Computed:    true,
							Description: "Security system associated with the endpoint",
						},
						"endpointname": schema.StringAttribute{
							Computed:    true,
							Description: "Logical name of the endpoint",
						},
						"updated_by": schema.StringAttribute{
							Computed:    true,
							Description: "User who last updated the endpoint",
						},
						"accessquery": schema.StringAttribute{
							Computed:    true,
							Description: "Query to restrict endpoint visibility",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Status of the endpoint",
						},
						"display_name": schema.StringAttribute{
							Computed:    true,
							Description: "User-friendly display name for the endpoint",
						},
						"update_date": schema.StringAttribute{
							Computed:    true,
							Description: "Date when the endpoint was last updated",
						},
						"allow_remove_all_role_on_request": schema.StringAttribute{
							Computed:    true,
							Description: "Whether remove all roles is allowed in request",
						},
						"role_type_as_json": schema.StringAttribute{
							Computed:    true,
							Description: "Role types in JSON format",
						},
						"ents_with_new_account": schema.StringAttribute{
							Computed:    true,
							Description: "Entitlements associated with new account",
						},
						"connectionconfig_as_json": schema.StringAttribute{
							Computed:    true,
							Description: "Connection configuration in JSON",
						},
						"connectionconfig": schema.StringAttribute{
							Computed:    true,
							Description: "Connection configuration",
						},
						"account_name_rule": schema.StringAttribute{
							Computed:    true,
							Description: "Rule to generate account names",
						},
						"change_password_access_query": schema.StringAttribute{
							Computed:    true,
							Description: "Query to restrict password change",
						},
						"disableaccountrequest": schema.StringAttribute{
							Computed:    true,
							Description: "Disable account request",
						},
						"plugin_configs": schema.StringAttribute{
							Computed:    true,
							Description: "Plugin configuration for SmartAssist",
						},
						"disableaccountrequest_service_account": schema.StringAttribute{
							Computed:    true,
							Description: "Disable account request for service accounts",
						},
						"requestableapplication": schema.StringAttribute{
							Computed:    true,
							Description: "Associated requestable application",
						},
						"created_from": schema.StringAttribute{
							Computed:    true,
							Description: "Source of creation",
						},
						"created_by": schema.StringAttribute{
							Computed:    true,
							Description: "User who created the endpoint",
						},
						"create_date": schema.StringAttribute{
							Computed:    true,
							Description: "Date of creation",
						},
						"parent_endpoint": schema.StringAttribute{
							Computed:    true,
							Description: "Parent endpoint",
						},
						"base_line_config": schema.StringAttribute{
							Computed:    true,
							Description: "Baseline configuration",
						},
						"requestownertype": schema.StringAttribute{
							Computed:    true,
							Description: "Type of request owner",
						},
						"create_ent_taskfor_remove_acc": schema.StringAttribute{
							Computed:    true,
							Description: "Whether entitlement task is created for remove account",
						},
						"enable_copy_access": schema.StringAttribute{
							Computed:    true,
							Description: "Whether copy access is enabled",
						},
						"account_type_no_deprovision": schema.StringAttribute{
							Computed:    true,
							Description: "Account types not allowed for deprovision",
						},
						"endpoint_config": schema.StringAttribute{
							Computed:    true,
							Description: "Endpoint configuration",
						},
						"taskemailtemplates": schema.StringAttribute{
							Computed:    true,
							Description: "Task email templates",
						},
						"ownerkey": schema.StringAttribute{
							Computed:    true,
							Description: "Key of the owner",
						},
						"service_account_access_query": schema.StringAttribute{
							Computed:    true,
							Description: "Query to filter service account access",
						},
						"user_account_correlation_rule": schema.StringAttribute{
							Computed:    true,
							Description: "Rule to correlate user and account",
						},
						"status_config": schema.StringAttribute{
							Computed:    true,
							Description: "Status configuration for account operations",
						},
						"custom_property_1":  schema.StringAttribute{Computed: true, Description: "Custom property 1 value for the endpoint."},
						"custom_property_2":  schema.StringAttribute{Computed: true, Description: "Custom property 2 value for the endpoint."},
						"custom_property_3":  schema.StringAttribute{Computed: true, Description: "Custom property 3 value for the endpoint."},
						"custom_property_4":  schema.StringAttribute{Computed: true, Description: "Custom property 4 value for the endpoint."},
						"custom_property_5":  schema.StringAttribute{Computed: true, Description: "Custom property 5 value for the endpoint."},
						"custom_property_6":  schema.StringAttribute{Computed: true, Description: "Custom property 6 value for the endpoint."},
						"custom_property_7":  schema.StringAttribute{Computed: true, Description: "Custom property 7 value for the endpoint."},
						"custom_property_8":  schema.StringAttribute{Computed: true, Description: "Custom property 8 value for the endpoint."},
						"custom_property_9":  schema.StringAttribute{Computed: true, Description: "Custom property 9 value for the endpoint."},
						"custom_property_10": schema.StringAttribute{Computed: true, Description: "Custom property 10 value for the endpoint."},
						"custom_property_11": schema.StringAttribute{Computed: true, Description: "Custom property 11 value for the endpoint."},
						"custom_property_12": schema.StringAttribute{Computed: true, Description: "Custom property 12 value for the endpoint."},
						"custom_property_13": schema.StringAttribute{Computed: true, Description: "Custom property 13 value for the endpoint."},
						"custom_property_14": schema.StringAttribute{Computed: true, Description: "Custom property 14 value for the endpoint."},
						"custom_property_15": schema.StringAttribute{Computed: true, Description: "Custom property 15 value for the endpoint."},
						"custom_property_16": schema.StringAttribute{Computed: true, Description: "Custom property 16 value for the endpoint."},
						"custom_property_17": schema.StringAttribute{Computed: true, Description: "Custom property 17 value for the endpoint."},
						"custom_property_18": schema.StringAttribute{Computed: true, Description: "Custom property 18 value for the endpoint."},
						"custom_property_19": schema.StringAttribute{Computed: true, Description: "Custom property 19 value for the endpoint."},
						"custom_property_20": schema.StringAttribute{Computed: true, Description: "Custom property 20 value for the endpoint."},
						"custom_property_21": schema.StringAttribute{Computed: true, Description: "Custom property 21 value for the endpoint."},
						"custom_property_22": schema.StringAttribute{Computed: true, Description: "Custom property 22 value for the endpoint."},
						"custom_property_23": schema.StringAttribute{Computed: true, Description: "Custom property 23 value for the endpoint."},
						"custom_property_24": schema.StringAttribute{Computed: true, Description: "Custom property 24 value for the endpoint."},
						"custom_property_25": schema.StringAttribute{Computed: true, Description: "Custom property 25 value for the endpoint."},
						"custom_property_26": schema.StringAttribute{Computed: true, Description: "Custom property 26 value for the endpoint."},
						"custom_property_27": schema.StringAttribute{Computed: true, Description: "Custom property 27 value for the endpoint."},
						"custom_property_28": schema.StringAttribute{Computed: true, Description: "Custom property 28 value for the endpoint."},
						"custom_property_29": schema.StringAttribute{Computed: true, Description: "Custom property 29 value for the endpoint."},
						"custom_property_30": schema.StringAttribute{Computed: true, Description: "Custom property 30 value for the endpoint."},
						"custom_property_31": schema.StringAttribute{Computed: true, Description: "Custom property 31 value for the endpoint."},
						"custom_property_32": schema.StringAttribute{Computed: true, Description: "Custom property 32 value for the endpoint."},
						"custom_property_33": schema.StringAttribute{Computed: true, Description: "Custom property 33 value for the endpoint."},
						"custom_property_34": schema.StringAttribute{Computed: true, Description: "Custom property 34 value for the endpoint."},
						"custom_property_35": schema.StringAttribute{Computed: true, Description: "Custom property 35 value for the endpoint."},
						"custom_property_36": schema.StringAttribute{Computed: true, Description: "Custom property 36 value for the endpoint."},
						"custom_property_37": schema.StringAttribute{Computed: true, Description: "Custom property 37 value for the endpoint."},
						"custom_property_38": schema.StringAttribute{Computed: true, Description: "Custom property 38 value for the endpoint."},
						"custom_property_39": schema.StringAttribute{Computed: true, Description: "Custom property 39 value for the endpoint."},
						"custom_property_40": schema.StringAttribute{Computed: true, Description: "Custom property 40 value for the endpoint."},
						"custom_property_41": schema.StringAttribute{Computed: true, Description: "Custom property 41 value for the endpoint."},
						"custom_property_42": schema.StringAttribute{Computed: true, Description: "Custom property 42 value for the endpoint."},
						"custom_property_43": schema.StringAttribute{Computed: true, Description: "Custom property 43 value for the endpoint."},
						"custom_property_44": schema.StringAttribute{Computed: true, Description: "Custom property 44 value for the endpoint."},
						"custom_property_45": schema.StringAttribute{Computed: true, Description: "Custom property 45 value for the endpoint."},

						"account_custom_property_1_label":  schema.StringAttribute{Computed: true, Description: "Label for account custom property 1."},
						"account_custom_property_2_label":  schema.StringAttribute{Computed: true, Description: "Label for account custom property 2."},
						"account_custom_property_3_label":  schema.StringAttribute{Computed: true, Description: "Label for account custom property 3."},
						"account_custom_property_4_label":  schema.StringAttribute{Computed: true, Description: "Label for account custom property 4."},
						"account_custom_property_5_label":  schema.StringAttribute{Computed: true, Description: "Label for account custom property 5."},
						"account_custom_property_6_label":  schema.StringAttribute{Computed: true, Description: "Label for account custom property 6."},
						"account_custom_property_7_label":  schema.StringAttribute{Computed: true, Description: "Label for account custom property 7."},
						"account_custom_property_8_label":  schema.StringAttribute{Computed: true, Description: "Label for account custom property 8."},
						"account_custom_property_9_label":  schema.StringAttribute{Computed: true, Description: "Label for account custom property 9."},
						"account_custom_property_10_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 10."},
						"account_custom_property_11_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 11."},
						"account_custom_property_12_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 12."},
						"account_custom_property_13_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 13."},
						"account_custom_property_14_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 14."},
						"account_custom_property_15_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 15."},
						"account_custom_property_16_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 16."},
						"account_custom_property_17_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 17."},
						"account_custom_property_18_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 18."},
						"account_custom_property_19_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 19."},
						"account_custom_property_20_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 20."},
						"account_custom_property_21_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 21."},
						"account_custom_property_22_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 22."},
						"account_custom_property_23_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 23."},
						"account_custom_property_24_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 24."},
						"account_custom_property_25_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 25."},
						"account_custom_property_26_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 26."},
						"account_custom_property_27_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 27."},
						"account_custom_property_28_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 28."},
						"account_custom_property_29_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 29."},
						"account_custom_property_30_label": schema.StringAttribute{Computed: true, Description: "Label for account custom property 30."},

						"custom_property_31_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 31."},
						"custom_property_32_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 32."},
						"custom_property_33_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 33."},
						"custom_property_34_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 34."},
						"custom_property_35_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 35."},
						"custom_property_36_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 36."},
						"custom_property_37_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 37."},
						"custom_property_38_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 38."},
						"custom_property_39_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 39."},
						"custom_property_40_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 40."},
						"custom_property_41_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 41."},
						"custom_property_42_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 42."},
						"custom_property_43_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 43."},
						"custom_property_44_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 44."},
						"custom_property_45_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 45."},
						"custom_property_46_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 46."},
						"custom_property_47_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 47."},
						"custom_property_48_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 48."},
						"custom_property_49_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 49."},
						"custom_property_50_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 50."},
						"custom_property_51_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 51."},
						"custom_property_52_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 52."},
						"custom_property_53_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 53."},
						"custom_property_54_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 54."},
						"custom_property_55_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 55."},
						"custom_property_56_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 56."},
						"custom_property_57_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 57."},
						"custom_property_58_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 58."},
						"custom_property_59_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 59."},
						"custom_property_60_label": schema.StringAttribute{Computed: true, Description: "Label for custom property 60."},

					},
				},
			},
		},
	}
}


func (d *EndpointsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *EndpointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EndpointsDataSourceModel

	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)

	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(d.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+d.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)

	areq := openapi.NewGetEndpointsRequest()

	if !state.EndpointName.IsNull() && state.EndpointName.ValueString() != "" {
		endpointName := state.EndpointName.ValueString()
		areq.SetEndpointname(endpointName)
	}

	if !state.ConnectionType.IsNull() && state.ConnectionType.ValueString() != "" {
		connectionType := state.ConnectionType.ValueString()
		areq.SetConnectionType(connectionType)
	}

	if !state.Displayname.IsNull() && state.Displayname.ValueString() != "" {
		displayName := state.Displayname.ValueString()
		areq.SetDisplayName(displayName)
	}

	if !state.Owner.IsNull() && state.Owner.ValueString() != "" {
		owner := state.Owner.ValueString()
		areq.SetOwner(owner)
	}

	if !state.Max.IsNull() && state.Max.ValueString() != "" {
		max := state.Max.ValueString()
		areq.SetMax(max)
	}

	if !state.FilterCriteria.IsNull() {
		var filterMap map[string]string
		diags := state.FilterCriteria.ElementsAs(ctx, &filterMap, true)

		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}

		filterCriteria := make(map[string]interface{}, len(filterMap))
		for k, v := range filterMap {
			filterCriteria[k] = v
		}

		areq.SetFilterCriteria(filterCriteria)
	}

	apiReq := apiClient.EndpointsAPI.GetEndpoints(ctx).GetEndpointsRequest(*areq)

	endpointsResponse, httpResp, err := apiReq.Execute()
	if err != nil {
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Message = types.StringValue(*endpointsResponse.Message)
	state.DisplayCount = types.Int64Value(int64(*endpointsResponse.DisplayCount))
	state.ErrorCode = types.StringValue(*endpointsResponse.ErrorCode)
	state.TotalCount = types.Int64Value(int64(*endpointsResponse.TotalCount))

	if endpointsResponse.Endpoints != nil {
		for _, item := range endpointsResponse.Endpoints {
			endpointState := Endpoint{
				Id:                                 util.SafeString(item.Id),
				Description:                        util.SafeString(item.Description),
				StatusForUniqueAccount:             util.SafeString(item.StatusForUniqueAccount),
				Requestowner:                       util.SafeString(item.Requestowner),
				Requestable:                        util.SafeString(item.Requestable),
				PrimaryAccountType:                 util.SafeString(item.PrimaryAccountType),
				AccountTypeNoPasswordChange:        util.SafeString(item.AccountTypeNoPasswordChange),
				ServiceAccountNameRule:             util.SafeString(item.ServiceAccountNameRule),
				AccountNameValidatorRegex:          util.SafeString(item.AccountNameValidatorRegex),
				AllowChangePasswordSqlquery:        util.SafeString(item.AllowChangePasswordSqlquery),
				ParentAccountPattern:               util.SafeString(item.ParentAccountPattern),
				OwnerType:                          util.SafeString(item.OwnerType),
				Securitysystem:                     util.SafeString(item.Securitysystem),
				Endpointname:                       util.SafeString(item.Endpointname),
				UpdatedBy:                          util.SafeString(item.UpdatedBy),
				Accessquery:                        util.SafeString(item.Accessquery),
				Status:                             util.SafeString(item.Status),
				DisplayName:                        util.SafeString(item.DisplayName),
				UpdateDate:                         util.SafeString(item.UpdateDate),
				AllowRemoveAllRoleOnRequest:        util.SafeString(item.AllowRemoveAllRoleOnRequest),
				RoleTypeAsJson:                     util.SafeString(item.RoleTypeAsJson),
				EntsWithNewAccount:                 util.SafeString(item.EntsWithNewAccount),
				ConnectionconfigAsJson:             util.SafeString(item.ConnectionconfigAsJson),
				Connectionconfig:                   util.SafeString(item.Connectionconfig),
				AccountNameRule:                    util.SafeString(item.AccountNameRule),
				ChangePasswordAccessQuery:          util.SafeString(item.ChangePasswordAccessQuery),
				Disableaccountrequest:              util.SafeString(item.Disableaccountrequest),
				PluginConfigs:                      util.SafeString(item.PluginConfigs),
				DisableaccountrequestServiceAccount: util.SafeString(item.DisableaccountrequestServiceAccount),
				Requestableapplication:             util.SafeString(item.Requestableapplication),
				CreatedFrom:                        util.SafeString(item.CreatedFrom),
				CreatedBy:                          util.SafeString(item.CreatedBy),
				CreateDate:                         util.SafeString(item.CreateDate),
				ParentEndpoint:                     util.SafeString(item.ParentEndpoint),
				BaseLineConfig:                     util.SafeString(item.BaseLineConfig),
				Requestownertype:                   util.SafeString(item.Requestownertype),
				CreateEntTaskforRemoveAcc:          util.SafeString(item.CreateEntTaskforRemoveAcc),
				EnableCopyAccess:                   util.SafeString(item.EnableCopyAccess),
				AccountTypeNoDeprovision:           util.SafeString(item.AccountTypeNoDeprovision),
				EndpointConfig:                     util.SafeString(item.EndpointConfig),
				Taskemailtemplates:                 util.SafeString(item.Taskemailtemplates),
				Ownerkey:                           util.SafeString(item.Ownerkey),
				ServiceAccountAccessQuery:          util.SafeString(item.ServiceAccountAccessQuery),
				UserAccountCorrelationRule:         util.SafeString(item.UserAccountCorrelationRule),
				StatusConfig:                       util.SafeString(item.StatusConfig),

				CustomProperty1:  util.SafeString(item.CustomProperty1),
				CustomProperty2:  util.SafeString(item.CustomProperty2),
				CustomProperty3:  util.SafeString(item.CustomProperty3),
				CustomProperty4:  util.SafeString(item.CustomProperty4),
				CustomProperty5:  util.SafeString(item.CustomProperty5),
				CustomProperty6:  util.SafeString(item.CustomProperty6),
				CustomProperty7:  util.SafeString(item.CustomProperty7),
				CustomProperty8:  util.SafeString(item.CustomProperty8),
				CustomProperty9:  util.SafeString(item.CustomProperty9),
				CustomProperty10: util.SafeString(item.CustomProperty10),
				CustomProperty11: util.SafeString(item.CustomProperty11),
				CustomProperty12: util.SafeString(item.CustomProperty12),
				CustomProperty13: util.SafeString(item.CustomProperty13),
				CustomProperty14: util.SafeString(item.CustomProperty14),
				CustomProperty15: util.SafeString(item.CustomProperty15),
				CustomProperty16: util.SafeString(item.CustomProperty16),
				CustomProperty17: util.SafeString(item.CustomProperty17),
				CustomProperty18: util.SafeString(item.CustomProperty18),
				CustomProperty19: util.SafeString(item.CustomProperty19),
				CustomProperty20: util.SafeString(item.CustomProperty20),
				CustomProperty21: util.SafeString(item.CustomProperty21),
				CustomProperty22: util.SafeString(item.CustomProperty22),
				CustomProperty23: util.SafeString(item.CustomProperty23),
				CustomProperty24: util.SafeString(item.CustomProperty24),
				CustomProperty25: util.SafeString(item.CustomProperty25),
				CustomProperty26: util.SafeString(item.CustomProperty26),
				CustomProperty27: util.SafeString(item.CustomProperty27),
				CustomProperty28: util.SafeString(item.CustomProperty28),
				CustomProperty29: util.SafeString(item.CustomProperty29),
				CustomProperty30: util.SafeString(item.CustomProperty30),
				CustomProperty31: util.SafeString(item.CustomProperty31),
				CustomProperty32: util.SafeString(item.CustomProperty32),
				CustomProperty33: util.SafeString(item.CustomProperty33),
				CustomProperty34: util.SafeString(item.CustomProperty34),
				CustomProperty35: util.SafeString(item.CustomProperty35),
				CustomProperty36: util.SafeString(item.CustomProperty36),
				CustomProperty37: util.SafeString(item.CustomProperty37),
				CustomProperty38: util.SafeString(item.CustomProperty38),
				CustomProperty39: util.SafeString(item.CustomProperty39),
				CustomProperty40: util.SafeString(item.CustomProperty40),
				CustomProperty41: util.SafeString(item.CustomProperty41),
				CustomProperty42: util.SafeString(item.CustomProperty42),
				CustomProperty43: util.SafeString(item.CustomProperty43),
				CustomProperty44: util.SafeString(item.CustomProperty44),
				CustomProperty45: util.SafeString(item.CustomProperty45),

				AccountCustomProperty1Label: util.SafeString(item.AccountCustomProperty1Label),
				AccountCustomProperty2Label:  util.SafeString(item.AccountCustomProperty2Label),
				AccountCustomProperty3Label:  util.SafeString(item.AccountCustomProperty3Label),
				AccountCustomProperty4Label:  util.SafeString(item.AccountCustomProperty4Label),
				AccountCustomProperty5Label:  util.SafeString(item.AccountCustomProperty5Label),
				AccountCustomProperty6Label:  util.SafeString(item.AccountCustomProperty6Label),
				AccountCustomProperty7Label:  util.SafeString(item.AccountCustomProperty7Label),
				AccountCustomProperty8Label:  util.SafeString(item.AccountCustomProperty8Label),
				AccountCustomProperty9Label:  util.SafeString(item.AccountCustomProperty9Label),
				AccountCustomProperty10Label: util.SafeString(item.AccountCustomProperty10Label),
				AccountCustomProperty11Label: util.SafeString(item.AccountCustomProperty11Label),
				AccountCustomProperty12Label: util.SafeString(item.AccountCustomProperty12Label),
				AccountCustomProperty13Label: util.SafeString(item.AccountCustomProperty13Label),
				AccountCustomProperty14Label: util.SafeString(item.AccountCustomProperty14Label),
				AccountCustomProperty15Label: util.SafeString(item.AccountCustomProperty15Label),
				AccountCustomProperty16Label: util.SafeString(item.AccountCustomProperty16Label),
				AccountCustomProperty17Label: util.SafeString(item.AccountCustomProperty17Label),
				AccountCustomProperty18Label: util.SafeString(item.AccountCustomProperty18Label),
				AccountCustomProperty19Label: util.SafeString(item.AccountCustomProperty19Label),
				AccountCustomProperty20Label: util.SafeString(item.AccountCustomProperty20Label),
				AccountCustomProperty21Label: util.SafeString(item.AccountCustomProperty21Label),
				AccountCustomProperty22Label: util.SafeString(item.AccountCustomProperty22Label),
				AccountCustomProperty23Label: util.SafeString(item.AccountCustomProperty23Label),
				AccountCustomProperty24Label: util.SafeString(item.AccountCustomProperty24Label),
				AccountCustomProperty25Label: util.SafeString(item.AccountCustomProperty25Label),
				AccountCustomProperty26Label: util.SafeString(item.AccountCustomProperty26Label),
				AccountCustomProperty27Label: util.SafeString(item.AccountCustomProperty27Label),
				AccountCustomProperty28Label: util.SafeString(item.AccountCustomProperty28Label),
				AccountCustomProperty29Label: util.SafeString(item.AccountCustomProperty29Label),
				AccountCustomProperty30Label: util.SafeString(item.AccountCustomProperty30Label),

				CustomProperty31Label: util.SafeString(item.Customproperty31Label),
				CustomProperty32Label: util.SafeString(item.Customproperty32Label),
				CustomProperty33Label: util.SafeString(item.Customproperty33Label),
				CustomProperty34Label: util.SafeString(item.Customproperty34Label),
				CustomProperty35Label: util.SafeString(item.Customproperty35Label),
				CustomProperty36Label: util.SafeString(item.Customproperty36Label),
				CustomProperty37Label: util.SafeString(item.Customproperty37Label),
				CustomProperty38Label: util.SafeString(item.Customproperty38Label),
				CustomProperty39Label: util.SafeString(item.Customproperty39Label),
				CustomProperty40Label: util.SafeString(item.Customproperty40Label),
				CustomProperty41Label: util.SafeString(item.Customproperty41Label),
				CustomProperty42Label: util.SafeString(item.Customproperty42Label),
				CustomProperty43Label: util.SafeString(item.Customproperty43Label),
				CustomProperty44Label: util.SafeString(item.Customproperty44Label),
				CustomProperty45Label: util.SafeString(item.Customproperty45Label),
				CustomProperty46Label: util.SafeString(item.Customproperty46Label),
				CustomProperty47Label: util.SafeString(item.Customproperty47Label),
				CustomProperty48Label: util.SafeString(item.Customproperty48Label),
				CustomProperty49Label: util.SafeString(item.Customproperty49Label),
				CustomProperty50Label: util.SafeString(item.Customproperty50Label),
				CustomProperty51Label: util.SafeString(item.Customproperty51Label),
				CustomProperty52Label: util.SafeString(item.Customproperty52Label),
				CustomProperty53Label: util.SafeString(item.Customproperty53Label),
				CustomProperty54Label: util.SafeString(item.Customproperty54Label),
				CustomProperty55Label: util.SafeString(item.Customproperty55Label),
				CustomProperty56Label: util.SafeString(item.Customproperty56Label),
				CustomProperty57Label: util.SafeString(item.Customproperty57Label),
				CustomProperty58Label: util.SafeString(item.Customproperty58Label),
				CustomProperty59Label: util.SafeString(item.Customproperty59Label),
				CustomProperty60Label: util.SafeString(item.Customproperty60Label),

			}
			state.Results = append(state.Results, endpointState)
		}
	}

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)

	if resp.Diagnostics.HasError() {
		return
	}
}