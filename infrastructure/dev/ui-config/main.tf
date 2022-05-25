terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "3.15.0"
    }
  }
}

provider "cloudflare" {
  account_id = var.CF_ACCOUNT_ID
  api_token  = var.CF_API_TOKEN
  email      = var.CF_SERVICE_ACCOUNT_EMAIL
}

module "ui_onboarding_config" {
  source = "../modules/cloudflare-onboarding-worker-kv"

  step_one = {
    homeschooler = "I am representing my household",
    organization = "I am representing an organization of more than 250 people",
    small_group  = "I am representing a small group up to 250 people"
  }

  step_one_header = "Please tell us, what is your situation?"
}
