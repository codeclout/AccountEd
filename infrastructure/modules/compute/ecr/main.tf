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
  name                 = "${var.environment}-${var.app}-${var.resource_purpose}"
  image_tag_mutability = var.image_tag_mutability

  encryption_configuration {
    encryption_type = "AES256"
  }
  image_scanning_configuration {
    scan_on_push = var.should_scan_image_on_push
  }
}

resource "aws_ecr_lifecycle_policy" "ecr_lifecycle_policy" {
  repository = aws_ecr_repository.container_repo.name

  policy = jsonencode({
    rules = [{
      action = {
        type = "expire"
      }

      description  = "keep last 2 images"
      rulePriority = 1

      selection = {
        countNumber = 2
        countType   = "imageCountMoreThan"
        tagStatus   = "any"
      }
    }]
  })
}
