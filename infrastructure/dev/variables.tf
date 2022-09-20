variable "aws_region" {
  type    = string
  default = "us-east-2"
}

variable "AWS_ADM_ACCOUNT_EMAIL" {
  type        = string
  description = "AWS ADM Account Creation Email"
}

variable "CF_ACCOUNT_ID" {
  type        = string
  description = "Cloudflare Account ID"
}

variable "CF_API_TOKEN" {
  type        = string
  description = "Cloudflare API token"
}

variable "CF_TUNNEL_SECRET" {
  type        = string
  description = "32 or more bytes encoded as a base64 string"
}

variable "CF_ZONE_ID" {
  type        = string
  description = "Cloudflare zone id"
}

variable "GITHUB_TOKEN" {
  type = string
}
