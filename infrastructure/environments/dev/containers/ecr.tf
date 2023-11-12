module "workflow_token_generation" {
  source = "../../../modules/containers"

  ecr_name = "${var.environment}-workflow-token-generation"
}

module "notifications" {
  source = "../../../modules/containers"

  ecr_name = "${var.environment}-notifications"
}
