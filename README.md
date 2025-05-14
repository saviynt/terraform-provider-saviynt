[![Release build for Saviynt Terraform Provider](https://github.com/saviynt/terraform-provider-saviynt/actions/workflows/release.yml/badge.svg)](https://github.com/saviynt/terraform-provider-saviynt/actions/workflows/release.yml)
<br/><br/>

<a href="https://terraform.io">
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="assets/hashicorp-terraform-dark.svg">
        <source media="(prefers-color-scheme: light)" srcset="assets/hashicorp-terraform-light.svg">
        <img alt="Terraform logo" title="Terraform" height="60" src="assets/hashicorp-terraform-dark.svg">
    </picture>
</a>

<a href="https://saviynt.com/">
    <img src="assets/s-platform-icon-01.svg" alt="Saviynt logo" title="Saviynt" height="75" />
</a>

# Terraform Provider for Saviynt

The Saviynt Terraform provider empowers you to leverage Terraform's declarative Infrastructure-as-Code (IaC) capabilities to provision, configure, and manage resources within the Saviynt Identity Cloud.

New to Terraform? Check out the [official Terraform introduction by HashiCorp](https://developer.hashicorp.com/terraform/intro) to get up to speed with the basics.

---

##  Requirements

- Terraform version `>= 1.8+`
- Go programming language `>= 1.21+`
- Saviynt Identity Cloud instance and credentials.

---

##  Documentation

Check out the [Latest Saviynt Provider Docs](https://registry.terraform.io/providers/saviynt/saviynt/latest/docs) to know more.


---

##  Examples

Examples are available for all resources. Follow the following steps to try out the examples

1. Uncomment the code block corresponding to the object for which you want to try an operation (say create ad connection) in [provider.tf](provider.tf)
2. Navigate to the file corresponding to the resource that you uncommented (the uncommented code block contains the path) and update the values.
3. Review the changes using the following
   ```
   terraform plan
   ```
5. If everything works fine, apply the changes using the following
   ```
   terraform apply -var-file=<main tf file>
   ```

---
## Known Limitations

The following limitations are present in the latest version of the provider. These are being prioritized for resolution in the upcoming release alongside new feature additions:

### 1. All Resource objects
 - `terraform destroy` is not supported.

### 2. Endpoints

- **State management is not supported** for the following attributes:
  - `Owner`
  - `ResourceOwner`
  - `Requestable`
  - `OutOfBandAccess`

- The `MappedEndpoints` field **cannot be configured during endpoint creation**; it must be managed after the endpoint is created.

- The `RequestableRoleType` attribute **can only be set during updates**, since the role must be assigned to the endpoint beforehand.

- For `saviynt_endpoint_resource.requestable_role_types.request_option`, the supported values for proper state tracking are:
  - `DropdownSingle`
  - `Table`
  - `TableOnlyAdd`

- The following service account settings are **not currently configurable via Terraform**:
  - `Disable Remove Service Account`
  - `Disable Modify Service Account`
  - `Disable New Account Request if Account Exists`

### 3. Connections
- `description` field can't be set from Terraform currently.
- **State management** is not supported for the following attributes due to their sensitive nature:
  - **AD**: `password`
  - **ADSI**: `password`
  - **DB**: `password`, `change_pass_json`
  - **EntraId**: `access_token`, `azure_mgmt_access_token`, `client_secret`, `windows_connector_json`, `connection_json`
  - **Github REST**: `connection_json`, `access_tokens`
  - **REST**: `connection_json`
  - **Salesforce**: `client_secret`, `refresh_token`
  - **SAP**: `password`, `prov_password`
  - **Unix**: `password`, `passphrase`, `ssh_key`, `ssh_pass_through_password`, `ssh_pass_through_sshkey`, `ssh_pass_through_passphrase`
  - **Workday**: `password`, `client_secret`, `refresh_token`
---

##  Contributing

> 👋 **Hey Developer!**
>
> We’re glad you’re here and excited that you're interested in contributing. Right now, we’re in the middle of setting up some core processes — like contribution guidelines, issue templates, and workflows — to make sure everything runs smoothly for everyone.
>
> While we’re not quite ready for external contributions just yet, we’re getting close! Hang tight, and keep an eye on this space — we’ll be opening things up soon, and we’d love to have you onboard when we do.

---

##  License

Licensed under Mozilla Public License 2.0. Refer to [LICENSE](LICENSE) for full license details.

---
