resource "aws_acm_certificate" "alb_cert" {
  domain_name       = "*.${local.app_name}"
  validation_method = "DNS"

  tags = {
    environment = "dev"
  }

  validation_option {
    domain_name       = "*.${local.app_name}"
    validation_domain = local.app_name
  }
}
