terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=4.1.0"
    }
  }
}

module "iam" {
  source  = "terraform-aws-modules/iam/aws"
  version = "4.13.2"
  # insert the 1 required variable here
}