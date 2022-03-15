terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=4.1.0"
    }
  }
}

provider "aws" {
  region = var.awsRegion
}

resource "aws_organizations_policy" "restrict_leave_org" {
  name = "PROD_OU_RESTRICT_LEAVE_ORG"

  content = <<CONTENT
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Deny",
        "Action": [
          "organizations:LeaveOrganization"
        ],
        "Resource": [ "*" ]
      }
    ]
  }
  CONTENT
}

resource "aws_organizations_policy" "restrict_root" {
  name = "ROOT_USER_RESTRICT"

  content = <<CONTENT
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Deny",
        "Action": [ "*" ],
        "Resource": [ "*" ],
        "Condition: {
          "ForAllValues:StringLike": {
            "aws:PrincipalArn": [
              "arn:aws:iam::*:root"
            ]
          }
        }
      }
    ]
  }
  CONTENT
}

resource "aws_organizations_policy" "prod_logging" {
  name = "PROD_OU_LOGGING"

  content = <<CONTENT
    {
      "Version": "2012-10-17",
      "Statement": [
        {
          "Action": [
            "cloudtrail:StopLogging",
            "cloudtrail:DeleteTrail"
          ],
          "Resource": "*",
          "Effect": "Deny"
        },
        {
          "Action": [
            "config:DeleteConfigRule",
            "config:DeleteConfigurationRecorder",
            "config:DeleteDeliveryChannel",
            "config:StopConfigurationRecorder"
          ],
          "Resource": "*",
          "Effect": "Deny"
        },
        {
          "Action": [
            "guardduty:DeleteDetector",
            "guardduty:DeleteInvitations",
            "guardduty:DeleteIPSet",
            "guardduty:DeleteMembers",
            "guardduty:DeleteThreatIntelSet",
            "guardduty:DisassociateFromMasterAccount",
            "guardduty:DisassociateMembers",
            "guardduty:StopMonitoringMembers",
            "guardduty:UpdateDetector"
          ],
          "Resource": "*",
          "Effect": "Deny"
        },
        {
          "Action": [
            "securityhub:DeleteInvitations",
            "securityhub:DisableSecurityHub",
            "securityhub:DisassociateFromMasterAccount",
            "securityhub:DeleteMembers",
            "securityhub:DisassociateMembers"
          ],
          "Resource": "*",
          "Effect": "Deny"
        }
      ]
    }
  CONTENT
}

resource "aws_organizations_policy" "restrict_billing" {
  name = "ROOT_RESTRICT_BILLING"

  content = <<CONTENT
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Deny",
        "Action": [
          "aws-portal:ModifyAccount",
          "aws-portal:ModifyBilling",
          "aws-portal:ModifyPaymentMethods"
        ],
        "Resource": [ "*" ]
      }
    ]
  }
  CONTENT
}

resource "aws_organizations_policy" "prod_monitoring" {
  name = "PROD_MONITORING"

  content = <<CONTENT
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Deny",
        "Action": [
          "aws-portal:ModifyAccount",
          "aws-portal:ModifyBilling",
          "aws-portal:ModifyPaymentMethods"
        ],
        "Resource": [ "*" ]
      }
    ]
  }
  CONTENT
}

resource "aws_organizations_policy_attachment" "root_restrict_billing" {
  policy_id = aws_organizations_policy.restrict_billing.id
  target_id = aws_organizations_organization.AccountEdManagement.roots[0].id
}

resource "aws_organizations_policy_attachment" "root_restrict_root" {
  policy_id = aws_organizations_policy.restrict_root.id
  target_id = aws_organizations_organization.AccountEdManagement.roots[0].id
}

resource "aws_organizations_policy_attachment" "ou_prod_restrict_leave_org" {
  policy_id = aws_organizations_policy.restrict_leave_org.id
  target_id = aws_organizations_organizational_unit.prod.id
}

resource "aws_organizations_policy_attachment" "ou_prod_logging" {
  policy_id = aws_organizations_policy.prod_logging.id
  target_id = aws_organizations_organizational_unit.prod.id
}

resource "aws_organizations_policy_attachment" "ou_prod_monitoring" {
  policy_id = aws_organizations_policy.prod_monitoring.id
  target_id = aws_organizations_organizational_unit.prod.id
}

resource "aws_organizations_organization" "AccountEdManagement" {
  aws_service_access_principals = [
    "account.amazonaws.com",
    "cloudtrail.amazonaws.com",
    "servicecatalog.amazonaws.com"
  ]
}

resource "aws_organizations_organizational_unit" "prod" {
  name      = "production"
  parent_id = aws_organizations_organization.AccountEdManagement.roots[0].id
}

resource "aws_organizations_organizational_unit" "prodLMS" {
  name      = "prooduction LMS"
  parent_id = aws_organizations_organizational_unit.prod.id
}

resource "aws_organizations_organizational_unit" "prodBibleCollege" {
  name      = "production Bible College"
  parent_id = aws_organizations_organizational_unit.prodLMS.id
}

resource "aws_organization_account" "lms_users" {
  name      = "LMS Users"
  parent_id = aws_organizations_organizational_unit.prodLMS.id
  email     = var.LMS_ACCOUNT_EMAIL
  role_name = var.LMS_ACCOUNT_ROLE_NAME
}

resource "aws_organization_account" "bible_college_users" {
  name      = "Bible College Users"
  parent_id = aws_organizations_organizational_unit.prodBibleCollege.id
  email     = var.PROXY_ACCOUNT_USERS_EMAIL
  role_name = var.PROXY_ACCOUNT_ROLE_NAME
}

resource "aws_organizations_organizational_unit" "stage" {
  name      = "staging"
  parent_id = aws_organizations_organization.AccountEdManagement.roots[0].id
}

resource "aws_organizations_organizational_unit" "ref" {
  name      = "qa-reference"
  parent_id = aws_organizations_organization.AccountEdManagement.roots[0].id
}

resource "aws_organizations_organizational_unit" "dev" {
  name      = "development"
  parent_id = aws_organizations_organization.AccountEdManagement.roots[0].id
}

resource "aws_organizations_organizational_unit" "AccountEdAppProd" {
  name      = "prodAccountEdApp"
  parent_id = aws_organizations_organizational_unit.prod.id
}

resource "aws_organizations_organizational_unit" "AccountEdAppStage" {
  name      = "stageAccountEdApp"
  parent_id = aws_organizations_organizational_unit.stage.id
}

resource "aws_organizations_organizational_unit" "AccountEdAppRef" {
  name      = "refAccountEdApp"
  parent_id = aws_organizations_organizational_unit.ref.id
}

resource "aws_organizations_organizational_unit" "AccountEdAppDev" {
  name      = "devAccountEdApp"
  parent_id = aws_organizations_organizational_unit.dev.id
}