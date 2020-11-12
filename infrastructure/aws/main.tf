provider "aws" {
  profile = var.aws_profile
  region  = var.region
}

terraform {
  required_version = "~> 0.13"
  required_providers {
    local = "1.4.0"
    aws = {
      source  = "hashicorp/aws"
      version = "3.14.1"
    }
    cloudinit = "1.0.0"
    random    = "2.3.0"
  }
}
