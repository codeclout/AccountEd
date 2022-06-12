output "access_key_id" {
  value = aws_iam_access_key.ci_gh_key.id
}

output "access_key_secret" {
  value = aws_iam_access_key.ci_gh_key.encrypted_secret
}