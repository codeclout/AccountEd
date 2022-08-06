terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1"
    }
  }
}

resource "aws_ecs_cluster" "app_cluster" {
  name = "${var.environment}-${var.app}-${var.resource_purpose}"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }

  tags = var.tags
}

resource "aws_ecs_task_definition" "fargate_task_definition" {
  cpu                      = var.task_cpu
  execution_role_arn       = var.task_execution_role_arn
  family                   = "api-core"
  memory                   = var.task_memory
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "ARM64"
  }
  task_role_arn = var.task_role_arn

  container_definitions = jsonencode([
    {
      name      = var.task_container_name
      essential = true
      environment = [
        { "name" : "ENVIRONMENT", "value" : "${var.environment}" },
        { "name" : "HEALTHCHECK_INTERVAL", "value" : "${var.task_container_hc_interval}" },
        { "name" : "PORT", "value" : var.task_container_port }
      ],
      image = var.task_image
      logConfiguration = {
        logDriver = "awslogs",
        options = {
          awslogs-group         = "/ecs/fargate/${var.app}-${var.environment}"
          awslogs-region        = var.aws_region
          awslogs-stream-prefix = "ecs"
        }
      },
      portMappings = [
        {
          containerPort = var.task_container_port
          protocol      = "tcp"
        }
      ],
      secrets = var.task_container_secrets
    }
  ])

  tags = var.tags
}

resource "aws_ecs_service" "fargate_service" {
  name            = "${var.environment}-${var.app}-${var.resource_purpose}"
  cluster         = aws_ecs_cluster.app_cluster.id
  task_definition = aws_ecs_task_definition.fargate_task_definition.arn
  desired_count   = var.task_desired_count

  force_new_deployment              = true
  health_check_grace_period_seconds = parseint(var.task_container_hc_interval, 10)
  launch_type                       = "FARGATE"

  lifecycle {
    ignore_changes = [desired_count]
  }

  network_configuration {
    subnets = var.service_subnets
  }
}
