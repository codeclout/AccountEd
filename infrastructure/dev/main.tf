terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1"
    }

    github = {
      source  = "integrations/github"
      version = ">= 4.28.0"
    }
  }
}

provider "aws" {
  profile = "default"
  region  = var.aws_region
}

# GITHUB_TOKEN required in the environment to authenticate with GitHub
provider "github" {
  token = var.GITHUB_TOKEN
}

locals {
  app_name = "sch00l.io"
}

data "github_ref" "dev" {
  repository = "AccountEd"
  ref        = "heads/develop"
}

module "iam" {
  source = "../global/iam"
}

module "network" {
  source = "../modules/network"

  app                     = local.app_name
  availability_zone_count = 2
  aws_region              = var.aws_region
  environment             = "dev"

  tags = {
    environment = "dev"
  }
}

module "ecr" {
  source = "../modules/compute/ecr"

  app         = local.app_name
  aws_region  = var.aws_region
  environment = "dev"
}

data "aws_ecr_image" "svc_image" {
  image_tag       = substr(data.github_ref.dev.sha, 0, 11)
  repository_name = module.ecr.container_repository_url
}

module "ecs_compute" {
  source = "../modules/compute/fargate"

  app                     = local.app_name
  aws_region              = var.aws_region
  environment             = "dev"
  task_container_port     = "8088"
  task_cpu                = 256
  task_execution_role_arn = module.iam.ecs_task_execution_role_arn
  task_memory             = 512
  task_role_arn           = module.iam.ecs_task_role_arn

  task_container_hc_interval = 5
  task_container_name        = "my container"
  task_desired_count         = 1
  task_image                 = data.aws_ecr_image.svc_image.id

  service_subnets = [module.network.compute_subnet_az_0, module.network.compute_subnet_az_1]

  task_container_secrets = [
    { name = "DB_STRING", valueFrom = "db-string-secret" }
  ]

  tags = {
    environment = "dev"
  }

}
