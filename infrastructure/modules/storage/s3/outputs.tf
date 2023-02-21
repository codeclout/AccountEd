output "bucket_name" {
  value = aws_s3_bucket.newBucket_unencrypted[*].id
}
