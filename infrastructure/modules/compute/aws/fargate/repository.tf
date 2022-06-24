resource "aws_ecr_repository" "container_repo" {
  name                 = "${var.app}-${var.environment}"
  image_tag_mutability = var.image_tag_mutability

  image_scanning_configuration {
    scan_on_push = var.should_scan_image_on_push
  }
}
