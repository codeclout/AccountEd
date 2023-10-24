terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.21.0, < 6.0.0"
    }
    time = {
      source  = "hashicorp/time"
      version = ">= 0.9.1, < 1.0.0"
    }
  }
}

resource "time_static" "ts" {}

provider "aws" {
  assume_role {
    role_arn     = var.AWS_CI_ROLE_TO_ASSUME
    session_name = "TF_CLOUD_RUN_${time_static.ts.unix}"
  }
}