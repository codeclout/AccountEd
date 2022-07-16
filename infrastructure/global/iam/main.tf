terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1"
    }
  }
}

provider "aws" {
  region = var.aws_region
  assume_role {
    duration = "15m"
    role_arn = "arn:aws:iam::*:role/ci-svc-usr-role"
  }
}
