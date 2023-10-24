terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.21.0, < 6.0.0"
    }
    tls = {
      source = "hashicorp/tls"
      version = "4.0.4, < 4.1.0"
    }
  }
}