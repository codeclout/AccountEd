data "aws_iam_user" "build_user" {
  user_name = "ci-svc-build-usr"
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
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          AWS = data.aws_iam_user.build_user.arn
        }
        Sid = ""
      },
      {
        Action = "sts:TagSession"
        Effect = "Allow"
        Principal = {
          AWS = data.aws_iam_user.build_user.arn
        }
        Sid = ""
      }
    ]
  })
}

resource "aws_iam_user_policy" "ci_svc_build_usr" {
  name = "buildUserPolicy"
  user = "ci-svc-build-usr"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = "sts:AssumeRole"
        Effect   = "Allow"
        Sid      = ""
        Resource = aws_iam_role.ecr_build_role.arn
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
