output "connection_strings" {
  value = mongodbatlas_cluster.atlas_cluster.connection_strings.0.standard_srv
}
