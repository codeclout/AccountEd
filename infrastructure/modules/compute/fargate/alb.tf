resource "aws_lb" "core_app_lb" {
  name = "${var.environment}-${var.task_container_name}-alb"

  drop_invalid_header_fields = false
  enable_deletion_protection = false
  enable_http2               = true
  enable_waf_fail_open       = false
  internal                   = false

  idle_timeout = 30

  desync_mitigation_mode = "defensive"
  ip_address_type        = "ipv4"
  load_balancer_type     = "application"

  security_groups = var.alb_security_groups
  subnets         = var.alb_subnets
}

resource "aws_lb_target_group" "core_app_target_group_fargate_ip" {
  name   = "${var.environment}-${var.resource_purpose}-tg"
  port   = parseint(var.task_container_port, 10)
  vpc_id = var.alb_vpc_id

  deregistration_delay = 30

  protocol    = "HTTP"
  target_type = "ip"

  health_check {
    enabled             = true
    healthy_threshold   = 2
    interval            = 15
    matcher             = "200-299"
    path                = var.health_check_path
    port                = parseint(var.task_container_port, 10)
    protocol            = "HTTP"
    timeout             = 10
    unhealthy_threshold = 3
  }

  depends_on = [
    aws_lb.core_app_lb
  ]

  # lifecycle {
  #   create_before_destroy = true
  # }
}

resource "aws_lb_listener" "core_app_listener_secure" {
  certificate_arn   = var.alb_certificate_arn
  load_balancer_arn = aws_lb.core_app_lb.arn

  port       = "443"
  protocol   = "HTTPS"
  ssl_policy = "ELBSecurityPolicy-2016-08"

  default_action {
    target_group_arn = aws_lb_target_group.core_app_target_group_fargate_ip.arn
    type             = "forward"
  }
}

resource "aws_lb_listener" "core_alb_listener_redirect" {
  load_balancer_arn = aws_lb.core_app_lb.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_listener_rule" "health_check" {
  listener_arn = aws_lb_listener.core_app_listener_secure.arn

  action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      message_body = "OK"
      status_code  = "200"
    }
  }

  condition {
    path_pattern {
      values = [var.health_check_path]
    }
  }
}

resource "aws_appautoscaling_target" "ecs_autoscaling_target" {
  max_capacity       = var.task_desired_count * 2
  min_capacity       = 1
  resource_id        = "service/${aws_ecs_cluster.app_cluster.name}/${aws_ecs_service.fargate_service.name}"
  scalable_dimension = "ecs:service:DesiredCount"
  service_namespace  = "ecs"

  lifecycle {
    ignore_changes = [
      role_arn,
    ]
  }
}

resource "aws_appautoscaling_policy" "ecs_autoscaling_policy_memory" {
  name               = "ecs-memory-autoscaling"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.ecs_autoscaling_target.resource_id
  scalable_dimension = aws_appautoscaling_target.ecs_autoscaling_target.scalable_dimension
  service_namespace  = aws_appautoscaling_target.ecs_autoscaling_target.service_namespace

  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageMemoryUtilization"
    }

    target_value = 80
  }
}

resource "aws_appautoscaling_policy" "ecs_autoscaling_policy_cpu" {
  name               = "ecs-cpu-autoscaling"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.ecs_autoscaling_target.resource_id
  scalable_dimension = aws_appautoscaling_target.ecs_autoscaling_target.scalable_dimension
  service_namespace  = aws_appautoscaling_target.ecs_autoscaling_target.service_namespace

  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageCPUUtilization"
    }

    target_value = 60
  }
}

# ------- High memory alarm -------
resource "aws_cloudwatch_metric_alarm" "high-memory-policy-alarm" {
  alarm_name          = "${var.environment}-high-memory-ecs-service-${aws_ecs_service.fargate_service.name}"
  alarm_description   = "High Memory for ecs service-${aws_ecs_service.fargate_service.name}"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = "2"
  metric_name         = "MemoryUtilization"
  namespace           = "AWS/ECS"
  period              = "60"
  statistic           = "Maximum"
  threshold           = 50

  dimensions = {
    ServiceName = aws_ecs_service.fargate_service.name,
    ClusterName = aws_ecs_cluster.app_cluster.name
  }

}
