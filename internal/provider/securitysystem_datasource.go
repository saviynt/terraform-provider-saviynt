// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
	openapi "github.com/saviynt/saviynt-api-go-client/securitysystems"
)

// SecuritySystemsDataSource defines the data source
type SecuritySystemsDataSource struct {
	client *s.Client
	token  string
}

// SecuritySystemsDataSourceModel maps the API response and user inputs
type SecuritySystemsDataSourceModel struct {
	Systemname     types.String `tfsdk:"systemname"`
	Max            types.Int64  `tfsdk:"max"`
	Offset         types.Int64  `tfsdk:"offset"`
	Connectionname types.String `tfsdk:"connectionname"`
	ConnectionType types.String `tfsdk:"connection_type"`
	Msg            types.String `tfsdk:"msg"`
	DisplayCount   types.Int64  `tfsdk:"display_count"`
	ErrorCode      types.String `tfsdk:"error_code"`
	TotalCount     types.Int64  `tfsdk:"total_count"`

	Results types.List `tfsdk:"results"`
}

// SecuritySystemDetails represents a single security system details object.
type SecuritySystemDetails struct {
	Systemname                         types.String   `tfsdk:"systemname1"`
	ConnectionType                     types.String   `tfsdk:"connection_type_1"`
	DisplayName                        types.String   `tfsdk:"display_name"`
	Hostname                           types.String   `tfsdk:"hostname"`
	Port                               types.String   `tfsdk:"port"`
	AccessAddWorkflow                  types.String   `tfsdk:"access_add_workflow"`
	AccessRemoveWorkflow               types.String   `tfsdk:"access_remove_workflow"`
	AddServiceAccountWorkflow          types.String   `tfsdk:"add_service_account_workflow"`
	RemoveServiceAccountWorkflow       types.String   `tfsdk:"remove_service_account_workflow"`
	Connectionparameters               types.String   `tfsdk:"connection_parameters"`
	AutomatedProvisioning              types.String   `tfsdk:"automated_provisioning"`
	UseOpenConnector                   types.String   `tfsdk:"use_open_connector"`
	ManageEntity                       types.String   `tfsdk:"manage_entity"`
	PersistentData                     types.String   `tfsdk:"persistent_data"`
	DefaultSystem                      types.String   `tfsdk:"default_system"`
	ReconApplication                   types.String   `tfsdk:"recon_application"`
	InstantProvision                   types.String   `tfsdk:"instant_provision"`
	ProvisioningTries                  types.String   `tfsdk:"provisioning_tries"`
	ProvisioningComments               types.String   `tfsdk:"provisioning_comments"`
	ProposedAccountOwnersWorkflow      types.String   `tfsdk:"proposed_account_owners_workflow"`
	FirefighterIDWorkflow              types.String   `tfsdk:"firefighterid_workflow"`
	FirefighterIDRequestAccessWorkflow types.String   `tfsdk:"firefighterid_request_access_workflow"`
	PolicyRule                         types.String   `tfsdk:"policy_rule"`
	PolicyRuleServiceAccount           types.String   `tfsdk:"policy_rule_service_account"`
	Connectionname                     types.String   `tfsdk:"connectionname1"`
	ProvisioningConnection             types.String   `tfsdk:"provisioning_connection"`
	ServiceDeskConnection              types.String   `tfsdk:"service_desk_connection"`
	ExternalRiskConnectionJson         types.String   `tfsdk:"external_risk_connection_json"`
	Connection                         types.String   `tfsdk:"connection"`
	CreateDate                         types.String   `tfsdk:"create_date"`
	UpdateDate                         types.String   `tfsdk:"update_date"`
	Endpoints                          types.String   `tfsdk:"endpoints"`
	CreatedBy                          types.String   `tfsdk:"created_by"`
	UpdatedBy                          types.String   `tfsdk:"updated_by"`
	Status                             types.String   `tfsdk:"status"`
	CreatedFrom                        types.String   `tfsdk:"created_from"`
	InherentSodReportFields            []types.String `tfsdk:"inherent_sod_report_fields"`
}

