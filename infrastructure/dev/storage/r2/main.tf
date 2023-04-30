terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.20.1, < 5.0.0"
    }
  }
}

provider "aws" {
  access_key                  = var.R2_ACCESS_KEY
  secret_key                  = var.R2_SECRET_KEY
  skip_credentials_validation = true
  skip_region_validation      = true
  skip_requesting_account_id  = true
  endpoints {
    s3 = var.CF_ACCOUNT_R2_ENDPOINT // "https://<account id>.r2.cloudflarestorage.com"
  }
}

module "edge_storage" {
  source = "../../../modules/storage/r2"

}
