output "connection_strings" {
  value = mongodbatlas_advanced_cluster.atlas_cluster.connection_strings.0.standard_srv
}

output "cluster_id" {
  value = mongodbatlas_advanced_cluster.atlas_cluster.cluster_id
}