// Ensure the implementation satisfies Terraform framework interface
var _ datasource.DataSource = &SecuritySystemsDataSource{}

// NewSecuritySystemsDataSource returns a new instance
func NewSecuritySystemsDataSource() datasource.DataSource {
	return &SecuritySystemsDataSource{}
}

// Metadata defines the data source name
func (d *SecuritySystemsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_security_systems_datasource"
}

// Schema defines the attributes for the data source
func (d *SecuritySystemsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"systemname": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the security systeme.",
			},
			"max": schema.Int64Attribute{
				Optional:    true,
				Description: "Name for the security system that will be displayed in the user interface.",
			},
			"offset": schema.Int64Attribute{
				Optional:    true,
				Description: "Security system for which you want to create an endpoint.",
			},
			"connectionname": schema.StringAttribute{
				Optional:    true,
				Description: "Owner type of the endpoint. It could be User or Usergroup.",
			},
			"connection_type": schema.StringAttribute{
				Optional:    true,
				Description: "Owner of the endpoint. If ownerType is User, specify the username of the owner. If ownerType is Usergroup, sepecify the name of the User group.",
			},
			"msg": schema.StringAttribute{
				Computed:    true,
				Description: "A message indicating the outcome of the operation.",
			},
			"display_count": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of items currently displayed (e.g., on the current page or view).",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
			},
			"total_count": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of items available in the dataset, irrespective of the current display settings.",
			},
			"results": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of security systems retrieved",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"display_name":                          schema.StringAttribute{Computed: true, Description: "Specify a user-friendly display name that is shown on the the user interface."},
						"hostname":                              schema.StringAttribute{Computed: true, Description: "Security system for which you want to create an endpoint."},
						"connection_type_1":                     schema.StringAttribute{Computed: true, Description: "Specify a connection type to view all connections in EIC for the connection type."},
						"systemname1":                           schema.StringAttribute{Computed: true, Description: "Specify the security system name."},
						"access_add_workflow":                   schema.StringAttribute{Computed: true, Description: "Specify the workflow used for approvals for an access request (account, entitlements, role, etc.)."},
						"access_remove_workflow":                schema.StringAttribute{Computed: true, Description: "Workflow used when revoking access from accounts, entitlements, or performing other de-provisioning tasks."},
						"add_service_account_workflow":          schema.StringAttribute{Computed: true, Description: "Workflow for adding a service account."},
						"remove_service_account_workflow":       schema.StringAttribute{Computed: true, Description: "Workflow for removing a service account."},
						"connection_parameters":                 schema.StringAttribute{Computed: true, Description: "Query or parameters to restrict endpoint access to specific users."},
						"automated_provisioning":                schema.StringAttribute{Computed: true, Description: "Enables automated provisioning if set to true."},
						"use_open_connector":                    schema.StringAttribute{Computed: true, Description: "Enables connectivity using open-source connectors such as REST if set to true."},
						"manage_entity":                         schema.StringAttribute{Computed: true, Description: "Indicates if entity management is enabled for the security system."},
						"persistent_data":                       schema.StringAttribute{Computed: true, Description: "Indicates whether persistent data storage is enabled for the security system."},
						"default_system":                        schema.StringAttribute{Computed: true, Description: "Sets this security system as the default system for account searches when set to true."},
						"recon_application":                     schema.StringAttribute{Computed: true, Description: "Enables importing data from endpoints associated with the security system."},
						"instant_provision":                     schema.StringAttribute{Computed: true, Description: "Prevents users from submitting duplicate provisioning requests if set to true."},
						"provisioning_tries":                    schema.StringAttribute{Computed: true, Description: "Number of attempts allowed for provisioning actions."},
						"provisioning_comments":                 schema.StringAttribute{Computed: true, Description: "Comments relevant to provisioning actions."},
						"proposed_account_owners_workflow":      schema.StringAttribute{Computed: true, Description: "Workflow for assigning proposed account owners."},
						"firefighterid_workflow":                schema.StringAttribute{Computed: true, Description: "Workflow for handling firefighter ID requests."},
						"firefighterid_request_access_workflow": schema.StringAttribute{Computed: true, Description: "Workflow for requesting access to firefighter IDs."},
						"policy_rule":                           schema.StringAttribute{Computed: true, Description: "Password policy assigned for the security system."},
						"policy_rule_service_account":           schema.StringAttribute{Computed: true, Description: "Password policy applied to service accounts."},
						"connectionname1":                       schema.StringAttribute{Computed: true, Description: "Name of connection used for reconciling identity objects from third-party applications."},
						"provisioning_connection":               schema.StringAttribute{Computed: true, Description: "Dedicated connection for provisioning and de-provisioning tasks."},
						"service_desk_connection":               schema.StringAttribute{Computed: true, Description: "Connection to service desk or ticketing system integration."},
						"external_risk_connection_json":         schema.StringAttribute{Computed: true, Description: "JSON configuration for external risk connections (e.g., SAP)."},
						"connection":                            schema.StringAttribute{Computed: true, Description: "Primary connection used by the security system."},
						"create_date":                           schema.StringAttribute{Computed: true, Description: "Timestamp indicating when the security system was created."},
						"update_date":                           schema.StringAttribute{Computed: true, Description: "Timestamp indicating the last update to the security system."},
						"endpoints":                             schema.StringAttribute{Computed: true, Description: "Endpoints associated with the security system."},
						"created_by":                            schema.StringAttribute{Computed: true, Description: "Identifier of the user who created the security system."},
						"updated_by":                            schema.StringAttribute{Computed: true, Description: "Identifier of the user who last updated the security system."},
						"status":                                schema.StringAttribute{Computed: true, Description: "Current status of the security system (e.g., enabled, disabled)."},
						"created_from":                          schema.StringAttribute{Computed: true, Description: "Origin or method through which the security system was created."},
						"port":                                  schema.StringAttribute{Computed: true, Description: "Port information or description for the endpoint."},
						"inherent_sod_report_fields": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "List of fields used in filtering Segregation of Duties (SOD) reports.",
						},
					},
				},
			},
		},
	}
}

