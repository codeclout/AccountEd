resource "aws_iam_openid_connect_provider" "main" {
  url             = "https://token.actions.githubusercontent.com"
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = ["6938fd4d98bab03faadb97b34396831e3780aea1"]
}

data "aws_ecr_repository" "service" {
  name = "sch00l.io-dev"
}

resource "aws_iam_role" "ecr_build_role" {
  name = "ecrBuildRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRoleWithIdentity"
        Effect = "Allow"
        Principal = {
          Federated = aws_iam_openid_connect_provider.main.arn
        }
        Sid = ""
        Condition = {
          StringEquals = {
            "token.actions.githubusercontent.com:aud" : "sts.amazonaws.com",
            "token.actions.githubusercontent.com:sub" : "repo:codeclout/AccountEd:*"
          }
        }
      }
    ]
  })
}

resource "aws_iam_policy" "ecr_push_private" {
  name = "ecrPushPrivate"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ecr:BatchCheckLayerAvailability",
          "ecr:CompleteLayerUpload",
          "ecr:GetAuthorizationToken",
          "ecr:InitiateLayerUpload",
          "ecr:PutImage",
          "ecr:UploadLayerPart"
        ]
        Effect   = "Allow"
        Resource = data.aws_ecr_repository.service.arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecr_push_private" {
  role       = aws_iam_role.ecr_build_role.name
  policy_arn = aws_iam_policy.ecr_push_private.arn
}
