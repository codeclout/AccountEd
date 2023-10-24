resource "aws_iam_role" "oidc_role_tfc" {
  name = "oidc-role"

  assume_role_policy = <<EOF
    {
      "Version": "2012-10-17",
      "Statement": [
        {
           "Effect": "Allow",
           "Principal": {
             "Federated": "${module.oidc_terraform_cloud.openid-connect-provider-arn}"
           },
           "Action": "sts:AssumeRoleWithWebIdentity",
           "Condition": {
             "StringEquals": {
               "${module.oidc_terraform_cloud.openid-connect-provider-hostname}:aud": "${one(module.oidc_terraform_cloud.openid-connect-provider-client-id-list)}"
             },
             "StringLike": {
               "${module.oidc_terraform_cloud.openid-connect-provider-hostname}:sub": "organization:${var.tfc_organization_name}:project:${var.tfc_project_name}:workspace:${var.tfc_workspace_name}:run_phase:*"
             }
          }
        }
      ]
    }
  EOF
}

resource "aws_iam_role" "oidc_role_github" {
  assume_role_policy = <<EOF
    {
      "Version": "2012-10-17",
      "Statement": [
        {
          "Effect": "Allow",
          "Principal": {
            "Federated": "${module.oidc_github.openid-connect-provider-arn}"
          },
          "Action": "sts:AssumeRoleWithWebIdentity",
          "Condtion": {
            "StringEquals": {
              "${module.oidc_github.openid-connect-provider-hostname}:aud": "${one(module.oidc_github.openid-connect-provider-client-id-list)}"
              "${module.oidc_github.openid-connect-provider-hostname}:sub": [
                "repo:codeclout/AccountEd:ref:refs/heads/alpha
              ]
            }
          }
        }
      ]
    }
  EOF
}