func (d *SecuritySystemsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = prov.client
	d.token = prov.accessToken
}

// securitySystemObjectType defines the Terraform object type for a single security system.
var securitySystemObjectType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"display_name":                          types.StringType,
		"hostname":                              types.StringType,
		"connection_type_1":                     types.StringType,
		"systemname1":                           types.StringType,
		"access_add_workflow":                   types.StringType,
		"access_remove_workflow":                types.StringType,
		"add_service_account_workflow":          types.StringType,
		"remove_service_account_workflow":       types.StringType,
		"connection_parameters":                 types.StringType,
		"automated_provisioning":                types.StringType,
		"use_open_connector":                    types.StringType,
		"manage_entity":                         types.StringType,
		"persistent_data":                       types.StringType,
		"default_system":                        types.StringType,
		"recon_application":                     types.StringType,
		"instant_provision":                     types.StringType,
		"provisioning_tries":                    types.StringType,
		"provisioning_comments":                 types.StringType,
		"proposed_account_owners_workflow":      types.StringType,
		"firefighterid_workflow":                types.StringType,
		"firefighterid_request_access_workflow": types.StringType,
		"policy_rule":                           types.StringType,
		"policy_rule_service_account":           types.StringType,
		"connectionname1":                       types.StringType,
		"provisioning_connection":               types.StringType,
		"service_desk_connection":               types.StringType,
		"external_risk_connection_json":         types.StringType,
		"connection":                            types.StringType,
		"create_date":                           types.StringType,
		"update_date":                           types.StringType,
		"endpoints":                             types.StringType,
		"created_by":                            types.StringType,
		"updated_by":                            types.StringType,
		"status":                                types.StringType,
		"created_from":                          types.StringType,
		"port":                                  types.StringType,
		"inherent_sod_report_fields":            types.ListType{ElemType: types.StringType},
	},
}

