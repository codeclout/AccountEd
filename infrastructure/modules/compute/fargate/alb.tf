resource "aws_lb" "core_app_lb" {
  name = "${var.environment}-${var.app}-core-app"

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

resource "aws_lb_listener" "front_end" {
  load_balancer_arn = aws_lb.core_app_lb.arn
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = var.alb_certificate_arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.front_end.arn
  }
}
