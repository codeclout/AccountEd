output "policy_unencrypted_bucket" {
  value = one(aws_iam_policy.bucket_role_access[*].policy)
}
