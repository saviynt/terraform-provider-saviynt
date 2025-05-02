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
	ID             types.String `tfsdk:"id"`
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
	Id1                         types.String `tfsdk:"id1"`
	EndpointName                types.String `tfsdk:"endpoint_name"`
	DisplayName                 types.String `tfsdk:"display_name"`
	SecuritySystem              types.String `tfsdk:"security_system"`
	AccessQuery                 types.String `tfsdk:"access_query"`
	EnableCopyAccess            types.String `tfsdk:"enable_copy_access"`
	UpdatedBy                   types.String `tfsdk:"updated_by"`
	Status                      types.String `tfsdk:"status"`
	UpdateDate                  types.String `tfsdk:"update_date"`
	AllowRemoveAllRoleOnRequest types.String `tfsdk:"allow_remove_all_role_on_request"`
	RoleTypeAsJson              types.String `tfsdk:"role_type_as_json"`
	EntsWithNewAccount          types.String `tfsdk:"ents_with_new_account"`
	ConnectionConfigAsJson      types.String `tfsdk:"connection_config_as_json"`
	ConnectionConfig            types.String `tfsdk:"connectionconfig"`
	AccountNameRule             types.String `tfsdk:"account_name_rule"`

	ChangePasswordAccessQuery           types.String `tfsdk:"change_password_access_query"`
	ServiceAccountAccessQuery           types.String `tfsdk:"service_account_access_query"`
	CreateEntTaskForRemoveAcc           types.String `tfsdk:"create_ent_task_for_remove_acc"`
	UserAccountCorrelationRule          types.String `tfsdk:"user_account_correlation_rule"`
	DisableAccountRequest               types.String `tfsdk:"disable_account_request"`
	PluginConfigs                       types.String `tfsdk:"plugin_configs"`
	DisableAccountRequestServiceAccount types.String `tfsdk:"disable_account_request_service_account"`
	RequestableApplication              types.String `tfsdk:"requestable_application"`
	CreatedFrom                         types.String `tfsdk:"created_from"`
	CreatedBy                           types.String `tfsdk:"created_by"`
	CreateDate                          types.String `tfsdk:"create_date"`
	ParentEndpoint                      types.String `tfsdk:"parent_endpoint"`
	BaseLineConfig                      types.String `tfsdk:"base_line_config"`
	EndpointConfig                      types.String `tfsdk:"endpoint_config"`
	TaskEmailTemplates                  types.String `tfsdk:"task_email_templates"`
	StatusConfig                        types.String `tfsdk:"status_config"`

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
}

