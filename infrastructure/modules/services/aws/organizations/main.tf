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
  email     = var.LMS_ACCOUNT_EMAIL
  role_name = var.LMS_ACCOUNT_ROLE
}

resource "aws_organization_account" "bible_college_users" {
  name      = "Bible College Users"
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