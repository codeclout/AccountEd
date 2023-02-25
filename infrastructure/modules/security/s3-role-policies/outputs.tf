output "policy_unencrypted_bucket" {
  value = toSet(aws_iam_policy.bucket_role_access[*].policy)
}
