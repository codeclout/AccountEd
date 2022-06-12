terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.15.1"
    }
  }
}

provider "aws" {
  # Configuration options
}

resource "aws_cognito_user_pool" "accountEd_user_pool" {
  name = var.accountEd_user_pool

  account_recovery_setting {
    dynamic "recovery_mechanism" {
      for_each = var.account_recovery_settings
      content {
        name     = recovery_mechanism.value["name"]
        priority = recovery_mechanism.value["priority"]
      }
    }
  }

  admin_create_user_config {
    allow_admin_create_user_only = var.allow_admin_create_user_only
  }


}
