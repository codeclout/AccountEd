terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.2.0"
    }
  }
}

provider "aws" {
  region = "us-west-2"
}

module "iam" {
  source  = "terraform-aws-modules/iam/aws"
  version = "4.13.1"
}