func (d *EndpointsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_endpoints_datasource"
}
func (d *EndpointsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.EndpointDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
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
						"id1":                              schema.StringAttribute{Computed: true},
						"endpoint_name":                    schema.StringAttribute{Computed: true},
						"display_name":                     schema.StringAttribute{Computed: true},
						"security_system":                  schema.StringAttribute{Computed: true},
						"access_query":                     schema.StringAttribute{Computed: true},
						"enable_copy_access":               schema.StringAttribute{Computed: true},
						"updated_by":                       schema.StringAttribute{Computed: true},
						"status":                           schema.StringAttribute{Computed: true},
						"update_date":                      schema.StringAttribute{Computed: true},
						"allow_remove_all_role_on_request": schema.StringAttribute{Computed: true},
						"role_type_as_json":                schema.StringAttribute{Computed: true},
						"ents_with_new_account":            schema.StringAttribute{Computed: true},
						"connection_config_as_json":        schema.StringAttribute{Computed: true},
						"account_name_rule":                schema.StringAttribute{Computed: true},
						"disable_account_request":          schema.StringAttribute{Computed: true},
						"disable_account_request_service_account": schema.StringAttribute{Computed: true},
						"change_password_access_query":            schema.StringAttribute{Computed: true},
						"service_account_access_query":            schema.StringAttribute{Computed: true},
						"create_ent_task_for_remove_acc":          schema.StringAttribute{Computed: true},
						"user_account_correlation_rule":           schema.StringAttribute{Computed: true},
						"plugin_configs":                          schema.StringAttribute{Computed: true},
						"requestable_application":                 schema.StringAttribute{Computed: true},
						"created_from":                            schema.StringAttribute{Computed: true},
						"created_by":                              schema.StringAttribute{Computed: true},
						"create_date":                             schema.StringAttribute{Computed: true},
						"parent_endpoint":                         schema.StringAttribute{Computed: true},
						"base_line_config":                        schema.StringAttribute{Computed: true},
						"endpoint_config":                         schema.StringAttribute{Computed: true},
						"task_email_templates":                    schema.StringAttribute{Computed: true},
						"status_config":                           schema.StringAttribute{Computed: true},
						"connectionconfig":                        schema.StringAttribute{Computed: true},

						"custom_property_1":  schema.StringAttribute{Computed: true},
						"custom_property_2":  schema.StringAttribute{Computed: true},
						"custom_property_3":  schema.StringAttribute{Computed: true},
						"custom_property_4":  schema.StringAttribute{Computed: true},
						"custom_property_5":  schema.StringAttribute{Computed: true},
						"custom_property_6":  schema.StringAttribute{Computed: true},
						"custom_property_7":  schema.StringAttribute{Computed: true},
						"custom_property_8":  schema.StringAttribute{Computed: true},
						"custom_property_9":  schema.StringAttribute{Computed: true},
						"custom_property_10": schema.StringAttribute{Computed: true},
						"custom_property_11": schema.StringAttribute{Computed: true},
						"custom_property_12": schema.StringAttribute{Computed: true},
						"custom_property_13": schema.StringAttribute{Computed: true},
						"custom_property_14": schema.StringAttribute{Computed: true},
						"custom_property_15": schema.StringAttribute{Computed: true},
						"custom_property_16": schema.StringAttribute{Computed: true},
						"custom_property_17": schema.StringAttribute{Computed: true},
						"custom_property_18": schema.StringAttribute{Computed: true},
						"custom_property_19": schema.StringAttribute{Computed: true},
						"custom_property_20": schema.StringAttribute{Computed: true},
						"custom_property_21": schema.StringAttribute{Computed: true},
						"custom_property_22": schema.StringAttribute{Computed: true},
						"custom_property_23": schema.StringAttribute{Computed: true},
						"custom_property_24": schema.StringAttribute{Computed: true},
						"custom_property_25": schema.StringAttribute{Computed: true},
						"custom_property_26": schema.StringAttribute{Computed: true},
						"custom_property_27": schema.StringAttribute{Computed: true},
						"custom_property_28": schema.StringAttribute{Computed: true},
						"custom_property_29": schema.StringAttribute{Computed: true},
						"custom_property_30": schema.StringAttribute{Computed: true},
						"custom_property_31": schema.StringAttribute{Computed: true},
						"custom_property_32": schema.StringAttribute{Computed: true},
						"custom_property_33": schema.StringAttribute{Computed: true},
						"custom_property_34": schema.StringAttribute{Computed: true},
						"custom_property_35": schema.StringAttribute{Computed: true},
						"custom_property_36": schema.StringAttribute{Computed: true},
						"custom_property_37": schema.StringAttribute{Computed: true},
						"custom_property_38": schema.StringAttribute{Computed: true},
						"custom_property_39": schema.StringAttribute{Computed: true},
						"custom_property_40": schema.StringAttribute{Computed: true},
						"custom_property_41": schema.StringAttribute{Computed: true},
						"custom_property_42": schema.StringAttribute{Computed: true},
						"custom_property_43": schema.StringAttribute{Computed: true},
						"custom_property_44": schema.StringAttribute{Computed: true},
						"custom_property_45": schema.StringAttribute{Computed: true},
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
				Id1:                                 util.SafeString(item.Id),
				EndpointName:                        util.SafeString(item.Endpointname),
				DisplayName:                         util.SafeString(item.DisplayName),
				SecuritySystem:                      util.SafeString(item.Securitysystem),
				AccessQuery:                         util.SafeString(item.Accessquery),
				EnableCopyAccess:                    util.SafeString(item.EnableCopyAccess),
				UpdatedBy:                           util.SafeString(item.UpdatedBy),
				Status:                              util.SafeString(item.Status),
				UpdateDate:                          util.SafeString(item.UpdateDate),
				AllowRemoveAllRoleOnRequest:         util.SafeString(item.AllowRemoveAllRoleOnRequest),
				RoleTypeAsJson:                      util.SafeString(item.RoleTypeAsJson),
				EntsWithNewAccount:                  util.SafeString(item.EntsWithNewAccount),
				ConnectionConfigAsJson:              util.SafeString(item.ConnectionconfigAsJson),
				ConnectionConfig:                    util.SafeString(item.Connectionconfig),
				AccountNameRule:                     util.SafeString(item.AccountNameRule),
				ChangePasswordAccessQuery:           util.SafeString(item.ChangePasswordAccessQuery),
				ServiceAccountAccessQuery:           util.SafeString(item.ServiceAccountAccessQuery),
				CreateEntTaskForRemoveAcc:           util.SafeString(item.CreateEntTaskforRemoveAcc),
				UserAccountCorrelationRule:          util.SafeString(item.UserAccountCorrelationRule),
				DisableAccountRequest:               util.SafeString(item.Disableaccountrequest),
				PluginConfigs:                       util.SafeString(item.PluginConfigs),
				DisableAccountRequestServiceAccount: util.SafeString(item.DisableaccountrequestServiceAccount),
				RequestableApplication:              util.SafeString(item.Requestableapplication),
				CreatedFrom:                         util.SafeString(item.CreatedFrom),
				CreatedBy:                           util.SafeString(item.CreatedBy),
				CreateDate:                          util.SafeString(item.CreateDate),
				ParentEndpoint:                      util.SafeString(item.ParentEndpoint),
				BaseLineConfig:                      util.SafeString(item.BaseLineConfig),
				EndpointConfig:                      util.SafeString(item.EndpointConfig),
				TaskEmailTemplates:                  util.SafeString(item.Taskemailtemplates),
				StatusConfig:                        util.SafeString(item.StatusConfig),

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
