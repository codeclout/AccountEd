terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=4.1.0"
    }
  }
}

resource "aws_iam_user" "ci_gh" {
  name = "ci-gh-user"
  path = "/ci/"
}

resource "aws_iam_access_key" "ci_gh_key" {
  pgp_key = "keybase:codeclout"
  user    = aws_iam_user.ci_gh.name

}