// ToMap converts SecuritySystemDetails to a map[string]attr.Value matching the Terraform object schema.
func (r SecuritySystemDetails) ToMap() (map[string]attr.Value, diag.Diagnostics) {
	inherentSodList, diags := util.SafeList(util.ConvertTypesStringToStrings(r.InherentSodReportFields))
	if diags.HasError() {
		inherentSodList = types.ListNull(types.StringType)
	}

	return map[string]attr.Value{
		"display_name":                          r.DisplayName,
		"hostname":                              r.Hostname,
		"connection_type_1":                     r.ConnectionType,
		"systemname1":                           r.Systemname,
		"access_add_workflow":                   r.AccessAddWorkflow,
		"access_remove_workflow":                r.AccessRemoveWorkflow,
		"add_service_account_workflow":          r.AddServiceAccountWorkflow,
		"remove_service_account_workflow":       r.RemoveServiceAccountWorkflow,
		"connection_parameters":                 r.Connectionparameters,
		"automated_provisioning":                r.AutomatedProvisioning,
		"use_open_connector":                    r.UseOpenConnector,
		"manage_entity":                         r.ManageEntity,
		"persistent_data":                       r.PersistentData,
		"default_system":                        r.DefaultSystem,
		"recon_application":                     r.ReconApplication,
		"instant_provision":                     r.InstantProvision,
		"provisioning_tries":                    r.ProvisioningTries,
		"provisioning_comments":                 r.ProvisioningComments,
		"proposed_account_owners_workflow":      r.ProposedAccountOwnersWorkflow,
		"firefighterid_workflow":                r.FirefighterIDWorkflow,
		"firefighterid_request_access_workflow": r.FirefighterIDRequestAccessWorkflow,
		"policy_rule":                           r.PolicyRule,
		"policy_rule_service_account":           r.PolicyRuleServiceAccount,
		"connectionname1":                       r.Connectionname,
		"provisioning_connection":               r.ProvisioningConnection,
		"service_desk_connection":               r.ServiceDeskConnection,
		"external_risk_connection_json":         r.ExternalRiskConnectionJson,
		"connection":                            r.Connection,
		"create_date":                           r.CreateDate,
		"update_date":                           r.UpdateDate,
		"endpoints":                             r.Endpoints,
		"created_by":                            r.CreatedBy,
		"updated_by":                            r.UpdatedBy,
		"status":                                r.Status,
		"created_from":                          r.CreatedFrom,
		"port":                                  r.Port,
		"inherent_sod_report_fields":            inherentSodList,
	}, diags
}

