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
  source = "../../modules/services/cloudflare/kv "

  step_one = {
    homeschooler = "I am representing my household",
    organization = "I am representing an organization of more than 250 people",
    small_group  = "I am representing a small group up to 250 people"
  }

  step_two = {
    include_online_portfolio_heading     = "Include online portfolio?"
    include_online_portfolio_description = "Your profile can include a portfolio comprised of all coursework"
  }

  step_one_header   = "Please tell us, what is your situation?"
  step_two_header   = "Tell us about your company"
  step_three_header = "Information"

  worker_kv_ns = "NS_UI_ONBOARDING_CONFIG"
}
