<!-- BEGIN_TF_DOCS -->
## Requirements

| Name                                                                               | Version           |
| ---------------------------------------------------------------------------------- | ----------------- |
| <a name="requirement_mongodbatlas"></a> [mongodbatlas](#requirement\_mongodbatlas) | >= 1.6.0, < 2.0.0 |

## Providers

| Name                                                                         | Version           |
| ---------------------------------------------------------------------------- | ----------------- |
| <a name="provider_mongodbatlas"></a> [mongodbatlas](#provider\_mongodbatlas) | >= 1.6.0, < 2.0.0 |

## Modules

No modules.

## Resources

| Name                                                                                                                                               | Type     |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| [mongodbatlas_advanced_cluster.atlas_cluster](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/advanced_cluster) | resource |
| [mongodbatlas_database_user.db_user](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/database_user)             | resource |
| [mongodbatlas_project.atlas_project](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/project)                   | resource |

## Inputs

| Name                                                                                                                      | Description                                  | Type           | Default       | Required |
| ------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------- | -------------- | ------------- | :------: |
| <a name="input_ATLAS_API_KEY_ID"></a> [ATLAS\_API\_KEY\_ID](#input\_ATLAS\_API\_KEY\_ID)                                  | Atlas org API key ID                         | `string`       | n/a           |   yes    |
| <a name="input_ATLAS_ORG_ID"></a> [ATLAS\_ORG\_ID](#input\_ATLAS\_ORG\_ID)                                                | Atlas Organization ID                        | `string`       | n/a           |   yes    |
| <a name="input_ATLAS_PROJECT_NAME"></a> [ATLAS\_PROJECT\_NAME](#input\_ATLAS\_PROJECT\_NAME)                              | Atlas Project Name                           | `string`       | n/a           |   yes    |
| <a name="input_atlas_cluster_instance_size"></a> [atlas\_cluster\_instance\_size](#input\_atlas\_cluster\_instance\_size) | n/a                                          | `string`       | `"M10"`       |    no    |
| <a name="input_atlas_region"></a> [atlas\_region](#input\_atlas\_region)                                                  | Atlas region where resources will be created | `string`       | n/a           |   yes    |
| <a name="input_cloud_provider"></a> [cloud\_provider](#input\_cloud\_provider)                                            | AWS, GCP or Azure                            | `string`       | `"AWS"`       |    no    |
| <a name="input_environment"></a> [environment](#input\_environment)                                                       | The environment to be built                  | `string`       | n/a           |   yes    |
| <a name="input_ip_address"></a> [ip\_address](#input\_ip\_address)                                                        | IP address used to access Atlas cluster      | `list(string)` | n/a           |   yes    |
| <a name="input_mongo_db"></a> [mongo\_db](#input\_mongo\_db)                                                              | n/a                                          | `string`       | `"accountEd"` |    no    |
| <a name="input_mongo_db_cluster_name"></a> [mongo\_db\_cluster\_name](#input\_mongo\_db\_cluster\_name)                   | n/a                                          | `string`       | n/a           |   yes    |
| <a name="input_mongo_db_role_arn"></a> [mongo\_db\_role\_arn](#input\_mongo\_db\_role\_arn)                               | ARN of role to assume for access to the db   | `string`       | n/a           |   yes    |
| <a name="input_mongodb_version"></a> [mongodb\_version](#input\_mongodb\_version)                                         | MongoDB Version                              | `string`       | `"6.0"`       |    no    |

## Outputs

| Name                                                                                         | Description |
| -------------------------------------------------------------------------------------------- | ----------- |
| <a name="output_cluster_id"></a> [cluster\_id](#output\_cluster\_id)                         | n/a         |
| <a name="output_connection_strings"></a> [connection\_strings](#output\_connection\_strings) | n/a         |
<!-- END_TF_DOCS -->