resource "aws_iam_policy" "tfc_policy" {
  name = "dev-terraform-cloud-ci-policy"

  policy = jsondecode({
    Version: "2012-10-17"
    Statement: [
      {
        Action = [
          "iam:CreateOpenIDConnectProvider",
          "iam:GetOpenIDConnectProvider",
        ]
        Effect: "Allow"
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "tfc_policy_attachment" {
  role       = aws_iam_role.oidc_role_tfc.name
  policy_arn = aws_iam_policy.tfc_policy.arn
}