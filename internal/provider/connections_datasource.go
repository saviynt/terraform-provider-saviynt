// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	// "encoding/json"
	"fmt"
	"net/http"
	"strings"

	// openapi "github.com/saviynt/saviynt-api-go-client/connections"
	openapi "connections"

	s "github.com/saviynt/saviynt-api-go-client"

	// "github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConnectionsDataSource struct {
	client *s.Client
	token  string
}

var _ datasource.DataSource = &ConnectionsDataSource{}
var _ datasource.DataSourceWithConfigure = &ConnectionsDataSource{}

func NewConnectionsDataSource() datasource.DataSource {
	return &ConnectionsDataSource{}
}

type ConnectionsDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	// Results        types.List   `tfsdk:"results"`
	Results        []Connection `tfsdk:"results"`
	ConnectionName types.String `tfsdk:"connection_name"`
	Offset         types.String `tfsdk:"offset"`
	DisplayCount   types.Int64  `tfsdk:"display_count"`
	ErrorCode      types.String `tfsdk:"error_code"`
	TotalCount     types.Int64  `tfsdk:"total_count"`
	Msg            types.String `tfsdk:"msg"`
	ConnectionType types.String `tfsdk:"connection_type"`
	Max            types.String `tfsdk:"max"`
}

type Connection struct {
	ConnectionName        types.String `tfsdk:"connectionname"`
	ConnectionType        types.String `tfsdk:"connectiontype"`
	ConnectionDescription types.String `tfsdk:"connectiondescription"`
	Status                types.Int32  `tfsdk:"status"`
	CreatedBy             types.String `tfsdk:"createdby"`
	CreatedOn             types.String `tfsdk:"createdon"`
	UpdatedBy             types.String `tfsdk:"updatedby"`
	UpdatedOn             types.String `tfsdk:"updatedon"`
}

func (d *ConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_connections_datasource"
}

func (d *ConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Datasource to retrieve all connections",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"connection_name": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by connection name",
			},
			"connection_type": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by connection type",
			},
			"max": schema.StringAttribute{
				Optional:    true,
				Description: "Maximum number of connections to retrieve",
			},
			"offset": schema.StringAttribute{
				Optional:    true,
				Description: "Offset",
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
				Description: "Total count of available connections",
			},
			"msg": schema.StringAttribute{
				Computed:    true,
				Description: "API response message",
			},
			"results": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of connections retrieved",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// "id1": schema.StringAttribute{
						// 	Computed: true,
						// 	Description: "Connection ID",
						// },
						"connectionname": schema.StringAttribute{
							Computed:    true,
							Description: "Connection Name",
						},
						"connectiontype": schema.StringAttribute{
							Computed:    true,
							Description: "Type of connection",
						},
						"connectiondescription": schema.StringAttribute{
							Computed:    true,
							Description: "Description of the connection",
						},
						"status": schema.Int64Attribute{
							Computed:    true,
							Description: "Status of the connection",
						},
						"createdby": schema.StringAttribute{
							Computed:    true,
							Description: "User who created the connection",
						},
						"createdon": schema.StringAttribute{
							Computed:    true,
							Description: "Timestamp when the connection was created",
						},
						"updatedby": schema.StringAttribute{
							Computed:    true,
							Description: "User who last updated the connection",
						},
						"updatedon": schema.StringAttribute{
							Computed:    true,
							Description: "Timestamp when the connection was last updated",
						},
					},
				},
			},
		},
	}
}

