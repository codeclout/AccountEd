variable "aws_region" {
  type    = string
  default = "us-east-2"
}

variable "container_registry_arn" {
  type    = string
  default = "arn:aws:ecr:us-east-2:*"
}
