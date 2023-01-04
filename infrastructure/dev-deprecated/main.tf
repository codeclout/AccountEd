terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1, < 5.0.0"
    }

    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = ">= 3.22.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

provider "cloudflare" {
  api_token = var.CF_API_TOKEN
}

# GITHUB_TOKEN required in the environment to authenticate with GitHub
provider "github" {
  token = var.GITHUB_TOKEN
}

locals {
  app_name         = "sch00l.io"
  container_port   = "8088"
  environment      = "dev"
  resource_purpose = "core"
}

data "aws_ecr_repository" "data" {
  name = "${local.environment}-${split(".", local.app_name)[0]}-${local.resource_purpose}"
}

module "iam" {
  source = "../global/iam"

  container_registry_arn = data.aws_ecr_repository.data.arn
}

module "network" {
  source = "../modules/network"

  environment = "dev"
  vpc_cidr    = "10.1.0.0/16"

  app            = local.app_name
  aws_region     = var.aws_region
  container_port = local.container_port

  availability_zone_count = 2
  nat_gateway_count       = 1
  tags = {
    environment = "dev"
  }
}

module "ecs_compute" {
  source = "../modules/compute/fargate"

  environment         = "dev"
  health_check_path   = "/hc"
  resource_purpose    = "core-account-management"
  task_container_name = "core-api"

  alb_certificate_arn     = aws_acm_certificate.alb_cert.arn
  alb_security_groups     = [module.network.alb_security_grp]
  alb_vpc_id              = module.network.alb_vpc_id
  app                     = local.app_name
  aws_region              = var.aws_region
  ecs_security_groups     = [module.network.ecs_security_grp]
  task_container_port     = local.container_port
  task_execution_role_arn = module.iam.ecs_task_execution_role_arn
  task_image              = "scratch"
  task_role_arn           = module.iam.ecs_task_role_arn

  task_container_hc_interval = 5
  task_cpu                   = 256
  task_desired_count         = 1
  task_memory                = 512

  alb_subnets     = [module.network.public_compute_subnet_az_4, module.network.public_compute_subnet_az_5]
  service_subnets = [module.network.compute_subnet_az_0, module.network.compute_subnet_az_1]

  task_container_secrets = [
    { name = "DB_STRING", valueFrom = "db-string-secret" }
  ]

  tags = {
    environment = "dev"
  }

}