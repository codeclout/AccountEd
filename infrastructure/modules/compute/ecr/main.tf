terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

resource "aws_ecr_repository" "container_repo" {
  name                 = "${var.app}-${var.environment}"
  image_tag_mutability = var.image_tag_mutability

  image_scanning_configuration {
    scan_on_push = var.should_scan_image_on_push
  }
}
