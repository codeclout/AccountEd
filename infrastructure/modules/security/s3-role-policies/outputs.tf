output "policy_unencrypted_bucket" {
  value = toset(aws_iam_policy.bucket_role_access[*].policy)
}
