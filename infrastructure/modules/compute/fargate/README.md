<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 4.20.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 4.20.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_appautoscaling_policy.ecs_autoscaling_policy_cpu](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/appautoscaling_policy) | resource |
| [aws_appautoscaling_policy.ecs_autoscaling_policy_memory](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/appautoscaling_policy) | resource |
| [aws_appautoscaling_target.ecs_autoscaling_target](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/appautoscaling_target) | resource |
| [aws_cloudwatch_metric_alarm.high-memory-policy-alarm](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_metric_alarm) | resource |
| [aws_ecs_cluster.app_cluster](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_cluster) | resource |
| [aws_ecs_service.fargate_service](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_service) | resource |
| [aws_ecs_task_definition.fargate_task_definition](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_task_definition) | resource |
| [aws_lb.core_app_lb](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb) | resource |
| [aws_lb_listener.core_alb_listener_redirect](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_listener) | resource |
| [aws_lb_listener.core_app_listener_secure](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_listener) | resource |
| [aws_lb_listener_rule.health_check](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_listener_rule) | resource |
| [aws_lb_target_group.core_app_target_group_fargate_ip](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_target_group) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_alb_certificate_arn"></a> [alb\_certificate\_arn](#input\_alb\_certificate\_arn) | ARN of the ACM Certificate resource | `string` | n/a | yes |
| <a name="input_alb_security_groups"></a> [alb\_security\_groups](#input\_alb\_security\_groups) | Security group IDs for ALB | `list(string)` | n/a | yes |
| <a name="input_alb_subnets"></a> [alb\_subnets](#input\_alb\_subnets) | Subnets associated with the ALB | `list(string)` | n/a | yes |
| <a name="input_alb_vpc_id"></a> [alb\_vpc\_id](#input\_alb\_vpc\_id) | n/a | `string` | n/a | yes |
| <a name="input_app"></a> [app](#input\_app) | Name of the application | `string` | `"sch00l.io"` | no |
| <a name="input_aws_profile"></a> [aws\_profile](#input\_aws\_profile) | n/a | `string` | `"default"` | no |
| <a name="input_aws_region"></a> [aws\_region](#input\_aws\_region) | n/a | `string` | `"us-east-2"` | no |
| <a name="input_ecs_security_groups"></a> [ecs\_security\_groups](#input\_ecs\_security\_groups) | n/a | `list(string)` | n/a | yes |
| <a name="input_environment"></a> [environment](#input\_environment) | n/a | `string` | n/a | yes |
| <a name="input_health_check_path"></a> [health\_check\_path](#input\_health\_check\_path) | n/a | `string` | n/a | yes |
| <a name="input_image_tag_mutability"></a> [image\_tag\_mutability](#input\_image\_tag\_mutability) | n/a | `string` | `"IMMUTABLE"` | no |
| <a name="input_resource_purpose"></a> [resource\_purpose](#input\_resource\_purpose) | Answers the purpose the resource serves - e.g. core-account-management | `string` | `"ephemeral"` | no |
| <a name="input_service_subnets"></a> [service\_subnets](#input\_service\_subnets) | Subnets associated with the service | `list(string)` | n/a | yes |
| <a name="input_should_scan_image_on_push"></a> [should\_scan\_image\_on\_push](#input\_should\_scan\_image\_on\_push) | n/a | `bool` | `true` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `map(string)` | n/a | yes |
| <a name="input_task_container_hc_interval"></a> [task\_container\_hc\_interval](#input\_task\_container\_hc\_interval) | ECS task definition container healthcheck interval | `string` | n/a | yes |
| <a name="input_task_container_name"></a> [task\_container\_name](#input\_task\_container\_name) | ECS task definition container name | `string` | n/a | yes |
| <a name="input_task_container_port"></a> [task\_container\_port](#input\_task\_container\_port) | ECS task defintion container port number bound to the host port | `string` | n/a | yes |
| <a name="input_task_container_secrets"></a> [task\_container\_secrets](#input\_task\_container\_secrets) | List of secrets from parameter store or secrets manager | <pre>list(object({<br>    name      = string<br>    valueFrom = string<br>  }))</pre> | n/a | yes |
| <a name="input_task_cpu"></a> [task\_cpu](#input\_task\_cpu) | ECS task CPU | `number` | n/a | yes |
| <a name="input_task_desired_count"></a> [task\_desired\_count](#input\_task\_desired\_count) | Number of instances of the task definition to place and keep running | `number` | n/a | yes |
| <a name="input_task_execution_role_arn"></a> [task\_execution\_role\_arn](#input\_task\_execution\_role\_arn) | n/a | `string` | n/a | yes |
| <a name="input_task_image"></a> [task\_image](#input\_task\_image) | The image URI & tag used to start the container | `string` | n/a | yes |
| <a name="input_task_memory"></a> [task\_memory](#input\_task\_memory) | ECS task memory allocation | `number` | n/a | yes |
| <a name="input_task_role_arn"></a> [task\_role\_arn](#input\_task\_role\_arn) | IAM role for API requests to AWS services | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_ecs_alb_arn"></a> [ecs\_alb\_arn](#output\_ecs\_alb\_arn) | n/a |
| <a name="output_ecs_alb_dns_name"></a> [ecs\_alb\_dns\_name](#output\_ecs\_alb\_dns\_name) | n/a |
| <a name="output_ecs_alb_zone_id"></a> [ecs\_alb\_zone\_id](#output\_ecs\_alb\_zone\_id) | n/a |
<!-- END_TF_DOCS -->