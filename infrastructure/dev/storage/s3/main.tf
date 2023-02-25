terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1, < 5.0.0"
    }
    time = {
      source  = "hashicorp/time"
      version = ">= 0.9.1, < 1.0.0"
    }
  }
}

resource "time_static" "ts" {}

provider "aws" {
  access_key = var.AWS_ACCESS_KEY_NO_CREDS
  region     = var.aws_region
  secret_key = var.AWS_SECRET_KEY_NO_CREDS

  assume_role {
    role_arn     = var.AWS_CI_ROLE_TO_ASSUME
    session_name = "TF_CLOUD_RUN_${time_static.ts.unix}"
  }
}

# GITHUB_TOKEN required in the environment to authenticate with GitHub
provider "github" {
  token = var.GITHUB_TOKEN
}

locals {
  app_codename = "accountEd"
  environment  = "dev"

  unencrypted_buckets = [
    "${local.environment}-${var.aws_region}-${local.app_codename}"
  ]
}

# get the role friendly name
data "aws_iam_role" "s3_access_role_id" {
  name = var.AWS_S3_ACCESS_ROLE_NAME
}

data "aws_caller_identity" "current" {}

module "storage" {
  source = "../../../modules/storage/s3"

  buckets_unencrypted = local.unencrypted_buckets
  is_object_locked    = [false]
}

module "s3_role_policies" {
  count  = length(local.unencrypted_buckets)
  source = "../../../modules/security/s3-role-policies"

  acccount_id             = data.aws_caller_identity.current.account_id
  bucket_role_id          = data.aws_iam_role.s3_access_role_id.unique_id
  role_access_bucket_name = module.storage.bucket_name[count.index]
}

resource "aws_s3_bucket_policy" "unencrypted_bucket_policy" {
  count = length(local.unencrypted_buckets)

  bucket = module.storage.bucket_name[count.index]
  policy = module.s3_role_policies.policy_unencrypted_bucket[*]
}
