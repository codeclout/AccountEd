terraform {
  required_providers {
    mongodbatlas = {
      source  = "mongodb/mongodbatlas"
      version = ">= 1.6.0, < 2.0.0"
    }
  }
}

resource "mongodbatlas_project" "atlas_project" {
  name   = var.ATLAS_PROJECT_NAME
  org_id = var.ATLAS_ORG_ID
}

# Create an Atlas Admin Database User
resource "mongodbatlas_database_user" "db_user" {
  auth_database_name = "$external"
  aws_iam_type       = "ROLE"
  project_id         = mongodbatlas_project.atlas_project.id
  username           = var.mongo_db_role_arn

  roles {
    role_name     = "readWrite"
    database_name = var.mongo_db
  }
}

# Create a Shared Tier Cluster
resource "mongodbatlas_advanced_cluster" "atlas_cluster" {
  cluster_type = "REPLICASET"
  name         = "${var.environment}-${var.mongo_db_cluster_name}-cluster"
  project_id   = mongodbatlas_project.atlas_project.id

  backup_enabled = true

  replication_specs {
    regions_config {

      electable_specs {
        instance_size = var.atlas_cluster_instance_size
        node_count    = 3
      }

      analytics_specs {
        instance_size = var.atlas_cluster_instance_size
        node_count    = 1
      }

      provider_name = "AWS"
      priority      = 1
      region_name   = upper(var.aws_region)
    }

  }

  mongo_db_major_version = "6.0"
}
