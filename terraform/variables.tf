variable "subscription_id" {
  sensitive = true
}
variable "client_secret" {
  sensitive = true
}

variable "azure_tenant_id" {
  sensitive = true
}

variable "location" {
  default = "westeurope"
}

variable "github_access_token" {
  description = "The access token for the GitHub repository"
  type        = string
  sensitive   = true
}
