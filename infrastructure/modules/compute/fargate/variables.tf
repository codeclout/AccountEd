variable "alb_certificate_arn" {
  type        = string
  description = "ARN of the ACM Certificate resource"
}

variable "alb_security_groups" {
  type        = list(string)
  description = "Security group IDs for ALB"
}

variable "alb_subnets" {
  type        = list(string)
  description = "Subnets associated with the ALB"
}

variable "alb_vpc_id" {
  type = string
}

# Name of the application
variable "app" {
  type    = string
  default = "sch00l.io"
}

variable "aws_profile" {
  type    = string
  default = "default"
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

variable "health_check_path" {
  type = list(string)
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

variable "service_subnets" {
  type        = list(string)
  description = "Subnets associated with the service"
}

variable "should_scan_image_on_push" {
  type    = bool
  default = true
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

variable "task_container_hc_interval" {
  type        = string
  description = "ECS task definition container healthcheck interval"
}
variable "task_container_name" {
  type        = string
  description = "ECS task definition container name"
}

variable "task_container_port" {
  type        = number
  description = "ECS task defintion container port number bound to the host port"
}

variable "task_container_secrets" {
  type = list(object({
    name      = string
    valueFrom = string
  }))
  description = "List of secrets from parameter store or secrets manager"
}

variable "task_cpu" {
  type        = number
  description = "ECS task CPU"
}

variable "task_desired_count" {
  type        = number
  description = "Number of instances of the task definition to place and keep running"

  # validation {
  #   condition     = var.environment == "prod" ? var.task_desired_count >= 1 && var.task_desired_count <= 4 : var.task_desired_count <= 2
  #   error_message = "Prod number of task instances must be greater than 1 and less than 4 | dev task instances, 2 or less"
  # }
}

variable "task_execution_role_arn" {
  type        = string
  description = ""
}

variable "task_image" {
  type        = string
  description = "The image URI & tag used to start the container"
}

variable "task_memory" {
  type        = number
  description = "ECS task memory allocation"
}

variable "task_role_arn" {
  type        = string
  description = "IAM role for API requests to AWS services"
}


