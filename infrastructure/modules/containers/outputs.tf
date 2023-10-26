output "registry_arn" {
  value = aws_ecr_repository.image_registry.arn
}

output "registry_url" {
  value = aws_ecr_repository.image_registry.repository_url
}