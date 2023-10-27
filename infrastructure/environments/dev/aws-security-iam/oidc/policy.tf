resource "aws_iam_policy" "tfc_policy" {
  name = "dev-terraform-cloud-ci-policy"

  policy = jsonencode({
    Version: "2012-10-17"
    Statement: [
      {
        Action = [
          "iam:CreateOpenIDConnectProvider",
          "iam:GetOpenIDConnectProvider",
        ]
        Effect = "Allow"
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "tfc_policy_attachment" {
  policy_arn = aws_iam_policy.tfc_policy.arn
  role       = aws_iam_role.oidc_role_tfc.name
}

resource "aws_iam_policy" "github_policy" {
  name = "dev-github-action-ci-policy"

  policy = jsonencode({
    Version: "2012-10-17"
    Statement: [
      {
        Action = [
          "ecr:PutImage"
        ]
        Effect = "Allow"
        Resource = "arn:aws:ecr:${var.AWS_REGION}:*:repository/${var.ENVIRONMENT}-*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "github_policy_attachment" {
  policy_arn = aws_iam_policy.github_policy
  role       = aws_iam_role.oidc_role_github
}