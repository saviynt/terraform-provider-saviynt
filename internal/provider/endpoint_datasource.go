package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/endpoints"

	s "github.com/saviynt/saviynt-api-go-client"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EndpointsDataSource struct {
	client *s.Client
	token  string
}

var _ datasource.DataSource = &EndpointsDataSource{}

func NewEndpointsDataSource(client *s.Client, token string) datasource.DataSource {
	return &EndpointsDataSource{
		client: client,
		token:  token,
	}
}

type EndpointsDataSourceModel struct {
	ID             types.String `tfsdk:"id"`
	Results        types.List   `tfsdk:"results"`
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
	resp.TypeName = "saviynt_datasource_endpoints"
}
func (d *EndpointsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Datasource to retrieve all endpoints",
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

var endpointObjectType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"id1":                                     types.StringType,
		"endpoint_name":                           types.StringType,
		"display_name":                            types.StringType,
		"security_system":                         types.StringType,
		"access_query":                            types.StringType,
		"enable_copy_access":                      types.StringType,
		"user_account_correlation_rule":           types.StringType,
		"connectionconfig":                        types.StringType,
		"updated_by":                              types.StringType,
		"status":                                  types.StringType,
		"update_date":                             types.StringType,
		"allow_remove_all_role_on_request":        types.StringType,
		"role_type_as_json":                       types.StringType,
		"ents_with_new_account":                   types.StringType,
		"connection_config_as_json":               types.StringType,
		"account_name_rule":                       types.StringType,
		"change_password_access_query":            types.StringType,
		"plugin_configs":                          types.StringType,
		"disable_account_request_service_account": types.StringType,
		"requestable_application":                 types.StringType,
		"created_from":                            types.StringType,
		"created_by":                              types.StringType,
		"create_date":                             types.StringType,
		"parent_endpoint":                         types.StringType,
		"base_line_config":                        types.StringType,
		"create_ent_task_for_remove_acc":          types.StringType,
		"endpoint_config":                         types.StringType,
		"task_email_templates":                    types.StringType,
		"service_account_access_query":            types.StringType,
		"status_config":                           types.StringType,
		"disable_account_request":                 types.StringType,

		"custom_property_1":  types.StringType,
		"custom_property_2":  types.StringType,
		"custom_property_3":  types.StringType,
		"custom_property_4":  types.StringType,
		"custom_property_5":  types.StringType,
		"custom_property_6":  types.StringType,
		"custom_property_7":  types.StringType,
		"custom_property_8":  types.StringType,
		"custom_property_9":  types.StringType,
		"custom_property_10": types.StringType,
		"custom_property_11": types.StringType,
		"custom_property_12": types.StringType,
		"custom_property_13": types.StringType,
		"custom_property_14": types.StringType,
		"custom_property_15": types.StringType,
		"custom_property_16": types.StringType,
		"custom_property_17": types.StringType,
		"custom_property_18": types.StringType,
		"custom_property_19": types.StringType,
		"custom_property_20": types.StringType,
		"custom_property_21": types.StringType,
		"custom_property_22": types.StringType,
		"custom_property_23": types.StringType,
		"custom_property_24": types.StringType,
		"custom_property_25": types.StringType,
		"custom_property_26": types.StringType,
		"custom_property_27": types.StringType,
		"custom_property_28": types.StringType,
		"custom_property_29": types.StringType,
		"custom_property_30": types.StringType,
		"custom_property_31": types.StringType,
		"custom_property_32": types.StringType,
		"custom_property_33": types.StringType,
		"custom_property_34": types.StringType,
		"custom_property_35": types.StringType,
		"custom_property_36": types.StringType,
		"custom_property_37": types.StringType,
		"custom_property_38": types.StringType,
		"custom_property_39": types.StringType,
		"custom_property_40": types.StringType,
		"custom_property_41": types.StringType,
		"custom_property_42": types.StringType,
		"custom_property_43": types.StringType,
		"custom_property_44": types.StringType,
		"custom_property_45": types.StringType,
	},
}

