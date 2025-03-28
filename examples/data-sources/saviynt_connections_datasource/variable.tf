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
# variable "IP_ADDRESS" {
#   type        = string
#   description = "Saviynt host server"
# }
# variable "LDAP_PORT" {
#   type        = string
#   description = "Port for the connection"
# }
# variable "LDAP_PROTOCOL" {
#   type        = string
#   description = "Protocol type (e.g., LDAP, HTTP, etc.)"
# }
# variable "PASSWORD" {
#   type        = string
#   description = "Connection password"
#   sensitive   = true
# }
# variable "BIND_USER" {
#   type        = string
#   description = "Connection username"
# }