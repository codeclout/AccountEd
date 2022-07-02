terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1"
    }
  }
}

provider "aws" {
  profile = "default"
  region  = var.aws_region
}

locals {
  app_name = "sch00l.io"
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
  task_image                 = "my image"

  service_subnets = [module.network.compute_subnet_az_0, module.network.compute_subnet_az_1]

  task_container_secrets = [
    { name = "DB_STRING", valueFrom = "db-string-secret" }
  ]

  tags = {
    environment = "dev"
  }

}
