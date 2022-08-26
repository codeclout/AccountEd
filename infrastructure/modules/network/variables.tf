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

variable "container_port" {
  type        = string
  description = "Allow ingress to the port exposed by the task"
}

variable "environment" {
  type = string

  validation {
    condition     = can(regex("^dev$|^prod$", var.environment))
    error_message = "Error: Only 2 environments are supported - dev & prod"
  }
}

variable "nat_gateway_count" {
  type        = number
  description = "Number of NAT Gateways to deploy"

  validation {
    condition     = var.nat_gateway_count >= 1 && var.nat_gateway_count <= 2
    error_message = "A minimum of 1 and maximum of 2 NAT Gateway(s) are required"
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
