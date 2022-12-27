<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 4.20.1, < 5.0.0 |
| <a name="requirement_cloudflare"></a> [cloudflare](#requirement\_cloudflare) | >= 3.22.0 |
| <a name="requirement_github"></a> [github](#requirement\_github) | >= 4.28.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 4.28.0 |
| <a name="provider_cloudflare"></a> [cloudflare](#provider\_cloudflare) | 3.22.0 |
| <a name="provider_github"></a> [github](#provider\_github) | 4.29.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_ecr"></a> [ecr](#module\_ecr) | ../modules/compute/ecr | n/a |
| <a name="module_ecs_compute"></a> [ecs\_compute](#module\_ecs\_compute) | ../modules/compute/fargate | n/a |
| <a name="module_iam"></a> [iam](#module\_iam) | ../global/iam | n/a |
| <a name="module_network"></a> [network](#module\_network) | ../modules/network | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_acm_certificate.alb_cert](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate) | resource |
| [aws_route53_record.dev_zone_ns](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record) | resource |
| [aws_route53_zone.dev_zone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_zone) | resource |
| [cloudflare_record.dev_zone](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/record) | resource |
| [aws_ecr_image.svc_image](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ecr_image) | data source |
| [github_ref.dev](https://registry.terraform.io/providers/integrations/github/latest/docs/data-sources/ref) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_CF_API_TOKEN"></a> [CF\_API\_TOKEN](#input\_CF\_API\_TOKEN) | Cloudflare API token | `string` | n/a | yes |
| <a name="input_CF_ZONE_ID"></a> [CF\_ZONE\_ID](#input\_CF\_ZONE\_ID) | Cloudflare zone id | `string` | n/a | yes |
| <a name="input_GITHUB_TOKEN"></a> [GITHUB\_TOKEN](#input\_GITHUB\_TOKEN) | n/a | `string` | n/a | yes |
| <a name="input_aws_region"></a> [aws\_region](#input\_aws\_region) | n/a | `string` | `"us-east-2"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_dns_name_servers"></a> [dns\_name\_servers](#output\_dns\_name\_servers) | n/a |
<!-- END_TF_DOCS -->