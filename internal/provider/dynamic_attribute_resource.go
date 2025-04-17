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
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
	openapi "github.com/saviynt/saviynt-api-go-client/dynamicattributes"
)

type DynamicAttributeResourceModel struct {
	ID               types.String      `tfsdk:"id"` // Typically generated/managed by the API
	Securitysystem   types.String      `tfsdk:"security_system"`
	Endpoint         types.String      `tfsdk:"endpoint"`
	Username         types.String      `tfsdk:"user_name"`
	Updateuser       types.String      `tfsdk:"update_user"`
	Dynamicattribute *DynamicAttribute `tfsdk:"dynamic_attribute"`
	Msg              types.String      `tfsdk:"msg"`
	ErrorCode        types.String      `tfsdk:"error_code"`
}

type DynamicAttribute struct {
	Attributename                             types.String `tfsdk:"attribute_name"`
	Requesttype                               types.String `tfsdk:"request_type"`
	Attributetype                             types.String `tfsdk:"attribute_type"`
	Attributegroup                            types.String `tfsdk:"attribute_group"`
	Orderindex                                types.String `tfsdk:"order_index"`
	Attributelable                            types.String `tfsdk:"attribute_lable"`
	Accountscolumn                            types.String `tfsdk:"accounts_column"`
	Hideoncreate                              types.String `tfsdk:"hide_on_create"`
	Actionstring                              types.String `tfsdk:"action_string"`
	Editable                                  types.String `tfsdk:"editable"`
	Hideonupdate                              types.String `tfsdk:"hide_on_update"`
	Actiontoperformwhenparentattributechanges types.String `tfsdk:"actiontoperformwhenparentattributechanges"`
	Defaultvalue                              types.String `tfsdk:"default_value"`
	Required                                  types.String `tfsdk:"required"`
	Regex                                     types.String `tfsdk:"regex"`
	Attributevalue                            types.String `tfsdk:"attribute_value"`
	Showonchild                               types.String `tfsdk:"showonchild"`
	Parentattribute                           types.String `tfsdk:"parentattribute"`
	Descriptionascsv                          types.String `tfsdk:"descriptionascsv"`
}

type DynamicAttributeResource struct {
	client *s.Client
	token  string
}

func NewDynamicAttributeResource() resource.Resource {
	return &DynamicAttributeResource{}
}

func (r *DynamicAttributeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_dynamic_attribute_resource"
}

func (r *DynamicAttributeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and Manage Dynamic Attributes",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the resource, typically managed by the API.",
			},
			"security_system": schema.StringAttribute{
				Required:    true,
				Description: "Security system associated with the dynamic attribute.",
			},
			"endpoint": schema.StringAttribute{
				Required:    true,
				Description: "Endpoint associated with the dynamic attribute.",
			},
			"user_name": schema.StringAttribute{
				Required:    true,
				Description: "Username of the user creating or managing the dynamic attribute.",
			},
			"update_user": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User who last updated the dynamic attribute.",
			},
			"msg": schema.StringAttribute{
				Computed:    true,
				Description: "Response message from the API.",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code from the API response.",
			},
			"dynamic_attribute": schema.SingleNestedAttribute{
				Required:    true,
				Description: "Configuration block for the dynamic attribute itself.",
				Attributes: map[string]schema.Attribute{
					"attribute_name": schema.StringAttribute{
						Required:    true,
						Description: "Specify the dynamic attribute name.",
					},
					"request_type": schema.StringAttribute{
						Required:    true,
						Description: "Type of request.",
					},
					"attribute_type": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Attribute type used for filtering and display.",
					},
					"attribute_group": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Group or categorize the attribute in the request form.",
					},
					"order_index": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Sequence for display of the dynamic attribute.",
					},
					"attribute_lable": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Name to be shown in the Access Requests form.",
					},
					"accounts_column": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Accounts column mapping.",
					},
					"hide_on_create": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Whether to hide this attribute on create.",
					},
					"action_string": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Action string value.",
					},
					"editable": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Whether the attribute is editable.",
					},
					"hide_on_update": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Whether to hide this attribute on update.",
					},
					"actiontoperformwhenparentattributechanges": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Action to perform when the parent attribute changes.",
					},
					"default_value": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Default value for the attribute.",
					},
					"required": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Whether this attribute is required.",
					},
					"regex": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Regex for validation.",
					},
					"attribute_value": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Value options or query for the attribute.",
					},
					"showonchild": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Whether to show this on child requests.",
					},
					"parentattribute": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Parent attribute this one depends on.",
					},
					"descriptionascsv": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Description of values as CSV.",
					},
				},
			},
		},
	}
}

