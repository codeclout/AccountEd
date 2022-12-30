<!-- BEGIN_TF_DOCS -->
## Requirements

| Name                                                    | Version            |
| ------------------------------------------------------- | ------------------ |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 4.20.1, < 5.0.0 |

## Providers

| Name                                              | Version |
| ------------------------------------------------- | ------- |
| <a name="provider_aws"></a> [aws](#provider\_aws) | 4.48.0  |

## Modules

| Name                                                         | Source        | Version |
| ------------------------------------------------------------ | ------------- | ------- |
| <a name="module_database"></a> [database](#module\_database) | ../modules/db | n/a     |

## Resources

| Name                                                                                                                                                     | Type     |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| [aws_secretsmanager_secret.db_secret_name](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager_secret)            | resource |
| [aws_secretsmanager_secret_version.db_secret](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager_secret_version) | resource |

## Inputs

| Name                                                                                                                                       | Description                                | Type           | Default       | Required |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------ | -------------- | ------------- | :------: |
| <a name="input_ATLAS_API_KEY_ID"></a> [ATLAS\_API\_KEY\_ID](#input\_ATLAS\_API\_KEY\_ID)                                                   | n/a                                        | `string`       | n/a           |   yes    |
| <a name="input_ATLAS_ORG_ID"></a> [ATLAS\_ORG\_ID](#input\_ATLAS\_ORG\_ID)                                                                 | Atlas organization name                    | `string`       | n/a           |   yes    |
| <a name="input_ATLAS_PROJECT_NAME"></a> [ATLAS\_PROJECT\_NAME](#input\_ATLAS\_PROJECT\_NAME)                                               | Atlas project name                         | `string`       | n/a           |   yes    |
| <a name="input_ATLAS_REGION"></a> [ATLAS\_REGION](#input\_ATLAS\_REGION)                                                                   | n/a                                        | `string`       | n/a           |   yes    |
| <a name="input_AWS_ACCESS_KEY_NO_CREDS"></a> [AWS\_ACCESS\_KEY\_NO\_CREDS](#input\_AWS\_ACCESS\_KEY\_NO\_CREDS)                            | IAM user with no permissions               | `string`       | n/a           |   yes    |
| <a name="input_AWS_CI_ROLE_TO_ASSUME"></a> [AWS\_CI\_ROLE\_TO\_ASSUME](#input\_AWS\_CI\_ROLE\_TO\_ASSUME)                                  | Role for the AWS provider to assume        | `string`       | n/a           |   yes    |
| <a name="input_AWS_SECRET_KEY_NO_CREDS"></a> [AWS\_SECRET\_KEY\_NO\_CREDS](#input\_AWS\_SECRET\_KEY\_NO\_CREDS)                            | IAM user with no permissions               | `string`       | n/a           |   yes    |
| <a name="input_GITHUB_TOKEN"></a> [GITHUB\_TOKEN](#input\_GITHUB\_TOKEN)                                                                   | n/a                                        | `string`       | n/a           |   yes    |
| <a name="input_MONGO_CLUSTER_NAME"></a> [MONGO\_CLUSTER\_NAME](#input\_MONGO\_CLUSTER\_NAME)                                               | n/a                                        | `string`       | n/a           |   yes    |
| <a name="input_atlas_cluster_instance_size"></a> [atlas\_cluster\_instance\_size](#input\_atlas\_cluster\_instance\_size)                  | n/a                                        | `string`       | `"M10"`       |    no    |
| <a name="input_aws_region"></a> [aws\_region](#input\_aws\_region)                                                                         | n/a                                        | `string`       | `"us-east-2"` |    no    |
| <a name="input_db_connection_string_secret_name"></a> [db\_connection\_string\_secret\_name](#input\_db\_connection\_string\_secret\_name) | Name of the DB connection string secret    | `string`       | n/a           |   yes    |
| <a name="input_environment"></a> [environment](#input\_environment)                                                                        | n/a                                        | `string`       | `"dev"`       |    no    |
| <a name="input_ip_access_list"></a> [ip\_access\_list](#input\_ip\_access\_list)                                                           | List of ip addresses with access to the db | `list(string)` | n/a           |   yes    |
| <a name="input_mongo_db"></a> [mongo\_db](#input\_mongo\_db)                                                                               | n/a                                        | `string`       | `"accountEd"` |    no    |
| <a name="input_mongo_db_cluster_name"></a> [mongo\_db\_cluster\_name](#input\_mongo\_db\_cluster\_name)                                    | Atlas db cluster name                      | `string`       | n/a           |   yes    |

## Outputs

No outputs.
<!-- END_TF_DOCS -->