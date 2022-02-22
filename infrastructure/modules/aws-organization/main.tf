terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=4.1.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

resource "aws_organizations_organization" "org_manager" {
  aws_service_access_principals = [
    "account.amazonaws.com",
    "cloudtrail.amazonaws.com",
    "servicecatalog.amazonaws.com"
  ]
}
