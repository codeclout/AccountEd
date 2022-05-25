terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "3.15.0"
    }
  }
}

resource "cloudflare_workers_kv_namespace" "ui_config" {
  title = "NS_UI_ONBOARDING_CONFIG"
}

resource "cloudflare_workers_kv" "onboarding_step_one" {
  namespace_id = cloudflare_workers_kv_namespace.ui_config.id

  for_each = var.step_one
  key      = each.key
  value    = each.value
}

resource "cloudflare_workers_kv" "onboarding_step_one_header" {
  namespace_id = cloudflare_workers_kv_namespace.ui_config.id

  key   = "step_one_header"
  value = var.step_one_header
}
