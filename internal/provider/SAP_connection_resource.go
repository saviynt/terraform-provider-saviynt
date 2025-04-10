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

type SAPConnectorResourceModel struct {
	BaseConnector
	ID                             types.String `tfsdk:"id"`
	Messageserver                  types.String `tfsdk:"message_server"`
	JcoAshost                      types.String `tfsdk:"jco_ashost"`
	JcoSysnr                       types.String `tfsdk:"jco_sysnr"`
	JcoClient                      types.String `tfsdk:"jco_client"`
	JcoUser                        types.String `tfsdk:"jco_user"`
	Password                       types.String `tfsdk:"password"`
	JcoLang                        types.String `tfsdk:"jco_lang"`
	JcoR3Name                      types.String `tfsdk:"jco_r3name"`
	JcoMshost                      types.String `tfsdk:"jco_mshost"`
	JcoMsserv                      types.String `tfsdk:"jco_msserv"`
	JcoGroup                       types.String `tfsdk:"jco_group"`
	Snc                            types.String `tfsdk:"snc"`
	JcoSncMode                     types.String `tfsdk:"jco_snc_mode"`
	JcoSncPartnername              types.String `tfsdk:"jco_snc_partnername"`
	JcoSncMyname                   types.String `tfsdk:"jco_snc_myname"`
	JcoSncLibrary                  types.String `tfsdk:"jco_snc_library"`
	JcoSncQop                      types.String `tfsdk:"jco_snc_qop"`
	Tables                         types.String `tfsdk:"tables"`
	Systemname                     types.String `tfsdk:"system_name"`
	Terminatedusergroup            types.String `tfsdk:"terminated_user_group"`
	TerminatedUserRoleAction       types.String `tfsdk:"terminated_user_role_action"`
	Createaccountjson              types.String `tfsdk:"create_account_json"`
	ProvJcoAshost                  types.String `tfsdk:"prov_jco_ashost"`
	ProvJcoSysnr                   types.String `tfsdk:"prov_jco_sysnr"`
	ProvJcoClient                  types.String `tfsdk:"prov_jco_client"`
	ProvJcoUser                    types.String `tfsdk:"prov_jco_user"`
	ProvPassword                   types.String `tfsdk:"prov_password"`
	ProvJcoLang                    types.String `tfsdk:"prov_jco_lang"`
	ProvJcoR3Name                  types.String `tfsdk:"prov_jco_r3name"`
	ProvJcoMshost                  types.String `tfsdk:"prov_jco_mshost"`
	ProvJcoMsserv                  types.String `tfsdk:"prov_jco_msserv"`
	ProvJcoGroup                   types.String `tfsdk:"prov_jco_group"`
	ProvCuaEnabled                 types.String `tfsdk:"prov_cua_enabled"`
	ProvCuaSnc                     types.String `tfsdk:"prov_cua_snc"`
	ResetPwdForNewaccount          types.String `tfsdk:"reset_pwd_for_newaccount"`
	Enforcepasswordchange          types.String `tfsdk:"enforce_password_change"`
	PasswordMinLength              types.String `tfsdk:"password_min_length"`
	PasswordMaxLength              types.String `tfsdk:"password_max_length"`
	PasswordNoofcapsalpha          types.String `tfsdk:"password_no_of_caps_alpha"`
	PasswordNoofdigits             types.String `tfsdk:"password_no_of_digits"`
	PasswordNoofsplchars           types.String `tfsdk:"password_no_of_spl_chars"`
	Hanareftablejson               types.String `tfsdk:"hanareftablejson"`
	Enableaccountjson              types.String `tfsdk:"enable_account_json"`
	Updateaccountjson              types.String `tfsdk:"update_account_json"`
	Userimportjson                 types.String `tfsdk:"user_import_json"`
	StatusThresholdConfig          types.String `tfsdk:"status_threshold_config"`
	Setcuasystem                   types.String `tfsdk:"set_cua_system"`
	FirefighteridGrantAccessJson   types.String `tfsdk:"fire_fighter_id_grant_access_json"`
	FirefighteridRevokeAccessJson  types.String `tfsdk:"fire_fighter_id_revoke_access_json"`
	Modifyuserdatajson             types.String `tfsdk:"modify_user_data_json"`
	ExternalSodEvalJson            types.String `tfsdk:"external_sod_eval_json"`
	ExternalSodEvalJsonDetail      types.String `tfsdk:"external_sod_eval_json_detail"`
	LogsTableFilter                types.String `tfsdk:"logs_table_filter"`
	PamConfig                      types.String `tfsdk:"pam_config"`
	SaptableFilterLang             types.String `tfsdk:"saptable_filter_lang"`
	AlternateOutputParameterEtData types.String `tfsdk:"alternate_output_parameter_et_data"`
	AuditLogJson                   types.String `tfsdk:"audit_log_json"`
	EccOrS4Hana                    types.String `tfsdk:"ecc_or_s4hana"`
	DataImportFilter               types.String `tfsdk:"data_import_filter"`
	Configjson                     types.String `tfsdk:"config_json"`
}

