output "container_repository_arn" {
  value = aws_ecr_repository.container_repo.arn
}

output "container_repository_url" {
  value = aws_ecr_repository.container_repo.repository_url
}