func (r Endpoint) ToMap() map[string]attr.Value {
	return map[string]attr.Value{
		"id1":                                     r.Id1,
		"endpoint_name":                           r.EndpointName,
		"display_name":                            r.DisplayName,
		"security_system":                         r.SecuritySystem,
		"access_query":                            r.AccessQuery,
		"enable_copy_access":                      r.EnableCopyAccess,
		"updated_by":                              r.UpdatedBy,
		"status":                                  r.Status,
		"update_date":                             r.UpdateDate,
		"allow_remove_all_role_on_request":        r.AllowRemoveAllRoleOnRequest,
		"role_type_as_json":                       r.RoleTypeAsJson,
		"ents_with_new_account":                   r.EntsWithNewAccount,
		"connection_config_as_json":               r.ConnectionConfigAsJson,
		"connectionconfig":                        r.ConnectionConfig,
		"account_name_rule":                       r.AccountNameRule,
		"change_password_access_query":            r.ChangePasswordAccessQuery,
		"service_account_access_query":            r.ServiceAccountAccessQuery,
		"create_ent_task_for_remove_acc":          r.CreateEntTaskForRemoveAcc,
		"user_account_correlation_rule":           r.UserAccountCorrelationRule,
		"disable_account_request":                 r.DisableAccountRequest,
		"plugin_configs":                          r.PluginConfigs,
		"disable_account_request_service_account": r.DisableAccountRequestServiceAccount,
		"requestable_application":                 r.RequestableApplication,
		"created_from":                            r.CreatedFrom,
		"created_by":                              r.CreatedBy,
		"create_date":                             r.CreateDate,
		"parent_endpoint":                         r.ParentEndpoint,
		"base_line_config":                        r.BaseLineConfig,
		"endpoint_config":                         r.EndpointConfig,
		"task_email_templates":                    r.TaskEmailTemplates,
		"status_config":                           r.StatusConfig,

		"custom_property_1":  r.CustomProperty1,
		"custom_property_2":  r.CustomProperty2,
		"custom_property_3":  r.CustomProperty3,
		"custom_property_4":  r.CustomProperty4,
		"custom_property_5":  r.CustomProperty5,
		"custom_property_6":  r.CustomProperty6,
		"custom_property_7":  r.CustomProperty7,
		"custom_property_8":  r.CustomProperty8,
		"custom_property_9":  r.CustomProperty9,
		"custom_property_10": r.CustomProperty10,
		"custom_property_11": r.CustomProperty11,
		"custom_property_12": r.CustomProperty12,
		"custom_property_13": r.CustomProperty13,
		"custom_property_14": r.CustomProperty14,
		"custom_property_15": r.CustomProperty15,
		"custom_property_16": r.CustomProperty16,
		"custom_property_17": r.CustomProperty17,
		"custom_property_18": r.CustomProperty18,
		"custom_property_19": r.CustomProperty19,
		"custom_property_20": r.CustomProperty20,
		"custom_property_21": r.CustomProperty21,
		"custom_property_22": r.CustomProperty22,
		"custom_property_23": r.CustomProperty23,
		"custom_property_24": r.CustomProperty24,
		"custom_property_25": r.CustomProperty25,
		"custom_property_26": r.CustomProperty26,
		"custom_property_27": r.CustomProperty27,
		"custom_property_28": r.CustomProperty28,
		"custom_property_29": r.CustomProperty29,
		"custom_property_30": r.CustomProperty30,
		"custom_property_31": r.CustomProperty31,
		"custom_property_32": r.CustomProperty32,
		"custom_property_33": r.CustomProperty33,
		"custom_property_34": r.CustomProperty34,
		"custom_property_35": r.CustomProperty35,
		"custom_property_36": r.CustomProperty36,
		"custom_property_37": r.CustomProperty37,
		"custom_property_38": r.CustomProperty38,
		"custom_property_39": r.CustomProperty39,
		"custom_property_40": r.CustomProperty40,
		"custom_property_41": r.CustomProperty41,
		"custom_property_42": r.CustomProperty42,
		"custom_property_43": r.CustomProperty43,
		"custom_property_44": r.CustomProperty44,
		"custom_property_45": r.CustomProperty45,
	}
}

func safeString(s *string) types.String {
	if s == nil {
		return types.StringValue("")
	}
	return types.StringValue(*s)
}

