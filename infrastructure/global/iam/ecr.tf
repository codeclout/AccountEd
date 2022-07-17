resource "aws_iam_role" "ecr_build_role" {
  name = "ecrBuildRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::*:user/ci-svc-build-usr"
        }
        Sid = ""
      }
    ]
  })
}

resource "aws_iam_policy" "ecr_authorization_policy" {
  name = "ecrAuthorization"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ecr:GetAuthorizationToken"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_policy" "ecr_private_push_pull_policy" {
  name = "ecr-push-pull-private"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ecr:BatchGetImage",
          "ecr:BatchCheckLayerAvailability",
          "ecr:CompleteLayerUpload",
          "ecr:GetDownloadUrlForLayer",
          "ecr:InitiateLayerUpload",
          "ecr:PutImage",
          "ecr:UploadLayerPart"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecr_authorization" {
  role       = aws_iam_role.ecr_build_role.name
  policy_arn = aws_iam_policy.ecr_authorization_policy.arn
}

resource "aws_iam_role_policy_attachment" "ecr_private_push_pull" {
  role       = aws_iam_role.ecr_build_role.name
  policy_arn = aws_iam_policy.ecr_private_push_pull_policy.arn
}
