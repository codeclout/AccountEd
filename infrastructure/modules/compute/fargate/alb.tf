resource "aws_lb" "core_app_lb" {
  name = "${var.environment}-${var.task_container_name}-alb"

  drop_invalid_header_fields = false
  enable_deletion_protection = false
  enable_http2               = true
  enable_waf_fail_open       = false
  internal                   = false

  idle_timeout = 60

  desync_mitigation_mode = "defensive"
  ip_address_type        = "ipv4"
  load_balancer_type     = "application"

  security_groups = var.alb_security_groups
  subnets         = var.alb_subnets
}

resource "aws_lb_target_group" "core_app_target_group_fargate_ip" {
  name        = "${var.environment}-${var.resource_purpose}-tg"
  port        = parseint(var.task_container_port, 10)
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = var.alb_vpc_id
}

resource "aws_lb_listener" "core_app_listener_secure" {
  load_balancer_arn = aws_lb.core_app_lb.arn
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = var.alb_certificate_arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.core_app_target_group_fargate_ip.arn
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
      values = var.health_check_path
    }
  }
}
