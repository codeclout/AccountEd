output "policy_unencrypted_bucket" {
  value = data.aws_iam_policy_document.unencrypted_bucket_role_access[*].json
}
