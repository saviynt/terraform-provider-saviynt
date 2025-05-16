// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"terraform-provider-Saviynt/util"

	// "github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	// "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	// "github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	// "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"

	// openapi "github.com/saviynt/saviynt-api-go-client/dynamicattributes"
	openapi "dynamicattributes"
)

type DynamicAttributeResourceModel struct {
	ID                types.String       `tfsdk:"id"` // Typically generated/managed by the API
	Securitysystem    types.String       `tfsdk:"security_system"`
	Endpoint          types.String       `tfsdk:"endpoint"`
	Username          types.String       `tfsdk:"user_name"`
	Updateuser        types.String       `tfsdk:"update_user"`
	Dynamicattributes []Dynamicattribute `tfsdk:"dynamic_attributes"`
	// Dynamicattributes types.Set `tfsdk:"dynamic_attributes"`
	Msg               types.String       `tfsdk:"msg"`
	ErrorCode         types.String       `tfsdk:"error_code"`
}

type Dynamicattribute struct {
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
	Parentattribute                           types.String `tfsdk:"parent_attribute"`
	Descriptionascsv                          types.String `tfsdk:"description_as_csv"`
}

type dynamicAttributeResource struct {
	client *s.Client
	token  string
}

func NewDynamicAttributeResource() resource.Resource {
	return &dynamicAttributeResource{}
}

func (r *dynamicAttributeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_dynamic_attribute_resource"
}

func (r *dynamicAttributeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.DynamicAttrDescription,
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
			// "dynamic_attributes": schema.SetNestedAttribute{
			// 	Description: "Set of dynamic attribute configuration blocks.",
			// 	Required:    true,
			// 	// Nested object defines the structure of each element in the set
			// 	NestedObject: schema.NestedAttributeObject{
			// 		Attributes: map[string]schema.Attribute{
			// 			"attribute_name": schema.StringAttribute{
			// 				Required:    true,
			// 				Description: "Specify the dynamic attribute name.",
			// 				PlanModifiers: []planmodifier.String{
			// 					stringplanmodifier.RequiresReplace(),
			// 				},
			// 			},
			// 			"request_type": schema.StringAttribute{
			// 				Required:    true,
			// 				Description: "Type of request.",
			// 			},
			// 			"attribute_type": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Attribute type used for filtering and display.",
			// 			},
			// 			"attribute_group": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Group or categorize the attribute in the request form.",
			// 			},
			// 			"order_index": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Sequence for display of the dynamic attribute.",
			// 			},
			// 			"attribute_label": schema.StringAttribute{  // fixed label spelling here
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Name to be shown in the Access Requests form.",
			// 			},
			// 			"accounts_column": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Accounts column mapping.",
			// 			},
			// 			"hide_on_create": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Whether to hide this attribute on create.",
			// 			},
			// 			"action_string": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Action string value.",
			// 			},
			// 			"editable": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Whether the attribute is editable.",
			// 			},
			// 			"hide_on_update": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Whether to hide this attribute on update.",
			// 			},
			// 			"action_to_perform_when_parent_attribute_changes": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Action to perform when the parent attribute changes.",
			// 			},
			// 			"default_value": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Default value for the attribute.",
			// 			},
			// 			"required": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Whether this attribute is required.",
			// 			},
			// 			"regex": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Regex for validation.",
			// 			},
			// 			"attribute_value": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Value options or query for the attribute.",
			// 			},
			// 			"show_on_child": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Whether to show this on child requests.",
			// 			},
			// 			"parent_attribute": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Parent attribute this one depends on.",
			// 			},
			// 			"description_as_csv": schema.StringAttribute{
			// 				Optional:    true,
			// 				Computed:    true,
			// 				Description: "Description of values as CSV.",
			// 			},
			// 		},
			// 	},
			// 	// Use RequiresReplace if you want the entire resource to be replaced when this set changes
			// 	PlanModifiers: []planmodifier.Set{
			// 		setplanmodifier.RequiresReplace(),
			// 	},
			// },


			"dynamic_attributes": schema.ListNestedAttribute{
				Required:    true,
				Description: "List of dynamic attribute configuration blocks.",
				NestedObject: schema.NestedAttributeObject{
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
						"parent_attribute": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Parent attribute this one depends on.",
						},
						"description_as_csv": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Description of values as CSV.",
						},
					},
				},
			},
		},
	}
}

