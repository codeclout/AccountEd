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
  app_codename = "accountEd"
  environment  = "dev"
  param_suffix = "dbcs"
}

# GITHUB_TOKEN required in the environment to authenticate with GitHub
provider "github" {
  token = var.GITHUB_TOKEN
}

module "database" {
  source = "../modules/db"

  ATLAS_API_KEY_ID   = var.ATLAS_API_KEY_ID
  ATLAS_ORG_ID       = var.ATLAS_ORG_ID
  ATLAS_PROJECT_NAME = var.ATLAS_PROJECT_NAME

  atlas_region = var.ATLAS_REGION
  environment  = local.environment
  ip_address   = var.ip_access_list

  mongo_db_cluster_name = var.MONGO_CLUSTER_NAME
  mongo_db_role_arn     = var.AWS_CI_ROLE_TO_ASSUME

}

resource "aws_secretsmanager_secret" "db_secret_name" {
  name = "${var.environment}-${module.database.cluster_id}-${var.db_connection_string_secret_name}"
}

resource "aws_secretsmanager_secret_version" "db_secret" {
  secret_id     = aws_secretsmanager_secret.db_secret_name.id
  secret_string = module.database.connection_strings
}

resource "aws_ssm_parameter" "db_secret_param" {
  name  = "/${var.environment}/db/${local.app_codename}/${local.param_suffix}"
  type  = "SecureString"
  value = aws_secretsmanager_secret.db_secret_name.name
}