func (d *EndpointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EndpointsDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := d.client.APIBaseURL()
	cfg.Host = strings.TrimPrefix(apiBaseURL, "https://")
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+d.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)

	areq := openapi.NewGetEndpointsRequest()

	if !state.EndpointName.IsNull() && state.EndpointName.ValueString() != "" {
		endpointName := state.EndpointName.ValueString()
		areq.Endpointname = &endpointName
	}

	if !state.ConnectionType.IsNull() && state.ConnectionType.ValueString() != "" {
		connectionType := state.ConnectionType.ValueString()
		areq.ConnectionType = &connectionType
	}

	if !state.Displayname.IsNull() && state.Displayname.ValueString() != "" {
		displayName := state.Displayname.ValueString()
		areq.DisplayName = &displayName
	}

	if !state.Owner.IsNull() && state.Owner.ValueString() != "" {
		owner := state.Owner.ValueString()
		areq.Owner = &owner
	}

	if !state.Max.IsNull() && state.Max.ValueString() != "" {
		max := state.Max.ValueString()
		areq.Max = &max
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

		apiReq := apiClient.EndpointsAPI.GetEndpoints(ctx).GetEndpointsRequest(*areq)

		endpointsResponse, httpResp, err := apiReq.Execute()
		jsonBytes, err := json.MarshalIndent(endpointsResponse, "", "  ")
		if err != nil {
			fmt.Printf("Error marshalling apiResp: %v\n", err)
			return
		}
		fmt.Println("Marshalled API Response:")
		fmt.Println(string(jsonBytes))
		if err != nil {
			fmt.Printf("[ERROR] API Call Failed: %v\n", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
			return
		}
		fmt.Printf("[DEBUG] HTTP Status Code: %d\n", httpResp.StatusCode)

		state.Message = types.StringValue(*endpointsResponse.Message)
		state.DisplayCount = types.Int64Value(int64(*endpointsResponse.DisplayCount))
		state.ErrorCode = types.StringValue(*endpointsResponse.ErrorCode)
		state.TotalCount = types.Int64Value(int64(*endpointsResponse.TotalCount))

		var results []Endpoint

		if endpointsResponse.Endpoints != nil {
			for _, item := range endpointsResponse.Endpoints {
				results = append(results, Endpoint{
					Id1:                                 safeString(item.Id),
					EndpointName:                        safeString(item.Endpointname),
					DisplayName:                         safeString(item.DisplayName),
					SecuritySystem:                      safeString(item.Securitysystem),
					AccessQuery:                         safeString(item.Accessquery),
					EnableCopyAccess:                    safeString(item.EnableCopyAccess),
					UpdatedBy:                           safeString(item.UpdatedBy),
					Status:                              safeString(item.Status),
					UpdateDate:                          safeString(item.UpdateDate),
					AllowRemoveAllRoleOnRequest:         safeString(item.AllowRemoveAllRoleOnRequest),
					RoleTypeAsJson:                      safeString(item.RoleTypeAsJson),
					EntsWithNewAccount:                  safeString(item.EntsWithNewAccount),
					ConnectionConfigAsJson:              safeString(item.ConnectionconfigAsJson),
					ConnectionConfig:                    safeString(item.Connectionconfig),
					AccountNameRule:                     safeString(item.AccountNameRule),
					ChangePasswordAccessQuery:           safeString(item.ChangePasswordAccessQuery),
					ServiceAccountAccessQuery:           safeString(item.ServiceAccountAccessQuery),
					CreateEntTaskForRemoveAcc:           safeString(item.CreateEntTaskforRemoveAcc),
					UserAccountCorrelationRule:          safeString(item.UserAccountCorrelationRule),
					DisableAccountRequest:               safeString(item.Disableaccountrequest),
					PluginConfigs:                       safeString(item.PluginConfigs),
					DisableAccountRequestServiceAccount: safeString(item.DisableaccountrequestServiceAccount),
					RequestableApplication:              safeString(item.Requestableapplication),
					CreatedFrom:                         safeString(item.CreatedFrom),
					CreatedBy:                           safeString(item.CreatedBy),
					CreateDate:                          safeString(item.CreateDate),
					ParentEndpoint:                      safeString(item.ParentEndpoint),
					BaseLineConfig:                      safeString(item.BaseLineConfig),
					EndpointConfig:                      safeString(item.EndpointConfig),
					TaskEmailTemplates:                  safeString(item.Taskemailtemplates),
					StatusConfig:                        safeString(item.StatusConfig),

					CustomProperty1:  safeString(item.CustomProperty1),
					CustomProperty2:  safeString(item.CustomProperty2),
					CustomProperty3:  safeString(item.CustomProperty3),
					CustomProperty4:  safeString(item.CustomProperty4),
					CustomProperty5:  safeString(item.CustomProperty5),
					CustomProperty6:  safeString(item.CustomProperty6),
					CustomProperty7:  safeString(item.CustomProperty7),
					CustomProperty8:  safeString(item.CustomProperty8),
					CustomProperty9:  safeString(item.CustomProperty9),
					CustomProperty10: safeString(item.CustomProperty10),
					CustomProperty11: safeString(item.CustomProperty11),
					CustomProperty12: safeString(item.CustomProperty12),
					CustomProperty13: safeString(item.CustomProperty13),
					CustomProperty14: safeString(item.CustomProperty14),
					CustomProperty15: safeString(item.CustomProperty15),
					CustomProperty16: safeString(item.CustomProperty16),
					CustomProperty17: safeString(item.CustomProperty17),
					CustomProperty18: safeString(item.CustomProperty18),
					CustomProperty19: safeString(item.CustomProperty19),
					CustomProperty20: safeString(item.CustomProperty20),
					CustomProperty21: safeString(item.CustomProperty21),
					CustomProperty22: safeString(item.CustomProperty22),
					CustomProperty23: safeString(item.CustomProperty23),
					CustomProperty24: safeString(item.CustomProperty24),
					CustomProperty25: safeString(item.CustomProperty25),
					CustomProperty26: safeString(item.CustomProperty26),
					CustomProperty27: safeString(item.CustomProperty27),
					CustomProperty28: safeString(item.CustomProperty28),
					CustomProperty29: safeString(item.CustomProperty29),
					CustomProperty30: safeString(item.CustomProperty30),
					CustomProperty31: safeString(item.CustomProperty31),
					CustomProperty32: safeString(item.CustomProperty32),
					CustomProperty33: safeString(item.CustomProperty33),
					CustomProperty34: safeString(item.CustomProperty34),
					CustomProperty35: safeString(item.CustomProperty35),
					CustomProperty36: safeString(item.CustomProperty36),
					CustomProperty37: safeString(item.CustomProperty37),
					CustomProperty38: safeString(item.CustomProperty38),
					CustomProperty39: safeString(item.CustomProperty39),
					CustomProperty40: safeString(item.CustomProperty40),
					CustomProperty41: safeString(item.CustomProperty41),
					CustomProperty42: safeString(item.CustomProperty42),
					CustomProperty43: safeString(item.CustomProperty43),
					CustomProperty44: safeString(item.CustomProperty44),
					CustomProperty45: safeString(item.CustomProperty45),
				})
			}
		}

		fmt.Println("Debug: Printing Results Before Assigning to State:")
		for _, item := range results {
			fmt.Printf("Item: %+v\n", item.SecuritySystem)
		}

		var listValues []attr.Value
		for _, item := range results {
			objVal, objDiags := types.ObjectValue(endpointObjectType.AttrTypes, item.ToMap())
			if objDiags.HasError() {
				fmt.Println("[ERROR] ObjectValue conversion failed:", objDiags)
				resp.Diagnostics.Append(objDiags...)
				continue
			}
			listValues = append(listValues, objVal)
		}

		listValue, diags := types.ListValue(endpointObjectType, listValues)
		if diags.HasError() {
			fmt.Println("[ERROR] ListValue conversion failed:", diags)
			resp.Diagnostics.Append(diags...)
			return
		}
		state.Results = listValue
		fmt.Println("Diagnostics after ListValue:", diags)
		fmt.Println("Shaleen Shukla", state.Results)
		resp.Diagnostics.Append(diags...)
		err1 := resp.State.Set(ctx, state)
		if err1 != nil {
			fmt.Println("Error setting state:", err1)
		}
		fmt.Printf("Final state before set: %+v\n", state)
	}
}