func (r *dynamicAttributeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *dynamicAttributeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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

	fetchReq := apiClient.DynamicAttributesAPI.
		FetchDynamicAttribute(ctx).
		Securitysystem([]string{plan.Securitysystem.ValueString()}).
		Endpoint([]string{plan.Endpoint.ValueString()})
		// Dynamicattributes([]string{state.Dynamicattributes.Attributename.ValueString()})
	// fetchReq:=apiClient.DynamicAttributesAPI.FetchDynamicAttribute(ctx).Dynamicattributes([]string{state.Dynamicattributes.Attributename.ValueString()})

	apiResp, httpResp, err := fetchReq.Execute()

	if err != nil {
		log.Printf("[ERROR] Fetch API Call Failed: %v", err)
		resp.Diagnostics.AddError("Fetch API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] Read HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Read API Error code: %+v", *apiResp.Errorcode)
	log.Printf("[DEBUG] Read API Message: %+v", *apiResp.Msg)

	// Declare a temporary slice to store mapped attributes
	var fetchedAttrs []Dynamicattribute

	if len(apiResp.Dynamicattributes) != 0 {
		for _, item := range apiResp.Dynamicattributes {
			attr := Dynamicattribute{
				Attributename:                              util.SafeStringDatasource(item.Attributename),
				Requesttype:                                util.SafeStringDatasource(item.Requesttype),
				Attributetype:                              util.SafeStringDatasource(item.Attributetype),
				Attributegroup:                             util.SafeStringDatasource(item.Attributegroup),
				Orderindex:                                 util.SafeStringDatasource(item.Orderindex),
				Attributelable:                             util.SafeStringDatasource(item.Attributelable),
				Accountscolumn:                             util.SafeStringDatasource(item.Accountscolumn),
				Hideoncreate:                               util.SafeStringDatasource(item.Hideoncreate),
				Actionstring:                               util.SafeStringDatasource(item.Actionstring),
				Editable:                                   util.SafeStringDatasource(item.Editable),
				Hideonupdate:                               util.SafeStringDatasource(item.Hideonupdate),
				Actiontoperformwhenparentattributechanges:  util.SafeStringDatasource(item.Actiontoperformwhenparentattributechanges),
				Defaultvalue:                               util.SafeStringDatasource(item.Defaultvalue),
				Required:                                   util.SafeStringDatasource(item.Required),
				Regex:                                      util.SafeStringDatasource(item.Regex),
				Attributevalue:                             util.SafeStringDatasource(item.Attributevalue),
				Showonchild:                                util.SafeStringDatasource(item.Showonchild),
				Parentattribute:                            util.SafeStringDatasource(item.Parentattribute),
				Descriptionascsv:                           util.SafeStringDatasource(item.Descriptionascsv),
			}
			fetchedAttrs = append(fetchedAttrs, attr)
		}
	}


	// plan.ID = types.StringValue("dynamic-attr-" + plan.Endpoint.ValueString())
	// plan.Endpoint = util.SafeString(plan.Endpoint.ValueStringPointer())
	// plan.Securitysystem = util.SafeString(plan.Securitysystem.ValueStringPointer())
	// plan.Username = util.SafeString(plan.Username.ValueStringPointer())
	// plan.Updateuser = util.SafeString(plan.Updateuser.ValueStringPointer())

	// plan.Dynamicattributes = make([]Dynamicattribute, len(apiResp.Dynamicattributes))

	// for i, item := range apiResp.Dynamicattributes {
	// 	plan.Dynamicattributes[i].Attributename = util.SafeStringDatasource(item.Attributename)
	// 	plan.Dynamicattributes[i].Requesttype = util.SafeStringDatasource(item.Requesttype)
	// 	plan.Dynamicattributes[i].Attributetype = util.SafeStringDatasource(item.Attributetype)
	// 	plan.Dynamicattributes[i].Attributegroup = util.SafeStringDatasource(item.Attributegroup)
	// 	plan.Dynamicattributes[i].Orderindex = util.SafeStringDatasource(item.Orderindex)
	// 	plan.Dynamicattributes[i].Attributelable = util.SafeStringDatasource(item.Attributelable)
	// 	plan.Dynamicattributes[i].Accountscolumn = util.SafeStringDatasource(item.Accountscolumn)
	// 	plan.Dynamicattributes[i].Hideoncreate = util.SafeStringDatasource(item.Hideoncreate)
	// 	plan.Dynamicattributes[i].Actionstring = util.SafeStringDatasource(item.Actionstring)
	// 	plan.Dynamicattributes[i].Editable = util.SafeStringDatasource(item.Editable)
	// 	plan.Dynamicattributes[i].Hideonupdate = util.SafeStringDatasource(item.Hideonupdate)
	// 	plan.Dynamicattributes[i].Actiontoperformwhenparentattributechanges = util.SafeStringDatasource(item.Actiontoperformwhenparentattributechanges)
	// 	plan.Dynamicattributes[i].Defaultvalue = util.SafeStringDatasource(item.Defaultvalue)
	// 	plan.Dynamicattributes[i].Required = util.SafeStringDatasource(item.Required)
	// 	plan.Dynamicattributes[i].Regex = util.SafeStringDatasource(item.Regex)
	// 	plan.Dynamicattributes[i].Attributevalue = util.SafeStringDatasource(item.Attributevalue)
	// 	plan.Dynamicattributes[i].Showonchild = util.SafeStringDatasource(item.Showonchild)
	// 	plan.Dynamicattributes[i].Parentattribute = util.SafeStringDatasource(item.Parentattribute)
	// 	plan.Dynamicattributes[i].Descriptionascsv = util.SafeStringDatasource(item.Descriptionascsv)
	// }

	var dynamicAttrs []openapi.CreateDynamicAttributesPayloadInner

	for _, attr := range plan.Dynamicattributes {
		dynamicAttr := openapi.NewCreateDynamicAttributesPayloadInner(
			attr.Attributename.ValueString(),
			attr.Requesttype.ValueString(),
		)

		dynamicAttr.Attributetype = util.StringPointerOrEmpty(attr.Attributetype)
		dynamicAttr.Attributegroup = util.StringPointerOrEmpty(attr.Attributegroup)
		dynamicAttr.Orderindex = util.StringPointerOrEmpty(attr.Orderindex)
		dynamicAttr.Attributelable = util.StringPointerOrEmpty(attr.Attributelable)
		dynamicAttr.Accountscolumn = util.StringPointerOrEmpty(attr.Accountscolumn)
		dynamicAttr.Hideoncreate = util.StringPointerOrEmpty(attr.Hideoncreate)
		dynamicAttr.Actionstring = util.StringPointerOrEmpty(attr.Actionstring)
		dynamicAttr.Editable = util.StringPointerOrEmpty(attr.Editable)
		dynamicAttr.Hideonupdate = util.StringPointerOrEmpty(attr.Hideonupdate)
		dynamicAttr.Actiontoperformwhenparentattributechanges = util.StringPointerOrEmpty(attr.Actiontoperformwhenparentattributechanges)
		dynamicAttr.Defaultvalue = util.StringPointerOrEmpty(attr.Defaultvalue)
		dynamicAttr.Required = util.StringPointerOrEmpty(attr.Required)
		dynamicAttr.Regex = util.StringPointerOrEmpty(attr.Regex)
		dynamicAttr.Attributevalue = util.StringPointerOrEmpty(attr.Attributevalue)
		dynamicAttr.Showonchild = util.StringPointerOrEmpty(attr.Showonchild)
		dynamicAttr.Parentattribute = util.StringPointerOrEmpty(attr.Parentattribute)
		dynamicAttr.Descriptionascsv = util.StringPointerOrEmpty(attr.Descriptionascsv)
		// log.Print("[DEBUG] Dynamic Attribute: %+v", *dynamicAttr)
		dynamicAttrs = append(dynamicAttrs, *dynamicAttr)
		// fetchedAttrs = append(fetchedAttrs, dynamicAttr)
	}
	fetchedAttrs=append(fetchedAttrs, plan.Dynamicattributes...)

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
	}
	if createResp.Errorcode == nil {
		resp.Diagnostics.AddError(
			"Unexpected API Response",
			"Errorcode is nil in create response",
		)
		return
	}
	if *createResp.Errorcode != "0" {
		log.Printf("Error Creating Dynamic attribute: %v", *createResp.Errorcode)

		resp.Diagnostics.AddError(
			"Error Creating Dynamic Attribute",
			fmt.Sprintf("Error Code: %s\nMessage: %s", *createResp.Errorcode, *createResp.Msg),
		)

		log.Printf("[DEBUG] Create HTTP Status Code: %d", httpResp.StatusCode)
		log.Printf("[DEBUG] Create API Response: %+v", createResp)
		return
	}

	plan.ID = types.StringValue("dynamic-attr-" + plan.Endpoint.ValueString())
	// plan.Securitysystem = util.SafeString(plan.Securitysystem.ValueStringPointer())
	// plan.Endpoint = util.SafeString(plan.Endpoint.ValueStringPointer())
	// plan.Username = util.SafeString(plan.Username.ValueStringPointer())
	plan.Updateuser = util.SafeString(plan.Updateuser.ValueStringPointer())

	// if len(fetchedAttrs) > 0{
	// 	for i := range fetchedAttrs {
	// 		plan.Dynamicattributes[i].Attributename = util.SafeString(fetchedAttrs[i].Attributename.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Requesttype = util.SafeString(fetchedAttrs[i].Requesttype.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Attributetype = util.SafeString(fetchedAttrs[i].Attributetype.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Attributegroup = util.SafeString(fetchedAttrs[i].Attributegroup.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Orderindex = util.SafeString(fetchedAttrs[i].Orderindex.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Attributelable = util.SafeString(fetchedAttrs[i].Attributelable.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Accountscolumn = util.SafeString(fetchedAttrs[i].Accountscolumn.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Hideoncreate = util.SafeString(fetchedAttrs[i].Hideoncreate.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Actionstring = util.SafeString(fetchedAttrs[i].Actionstring.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Editable = util.SafeString(fetchedAttrs[i].Editable.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Hideonupdate = util.SafeString(fetchedAttrs[i].Hideonupdate.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Actiontoperformwhenparentattributechanges = util.SafeString(fetchedAttrs[i].Actiontoperformwhenparentattributechanges.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Defaultvalue = util.SafeString(fetchedAttrs[i].Defaultvalue.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Required = util.SafeString(fetchedAttrs[i].Required.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Regex = util.SafeString(fetchedAttrs[i].Regex.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Attributevalue = util.SafeString(fetchedAttrs[i].Attributevalue.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Showonchild = util.SafeString(fetchedAttrs[i].Showonchild.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Parentattribute = util.SafeString(fetchedAttrs[i].Parentattribute.ValueStringPointer())
	// 		plan.Dynamicattributes[i].Descriptionascsv = util.SafeString(fetchedAttrs[i].Descriptionascsv.ValueStringPointer())
	// 	}
	// }
	plan.Dynamicattributes=fetchedAttrs
	// for i := range plan.Dynamicattributes {
	// 	plan.Dynamicattributes[i].Attributename = util.SafeString(plan.Dynamicattributes[i].Attributename.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Requesttype = util.SafeString(plan.Dynamicattributes[i].Requesttype.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Attributetype = util.SafeString(plan.Dynamicattributes[i].Attributetype.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Attributegroup = util.SafeString(plan.Dynamicattributes[i].Attributegroup.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Orderindex = util.SafeString(plan.Dynamicattributes[i].Orderindex.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Attributelable = util.SafeString(plan.Dynamicattributes[i].Attributelable.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Accountscolumn = util.SafeString(plan.Dynamicattributes[i].Accountscolumn.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Hideoncreate = util.SafeString(plan.Dynamicattributes[i].Hideoncreate.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Actionstring = util.SafeString(plan.Dynamicattributes[i].Actionstring.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Editable = util.SafeString(plan.Dynamicattributes[i].Editable.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Hideonupdate = util.SafeString(plan.Dynamicattributes[i].Hideonupdate.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Actiontoperformwhenparentattributechanges = util.SafeString(plan.Dynamicattributes[i].Actiontoperformwhenparentattributechanges.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Defaultvalue = util.SafeString(plan.Dynamicattributes[i].Defaultvalue.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Required = util.SafeString(plan.Dynamicattributes[i].Required.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Regex = util.SafeString(plan.Dynamicattributes[i].Regex.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Attributevalue = util.SafeString(plan.Dynamicattributes[i].Attributevalue.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Showonchild = util.SafeString(plan.Dynamicattributes[i].Showonchild.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Parentattribute = util.SafeString(plan.Dynamicattributes[i].Parentattribute.ValueStringPointer())
	// 	plan.Dynamicattributes[i].Descriptionascsv = util.SafeString(plan.Dynamicattributes[i].Descriptionascsv.ValueStringPointer())
	// }

	msgValue := util.SafeDeref(createResp.Msg)
	errorCodeValue := util.SafeDeref(createResp.Errorcode)

	// Set the individual fields
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)

	stateCreateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateCreateDiagnostics...)
}