// Read fetches data from the API and converts it to Terraform state.
func (d *SecuritySystemsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SecuritySystemsDataSourceModel

	// Retrieve user-defined filters from configuration.
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare the API client configuration.
	cfg := openapi.NewConfiguration()
	apiBaseURL := d.client.APIBaseURL()
	cfg.Host = strings.TrimPrefix(apiBaseURL, "https://")
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+d.token)
	cfg.HTTPClient = http.DefaultClient

	// Initialize API client.
	apiClient := openapi.NewAPIClient(cfg)

	// Use provided Max filter or default value.
	apiReq := apiClient.SecuritySystemsAPI.GetSecuritySystems(ctx)

	// Only set Systemname if a non-null value is provided.
	if !state.Systemname.IsNull() && state.Systemname.ValueString() != "" {
		apiReq = apiReq.Systemname(state.Systemname.ValueString())
	}

	// Only set Max if provided.
	if !state.Max.IsNull() {
		apiReq = apiReq.Max(int32(state.Max.ValueInt64()))
	}

	// Only set Offset if provided.
	if !state.Offset.IsNull() {
		apiReq = apiReq.Offset(int32(state.Offset.ValueInt64()))
	}

	// Only set Connectionname if provided.
	if !state.Connectionname.IsNull() && state.Connectionname.ValueString() != "" {
		apiReq = apiReq.Connectionname(state.Connectionname.ValueString())
	}

	// Only set ConnectionType if provided.
	if !state.ConnectionType.IsNull() && state.ConnectionType.ValueString() != "" {
		apiReq = apiReq.ConnectionType(state.ConnectionType.ValueString())
	}

	// Execute the API request.
	apiResp, httpResp, err := apiReq.Execute()
	if err != nil {
		fmt.Printf("Error marshalling apiResp: %v\n", err)
		return
	}

	fmt.Printf("[DEBUG] HTTP Status Code: %d\n", httpResp.StatusCode)

	// Transform API response to a slice of SecuritySystemDetails.
	state.Msg = types.StringValue(*apiResp.Msg)
	state.DisplayCount = types.Int64Value(int64(*apiResp.DisplayCount))
	state.ErrorCode = types.StringValue(*apiResp.ErrorCode)
	state.TotalCount = types.Int64Value(int64(*apiResp.TotalCount))
	var results []SecuritySystemDetails
	if apiResp.SecuritySystemDetails != nil {
		for _, item := range apiResp.SecuritySystemDetails {
			results = append(results, SecuritySystemDetails{
				ConnectionType:                     util.SafeString(item.ConnectionType),
				Connectionname:                     util.SafeString(item.Connectionname),
				DefaultSystem:                      util.SafeString(item.DefaultSystem),
				DisplayName:                        util.SafeString(item.DisplayName),
				Hostname:                           util.SafeString(item.Hostname),
				Systemname:                         util.SafeString(item.Systemname),
				AccessAddWorkflow:                  util.SafeString(item.AccessAddWorkflow),
				AccessRemoveWorkflow:               util.SafeString(item.AccessRemoveWorkflow),
				AddServiceAccountWorkflow:          util.SafeString(item.AddServiceAccountWorkflow),
				RemoveServiceAccountWorkflow:       util.SafeString(item.RemoveServiceAccountWorkflow),
				Connectionparameters:               util.SafeString(item.Connectionparameters),
				ProvisioningConnection:             util.SafeString(item.ProvisioningConnection),
				Connection:                         util.SafeString(item.Connection),
				CreateDate:                         util.SafeString(item.CreateDate),
				UpdateDate:                         util.SafeString(item.UpdateDate),
				Endpoints:                          util.SafeString(item.Endpoints),
				UseOpenConnector:                   util.SafeString(item.Useopenconnector),
				ReconApplication:                   util.SafeString(item.ReconApplication),
				AutomatedProvisioning:              util.SafeString(item.AutomatedProvisioning),
				InstantProvision:                   util.SafeString(item.Instantprovision),
				ProvisioningComments:               util.SafeString(item.Provisioningcomments),
				ProvisioningTries:                  util.SafeString(item.ProvisioningTries),
				ProposedAccountOwnersWorkflow:      util.SafeString(item.ProposedAccountOwnersworkflow),
				FirefighterIDWorkflow:              util.SafeString(item.FirefighteridWorkflow),
				FirefighterIDRequestAccessWorkflow: util.SafeString(item.FirefighteridRequestAccessWorkflow),
				PolicyRuleServiceAccount:           util.SafeString(item.PolicyRuleServiceAccount),
				ServiceDeskConnection:              util.SafeString(item.ServiceDeskConnection),
				ExternalRiskConnectionJson:         util.SafeString(item.ExternalRiskConnectionJson),
				CreatedBy:                          util.SafeString(item.CreatedBy),
				UpdatedBy:                          util.SafeString(item.UpdatedBy),
				Status:                             util.SafeString(item.Status),
				CreatedFrom:                        util.SafeString(item.CreatedFrom),
				PolicyRule:                         util.SafeString(item.PolicyRule),
				Port:                               util.SafeString(item.Port),
				InherentSodReportFields:            util.ToTypesStringSlice(item.InherentSODReportFields),
			})
		}
	}

	// Convert the results slice to a list of Terraform object values.
	var listValues []attr.Value
	for _, item := range results {
		m, mDiags := item.ToMap()
		if mDiags.HasError() {
			resp.Diagnostics.Append(mDiags...)
			return
		}
		objVal, objDiags := types.ObjectValue(securitySystemObjectType.AttrTypes, m)

		if objDiags.HasError() {
			resp.Diagnostics.Append(objDiags...)
			return
		}
		listValues = append(listValues, objVal)
	}
	state.Results, diags = types.ListValue(securitySystemObjectType, listValues)
	resp.Diagnostics.Append(diags...)

	// Save the state.
	resp.State.Set(ctx, state)
}
