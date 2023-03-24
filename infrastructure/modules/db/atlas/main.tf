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

  api_keys {
    api_key_id = var.ATLAS_API_KEY_ID
    role_names = ["GROUP_OWNER"]
  }
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
    region_configs {

      electable_specs {
        instance_size = var.atlas_cluster_instance_size
        node_count    = 3
        priority      = 6
      }

      analytics_specs {
        instance_size = var.atlas_cluster_instance_size
        node_count    = 1
      }

      provider_name = var.cloud_provider
      priority      = 7
      region_name   = var.atlas_region
    }

  }

  mongo_db_major_version = var.mongodb_version
}
