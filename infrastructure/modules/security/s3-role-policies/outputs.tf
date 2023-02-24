output "policy_unencrypted_bucket" {
  value = aws_iam_policy.bucket_role_access[*]
}
