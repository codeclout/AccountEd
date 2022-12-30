variable "ATLAS_API_KEY_ID" {
  type        = string
  description = ""
}

variable "ATLAS_PROJECT_NAME" {
  type        = string
  description = "Atlas project name"
}

variable "ATLAS_ORG_ID" {
  type        = string
  description = "Atlas organization name"
}

variable "AWS_CI_ROLE_TO_ASSUME" {
  type        = string
  description = "Role for the AWS provider to assume"
}

variable "AWS_ACCESS_KEY_NO_CREDS" {
  type        = string
  description = "IAM user with no permissions"
}

variable "AWS_SECRET_KEY_NO_CREDS" {
  type        = string
  description = "IAM user with no permissions"
}

variable "ATLAS_REGION" {
  type        = string
  description = ""
}

variable "GITHUB_TOKEN" {
  type = string
}

variable "MONGO_CLUSTER_NAME" {
  type        = string
  description = ""
}

variable "atlas_cluster_instance_size" {
  type    = string
  default = "M10"
}

variable "aws_region" {
  default = "us-east-2"
  type    = string
}

variable "db_connection_string_secret_name" {
  type        = string
  description = "Name of the DB connection string secret"
}

variable "environment" {
  default = "dev"
  type    = string

  validation {
    condition     = can(regex("^dev$|^prod$", var.environment))
    error_message = "Error: Only 2 environments are supported - dev & prod"
  }
}

variable "ip_access_list" {
  type        = list(string)
  description = "List of ip addresses with access to the db"
}

variable "mongo_db" {
  type    = string
  default = "accountEd"
}

variable "mongo_db_cluster_name" {
  type        = string
  description = "Atlas db cluster name"
}