// func (r *dynamicAttributeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
//     var plan DynamicAttributeResourceModel

//     // Read Terraform plan into 'plan' struct
//     diags := req.Plan.Get(ctx, &plan)
//     resp.Diagnostics.Append(diags...)
//     if resp.Diagnostics.HasError() {
//         return
//     }

//     // Setup API client configuration
//     cfg := openapi.NewConfiguration()
//     apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
//     cfg.Host = apiBaseURL
//     cfg.Scheme = "https"
//     cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
//     cfg.HTTPClient = http.DefaultClient

//     apiClient := openapi.NewAPIClient(cfg)

//     // Optional: Fetch existing dynamic attributes (if needed)
//     fetchReq := apiClient.DynamicAttributesAPI.
//         FetchDynamicAttribute(ctx).
//         Securitysystem([]string{plan.Securitysystem.ValueString()}).
//         Endpoint([]string{plan.Endpoint.ValueString()})

//    apiResp, _, err := fetchReq.Execute()
// 	if err != nil {
// 		resp.Diagnostics.AddError("Fetch API Call Failed", err.Error())
// 		return
// 	}

// 	fetchedAttrs:=apiResp.Dynamicattributes

// 	// Map API response to fetchedAttrs slice (assumed done already)
// 	// var fetchedAttrs []Dynamicattribute
// 	// populate fetchedAttrs from apiResp.Dynamicattributes