func (r *DynamicAttributeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DynamicAttributeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DynamicAttributeResourceModel

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

	dynamicAttr := openapi.NewCreateDynamicAttributeRequestDynamicattributesInner(
		plan.Dynamicattribute.Attributename.ValueString(),
		plan.Dynamicattribute.Requesttype.ValueString(),
	)

	dynamicAttrs := []openapi.CreateDynamicAttributeRequestDynamicattributesInner{*dynamicAttr}
	dynamicAttr.Attributetype = util.StringPointerOrEmpty(plan.Dynamicattribute.Attributetype)
	dynamicAttr.Attributegroup = util.StringPointerOrEmpty(plan.Dynamicattribute.Attributegroup)
	dynamicAttr.Orderindex = util.StringPointerOrEmpty(plan.Dynamicattribute.Orderindex)
	dynamicAttr.Attributelable = util.StringPointerOrEmpty(plan.Dynamicattribute.Attributelable)
	dynamicAttr.Accountscolumn = util.StringPointerOrEmpty(plan.Dynamicattribute.Accountscolumn)
	dynamicAttr.Hideoncreate = util.StringPointerOrEmpty(plan.Dynamicattribute.Hideoncreate)
	dynamicAttr.Actionstring = util.StringPointerOrEmpty(plan.Dynamicattribute.Actionstring)
	dynamicAttr.Editable = util.StringPointerOrEmpty(plan.Dynamicattribute.Editable)
	dynamicAttr.Hideonupdate = util.StringPointerOrEmpty(plan.Dynamicattribute.Hideonupdate)
	dynamicAttr.Actiontoperformwhenparentattributechanges = util.StringPointerOrEmpty(plan.Dynamicattribute.Actiontoperformwhenparentattributechanges)
	dynamicAttr.Defaultvalue = util.StringPointerOrEmpty(plan.Dynamicattribute.Defaultvalue)
	dynamicAttr.Required = util.StringPointerOrEmpty(plan.Dynamicattribute.Required)
	dynamicAttr.Regex = util.StringPointerOrEmpty(plan.Dynamicattribute.Regex)
	dynamicAttr.Attributevalue = util.StringPointerOrEmpty(plan.Dynamicattribute.Attributevalue)
	dynamicAttr.Showonchild = util.StringPointerOrEmpty(plan.Dynamicattribute.Showonchild)
	dynamicAttr.Parentattribute = util.StringPointerOrEmpty(plan.Dynamicattribute.Parentattribute)
	dynamicAttr.Descriptionascsv = util.StringPointerOrEmpty(plan.Dynamicattribute.Descriptionascsv)

	createReq := openapi.NewCreateDynamicAttributeRequest(
		plan.Securitysystem.ValueString(),
		plan.Endpoint.ValueString(),
		plan.Username.ValueString(),
		dynamicAttrs,
	)

	createResp, httpResp, err := apiClient.DynamicAttributesAPI.
		CreateDynamicAttribute(ctx).
		CreateDynamicAttributeRequest(*createReq).
		Execute()

	if err != nil {
		log.Printf("Error Creating Dynamic attribute: %v, HTTP Response: %v", err, httpResp)
		resp.Diagnostics.AddError(
			"Error Creating Dynamic Attribute",
			"Check logs for details.",
		)
		return
	}
	log.Println("[DEBUG] Create HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Create API Response: %+v", createResp)

	plan.ID = types.StringValue("dynamic-attr-" + plan.Dynamicattribute.Attributename.ValueString())
	// plan.Securitysystem = util.SafeString(plan.Securitysystem.ValueStringPointer())
	// plan.Endpoint = util.SafeString(plan.Endpoint.ValueStringPointer())
	// plan.Username = util.SafeString(plan.Username.ValueStringPointer())
	plan.Updateuser = util.SafeString(plan.Updateuser.ValueStringPointer())
	plan.Dynamicattribute.Attributename = util.SafeString(plan.Dynamicattribute.Attributename.ValueStringPointer())
	plan.Dynamicattribute.Requesttype = util.SafeString(plan.Dynamicattribute.Requesttype.ValueStringPointer())
	plan.Dynamicattribute.Attributetype = util.SafeString(plan.Dynamicattribute.Attributetype.ValueStringPointer())
	plan.Dynamicattribute.Attributegroup = util.SafeString(plan.Dynamicattribute.Attributegroup.ValueStringPointer())
	plan.Dynamicattribute.Orderindex = util.SafeString(plan.Dynamicattribute.Orderindex.ValueStringPointer())
	plan.Dynamicattribute.Attributelable = util.SafeString(plan.Dynamicattribute.Attributelable.ValueStringPointer())
	plan.Dynamicattribute.Accountscolumn = util.SafeString(plan.Dynamicattribute.Accountscolumn.ValueStringPointer())
	plan.Dynamicattribute.Hideoncreate = util.SafeString(plan.Dynamicattribute.Hideoncreate.ValueStringPointer())
	plan.Dynamicattribute.Actionstring = util.SafeString(plan.Dynamicattribute.Actionstring.ValueStringPointer())
	plan.Dynamicattribute.Editable = util.SafeString(plan.Dynamicattribute.Editable.ValueStringPointer())
	plan.Dynamicattribute.Hideonupdate = util.SafeString(plan.Dynamicattribute.Hideonupdate.ValueStringPointer())
	plan.Dynamicattribute.Actiontoperformwhenparentattributechanges = util.SafeString(plan.Dynamicattribute.Actiontoperformwhenparentattributechanges.ValueStringPointer())
	plan.Dynamicattribute.Defaultvalue = util.SafeString(plan.Dynamicattribute.Defaultvalue.ValueStringPointer())
	plan.Dynamicattribute.Required = util.SafeString(plan.Dynamicattribute.Required.ValueStringPointer())
	plan.Dynamicattribute.Regex = util.SafeString(plan.Dynamicattribute.Regex.ValueStringPointer())
	plan.Dynamicattribute.Attributevalue = util.SafeString(plan.Dynamicattribute.Attributevalue.ValueStringPointer())
	plan.Dynamicattribute.Showonchild = util.SafeString(plan.Dynamicattribute.Showonchild.ValueStringPointer())
	plan.Dynamicattribute.Parentattribute = util.SafeString(plan.Dynamicattribute.Parentattribute.ValueStringPointer())
	plan.Dynamicattribute.Descriptionascsv = util.SafeString(plan.Dynamicattribute.Descriptionascsv.ValueStringPointer())

	msgValue := util.SafeDeref(createResp.Msg)
	errorCodeValue := util.SafeDeref(createResp.Errorcode)

	// Set the individual fields
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)

	stateCreateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateCreateDiagnostics...)
}