type sapConnectionResource struct {
	client *s.Client
	token  string
}

func SAPNewTestConnectionResource() resource.Resource {
	return &sapConnectionResource{}
}

func (r *sapConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_sap_connection_resource"
}

func (r *sapConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Description: "Description for the connection. Example: \"ORG_AD\"",
			},
			"defaultsavroles": schema.StringAttribute{
				Optional:    true,
				Description: "Default SAV roles for managing the connection. Example: \"ROLE_ORG\"",
			},
			"email_template": schema.StringAttribute{
				Optional:    true,
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
			"message_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Messageserver.",
			},
			"jco_ashost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcoashost.",
			},
			"jco_sysnr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcosysnr.",
			},
			"jco_client": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcoclient.",
			},
			"jco_user": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcouser.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Description: "Password.",
			},
			"jco_lang": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcolang.",
			},
			"jco_r3name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcor3name.",
			},
			"jco_mshost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcomshost.",
			},
			"jco_msserv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcomsserv.",
			},
			"jco_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcogroup.",
			},
			"snc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Snc.",
			},
			"jco_snc_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcosncmode.",
			},
			"jco_snc_partnername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcosncpartnername.",
			},
			"jco_snc_myname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcosncmyname.",
			},
			"jco_snc_library": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcosnclibrary.",
			},
			"jco_snc_qop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jcosncqop.",
			},
			"tables": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Tables.",
			},
			"system_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Systemname.",
			},
			"terminated_user_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Terminatedusergroup.",
			},
			"terminated_user_role_action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Terminateduserroleaction.",
			},
			"create_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Createaccountjson.",
			},
			"prov_jco_ashost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provjcoashost.",
			},
			"prov_jco_sysnr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provjcosysnr.",
			},
			"prov_jco_client": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provjcoclient.",
			},
			"prov_jco_user": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provjcouser.",
			},
			"prov_password": schema.StringAttribute{
				Optional:    true,
				Description: "Provpassword.",
			},
			"prov_jco_lang": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provjcolang.",
			},
			"prov_jco_r3name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provjcor3name.",
			},
			"prov_jco_mshost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provjcomshost.",
			},
			"prov_jco_msserv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provjcomsserv.",
			},
			"prov_jco_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provjcogroup.",
			},
			"prov_cua_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provcuaenabled.",
			},
			"prov_cua_snc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provcuasnc.",
			},
			"reset_pwd_for_newaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Resetpwdfornewaccount.",
			},
			"enforce_password_change": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enforcepasswordchange.",
			},
			"password_min_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Passwordminlength.",
			},
			"password_max_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Passwordmaxlength.",
			},
			"password_no_of_caps_alpha": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Passwordnoofcapsalpha.",
			},
			"password_no_of_digits": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Passwordnoofdigits.",
			},
			"password_no_of_spl_chars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Passwordnoofsplchars.",
			},
			"hanareftablejson": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hanareftablejson.",
			},
			"enable_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enableaccountjson.",
			},
			"update_account_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Updateaccountjson.",
			},
			"user_import_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Userimportjson.",
			},
			"status_threshold_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Statusthresholdconfig.",
			},
			"set_cua_system": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setcuasystem.",
			},
			"fire_fighter_id_grant_access_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Firefighteridgrantaccessjson.",
			},
			"fire_fighter_id_revoke_access_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Firefighteridrevokeaccessjson.",
			},
			"modify_user_data_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Modifyuserdatajson.",
			},
			"external_sod_eval_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Externalsodevaljson.",
			},
			"external_sod_eval_json_detail": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Externalsodevaljsondetail.",
			},
			"logs_table_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Logstablefilter.",
			},
			"pam_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pamconfig.",
			},
			"saptable_filter_lang": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Saptablefilterlang.",
			},
			"alternate_output_parameter_et_data": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Alternateoutputparameteretdata.",
			},
			"audit_log_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Auditlogjson.",
			},
			"ecc_or_s4hana": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Eccors4hana.",
			},
			"data_import_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Dataimportfilter.",
			},
			"config_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configjson.",
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
}

