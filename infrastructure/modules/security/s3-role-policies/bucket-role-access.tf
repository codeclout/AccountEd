resource "aws_iam_policy" "bucket_role_access" {
  name        = "s3bucketRoleAccess"
  description = "Allows access to and S3 bucket from the root user and a role"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect    = "Deny"
        Principal = "*"
        Action    = "s3:*"
        Resource = [
          "arn:aws:s3:::${var.role_access_bucket_name}",
          "arn:aws:s3:::${var.role_access_bucket_name}/*"
        ],
        Condition = {
          StringNotLike = {
            "aws:userId" = [
              var.bucket_role_id,
              var.acccount_id
            ]
          }
        }
      }
    ]
  })
}