// 	objs := make([]attr.Value, 0, len(apiResp.Dynamicattributes))

// 	// Get ElementType once and assert to ObjectType
// 	// elemType, ok := plan.Dynamicattributes.ElementType().(types.ObjectType)
// 	// if !ok {
// 	// 	resp.Diagnostics.AddError("Type Error", "Dynamicattributes element type is not ObjectType")
// 	// 	return
// 	// }

// 	for _, attr := range fetchedAttrs {
// 		objVal, diag := types.ObjectValue(
// 			elemType.AttrTypes,
// 			map[string]attr.Value{
// 				"attribute_name":                              attr.Attributename,
// 				"request_type":                               attr.Requesttype,
// 				"attribute_type":                             attr.Attributetype,
// 				"attribute_group":                            attr.Attributegroup,
// 				"order_index":                                attr.Orderindex,
// 				"attribute_lable":                            attr.Attributelable,
// 				"accounts_column":                            attr.Accountscolumn,
// 				"hide_on_create":                             attr.Hideoncreate,
// 				"action_string":                              attr.Actionstring,
// 				"editable":                                   attr.Editable,
// 				"hide_on_update":                             attr.Hideonupdate,
// 				"actiontoperformwhenparentattributechanges": attr.Actiontoperformwhenparentattributechanges,
// 				"default_value":                              attr.Defaultvalue,
// 				"required":                                   attr.Required,
// 				"regex":                                      attr.Regex,
// 				"attribute_value":                            attr.Attributevalue,
// 				"showonchild":                                attr.Showonchild,
// 				"parent_attribute":                           attr.Parentattribute,
// 				"description_as_csv":                         attr.Descriptionascsv,
// 			},
// 		)
// 		resp.Diagnostics.Append(diag...)
// 		if diag.HasError() {
// 			return
// 		}
// 		objs = append(objs, objVal)
// 	}

// 	setVal, diag := types.SetValue(elemType, objs)
// 	resp.Diagnostics.Append(diag...)
// 	if diag.HasError() {
// 		return
// 	}

// 	// Set the set value to the plan state
// 	plan.Dynamicattributes = setVal

// 	// Now build the API payload from the input config attrsSlice
// 	var createDynamicAttrs []openapi.CreateDynamicAttributesPayloadInner
// 	for _, attr := range attrsSlice {
// 		dynamicAttr := openapi.NewCreateDynamicAttributesPayloadInner(
// 			attr.Attributename.ValueString(),
// 			attr.Requesttype.ValueString(),
// 		)

// 		// Set optional fields safely using helper util converting types.String to *string or nil
// 		dynamicAttr.Attributetype = util.StringPointerOrEmpty(attr.Attributetype)
// 		dynamicAttr.Attributegroup = util.StringPointerOrEmpty(attr.Attributegroup)
// 		dynamicAttr.Orderindex = util.StringPointerOrEmpty(attr.Orderindex)
// 		dynamicAttr.Attributelable = util.StringPointerOrEmpty(attr.Attributelable)
// 		dynamicAttr.Accountscolumn = util.StringPointerOrEmpty(attr.Accountscolumn)
// 		dynamicAttr.Hideoncreate = util.StringPointerOrEmpty(attr.Hideoncreate)
// 		dynamicAttr.Actionstring = util.StringPointerOrEmpty(attr.Actionstring)
// 		dynamicAttr.Editable = util.StringPointerOrEmpty(attr.Editable)
// 		dynamicAttr.Hideonupdate = util.StringPointerOrEmpty(attr.Hideonupdate)
// 		dynamicAttr.Actiontoperformwhenparentattributechanges = util.StringPointerOrEmpty(attr.Actiontoperformwhenparentattributechanges)
// 		dynamicAttr.Defaultvalue = util.StringPointerOrEmpty(attr.Defaultvalue)
// 		dynamicAttr.Required = util.StringPointerOrEmpty(attr.Required)
// 		dynamicAttr.Regex = util.StringPointerOrEmpty(attr.Regex)
// 		dynamicAttr.Attributevalue = util.StringPointerOrEmpty(attr.Attributevalue)
// 		dynamicAttr.Showonchild = util.StringPointerOrEmpty(attr.Showonchild)
// 		dynamicAttr.Parentattribute = util.StringPointerOrEmpty(attr.Parentattribute)
// 		dynamicAttr.Descriptionascsv = util.StringPointerOrEmpty(attr.Descriptionascsv)

