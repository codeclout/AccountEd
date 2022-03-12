variable "awsRegion" {
  type        = string
  default     = "us-east-1"
  description = "AWS Region"
}

variable "LMS_USER_ACCOUNT_EMAIL" {
  type        = string
  description = "Email address used to create the production LMS account"
}

variable "LMS_ACCOUNT_ROLE" {
  type        = string
  description = "Role name used when creating the LMS account"
}

variable "PROXY_ACCOUNT_USERS_EMAIL" {
  type        = string
  description = "Email address used to create the LMS proxy account"
}

variable "PROXY_ACCOUNT_ROLE_NAME" {
  type        = string
  description = "Role name used when creating the LMS proxy account"
}