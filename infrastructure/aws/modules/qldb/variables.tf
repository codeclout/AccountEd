variable region {
  type        = string
  default     = "us-east-1"
  description = "AWS Region"
}

variable name {
  type        = string
  default     = "accountEd"
  description = "QLDB Name"
}

variable "app" {
  type = string
}

variable "aws_profile" {
  type = string
}

variable "environment" {
  type = map(string)
}

variable "namespace" {
  type = string
}

variable "region" {
  type = string
}

variable "tags" {
  type = map(string)
  default = {
    application   = "my_project"
    environment   = "dev"
    customer      = "customer_a"
    contact-email = "teamEmail_at_tattletale.io"
  }
}
