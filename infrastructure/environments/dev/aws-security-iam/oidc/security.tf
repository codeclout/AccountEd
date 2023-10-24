module "oidc_github" {
  source = "../../../../modules/security/oidc"

  identity_provider_audience = "sts.amazonaws.com"
  identity_provider_url      = "token.actions.githubusercontent.com"
}

module "oidc_terraform_cloud" {
  source = "../../../../modules/security/oidc"

  identity_provider_audience = "aws.workload.identity"
  identity_provider_url      = "app.terraform.io"
}