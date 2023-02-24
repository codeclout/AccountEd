output "policy" {
  value = aws_iam_policy.bucket_role_access[*].policy
}
