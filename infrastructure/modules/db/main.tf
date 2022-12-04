terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.44.0, < 5.0.0"
    }
    mongodbatlas = {
      source  = "mongodb/mongodbatlas"
      version = ">= 1.6.0, < 2.0.0"
    }
  }
}

# Create a Project
resource "mongodbatlas_project" "atlas-project" {
  org_id = var.atlas_org_id
  name   = var.atlas_project_name
}

# Create a Database User
resource "mongodbatlas_database_user" "db-user" {
  username           = "user-1"
  password           = random_password.db-user-password.result
  project_id         = mongodbatlas_project.atlas-project.id
  auth_database_name = "admin"
  roles {
    role_name     = "readWrite"
    database_name = "${var.atlas_project_name}-db"
  }
}

# Create a Database Password
resource "random_password" "db-user-password" {
  length           = 16
  special          = true
  override_special = "_%@"
}
