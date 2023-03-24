variable "ATLAS_API_KEY_ID" {
  type        = string
  description = "Atlas org API key ID"
}

# Atlas Organization ID 
variable "ATLAS_ORG_ID" {
  type        = string
  description = "Atlas Organization ID"
}

variable "ATLAS_PROJECT_NAME" {
  type        = string
  description = "Atlas Project Name"
}

# Atlas Region
variable "atlas_region" {
  type        = string
  description = "Atlas region where resources will be created"
}

# Atlas Project Environment
variable "environment" {
  type        = string
  description = "The environment to be built"
}

# Cloud Provider to Host Atlas Cluster
variable "cloud_provider" {
  type        = string
  default     = "AWS"
  description = "AWS, GCP or Azure"
}

variable "atlas_cluster_instance_size" {
  type        = string
  default     = "M10"
  description = ""
}

# IP Address Access
variable "ip_address" {
  type        = list(string)
  description = "IP address used to access Atlas cluster"
}

variable "mongo_db" {
  type    = string
  default = "accountEd"
}

# MongoDB Version 
variable "mongodb_version" {
  type        = string
  default     = "6.0"
  description = "MongoDB Version"
}

variable "mongo_db_cluster_name" {
  type        = string
  description = ""
}

variable "mongo_db_role_arn" {
  type        = string
  description = "ARN of role to assume for access to the db"
}
