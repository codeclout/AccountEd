# Name of the application
variable "app" {
  type    = string
  default = "sch00l.io"
}

variable "aws_region" {
  type    = string
  default = "us-east-2"
}

variable "environment" {
  type = string

  validation {
    condition     = can(regex("^dev$|^prod$", var.environment))
    error_message = "Error: Only 2 environments are supported - dev & prod"
  }
}

variable "image_tag_mutability" {
  type    = string
  default = "IMMUTABLE"

  validation {
    condition     = can(regex("^MUTABLE$|^IMMUTABLE$", var.image_tag_mutability))
    error_message = "Error: image tag mutability must be IMMUTABLE or MUTABLE"
  }
}

variable "resource_purpose" {
  type        = string
  description = "Answers the purpose the resource serves - e.g. core-account-management"
  default     = "ephemeral"
}

variable "should_scan_image_on_push" {
  type    = bool
  default = true
}


