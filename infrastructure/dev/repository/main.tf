terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1, < 5.0.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

locals {
  app_name   = "sch00l.io"
  aws_region = "us-east-2"
}

module "ecr" {
  source = "../../modules/compute/ecr"

  app              = split(".", local.app_name)[0]
  aws_region       = local.aws_region
  environment      = "dev"
  resource_purpose = "core"
}
