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

locals {
  environment = "dev"
}

# GITHUB_TOKEN required in the environment to authenticate with GitHub
provider "github" {
  token = var.GITHUB_TOKEN
}

module "database" {
  source = "../modules/db"

  ATLAS_ORG_ID       = var.ATLAS_ORG_ID
  ATLAS_PROJECT_NAME = var.ATLAS_PROJECT_NAME

  environment    = var.environment
  ip_access_list = var.ip_access_list

  mongo_db              = "accountEd"
  mongo_db_cluster_name = var.mongo_db_cluster_name
  mongo_db_role_arn     = var.mongo_db_role_arn
}