func (r *DynamicAttributeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DynamicAttributeResourceModel

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

	fetchReq := apiClient.DynamicAttributesAPI.
		FetchDynamicAttribute(ctx).
		Securitysystem([]string{state.Securitysystem.ValueString()}).
		Endpoint([]string{state.Endpoint.ValueString()}).
		Dynamicattributes([]string{state.Dynamicattribute.Attributename.ValueString()})

	apiResp, httpResp, err := fetchReq.Execute()

	if err != nil {
		log.Printf("[ERROR] Fetch API Call Failed: %v", err)
		resp.Diagnostics.AddError("Fetch API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] Read HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Read API Response: %+v", apiResp)

	var foundItem *openapi.FetchDynamicAttribute200ResponseDynamicattributesInner
	for _, item := range apiResp.Dynamicattributes {
		if item.Attributename != nil && *item.Attributename == state.Dynamicattribute.Attributename.ValueString() {
			foundItem = &item
			break
		}
	}
	if foundItem == nil {
		resp.State.RemoveResource(ctx)
	}

	log.Printf("[DEBUG] Read Found Item: %+v", foundItem)
	state.ID = types.StringValue("dynamic-attr-" + state.Dynamicattribute.Attributename.ValueString())
	state.Endpoint = util.SafeString(state.Endpoint.ValueStringPointer())
	state.Securitysystem = util.SafeString(state.Securitysystem.ValueStringPointer())
	state.Username = util.SafeString(state.Username.ValueStringPointer())
	state.Updateuser = util.SafeString(state.Updateuser.ValueStringPointer())
	state.Dynamicattribute.Attributename = util.SafeString(foundItem.Attributename)
	state.Dynamicattribute.Requesttype = util.SafeString(foundItem.Requesttype)
	state.Dynamicattribute.Attributetype = util.SafeString(foundItem.Attributetype)
	state.Dynamicattribute.Attributegroup = util.SafeString(foundItem.Attributegroup)
	state.Dynamicattribute.Orderindex = util.SafeString(foundItem.Orderindex)
	state.Dynamicattribute.Attributelable = util.SafeString(foundItem.Attributelable)
	state.Dynamicattribute.Accountscolumn = util.SafeString(foundItem.Accountscolumn)
	state.Dynamicattribute.Hideoncreate = util.SafeString(foundItem.Hideoncreate)
	state.Dynamicattribute.Actionstring = util.SafeString(foundItem.Actionstring)
	state.Dynamicattribute.Editable = util.SafeString(foundItem.Editable)
	state.Dynamicattribute.Hideonupdate = util.SafeString(foundItem.Hideonupdate)
	state.Dynamicattribute.Actiontoperformwhenparentattributechanges = util.SafeString(foundItem.Actiontoperformwhenparentattributechanges)
	state.Dynamicattribute.Defaultvalue = util.SafeString(foundItem.Defaultvalue)
	state.Dynamicattribute.Required = util.SafeString(foundItem.Required)
	state.Dynamicattribute.Regex = util.SafeString(foundItem.Regex)
	state.Dynamicattribute.Attributevalue = util.SafeString(foundItem.Attributevalue)
	state.Dynamicattribute.Showonchild = util.SafeString(foundItem.Showonchild)
	state.Dynamicattribute.Parentattribute = util.SafeString(foundItem.Parentattribute)
	state.Dynamicattribute.Descriptionascsv = util.SafeString(foundItem.Descriptionascsv)

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.Errorcode)
	state.Msg = types.StringValue(msgValue)
	state.ErrorCode = types.StringValue(errorCodeValue)

	stateSetDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateSetDiagnostics...)
}

