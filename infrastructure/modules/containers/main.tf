terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.21.0, < 6.0.0"
    }
  }

  required_version = ">= 1.5.5, < 2.0.0"
}