// 		createDynamicAttrs = append(createDynamicAttrs, *dynamicAttr)
// 	}

// 	// Build the create request with required fields and the payload
// 	createReq := openapi.NewCreateDynamicAttributeRequest(
// 		plan.Securitysystem.ValueString(),
// 		plan.Endpoint.ValueString(),
// 		plan.Username.ValueString(),
// 		createDynamicAttrs,
// 	)


//     // Call the API to create dynamic attributes
//     createResp, _, err := apiClient.DynamicAttributesAPI.
//         CreateDynamicAttribute(ctx).
//         CreateDynamicAttributeRequest(*createReq).
//         Execute()

//     if err != nil {
//         resp.Diagnostics.AddError("Error Creating Dynamic Attribute", err.Error())
//         return
//     }

//     if createResp.Errorcode == nil || *createResp.Errorcode != "0" {
//         resp.Diagnostics.AddError(
//             "Error Creating Dynamic Attribute",
//             fmt.Sprintf("Error Code: %v, Message: %v", createResp.Errorcode, createResp.Msg),
//         )
//         return
//     }

//     // Set the resource ID and update user (optional)
//     plan.ID = types.StringValue("dynamic-attr-" + plan.Endpoint.ValueString())
//     plan.Updateuser = util.SafeString(plan.Updateuser.ValueStringPointer())

//     // Save the final state
//     diags = resp.State.Set(ctx, plan)
//     resp.Diagnostics.Append(diags...)
// }

