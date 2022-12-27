terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1, < 5.0.0"
    }
  }
}

provider "aws" {
  access_key = var.AWS_ACCESS_KEY_NO_CREDS
  region     = var.aws_region
  secret_key = var.AWS_SECRET_KEY_NO_CREDS

  assume_role {
    role_arn = var.AWS_CI_ROLE_TO_ASSUME
  }
}

# GITHUB_TOKEN required in the environment to authenticate with GitHub
provider "github" {
  token = var.GITHUB_TOKEN
}
