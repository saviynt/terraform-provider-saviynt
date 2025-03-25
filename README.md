
# Terraform Provider from Saviynt

This Terraform provider enables efficient provisioning, configuration, and management of security systems, endpoints, and associated connections through the Saviynt API. Designed for seamless integration with Saviynt EIC, it brings Infrastructure-as-Code (IaC) principles to identity and access management.

---

##  Overview

With this provider, you can:

- Automate the creation and management of Security Systems and Endpoints
- Manage various Connections (e.g., AD, REST, Workday)
- Leverage Terraform’s declarative configuration to manage IAM resources
- Minimize manual intervention and reduce configuration drift

---

##  Features

- Full CRUD support for Saviynt security systems and endpoints
- Support for multiple Connection types: AD, ADSI, REST, Workday, etc.
- Rich filtering in data sources (e.g., filter security systems by systemname, connection_type, etc.)
- Advanced workflow integration (e.g., firefighter workflows, account ownership workflows)
- Advanced error handling, comprehensive logging, and enhanced debugging capabilities
- Pre-built templates for quick and effective implementation

---

##  Requirements

- Terraform version `>= 1.8+`
- Go programming language `>= 1.21+` (required for development and contributions)
- Saviynt credentials (url, username and password)

---

##  Installation

To use this provider, configure Terraform as follows:

### Terraform Configuration

```hcl
terraform {
  required_providers {
    saviynt = {
      source  = "registry.terraform.io/local/saviynt"
      version = "1.0.0"
    }
  }
}

provider "saviynt" {
  username   = "YOUR_SAVIYNT_USERNAME"
  password   = "YOUR_SAVIYNT_PASSWORD"
  api_token  = "YOUR_SAVIYNT_API_TOKEN"
}
```

---

##  Usage

Here's an example of defining and managing a resource:

```hcl
resource "saviynt_security_system_resource" "example" {
  systemname          = "hr_system"
  display_name        = "HR System"
  hostname            = "hr.example.com"
  port                = "443"
  access_add_workflow = "hr_access_add"
}
```

```hcl
data "saviynt_security_systems_datasource" "all" {
  connection_type = "REST"
  max             = 10
  offset          = 0
}

output "systems" {
  value = data.saviynt_security_systems_datasource.all.results
}
```

---

##  Available Resources

###  Resource

- `saviynt_security_system_resource`: Manages lifecycle (create, update, read) of security systems. Supports workflows, connectors, password policies and more.
- `saviynt_endpoints_resource`: For managing endpoints definitions used by security systems.
- `saviynt_connection_resouce`: For managing endpoints like AD, REST, etc. tied to security systems.

###  Data Source

- `saviynt_security_systems_datasource`: Retrieves a list of configured security systems filtered by systemname, connection_type, etc.

---

##  Development and Contribution

To contribute to this project or develop locally:

### 1. Clone the Repository

```bash
git clone https://github.com/saviynt/terraform-provider-saviynt.git
cd terraform-provider-saviynt
```

### 2. Build the Provider

```bash
go mod tidy
go build -o terraform-provider-saviynt
```

### 3. Locate Your `GOBIN` Path

```bash
go env GOBIN
```

If empty:

```bash
echo "$(go env GOPATH)/bin"
```

Examples:

- `/Users/<your-username>/go/bin` (macOS/Linux)  
- `C:\Users\<your-username>\go\bin` (Windows)

### 4. Configure `.terraformrc` or `terraform.rc`

Create the file at:

- **macOS/Linux**: `~/.terraformrc`
- **Windows**: `%APPDATA%\terraform.rc`

```hcl
provider_installation {
  dev_overrides {
    "yourorgname/saviynt" = "<PATH>"
  }
  direct {}
}
```

Replace `<PATH>` with your actual GOBIN path.

### 5. Test the Provider Locally

```hcl
provider "saviynt" {
  username  = "YOUR_USERNAME"
  password  = "YOUR_PASSWORD"
  server_url = "YOUR_SERVER_URL"
}

data "saviynt_security_systems_datasource" "example" {
  systemname = "MySystem"
}
```

Run:

```bash
terraform init
terraform plan
terraform apply
```

### 6. Run Tests

```bash
go test ./... -v
```

---

##  Contributions Welcome!

Contributions are warmly welcomed! Please follow these guidelines:

- Submit issues clearly describing bugs or enhancement suggestions.
- Provide pull requests that include relevant tests.
- Ensure existing tests are passed and functionality remains intact.

---

##  License

This project is licensed under the Apache License 2.0. Refer to [LICENSE](LICENSE) for full license details.

---

##  Support

If you encounter any issues or have questions, please open an issue on our GitHub page.

---