func (r *sapConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *sapConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SAPConnectorResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	sapConn := openapi.SAPConnector{
		BaseConnector: openapi.BaseConnector{
			//required field
			Connectiontype: "SAP",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional field
			Description:        util.StringPointerOrEmpty(plan.Description.ValueString()),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		//optional field
		MESSAGESERVER:                      util.StringPointerOrEmpty(plan.Messageserver.ValueString()),
		JCO_ASHOST:                         util.StringPointerOrEmpty(plan.JcoAshost.ValueString()),
		JCO_SYSNR:                          util.StringPointerOrEmpty(plan.JcoSysnr.ValueString()),
		JCO_CLIENT:                         util.StringPointerOrEmpty(plan.JcoClient.ValueString()),
		JCO_USER:                           util.StringPointerOrEmpty(plan.JcoUser.ValueString()),
		PASSWORD:                           util.StringPointerOrEmpty(plan.Password.ValueString()),
		JCO_LANG:                           util.StringPointerOrEmpty(plan.JcoLang.ValueString()),
		JCOR3NAME:                          util.StringPointerOrEmpty(plan.JcoR3Name.ValueString()),
		JCO_MSHOST:                         util.StringPointerOrEmpty(plan.JcoMshost.ValueString()),
		JCO_MSSERV:                         util.StringPointerOrEmpty(plan.JcoMsserv.ValueString()),
		JCO_GROUP:                          util.StringPointerOrEmpty(plan.JcoGroup.ValueString()),
		SNC:                                util.StringPointerOrEmpty(plan.Snc.ValueString()),
		JCO_SNC_MODE:                       util.StringPointerOrEmpty(plan.JcoSncMode.ValueString()),
		JCO_SNC_PARTNERNAME:                util.StringPointerOrEmpty(plan.JcoSncPartnername.ValueString()),
		JCO_SNC_MYNAME:                     util.StringPointerOrEmpty(plan.JcoSncMyname.ValueString()),
		JCO_SNC_LIBRARY:                    util.StringPointerOrEmpty(plan.JcoSncLibrary.ValueString()),
		JCO_SNC_QOP:                        util.StringPointerOrEmpty(plan.JcoSncQop.ValueString()),
		TABLES:                             util.StringPointerOrEmpty(plan.Tables.ValueString()),
		SYSTEMNAME:                         util.StringPointerOrEmpty(plan.Systemname.ValueString()),
		TERMINATEDUSERGROUP:                util.StringPointerOrEmpty(plan.Terminatedusergroup.ValueString()),
		TERMINATED_USER_ROLE_ACTION:        util.StringPointerOrEmpty(plan.TerminatedUserRoleAction.ValueString()),
		CREATEACCOUNTJSON:                  util.StringPointerOrEmpty(plan.Createaccountjson.ValueString()),
		PROV_JCO_ASHOST:                    util.StringPointerOrEmpty(plan.ProvJcoAshost.ValueString()),
		PROV_JCO_SYSNR:                     util.StringPointerOrEmpty(plan.ProvJcoSysnr.ValueString()),
		PROV_JCO_CLIENT:                    util.StringPointerOrEmpty(plan.ProvJcoClient.ValueString()),
		PROV_JCO_USER:                      util.StringPointerOrEmpty(plan.ProvJcoUser.ValueString()),
		PROV_PASSWORD:                      util.StringPointerOrEmpty(plan.ProvPassword.ValueString()),
		PROV_JCO_LANG:                      util.StringPointerOrEmpty(plan.ProvJcoLang.ValueString()),
		PROVJCOR3NAME:                      util.StringPointerOrEmpty(plan.ProvJcoR3Name.ValueString()),
		PROV_JCO_MSHOST:                    util.StringPointerOrEmpty(plan.ProvJcoMshost.ValueString()),
		PROV_JCO_MSSERV:                    util.StringPointerOrEmpty(plan.ProvJcoMsserv.ValueString()),
		PROV_JCO_GROUP:                     util.StringPointerOrEmpty(plan.ProvJcoGroup.ValueString()),
		PROV_CUA_ENABLED:                   util.StringPointerOrEmpty(plan.ProvCuaEnabled.ValueString()),
		PROV_CUA_SNC:                       util.StringPointerOrEmpty(plan.ProvCuaSnc.ValueString()),
		RESET_PWD_FOR_NEWACCOUNT:           util.StringPointerOrEmpty(plan.ResetPwdForNewaccount.ValueString()),
		ENFORCEPASSWORDCHANGE:              util.StringPointerOrEmpty(plan.Enforcepasswordchange.ValueString()),
		PASSWORD_MIN_LENGTH:                util.StringPointerOrEmpty(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:                util.StringPointerOrEmpty(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:             util.StringPointerOrEmpty(plan.PasswordNoofcapsalpha.ValueString()),
		PASSWORD_NOOFDIGITS:                util.StringPointerOrEmpty(plan.PasswordNoofdigits.ValueString()),
		PASSWORD_NOOFSPLCHARS:              util.StringPointerOrEmpty(plan.PasswordNoofsplchars.ValueString()),
		HANAREFTABLEJSON:                   util.StringPointerOrEmpty(plan.Hanareftablejson.ValueString()),
		ENABLEACCOUNTJSON:                  util.StringPointerOrEmpty(plan.Enableaccountjson.ValueString()),
		UPDATEACCOUNTJSON:                  util.StringPointerOrEmpty(plan.Updateaccountjson.ValueString()),
		USERIMPORTJSON:                     util.StringPointerOrEmpty(plan.Userimportjson.ValueString()),
		STATUS_THRESHOLD_CONFIG:            util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()),
		SETCUASYSTEM:                       util.StringPointerOrEmpty(plan.Setcuasystem.ValueString()),
		FIREFIGHTERID_GRANT_ACCESS_JSON:    util.StringPointerOrEmpty(plan.FirefighteridGrantAccessJson.ValueString()),
		FIREFIGHTERID_REVOKE_ACCESS_JSON:   util.StringPointerOrEmpty(plan.FirefighteridRevokeAccessJson.ValueString()),
		MODIFYUSERDATAJSON:                 util.StringPointerOrEmpty(plan.Modifyuserdatajson.ValueString()),
		EXTERNAL_SOD_EVAL_JSON:             util.StringPointerOrEmpty(plan.ExternalSodEvalJson.ValueString()),
		EXTERNAL_SOD_EVAL_JSON_DETAIL:      util.StringPointerOrEmpty(plan.ExternalSodEvalJsonDetail.ValueString()),
		LOGS_TABLE_FILTER:                  util.StringPointerOrEmpty(plan.LogsTableFilter.ValueString()),
		PAM_CONFIG:                         util.StringPointerOrEmpty(plan.PamConfig.ValueString()),
		SAPTABLE_FILTER_LANG:               util.StringPointerOrEmpty(plan.SaptableFilterLang.ValueString()),
		ALTERNATE_OUTPUT_PARAMETER_ET_DATA: util.StringPointerOrEmpty(plan.AlternateOutputParameterEtData.ValueString()),
		AUDIT_LOG_JSON:                     util.StringPointerOrEmpty(plan.AuditLogJson.ValueString()),
		ECCORS4HANA:                        util.StringPointerOrEmpty(plan.EccOrS4Hana.ValueString()),
		DATA_IMPORT_FILTER:                 util.StringPointerOrEmpty(plan.DataImportFilter.ValueString()),
		ConfigJSON:                         util.StringPointerOrEmpty(plan.Configjson.ValueString()),
	}
	sapConnRequest := openapi.CreateOrUpdateRequest{
		SAPConnector: &sapConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(sapConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.JcoAshost = types.StringValue(*util.StringPointerOrEmpty(plan.JcoAshost.ValueString()))
	plan.JcoSysnr = types.StringValue(*util.StringPointerOrEmpty(plan.JcoSysnr.ValueString()))
	plan.JcoClient = types.StringValue(*util.StringPointerOrEmpty(plan.JcoClient.ValueString()))
	plan.JcoUser = types.StringValue(*util.StringPointerOrEmpty(plan.JcoUser.ValueString()))
	plan.JcoLang = types.StringValue(*util.StringPointerOrEmpty(plan.JcoLang.ValueString()))
	plan.JcoR3Name = types.StringValue(*util.StringPointerOrEmpty(plan.JcoR3Name.ValueString()))
	plan.JcoMshost = types.StringValue(*util.StringPointerOrEmpty(plan.JcoMshost.ValueString()))
	plan.JcoMsserv = types.StringValue(*util.StringPointerOrEmpty(plan.JcoMsserv.ValueString()))
	plan.JcoGroup = types.StringValue(*util.StringPointerOrEmpty(plan.JcoGroup.ValueString()))
	plan.Snc = types.StringValue(*util.StringPointerOrEmpty(plan.Snc.ValueString()))
	plan.JcoSncMode = types.StringValue(*util.StringPointerOrEmpty(plan.JcoSncMode.ValueString()))
	plan.JcoSncPartnername = types.StringValue(*util.StringPointerOrEmpty(plan.JcoSncPartnername.ValueString()))
	plan.JcoSncMyname = types.StringValue(*util.StringPointerOrEmpty(plan.JcoSncMyname.ValueString()))
	plan.JcoSncLibrary = types.StringValue(*util.StringPointerOrEmpty(plan.JcoSncLibrary.ValueString()))
	plan.JcoSncQop = types.StringValue(*util.StringPointerOrEmpty(plan.JcoSncQop.ValueString()))
	plan.Tables = types.StringValue(*util.StringPointerOrEmpty(plan.Tables.ValueString()))
	plan.Systemname = types.StringValue(*util.StringPointerOrEmpty(plan.Systemname.ValueString()))
	plan.Terminatedusergroup = types.StringValue(*util.StringPointerOrEmpty(plan.Terminatedusergroup.ValueString()))
	plan.TerminatedUserRoleAction = types.StringValue(*util.StringPointerOrEmpty(plan.TerminatedUserRoleAction.ValueString()))
	plan.Createaccountjson = types.StringValue(*util.StringPointerOrEmpty(plan.Createaccountjson.ValueString()))
	plan.ProvJcoAshost = types.StringValue(*util.StringPointerOrEmpty(plan.ProvJcoAshost.ValueString()))
	plan.ProvJcoSysnr = types.StringValue(*util.StringPointerOrEmpty(plan.ProvJcoSysnr.ValueString()))
	plan.ProvJcoClient = types.StringValue(*util.StringPointerOrEmpty(plan.ProvJcoClient.ValueString()))
	plan.ProvJcoUser = types.StringValue(*util.StringPointerOrEmpty(plan.ProvJcoUser.ValueString()))
	plan.ProvJcoLang = types.StringValue(*util.StringPointerOrEmpty(plan.ProvJcoLang.ValueString()))
	plan.ProvJcoR3Name = types.StringValue(*util.StringPointerOrEmpty(plan.ProvJcoR3Name.ValueString()))
	plan.ProvJcoMshost = types.StringValue(*util.StringPointerOrEmpty(plan.ProvJcoMshost.ValueString()))
	plan.ProvJcoMsserv = types.StringValue(*util.StringPointerOrEmpty(plan.ProvJcoMsserv.ValueString()))
	plan.ProvJcoGroup = types.StringValue(*util.StringPointerOrEmpty(plan.ProvJcoGroup.ValueString()))
	plan.ProvCuaEnabled = types.StringValue(*util.StringPointerOrEmpty(plan.ProvCuaEnabled.ValueString()))
	plan.ProvCuaSnc = types.StringValue(*util.StringPointerOrEmpty(plan.ProvCuaSnc.ValueString()))
	plan.ResetPwdForNewaccount = types.StringValue(*util.StringPointerOrEmpty(plan.ResetPwdForNewaccount.ValueString()))
	plan.Enforcepasswordchange = types.StringValue(*util.StringPointerOrEmpty(plan.Enforcepasswordchange.ValueString()))
	plan.PasswordMinLength = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordMinLength.ValueString()))
	plan.PasswordMaxLength = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordMaxLength.ValueString()))
	plan.PasswordNoofcapsalpha = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordNoofcapsalpha.ValueString()))
	plan.PasswordNoofdigits = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordNoofdigits.ValueString()))
	plan.PasswordNoofsplchars = types.StringValue(*util.StringPointerOrEmpty(plan.PasswordNoofsplchars.ValueString()))
	plan.Hanareftablejson = types.StringValue(*util.StringPointerOrEmpty(plan.Hanareftablejson.ValueString()))
	plan.Enableaccountjson = types.StringValue(*util.StringPointerOrEmpty(plan.Enableaccountjson.ValueString()))
	plan.Updateaccountjson = types.StringValue(*util.StringPointerOrEmpty(plan.Updateaccountjson.ValueString()))
	plan.Userimportjson = types.StringValue(*util.StringPointerOrEmpty(plan.Userimportjson.ValueString()))
	plan.StatusThresholdConfig = types.StringValue(*util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()))
	plan.Setcuasystem = types.StringValue(*util.StringPointerOrEmpty(plan.Setcuasystem.ValueString()))
	plan.FirefighteridGrantAccessJson = types.StringValue(*util.StringPointerOrEmpty(plan.FirefighteridGrantAccessJson.ValueString()))
	plan.FirefighteridRevokeAccessJson = types.StringValue(*util.StringPointerOrEmpty(plan.FirefighteridRevokeAccessJson.ValueString()))
	plan.Modifyuserdatajson = types.StringValue(*util.StringPointerOrEmpty(plan.Modifyuserdatajson.ValueString()))
	plan.ExternalSodEvalJson = types.StringValue(*util.StringPointerOrEmpty(plan.ExternalSodEvalJson.ValueString()))
	plan.ExternalSodEvalJsonDetail = types.StringValue(*util.StringPointerOrEmpty(plan.ExternalSodEvalJsonDetail.ValueString()))
	plan.LogsTableFilter = types.StringValue(*util.StringPointerOrEmpty(plan.LogsTableFilter.ValueString()))
	plan.PamConfig = types.StringValue(*util.StringPointerOrEmpty(plan.PamConfig.ValueString()))
	plan.SaptableFilterLang = types.StringValue(*util.StringPointerOrEmpty(plan.SaptableFilterLang.ValueString()))
	plan.AlternateOutputParameterEtData = types.StringValue(*util.StringPointerOrEmpty(plan.AlternateOutputParameterEtData.ValueString()))
	plan.AuditLogJson = types.StringValue(*util.StringPointerOrEmpty(plan.AuditLogJson.ValueString()))
	plan.EccOrS4Hana = types.StringValue(*util.StringPointerOrEmpty(plan.EccOrS4Hana.ValueString()))
	plan.DataImportFilter = types.StringValue(*util.StringPointerOrEmpty(plan.DataImportFilter.ValueString()))
	plan.Configjson = types.StringValue(*util.StringPointerOrEmpty(plan.Configjson.ValueString()))
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	r.Read(ctx, resource.ReadRequest{State: resp.State}, &resource.ReadResponse{State: resp.State})
}

