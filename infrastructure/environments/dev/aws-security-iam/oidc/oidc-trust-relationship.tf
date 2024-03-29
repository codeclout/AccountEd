resource "aws_iam_role" "oidc_role_tfc" {
  name = "terraform-cloud-oidc-role"
  tags = {
    "terraform-cloud-scope": "workspace"
  }

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect : "Allow",
        Principal : {
          Federated : module.oidc_terraform_cloud.openid-connect-provider-arn
        },
        Action : "sts:AssumeRoleWithWebIdentity",
        Condition : {
          StringEquals : {
            "${module.oidc_terraform_cloud.openid-connect-provider-hostname}:aud" : one(module.oidc_terraform_cloud.openid-connect-provider-client-id-list)
          },
          StringLike : {
            "${module.oidc_terraform_cloud.openid-connect-provider-hostname}:sub" : "organization:${var.tfc_organization_name}:project:*:workspace:*:run_phase:*"
          }
        }
      }
    ]
  })
}

resource "aws_iam_role" "oidc_role_github" {
  name = "github-oidc-role"
  tags = {
    "github-scope": "gh-action"
  }

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect : "Allow",
        Principal : {
          Federated : module.oidc_github.openid-connect-provider-arn
        },
        Action : "sts:AssumeRoleWithWebIdentity",
        Condition : {
          StringLike : {
            "${module.oidc_github.openid-connect-provider-hostname}:sub" : [
              "repo:codeclout/AccountEd:*"
            ]
          }
          StringEquals : {
            "${module.oidc_github.openid-connect-provider-hostname}:aud" : one(module.oidc_github.openid-connect-provider-client-id-list)
          }
        }
      }
    ]
  })
}

resource "aws_iam_role" "oidc_role_gitlab" {
  name = "gitlab-oidc-role"
  tags = {
    "gitlab-scope": "gitlab-ci"
  }

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect : "Allow",
        Principal : {
          Federated : module.oidc_gitlab.openid-connect-provider-arn
        },
        Action : "sts:AssumeRoleWithWebIdentity",
        Condition : {
          StringLike : {
            "${module.oidc_gitlab.openid-connect-provider-hostname}:sub" : "project_path:sch00l.io/*:ref_type:branch:ref:main"
          }
        }
      }
    ]
  })
}
