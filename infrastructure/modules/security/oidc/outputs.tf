output "openid-connect-provider-arn" {
  value = aws_iam_openid_connect_provider.oidc_provider.arn
}

output "openid-connect-provider-audience" {
  value = var.identity_provider_audience
}

output "openid-connect-provider-client-id-list" {
  value = aws_iam_openid_connect_provider.oidc_provider.client_id_list
}

output "openid-connect-provider-hostname" {
  value = var.identity_provider_url
}