func (r *sapConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SAPConnectorResourceModel

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
	state.ConnectionKey = types.Int64Value(int64(*apiResp.SAPConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.SAPConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectiontype)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Emailtemplate)
	state.Createaccountjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	state.AuditLogJson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.AUDIT_LOG_JSON)
	state.SaptableFilterLang = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SAPTABLE_FILTER_LANG)
	state.PasswordNoofsplchars = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS)
	state.Terminatedusergroup = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.TERMINATEDUSERGROUP)
	state.LogsTableFilter = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.LOGS_TABLE_FILTER)
	state.EccOrS4Hana = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ECCORS4HANA)
	state.FirefighteridRevokeAccessJson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.FIREFIGHTERID_REVOKE_ACCESS_JSON)
	state.Configjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ConfigJSON)
	state.FirefighteridGrantAccessJson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.FIREFIGHTERID_GRANT_ACCESS_JSON)
	state.JcoSncLibrary = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_LIBRARY)
	state.JcoR3Name = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCOR3NAME)
	state.ExternalSodEvalJson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.EXTERNAL_SOD_EVAL_JSON)
	state.JcoAshost = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_ASHOST)
	state.PasswordNoofdigits = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS)
	state.ProvJcoMshost = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_MSHOST)
	state.PamConfig = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.JcoSncMyname = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_MYNAME)
	state.Enforcepasswordchange = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ENFORCEPASSWORDCHANGE)
	state.JcoUser = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_USER)
	state.JcoSncMode = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_MODE)
	state.ProvJcoMsserv = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_MSSERV)
	state.Hanareftablejson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.HANAREFTABLEJSON)
	state.PasswordMinLength = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH)
	state.JcoClient = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_CLIENT)
	state.TerminatedUserRoleAction = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.TERMINATED_USER_ROLE_ACTION)
	state.ResetPwdForNewaccount = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.RESET_PWD_FOR_NEWACCOUNT)
	state.ProvJcoClient = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_CLIENT)
	state.Snc = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SNC)
	state.JcoMsserv = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_MSSERV)
	state.ProvCuaSnc = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_CUA_SNC)
	state.ProvJcoUser = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_USER)
	state.JcoLang = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_LANG)
	state.JcoSncPartnername = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_PARTNERNAME)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.ProvJcoSysnr = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_SYSNR)
	state.Setcuasystem = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SETCUASYSTEM)
	state.Messageserver = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.MESSAGESERVER)
	state.ProvJcoAshost = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_ASHOST)
	state.ProvJcoGroup = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_GROUP)
	state.ProvCuaEnabled = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_CUA_ENABLED)
	state.JcoMshost = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_MSHOST)
	state.ProvJcoR3Name = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROVJCOR3NAME)
	state.PasswordNoofcapsalpha = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA)
	state.Modifyuserdatajson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.JcoSncQop = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_QOP)
	state.Tables = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.TABLES)
	state.ProvJcoLang = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_LANG)
	state.JcoSysnr = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SYSNR)
	state.ExternalSodEvalJsonDetail = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.EXTERNAL_SOD_EVAL_JSON_DETAIL)
	state.DataImportFilter = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.DATA_IMPORT_FILTER)
	state.Enableaccountjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	state.AlternateOutputParameterEtData = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ALTERNATE_OUTPUT_PARAMETER_ET_DATA)
	state.JcoGroup = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_GROUP)
	state.PasswordMaxLength = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH)
	state.Userimportjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.USERIMPORTJSON)
	state.Systemname = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SYSTEMNAME)
	state.Updateaccountjson = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	apiMessage := util.SafeDeref(apiResp.SAPConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.SAPConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sapConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SAPConnectorResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	sapConn := openapi.SAPConnector{
		BaseConnector: openapi.BaseConnector{
			//required field
			Connectiontype: "SAP",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional field
			Description:        util.StringPointerOrEmpty(plan.Description.ValueString()),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles.ValueString()),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate.ValueString()),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		//optional field
		MESSAGESERVER:                      util.StringPointerOrEmpty(plan.Messageserver.ValueString()),
		JCO_ASHOST:                         util.StringPointerOrEmpty(plan.JcoAshost.ValueString()),
		JCO_SYSNR:                          util.StringPointerOrEmpty(plan.JcoSysnr.ValueString()),
		JCO_CLIENT:                         util.StringPointerOrEmpty(plan.JcoClient.ValueString()),
		JCO_USER:                           util.StringPointerOrEmpty(plan.JcoUser.ValueString()),
		PASSWORD:                           util.StringPointerOrEmpty(plan.Password.ValueString()),
		JCO_LANG:                           util.StringPointerOrEmpty(plan.JcoLang.ValueString()),
		JCOR3NAME:                          util.StringPointerOrEmpty(plan.JcoR3Name.ValueString()),
		JCO_MSHOST:                         util.StringPointerOrEmpty(plan.JcoMshost.ValueString()),
		JCO_MSSERV:                         util.StringPointerOrEmpty(plan.JcoMsserv.ValueString()),
		JCO_GROUP:                          util.StringPointerOrEmpty(plan.JcoGroup.ValueString()),
		SNC:                                util.StringPointerOrEmpty(plan.Snc.ValueString()),
		JCO_SNC_MODE:                       util.StringPointerOrEmpty(plan.JcoSncMode.ValueString()),
		JCO_SNC_PARTNERNAME:                util.StringPointerOrEmpty(plan.JcoSncPartnername.ValueString()),
		JCO_SNC_MYNAME:                     util.StringPointerOrEmpty(plan.JcoSncMyname.ValueString()),
		JCO_SNC_LIBRARY:                    util.StringPointerOrEmpty(plan.JcoSncLibrary.ValueString()),
		JCO_SNC_QOP:                        util.StringPointerOrEmpty(plan.JcoSncQop.ValueString()),
		TABLES:                             util.StringPointerOrEmpty(plan.Tables.ValueString()),
		SYSTEMNAME:                         util.StringPointerOrEmpty(plan.Systemname.ValueString()),
		TERMINATEDUSERGROUP:                util.StringPointerOrEmpty(plan.Terminatedusergroup.ValueString()),
		TERMINATED_USER_ROLE_ACTION:        util.StringPointerOrEmpty(plan.TerminatedUserRoleAction.ValueString()),
		CREATEACCOUNTJSON:                  util.StringPointerOrEmpty(plan.Createaccountjson.ValueString()),
		PROV_JCO_ASHOST:                    util.StringPointerOrEmpty(plan.ProvJcoAshost.ValueString()),
		PROV_JCO_SYSNR:                     util.StringPointerOrEmpty(plan.ProvJcoSysnr.ValueString()),
		PROV_JCO_CLIENT:                    util.StringPointerOrEmpty(plan.ProvJcoClient.ValueString()),
		PROV_JCO_USER:                      util.StringPointerOrEmpty(plan.ProvJcoUser.ValueString()),
		PROV_PASSWORD:                      util.StringPointerOrEmpty(plan.ProvPassword.ValueString()),
		PROV_JCO_LANG:                      util.StringPointerOrEmpty(plan.ProvJcoLang.ValueString()),
		PROVJCOR3NAME:                      util.StringPointerOrEmpty(plan.ProvJcoR3Name.ValueString()),
		PROV_JCO_MSHOST:                    util.StringPointerOrEmpty(plan.ProvJcoMshost.ValueString()),
		PROV_JCO_MSSERV:                    util.StringPointerOrEmpty(plan.ProvJcoMsserv.ValueString()),
		PROV_JCO_GROUP:                     util.StringPointerOrEmpty(plan.ProvJcoGroup.ValueString()),
		PROV_CUA_ENABLED:                   util.StringPointerOrEmpty(plan.ProvCuaEnabled.ValueString()),
		PROV_CUA_SNC:                       util.StringPointerOrEmpty(plan.ProvCuaSnc.ValueString()),
		RESET_PWD_FOR_NEWACCOUNT:           util.StringPointerOrEmpty(plan.ResetPwdForNewaccount.ValueString()),
		ENFORCEPASSWORDCHANGE:              util.StringPointerOrEmpty(plan.Enforcepasswordchange.ValueString()),
		PASSWORD_MIN_LENGTH:                util.StringPointerOrEmpty(plan.PasswordMinLength.ValueString()),
		PASSWORD_MAX_LENGTH:                util.StringPointerOrEmpty(plan.PasswordMaxLength.ValueString()),
		PASSWORD_NOOFCAPSALPHA:             util.StringPointerOrEmpty(plan.PasswordNoofcapsalpha.ValueString()),
		PASSWORD_NOOFDIGITS:                util.StringPointerOrEmpty(plan.PasswordNoofdigits.ValueString()),
		PASSWORD_NOOFSPLCHARS:              util.StringPointerOrEmpty(plan.PasswordNoofsplchars.ValueString()),
		HANAREFTABLEJSON:                   util.StringPointerOrEmpty(plan.Hanareftablejson.ValueString()),
		ENABLEACCOUNTJSON:                  util.StringPointerOrEmpty(plan.Enableaccountjson.ValueString()),
		UPDATEACCOUNTJSON:                  util.StringPointerOrEmpty(plan.Updateaccountjson.ValueString()),
		USERIMPORTJSON:                     util.StringPointerOrEmpty(plan.Userimportjson.ValueString()),
		STATUS_THRESHOLD_CONFIG:            util.StringPointerOrEmpty(plan.StatusThresholdConfig.ValueString()),
		SETCUASYSTEM:                       util.StringPointerOrEmpty(plan.Setcuasystem.ValueString()),
		FIREFIGHTERID_GRANT_ACCESS_JSON:    util.StringPointerOrEmpty(plan.FirefighteridGrantAccessJson.ValueString()),
		FIREFIGHTERID_REVOKE_ACCESS_JSON:   util.StringPointerOrEmpty(plan.FirefighteridRevokeAccessJson.ValueString()),
		MODIFYUSERDATAJSON:                 util.StringPointerOrEmpty(plan.Modifyuserdatajson.ValueString()),
		EXTERNAL_SOD_EVAL_JSON:             util.StringPointerOrEmpty(plan.ExternalSodEvalJson.ValueString()),
		EXTERNAL_SOD_EVAL_JSON_DETAIL:      util.StringPointerOrEmpty(plan.ExternalSodEvalJsonDetail.ValueString()),
		LOGS_TABLE_FILTER:                  util.StringPointerOrEmpty(plan.LogsTableFilter.ValueString()),
		PAM_CONFIG:                         util.StringPointerOrEmpty(plan.PamConfig.ValueString()),
		SAPTABLE_FILTER_LANG:               util.StringPointerOrEmpty(plan.SaptableFilterLang.ValueString()),
		ALTERNATE_OUTPUT_PARAMETER_ET_DATA: util.StringPointerOrEmpty(plan.AlternateOutputParameterEtData.ValueString()),
		AUDIT_LOG_JSON:                     util.StringPointerOrEmpty(plan.AuditLogJson.ValueString()),
		ECCORS4HANA:                        util.StringPointerOrEmpty(plan.EccOrS4Hana.ValueString()),
		DATA_IMPORT_FILTER:                 util.StringPointerOrEmpty(plan.DataImportFilter.ValueString()),
		ConfigJSON:                         util.StringPointerOrEmpty(plan.Configjson.ValueString()),
	}
	sapConnRequest := openapi.CreateOrUpdateRequest{
		SAPConnector: &sapConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(sapConnRequest).Execute()
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
	plan.ConnectionKey = types.Int64Value(int64(*getResp.SAPConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.SAPConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionname)
	plan.Description = util.SafeStringDatasource(getResp.SAPConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.SAPConnectionResponse.Defaultsavroles)
	plan.ConnectionType = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectiontype)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.SAPConnectionResponse.Emailtemplate)
	plan.Createaccountjson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	plan.AuditLogJson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.AUDIT_LOG_JSON)
	plan.SaptableFilterLang = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.SAPTABLE_FILTER_LANG)
	plan.PasswordNoofsplchars = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS)
	plan.Terminatedusergroup = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.TERMINATEDUSERGROUP)
	plan.LogsTableFilter = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.LOGS_TABLE_FILTER)
	plan.EccOrS4Hana = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.ECCORS4HANA)
	plan.FirefighteridRevokeAccessJson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.FIREFIGHTERID_REVOKE_ACCESS_JSON)
	plan.Configjson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.ConfigJSON)
	plan.FirefighteridGrantAccessJson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.FIREFIGHTERID_GRANT_ACCESS_JSON)
	plan.JcoSncLibrary = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_LIBRARY)
	plan.JcoR3Name = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCOR3NAME)
	plan.ExternalSodEvalJson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.EXTERNAL_SOD_EVAL_JSON)
	plan.JcoAshost = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_ASHOST)
	plan.PasswordNoofdigits = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS)
	plan.ProvJcoMshost = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_MSHOST)
	plan.PamConfig = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PAM_CONFIG)
	plan.JcoSncMyname = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_MYNAME)
	plan.Enforcepasswordchange = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.ENFORCEPASSWORDCHANGE)
	plan.JcoUser = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_USER)
	plan.JcoSncMode = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_MODE)
	plan.ProvJcoMsserv = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_MSSERV)
	plan.Hanareftablejson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.HANAREFTABLEJSON)
	plan.PasswordMinLength = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH)
	plan.JcoClient = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_CLIENT)
	plan.TerminatedUserRoleAction = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.TERMINATED_USER_ROLE_ACTION)
	plan.ResetPwdForNewaccount = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.RESET_PWD_FOR_NEWACCOUNT)
	plan.ProvJcoClient = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_CLIENT)
	plan.Snc = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.SNC)
	plan.JcoMsserv = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_MSSERV)
	plan.ProvCuaSnc = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_CUA_SNC)
	plan.ProvJcoUser = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_USER)
	plan.JcoLang = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_LANG)
	plan.JcoSncPartnername = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_PARTNERNAME)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.ProvJcoSysnr = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_SYSNR)
	plan.Setcuasystem = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.SETCUASYSTEM)
	plan.Messageserver = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.MESSAGESERVER)
	plan.ProvJcoAshost = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_ASHOST)
	plan.ProvJcoGroup = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_GROUP)
	plan.ProvCuaEnabled = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_CUA_ENABLED)
	plan.JcoMshost = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_MSHOST)
	plan.ProvJcoR3Name = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROVJCOR3NAME)
	plan.PasswordNoofcapsalpha = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA)
	plan.Modifyuserdatajson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	plan.JcoSncQop = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_QOP)
	plan.Tables = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.TABLES)
	plan.ProvJcoLang = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_LANG)
	plan.JcoSysnr = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_SYSNR)
	plan.ExternalSodEvalJsonDetail = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.EXTERNAL_SOD_EVAL_JSON_DETAIL)
	plan.DataImportFilter = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.DATA_IMPORT_FILTER)
	plan.Enableaccountjson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	plan.AlternateOutputParameterEtData = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.ALTERNATE_OUTPUT_PARAMETER_ET_DATA)
	plan.JcoGroup = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.JCO_GROUP)
	plan.PasswordMaxLength = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH)
	plan.Userimportjson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.USERIMPORTJSON)
	plan.Systemname = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.SYSTEMNAME)
	plan.Updateaccountjson = util.SafeStringDatasource(getResp.SAPConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	apiMessage := util.SafeDeref(getResp.SAPConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.SAPConnectionResponse.Errorcode)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *sapConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