func (r *dynamicAttributeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DynamicAttributeResourceModel

	// Retrieve the current state
	diags := req.State.Get(ctx, &state)
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

	apiClient := openapi.NewAPIClient(cfg)

	fetchReq := apiClient.DynamicAttributesAPI.
		FetchDynamicAttribute(ctx).
		Securitysystem([]string{state.Securitysystem.ValueString()}).
		Endpoint([]string{state.Endpoint.ValueString()})

	fecthResp, httpResp, err := fetchReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Fetch API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	if httpResp.StatusCode == 404 || len(fecthResp.Dynamicattributes) == 0 {
		// Resource not found - remove from state
		resp.State.RemoveResource(ctx)
		return
	}

	apiAttrsMap := make(map[string]openapi.FetchDynamicAttributesPayloadInner)

	for _, attr := range fecthResp.Dynamicattributes {
		if attr.Attributename != nil {
			apiAttrsMap[*attr.Attributename] = attr
		}
	}

	var newStateAttrs []Dynamicattribute

// 2a. Retain attributes from state if they still exist in API
	existingNames := map[string]bool{}
	for _, oldAttr := range state.Dynamicattributes {
		name := oldAttr.Attributename.ValueString()
		if apiAttr, exists := apiAttrsMap[name]; exists {
			existingNames[name] = true // mark as handled

			newStateAttrs = append(newStateAttrs, Dynamicattribute{
				Attributename:                              types.StringValue(util.SafeDeref(apiAttr.Attributename)),
				Requesttype:                                types.StringValue(util.SafeDeref(apiAttr.Requesttype)),
				Attributetype:                              types.StringValue(util.SafeDeref(apiAttr.Attributetype)),
				Attributegroup:                             types.StringValue(util.SafeDeref(apiAttr.Attributegroup)),
				Orderindex:                                 types.StringValue(util.SafeDeref(apiAttr.Orderindex)),
				Attributelable:                             types.StringValue(util.SafeDeref(apiAttr.Attributelable)),
				Accountscolumn:                             types.StringValue(util.SafeDeref(apiAttr.Accountscolumn)),
				Hideoncreate:                               types.StringValue(util.SafeDeref(apiAttr.Hideoncreate)),
				Actionstring:                               types.StringValue(util.SafeDeref(apiAttr.Actionstring)),
				Editable:                                   types.StringValue(util.SafeDeref(apiAttr.Editable)),
				Hideonupdate:                               types.StringValue(util.SafeDeref(apiAttr.Hideonupdate)),
				Actiontoperformwhenparentattributechanges:  types.StringValue(util.SafeDeref(apiAttr.Actiontoperformwhenparentattributechanges)),
				Defaultvalue:                               types.StringValue(util.SafeDeref(apiAttr.Defaultvalue)),
				Required:                                   types.StringValue(util.SafeDeref(apiAttr.Required)),
				Regex:                                      types.StringValue(util.SafeDeref(apiAttr.Regex)),
				Attributevalue:                             types.StringValue(util.SafeDeref(apiAttr.Attributevalue)),
				Showonchild:                                types.StringValue(util.SafeDeref(apiAttr.Showonchild)),
				Parentattribute:                            types.StringValue(util.SafeDeref(apiAttr.Parentattribute)),
				Descriptionascsv:                           types.StringValue(util.SafeDeref(apiAttr.Descriptionascsv)),
			})
		}
	}

	for name, apiAttr := range apiAttrsMap {
	if _, handled := existingNames[name]; handled {
		continue
	}

		newStateAttrs = append(newStateAttrs, Dynamicattribute{
			Attributename:                              types.StringValue(util.SafeDeref(apiAttr.Attributename)),
			Requesttype:                                types.StringValue(util.SafeDeref(apiAttr.Requesttype)),
			Attributetype:                              types.StringValue(util.SafeDeref(apiAttr.Attributetype)),
			Attributegroup:                             types.StringValue(util.SafeDeref(apiAttr.Attributegroup)),
			Orderindex:                                 types.StringValue(util.SafeDeref(apiAttr.Orderindex)),
			Attributelable:                             types.StringValue(util.SafeDeref(apiAttr.Attributelable)),
			Accountscolumn:                             types.StringValue(util.SafeDeref(apiAttr.Accountscolumn)),
			Hideoncreate:                               types.StringValue(util.SafeDeref(apiAttr.Hideoncreate)),
			Actionstring:                               types.StringValue(util.SafeDeref(apiAttr.Actionstring)),
			Editable:                                   types.StringValue(util.SafeDeref(apiAttr.Editable)),
			Hideonupdate:                               types.StringValue(util.SafeDeref(apiAttr.Hideonupdate)),
			Actiontoperformwhenparentattributechanges:  types.StringValue(util.SafeDeref(apiAttr.Actiontoperformwhenparentattributechanges)),
			Defaultvalue:                               types.StringValue(util.SafeDeref(apiAttr.Defaultvalue)),
			Required:                                   types.StringValue(util.SafeDeref(apiAttr.Required)),
			Regex:                                      types.StringValue(util.SafeDeref(apiAttr.Regex)),
			Attributevalue:                             types.StringValue(util.SafeDeref(apiAttr.Attributevalue)),
			Showonchild:                                types.StringValue(util.SafeDeref(apiAttr.Showonchild)),
			Parentattribute:                            types.StringValue(util.SafeDeref(apiAttr.Parentattribute)),
			Descriptionascsv:                           types.StringValue(util.SafeDeref(apiAttr.Descriptionascsv)),
		})
	}

	state.Dynamicattributes = newStateAttrs

	// Map API response items into Go struct slice to convert into types.Set
	// var dynamicAttrsSlice []Dynamicattribute
	// for _, item := range apiResp.Dynamicattributes {
	// 	dynamicAttrsSlice = append(dynamicAttrsSlice, Dynamicattribute{
	// 		Attributename: util.SafeStringDatasource(item.Attributename),
	// 		Requesttype: util.SafeStringDatasource(item.Requesttype),
	// 		Attributetype: util.SafeStringDatasource(item.Attributetype),
	// 		Attributegroup: util.SafeStringDatasource(item.Attributegroup),
	// 		Orderindex: util.SafeStringDatasource(item.Orderindex),
	// 		Attributelable: util.SafeStringDatasource(item.Attributelable),
	// 		Accountscolumn: util.SafeStringDatasource(item.Accountscolumn),
	// 		Hideoncreate: util.SafeStringDatasource(item.Hideoncreate),
	// 		Actionstring: util.SafeStringDatasource(item.Actionstring),
	// 		Editable: util.SafeStringDatasource(item.Editable),
	// 		Hideonupdate: util.SafeStringDatasource(item.Hideonupdate),
	// 		Actiontoperformwhenparentattributechanges: util.SafeStringDatasource(item.Actiontoperformwhenparentattributechanges),
	// 		Defaultvalue: util.SafeStringDatasource(item.Defaultvalue),
	// 		Required: util.SafeStringDatasource(item.Required),
	// 		Regex: util.SafeStringDatasource(item.Regex),
	// 		Attributevalue: util.SafeStringDatasource(item.Attributevalue),
	// 		Showonchild: util.SafeStringDatasource(item.Showonchild),
	// 		Parentattribute: util.SafeStringDatasource(item.Parentattribute),
	// 		Descriptionascsv: util.SafeStringDatasource(item.Descriptionascsv),
	// 	})
	// }

	// // Convert []Dynamicattribute into types.Set and assign to state.Dynamicattributes
	// diags = state.Dynamicattributes.From(ctx, dynamicAttrsSlice)
	// resp.Diagnostics.Append(diags...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// Set other fields, keeping their values safe
	state.ID = types.StringValue("dynamic-attr-" + state.Endpoint.ValueString())
	state.Endpoint = util.SafeString(state.Endpoint.ValueStringPointer())
	state.Securitysystem = util.SafeString(state.Securitysystem.ValueStringPointer())
	state.Username = util.SafeString(state.Username.ValueStringPointer())
	state.Updateuser = util.SafeString(state.Updateuser.ValueStringPointer())

	state.Msg = types.StringValue(util.SafeDeref(fecthResp.Msg))
	state.ErrorCode = types.StringValue(util.SafeDeref(fecthResp.Errorcode))

	// Update state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}


// func (r *dynamicAttributeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
// 	var state DynamicAttributeResourceModel

// 	// Retrieve the current state
// 	stateRetrievalDiagnostics := req.State.Get(ctx, &state)
// 	resp.Diagnostics.Append(stateRetrievalDiagnostics...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	cfg := openapi.NewConfiguration()
// 	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
// 	cfg.Host = apiBaseURL
// 	cfg.Scheme = "https"
// 	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
// 	cfg.HTTPClient = http.DefaultClient

// 	apiClient := openapi.NewAPIClient(cfg)

// 	fetchReq := apiClient.DynamicAttributesAPI.
// 		FetchDynamicAttribute(ctx).
// 		Securitysystem([]string{state.Securitysystem.ValueString()}).
// 		Endpoint([]string{state.Endpoint.ValueString()})
// 		// Dynamicattributes([]string{state.Dynamicattributes.Attributename.ValueString()})
// 	// fetchReq:=apiClient.DynamicAttributesAPI.FetchDynamicAttribute(ctx).Dynamicattributes([]string{state.Dynamicattributes.Attributename.ValueString()})

// 	apiResp, httpResp, err := fetchReq.Execute()

// 	if err != nil {
// 		log.Printf("[ERROR] Fetch API Call Failed: %v", err)
// 		resp.Diagnostics.AddError("Fetch API Call Failed", fmt.Sprintf("Error: %v", err))
// 		return
// 	}
// 	log.Printf("[DEBUG] Read HTTP Status Code: %d", httpResp.StatusCode)
// 	log.Printf("[DEBUG] Read API Error code: %+v", *apiResp.Errorcode)
// 	log.Printf("[DEBUG] Read API Message: %+v", *apiResp.Msg)

// 	if len(apiResp.Dynamicattributes) == 0 {
// 		resp.State.RemoveResource(ctx)
// 		return
// 	}

// 	state.ID = types.StringValue("dynamic-attr-" + state.Endpoint.ValueString())
// 	state.Endpoint = util.SafeString(state.Endpoint.ValueStringPointer())
// 	state.Securitysystem = util.SafeString(state.Securitysystem.ValueStringPointer())
// 	state.Username = util.SafeString(state.Username.ValueStringPointer())
// 	state.Updateuser = util.SafeString(state.Updateuser.ValueStringPointer())

// 	state.Dynamicattributes = make([]Dynamicattribute, len(apiResp.Dynamicattributes))

// 	for i, item := range apiResp.Dynamicattributes {
// 		state.Dynamicattributes[i].Attributename = util.SafeStringDatasource(item.Attributename)
// 		state.Dynamicattributes[i].Requesttype = util.SafeStringDatasource(item.Requesttype)
// 		state.Dynamicattributes[i].Attributetype = util.SafeStringDatasource(item.Attributetype)
// 		state.Dynamicattributes[i].Attributegroup = util.SafeStringDatasource(item.Attributegroup)
// 		state.Dynamicattributes[i].Orderindex = util.SafeStringDatasource(item.Orderindex)
// 		state.Dynamicattributes[i].Attributelable = util.SafeStringDatasource(item.Attributelable)
// 		state.Dynamicattributes[i].Accountscolumn = util.SafeStringDatasource(item.Accountscolumn)
// 		state.Dynamicattributes[i].Hideoncreate = util.SafeStringDatasource(item.Hideoncreate)
// 		state.Dynamicattributes[i].Actionstring = util.SafeStringDatasource(item.Actionstring)
// 		state.Dynamicattributes[i].Editable = util.SafeStringDatasource(item.Editable)
// 		state.Dynamicattributes[i].Hideonupdate = util.SafeStringDatasource(item.Hideonupdate)
// 		state.Dynamicattributes[i].Actiontoperformwhenparentattributechanges = util.SafeStringDatasource(item.Actiontoperformwhenparentattributechanges)
// 		state.Dynamicattributes[i].Defaultvalue = util.SafeStringDatasource(item.Defaultvalue)
// 		state.Dynamicattributes[i].Required = util.SafeStringDatasource(item.Required)
// 		state.Dynamicattributes[i].Regex = util.SafeStringDatasource(item.Regex)
// 		state.Dynamicattributes[i].Attributevalue = util.SafeStringDatasource(item.Attributevalue)
// 		state.Dynamicattributes[i].Showonchild = util.SafeStringDatasource(item.Showonchild)
// 		state.Dynamicattributes[i].Parentattribute = util.SafeStringDatasource(item.Parentattribute)
// 		state.Dynamicattributes[i].Descriptionascsv = util.SafeStringDatasource(item.Descriptionascsv)
// 	}

// 	msgValue := util.SafeDeref(apiResp.Msg)
// 	errorCodeValue := util.SafeDeref(apiResp.Errorcode)
// 	state.Msg = types.StringValue(msgValue)
// 	state.ErrorCode = types.StringValue(errorCodeValue)

// 	stateSetDiagnostics := resp.State.Set(ctx, &state)
// 	resp.Diagnostics.Append(stateSetDiagnostics...)
// }

func (r *dynamicAttributeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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
	// 	plan.Dynamicattributes.Attributename.ValueString(),

	var dynamicAttrs []openapi.UpdateDynamicAttributesPayloadInner

	for _, attr := range plan.Dynamicattributes {
		dynamicAttr := openapi.NewUpdateDynamicAttributesPayloadInner(
			attr.Attributename.ValueString(),
		)

		dynamicAttr.Requesttype = util.StringPointerOrEmpty(attr.Requesttype)
		dynamicAttr.Attributetype = util.StringPointerOrEmpty(attr.Attributetype)
		dynamicAttr.Attributegroup = util.StringPointerOrEmpty(attr.Attributegroup)
		dynamicAttr.Orderindex = util.StringPointerOrEmpty(attr.Orderindex)
		dynamicAttr.Attributelable = util.StringPointerOrEmpty(attr.Attributelable)
		dynamicAttr.Accountscolumn = util.StringPointerOrEmpty(attr.Accountscolumn)
		dynamicAttr.Hideoncreate = util.StringPointerOrEmpty(attr.Hideoncreate)
		dynamicAttr.Actionstring = util.StringPointerOrEmpty(attr.Actionstring)
		dynamicAttr.Editable = util.StringPointerOrEmpty(attr.Editable)
		dynamicAttr.Hideonupdate = util.StringPointerOrEmpty(attr.Hideonupdate)
		dynamicAttr.Actiontoperformwhenparentattributechanges = util.StringPointerOrEmpty(attr.Actiontoperformwhenparentattributechanges)
		dynamicAttr.Defaultvalue = util.StringPointerOrEmpty(attr.Defaultvalue)
		dynamicAttr.Required = util.StringPointerOrEmpty(attr.Required)
		dynamicAttr.Regex = util.StringPointerOrEmpty(attr.Regex)
		dynamicAttr.Attributevalue = util.StringPointerOrEmpty(attr.Attributevalue)
		dynamicAttr.Showonchild = util.StringPointerOrEmpty(attr.Showonchild)
		dynamicAttr.Parentattribute = util.StringPointerOrEmpty(attr.Parentattribute)
		dynamicAttr.Descriptionascsv = util.StringPointerOrEmpty(attr.Descriptionascsv)

		dynamicAttrs = append(dynamicAttrs, *dynamicAttr)
	}

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

	if *updateResp.Errorcode != "0" {
		log.Printf("Error Updating Dynamic attribute: %v", *updateResp.Errorcode)

		// Print the entire response body for debugging
		respJSON, err := json.Marshal(updateResp)
		if err != nil {
			log.Printf("Error marshaling response: %v", err)
			resp.Diagnostics.AddError(
				"Error Updating Dynamic Attribute",
				"Failed to parse API response.",
			)
			return
		}

		// Log the entire response as a JSON string
		log.Printf("[DEBUG] Full API Response: %s", string(respJSON))

		resp.Diagnostics.AddError(
			"Error Updating Dynamic Attribute",
			fmt.Sprintf("Error Code: %s\nMessage: %s", *updateResp.Errorcode, *updateResp.Msg),
		)

		log.Printf("[DEBUG] Update HTTP Status Code: %d", httpResp.StatusCode)
		log.Printf("[DEBUG] Update API Response: %+v", updateResp)
		return
	}

	fetchReq := apiClient.DynamicAttributesAPI.
		FetchDynamicAttribute(ctx).
		Securitysystem([]string{plan.Securitysystem.ValueString()}).
		Endpoint([]string{plan.Endpoint.ValueString()})
		// Dynamicattributes([]string{plan.Dynamicattributes.Attributename.ValueString()})

	fetchResp, httpResp, err := fetchReq.Execute()

	if err != nil {
		log.Printf("[ERROR] Fetch API Call Failed: %v", err)
		resp.Diagnostics.AddError("Fetch API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Fetch API Response: %+v", fetchResp)

	if len(fetchResp.Dynamicattributes) == 0 {
		resp.State.RemoveResource(ctx)
		return
	}

	plan.ID = types.StringValue("dynamic-attr-" + plan.Endpoint.ValueString())
	plan.Endpoint = util.SafeString(plan.Endpoint.ValueStringPointer())
	plan.Securitysystem = util.SafeString(plan.Securitysystem.ValueStringPointer())
	plan.Username = util.SafeString(plan.Username.ValueStringPointer())
	plan.Updateuser = util.SafeString(plan.Updateuser.ValueStringPointer())

	plan.Dynamicattributes = make([]Dynamicattribute, len(fetchResp.Dynamicattributes))

	for i, item := range fetchResp.Dynamicattributes {
		plan.Dynamicattributes[i].Attributename = util.SafeStringDatasource(item.Attributename)
		plan.Dynamicattributes[i].Requesttype = util.SafeStringDatasource(item.Requesttype)
		plan.Dynamicattributes[i].Attributetype = util.SafeStringDatasource(item.Attributetype)
		plan.Dynamicattributes[i].Attributegroup = util.SafeStringDatasource(item.Attributegroup)
		plan.Dynamicattributes[i].Orderindex = util.SafeStringDatasource(item.Orderindex)
		plan.Dynamicattributes[i].Attributelable = util.SafeStringDatasource(item.Attributelable)
		plan.Dynamicattributes[i].Accountscolumn = util.SafeStringDatasource(item.Accountscolumn)
		plan.Dynamicattributes[i].Hideoncreate = util.SafeStringDatasource(item.Hideoncreate)
		plan.Dynamicattributes[i].Actionstring = util.SafeStringDatasource(item.Actionstring)
		plan.Dynamicattributes[i].Editable = util.SafeStringDatasource(item.Editable)
		plan.Dynamicattributes[i].Hideonupdate = util.SafeStringDatasource(item.Hideonupdate)
		plan.Dynamicattributes[i].Actiontoperformwhenparentattributechanges = util.SafeStringDatasource(item.Actiontoperformwhenparentattributechanges)
		plan.Dynamicattributes[i].Defaultvalue = util.SafeStringDatasource(item.Defaultvalue)
		plan.Dynamicattributes[i].Required = util.SafeStringDatasource(item.Required)
		plan.Dynamicattributes[i].Regex = util.SafeStringDatasource(item.Regex)
		plan.Dynamicattributes[i].Attributevalue = util.SafeStringDatasource(item.Attributevalue)
		plan.Dynamicattributes[i].Showonchild = util.SafeStringDatasource(item.Showonchild)
		plan.Dynamicattributes[i].Parentattribute = util.SafeStringDatasource(item.Parentattribute)
		plan.Dynamicattributes[i].Descriptionascsv = util.SafeStringDatasource(item.Descriptionascsv)
	}

	msgValue := util.SafeDeref(fetchResp.Msg)
	errorCodeValue := util.SafeDeref(fetchResp.Errorcode)

	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *dynamicAttributeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

	// Collect all attribute names from state.Dynamicattributes
	var attributeNames []string
	for _, attr := range state.Dynamicattributes {
		if !attr.Attributename.IsNull() && !attr.Attributename.IsUnknown() {
			attributeNames = append(attributeNames, attr.Attributename.ValueString())
		}
	}

	// Now pass it to the DeleteDynamicAttributeRequest
	deleteReq := openapi.DeleteDynamicAttributeRequest{
		Securitysystem:    state.Securitysystem.ValueString(),
		Endpoint:          state.Endpoint.ValueString(),
		Updateuser:        state.Updateuser.ValueString(),
		Dynamicattributes: attributeNames,
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

func (r *dynamicAttributeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	// resource.ImportStatePassthroughID(ctx, path.Root("dynamic_attribute").AtName("attribute_name"), req, resp)
	resource.ImportStatePassthroughID(ctx, path.Root("endpoint"), req, resp)
}
