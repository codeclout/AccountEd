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

| Name                                                                                                                   | Description                                  | Type           | Default | Required |
| ---------------------------------------------------------------------------------------------------------------------- | -------------------------------------------- | -------------- | ------- | :------: |
| <a name="input_atlas_org_id"></a> [atlas\_org\_id](#input\_atlas\_org\_id)                                             | Atlas Organization ID                        | `string`       | n/a     |   yes    |
| <a name="input_atlas_project_name"></a> [atlas\_project\_name](#input\_atlas\_project\_name)                           | Atlas Project Name                           | `string`       | n/a     |   yes    |
| <a name="input_atlas_region"></a> [atlas\_region](#input\_atlas\_region)                                               | Atlas region where resources will be created | `string`       | n/a     |   yes    |
| <a name="input_cloud_provider"></a> [cloud\_provider](#input\_cloud\_provider)                                         | AWS or GCP or Azure                          | `string`       | n/a     |   yes    |
| <a name="input_cluster_instance_size_name"></a> [cluster\_instance\_size\_name](#input\_cluster\_instance\_size\_name) | Cluster instance size name                   | `string`       | n/a     |   yes    |
| <a name="input_environment"></a> [environment](#input\_environment)                                                    | The environment to be built                  | `string`       | n/a     |   yes    |
| <a name="input_ip_address"></a> [ip\_address](#input\_ip\_address)                                                     | IP address used to access Atlas cluster      | `list(string)` | n/a     |   yes    |
| <a name="input_mongodb_version"></a> [mongodb\_version](#input\_mongodb\_version)                                      | MongoDB Version                              | `string`       | n/a     |   yes    |

## Outputs

| Name                                                                                         | Description                                         |
| -------------------------------------------------------------------------------------------- | --------------------------------------------------- |
| <a name="output_connection_strings"></a> [connection\_strings](#output\_connection\_strings) | Use terraform output to display connection strings. |
<!-- END_TF_DOCS -->