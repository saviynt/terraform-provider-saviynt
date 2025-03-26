// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
)

// Ensure SaviyntProvider satisfies Terraform's provider interfaces.
var (
	_ provider.Provider = &saviyntProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &saviyntProvider{
			version: version,
		}
	}
}

// SaviyntProvider defines the provider implementation.
type saviyntProvider struct {
	version      string
	client       *s.Client // your Go client SDK instance
	accessToken  string
	refreshToken string
	expiresIn    int64
}

// SaviyntProviderModel describes the provider data model.
type SaviyntProviderModel struct {
	ServerURL types.String `tfsdk:"server_url"`
	Username  types.String `tfsdk:"username"`
	Password  types.String `tfsdk:"password"`
}

func (p *saviyntProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "saviynt"
	resp.Version = p.version
}

func (p *saviyntProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with the Saviynt ECM",
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				Required:    true,
				Description: "URL of Saviynt server.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Username for authentication.",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "Password for user authentication.",
			},
		},
	}
}

// Configure prepares a Saviynt API client for data sources and resources.
func (p *saviyntProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config SaviyntProviderModel

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config.ServerURL.IsUnknown() || config.ServerURL.IsNull() ||
		config.Username.IsUnknown() || config.Username.IsNull() ||
		config.Password.IsUnknown() || config.Password.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Configuration",
			"server_url, username, and password must be set.",
		)
		return
	}
	ctx = context.Background()

	serverURL := config.ServerURL.ValueString()
	if strings.HasPrefix(serverURL, "https://") {
		serverURL = strings.TrimPrefix(serverURL, "https://")
	}

	client, err := s.NewClient(ctx, s.Credentials{
		ServerURL: "https://" + serverURL,
		Username:  config.Username.ValueString(),
		Password:  config.Password.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create Saviynt client",
			"Could not initialize Saviynt API client: "+err.Error(),
		)
		return
	}
	token := client.Token()
	if token == nil {
		resp.Diagnostics.AddError("Token Error", "Failed to fetch access token.")
		return
	}

	// Store the token details in the provider struct.
	p.client = client
	p.accessToken = token.AccessToken
	p.refreshToken = token.RefreshToken
	p.expiresIn = token.ExpiresIn

	//Storing in Resource and Datasource
	resp.ResourceData = p
	resp.DataSourceData = p

}

// DataSources defines the data sources implemented in the provider.
func (p *saviyntProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSecuritySystemsDataSource,
		NewEndpointsDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *saviyntProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSecuritySystemResource,
		ADNewTestConnectionResource,
		RestNewTestConnectionResource,
		NewEndpointResource,
	}
}