func (d *ConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// var connectionObjectType = types.ObjectType{
// 	AttrTypes: map[string]attr.Type{
// 		"connectionname":       types.StringType,
// 		"connectiontype":       types.StringType,
// 		"connectiondescription": types.StringType,
// 		"status":               types.Int64Type,
// 		"createdby":            types.StringType,
// 		"createdon":            types.StringType,
// 		"updatedby":            types.StringType,
// 		"updatedon":            types.StringType,
// 	},
// }

// func (c Connection) ToMap() map[string]attr.Value {
// 	return map[string]attr.Value{
// 		"connectionname":        c.ConnectionName,
// 		"connectiontype":        c.ConnectionType,
// 		"connectiondescription": c.ConnectionDescription,
// 		"status":                c.Status,
// 		"createdby":             c.CreatedBy,
// 		"createdon":             c.CreatedOn,
// 		"updatedby":             c.UpdatedBy,
// 		"updatedon":             c.UpdatedOn,
// 	}
// }

func (d *ConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ConnectionsDataSourceModel

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

	areq := openapi.NewGetConnectionsRequest()

	if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
		connectionName := state.ConnectionName.ValueString()
		areq.Connectionname = &connectionName
	}

	if !state.ConnectionType.IsNull() && state.ConnectionType.ValueString() != "" {
		connectionType := state.ConnectionType.ValueString()
		areq.Connectiontype = &connectionType
	}

	if !state.Offset.IsNull() && state.Offset.ValueString() != "" {
		offset := state.Offset.ValueString()
		areq.Offset = &offset
	}

	if !state.Max.IsNull() && state.Max.ValueString() != "" {
		max := state.Max.ValueString()
		areq.Max = &max
	}

	apiReq := apiClient.ConnectionsAPI.GetConnections(ctx).GetConnectionsRequest(*areq)

	connectionsResponse, httpResp, err := apiReq.Execute()
	if err != nil {
		fmt.Printf("[ERROR] API Call Failed: %v\n", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	fmt.Printf("[DEBUG] HTTP Status Code: %d\n", httpResp.StatusCode)

	// jsonBytes, err := json.MarshalIndent(connectionsResponse, "", "  ")
	// if err != nil {
	// 	fmt.Printf("Error marshalling API response: %v\n", err)
	// 	return
	// }
	// fmt.Println("Marshalled API Response:")
	// fmt.Println(string(jsonBytes))

	state.Msg = types.StringValue(*connectionsResponse.Msg)
	state.DisplayCount = types.Int64Value(int64(*connectionsResponse.DisplayCount))
	state.ErrorCode = types.StringValue(*connectionsResponse.ErrorCode)
	state.TotalCount = types.Int64Value(int64(*connectionsResponse.TotalCount))
	// state.Msg = safeString(connectionsResponse.Msg)
	// state.DisplayCount = safeInt64(pointerToInt64(connectionsResponse.DisplayCount))
	// state.ErrorCode = safeString(connectionsResponse.ErrorCode)
	// state.TotalCount = safeInt64(pointerToInt64(connectionsResponse.TotalCount))

	// var results []Connection

	// if connectionsResponse != nil && connectionsResponse.ConnectionList != nil {
	// 	for _, item := range connectionsResponse.ConnectionList {
	// 		results = append(results, Connection{
	// 			ConnectionName:        safeString(item.CONNECTIONNAME),
	// 			ConnectionType:        safeString(item.CONNECTIONTYPE),
	// 			ConnectionDescription: safeString(item.CONNECTIONDESCRIPTION),
	// 			Status:                safeInt32(item.STATUS),
	// 			CreatedBy:             safeString(item.CREATEDBY),
	// 			CreatedOn:             safeString(item.CREATEDON),
	// 			UpdatedBy:             safeString(item.UPDATEDBY),
	// 			UpdatedOn:             safeString(item.UPDATEDON),
	// 		})
	// 	}
	// }
	if connectionsResponse != nil && connectionsResponse.ConnectionList != nil {
		for _, item := range connectionsResponse.ConnectionList {
			resultState := Connection{
				ConnectionName:        safeString(item.CONNECTIONNAME),
				ConnectionType:        safeString(item.CONNECTIONTYPE),
				ConnectionDescription: safeString(item.CONNECTIONDESCRIPTION),
				Status:                safeInt32(item.STATUS),
				CreatedBy:             safeString(item.CREATEDBY),
				CreatedOn:             safeString(item.CREATEDON),
				UpdatedBy:             safeString(item.UPDATEDBY),
				UpdatedOn:             safeString(item.UPDATEDON),
			}

			state.Results = append(state.Results, resultState)
		}
	}

	diags1 := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags1...)
	if resp.Diagnostics.HasError() {
		return
	}

	// fmt.Println("Debug: Printing Results Before Assigning to State:")
	// for _, item := range results {
	// 	fmt.Printf("Item: %+v\n", item.ConnectionName)
	// }

	// var listValues []attr.Value
	// for _, item := range results {
	// 	objVal, objDiags := types.ObjectValue(connectionObjectType.AttrTypes, item.ToMap())
	// 	if objDiags.HasError() {
	// 		fmt.Println("[ERROR] ObjectValue conversion failed:", objDiags)
	// 		resp.Diagnostics.Append(objDiags...)
	// 		continue
	// 	}
	// 	listValues = append(listValues, objVal)
	// }

	// listValue, diags := types.ListValue(connectionObjectType, listValues)
	// if diags.HasError() {
	// 	fmt.Println("[ERROR] ListValue conversion failed:", diags)
	// 	resp.Diagnostics.Append(diags...)
	// 	return
	// }
	// state.Results = listValue
	// fmt.Println("Diagnostics after ListValue:", diags)

	// resp.Diagnostics.Append(diags...)
	// err1 := resp.State.Set(ctx, state)
	// if err1 != nil {
	// 	fmt.Println("Error setting state:", err1)
	// }
	// fmt.Printf("Final state before set: %+v\n", state)
}

func safeInt32(ptr *int32) types.Int32 {
	if ptr == nil {
		return types.Int32Null()
	}
	return types.Int32Value(*ptr)
}

func safeInt64(ptr *int64) types.Int64 {
	if ptr == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*ptr)
}

func pointerToInt64(ptr *int32) *int64 {
	if ptr == nil {
		return nil
	}
	val := int64(*ptr) // Convert int32 to int64
	return &val
}
