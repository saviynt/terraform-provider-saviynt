variable "saviynt_server_url" {
  type        = string
  description = "Saviynt API Server URL (without https://)"
}

variable "saviynt_username" {
  type        = string
  description = "Saviynt API Username"
}

variable "saviynt_password" {
  type        = string
  description = "Saviynt API Password"
  sensitive   = true
}
