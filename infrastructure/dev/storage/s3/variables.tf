variable "AWS_ACCESS_KEY_NO_CREDS" {
  type = string
}

variable "aws_region" {
  type    = string
  default = "us-east-2"
}

variable "AWS_S3_ACCESS_ROLE_NAME" {
  type = string
}

variable "AWS_SECRET_KEY_NO_CREDS" {
  type = string
}

variable "AWS_CI_ROLE_TO_ASSUME" {
  type = string
}

variable "GITHUB_TOKEN" {
  type = string
}
