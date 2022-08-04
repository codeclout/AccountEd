output "ecs_alb_arn" {
  value = aws_lb.core_app_lb.arn
}

output "ecs_alb_dns_name" {
  value = aws_lb.core_app_lb.dns_name
}

output "ecs_alb_zone_id" {
  value = aws_lb.core_app_lb.zone_id
}
