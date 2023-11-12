output "registry_arn" {
  value = module.workflow_token_generation.registry_arn
}

output "registry_url" {
  value = module.workflow_token_generation.registry_url
}

output "notifications_registry_arn" {
  value = module.notifications.registry_arn
}

output "notifications_registry_url" {
  value = module.notifications.registry_url
}
