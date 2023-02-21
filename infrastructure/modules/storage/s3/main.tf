terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.55.0, < 5.0.0"
    }
  }
}

resource "aws_s3_bucket" "newBucket_unencrypted" {
  count = length(var.buckets_unencrypted)

  bucket              = lower(var.buckets_unencrypted[count.index])
  object_lock_enabled = var.is_object_locked[count.index] // WORM - write once read many
}
