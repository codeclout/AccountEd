resource "aws_iam_openid_connect_provider" "oidc_provider" {
  url = "https://${var.identity_provider_url}"

  client_id_list = [
    var.identity_provider_audience,
  ]

  thumbprint_list = [data.tls_certificate.thumbprint.certificates[0].sha1_fingerprint]
}

data "tls_certificate" "thumbprint" {
  url = var.use_tls_url ? var.tls_url : "https://${var.identity_provider_url}"
}