func (r *DynamicAttributeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DynamicAttributeResourceModel

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
	// dynamicAttr := openapi.NewUpdateDynamicAttributeRequestDynamicattributesInner(
	// 	plan.Dynamicattribute.Attributename.ValueString(),
	// )

	dynamicAttr:=openapi.UpdateDynamicAttributeRequestDynamicattributesInner{
		Attributename: plan.Dynamicattribute.Attributename.ValueString(),
		Requesttype:   util.StringPointerOrEmpty(plan.Dynamicattribute.Requesttype),
		Accountscolumn: util.StringPointerOrEmpty(plan.Dynamicattribute.Accountscolumn),
		Attributetype: util.StringPointerOrEmpty(plan.Dynamicattribute.Attributetype),
		Attributegroup: util.StringPointerOrEmpty(plan.Dynamicattribute.Attributegroup),
		Orderindex: util.StringPointerOrEmpty(plan.Dynamicattribute.Orderindex),
		Attributelable: util.StringPointerOrEmpty(plan.Dynamicattribute.Attributelable),
		Hideoncreate: util.StringPointerOrEmpty(plan.Dynamicattribute.Hideoncreate),
		Actionstring: util.StringPointerOrEmpty(plan.Dynamicattribute.Actionstring),
		Editable: util.StringPointerOrEmpty(plan.Dynamicattribute.Editable),
		Hideonupdate: util.StringPointerOrEmpty(plan.Dynamicattribute.Hideonupdate),
		Actiontoperformwhenparentattributechanges: util.StringPointerOrEmpty(plan.Dynamicattribute.Actiontoperformwhenparentattributechanges),
		Defaultvalue: util.StringPointerOrEmpty(plan.Dynamicattribute.Defaultvalue),
		Required: util.StringPointerOrEmpty(plan.Dynamicattribute.Required),
		Regex: util.StringPointerOrEmpty(plan.Dynamicattribute.Regex),
		Attributevalue: util.StringPointerOrEmpty(plan.Dynamicattribute.Attributevalue),
		Showonchild: util.StringPointerOrEmpty(plan.Dynamicattribute.Showonchild),
		Parentattribute: util.StringPointerOrEmpty(plan.Dynamicattribute.Parentattribute),
		Descriptionascsv: util.StringPointerOrEmpty(plan.Dynamicattribute.Descriptionascsv),
	}

	dynamicAttrs := []openapi.UpdateDynamicAttributeRequestDynamicattributesInner{dynamicAttr}

	updateUser := plan.Updateuser.ValueString()
	if plan.Updateuser.IsNull() || updateUser == "" {
		updateUser = plan.Username.ValueString()
	}

	updateReq := openapi.NewUpdateDynamicAttributeRequest(
		plan.Securitysystem.ValueString(),
		plan.Endpoint.ValueString(),
		updateUser,
		dynamicAttrs,
	)

	log.Printf("[DEBUG] Update Request: %+v", updateReq)

	updateResp, httpResp, err := apiClient.DynamicAttributesAPI.
		UpdateDynamicAttribute(ctx).
		UpdateDynamicAttributeRequest(*updateReq).
		Execute()

	if err != nil {
		log.Printf("Error Updating Dynamic attribute: %v, HTTP Response: %v", err, httpResp)
		resp.Diagnostics.AddError(
			"Error updating Dynamic Attribute",
			"Check logs for details.",
		)
		return
	}

	log.Printf("[DEBUG] Update HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Update API Response: %+v", updateResp)

	fetchReq := apiClient.DynamicAttributesAPI.
		FetchDynamicAttribute(ctx).
		Securitysystem([]string{plan.Securitysystem.ValueString()}).
		Endpoint([]string{plan.Endpoint.ValueString()}).
		Dynamicattributes([]string{plan.Dynamicattribute.Attributename.ValueString()})

	fetchResp, httpResp, err := fetchReq.Execute()

	if err != nil {
		log.Printf("[ERROR] Fetch API Call Failed: %v", err)
		resp.Diagnostics.AddError("Fetch API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Fetch API Response: %+v", fetchResp)

	var foundItem *openapi.FetchDynamicAttribute200ResponseDynamicattributesInner
	for _, item := range fetchResp.Dynamicattributes {
		if item.Attributename != nil && *item.Attributename == plan.Dynamicattribute.Attributename.ValueString() {
			foundItem = &item
			break
		}
	}
	if foundItem == nil {
		resp.State.RemoveResource(ctx)
	}

	log.Printf("[DEBUG] Read Found Item: %+v", foundItem)
	plan.ID = types.StringValue("dynamic-attr-" + plan.Dynamicattribute.Attributename.ValueString())
	plan.Endpoint = util.SafeString(plan.Endpoint.ValueStringPointer())
	plan.Securitysystem = util.SafeString(plan.Securitysystem.ValueStringPointer())
	plan.Username = util.SafeString(plan.Username.ValueStringPointer())
	plan.Updateuser = util.SafeString(plan.Updateuser.ValueStringPointer())

	plan.Dynamicattribute.Attributename = util.SafeString(foundItem.Attributename)
	plan.Dynamicattribute.Requesttype = util.SafeString(foundItem.Requesttype)
	plan.Dynamicattribute.Attributetype = util.SafeString(foundItem.Attributetype)
	plan.Dynamicattribute.Attributegroup = util.SafeString(foundItem.Attributegroup)
	plan.Dynamicattribute.Orderindex = util.SafeString(foundItem.Orderindex)
	plan.Dynamicattribute.Attributelable = util.SafeString(foundItem.Attributelable)
	plan.Dynamicattribute.Accountscolumn = util.SafeString(foundItem.Accountscolumn)
	plan.Dynamicattribute.Hideoncreate = util.SafeString(foundItem.Hideoncreate)
	plan.Dynamicattribute.Actionstring = util.SafeString(foundItem.Actionstring)
	plan.Dynamicattribute.Editable = util.SafeString(foundItem.Editable)
	plan.Dynamicattribute.Hideonupdate = util.SafeString(foundItem.Hideonupdate)
	plan.Dynamicattribute.Actiontoperformwhenparentattributechanges = util.SafeString(foundItem.Actiontoperformwhenparentattributechanges)
	plan.Dynamicattribute.Defaultvalue = util.SafeString(foundItem.Defaultvalue)
	plan.Dynamicattribute.Required = util.SafeString(foundItem.Required)
	plan.Dynamicattribute.Regex = util.SafeString(foundItem.Regex)
	plan.Dynamicattribute.Attributevalue = util.SafeString(foundItem.Attributevalue)
	plan.Dynamicattribute.Showonchild = util.SafeString(foundItem.Showonchild)
	plan.Dynamicattribute.Parentattribute = util.SafeString(foundItem.Parentattribute)
	plan.Dynamicattribute.Descriptionascsv = util.SafeString(foundItem.Descriptionascsv)

	msgValue := util.SafeDeref(fetchResp.Msg)
	errorCodeValue := util.SafeDeref(fetchResp.Errorcode)

	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *DynamicAttributeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.State.RemoveResource(ctx)
	var state DynamicAttributeResourceModel

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

	deleteReq:=openapi.DeleteDynamicAttributeRequest{
		Securitysystem: state.Securitysystem.ValueString(),
		Endpoint: state.Endpoint.ValueString(),
		Updateuser: state.Updateuser.ValueString(),
		Dynamicattributes: []string{state.Dynamicattribute.Attributename.ValueString()},
	}

	deleteResp, httpResp, err := apiClient.DynamicAttributesAPI.
		DeleteDynamicAttribute(ctx).
		DeleteDynamicAttributeRequest(deleteReq).
		Execute()

	if err != nil {
		log.Printf("[ERROR] Delete API Call Failed: %v", err)
		resp.Diagnostics.AddError("Delete API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] Delete HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Delete API Response: %+v", deleteResp)
}