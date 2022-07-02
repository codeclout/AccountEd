variable "app" {
  type    = string
  default = "sch00l.io"
}

variable "availability_zone_count" {
  type    = number
  default = 2

  validation {
    condition     = var.availability_zone_count > 1 && var.availability_zone_count <= 4
    error_message = "Error: Requires more than 1 up to a max of 4 AZs"
  }
}

variable "aws_region" {
  type = string
}

variable "environment" {
  type = string

  validation {
    condition     = can(regex("^dev$|^prod$", var.environment))
    error_message = "Error: Only 2 environments are supported - dev & prod"
  }
}

variable "tags" {
  type = map(string)

  validation {
    condition = can([
      for k, v in var.tags : k
      if k == "environment" && can(regex("^dev$|^prod$", v))
    ])

    error_message = "Error: environment tag must be dev or prod"
  }
}
