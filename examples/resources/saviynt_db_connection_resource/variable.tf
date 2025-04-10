variable "SAVIYNT_SERVER_URL" {
  type        = string
  description = "Saviynt API Server URL (without https://)"
}
variable "SAVIYNT_USERNAME" {
  type        = string
  description = "Saviynt API Username"
}
variable "SAVIYNT_PASSWORD" {
  type        = string
  description = "Saviynt API Password"
  sensitive   = true
}
variable "PASSWORD" {
  type        = string
  description = "Connector Password"
  sensitive   = true
}