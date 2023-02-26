data "aws_iam_policy_document" "unencrypted_bucket_role_access" {
  statement {
    effect = "Deny"
    sid    = "s3UnencryptedBucketRoleAccess"

    actions = ["s3:*"]

    resources = [
      "arn:aws:s3:::${var.role_access_bucket_name}",
      "arn:aws:s3:::${var.role_access_bucket_name}/*"
    ]

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    condition {
      test     = "StringNotLike"
      variable = "aws:userid"

      values = [
        "${var.bucket_role_id}:*",
        "${var.acccount_id}"
      ]
    }
  }
}
