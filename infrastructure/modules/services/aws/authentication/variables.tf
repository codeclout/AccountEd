variable "accountEd_user_pool" {
  type = string
}

variable "account_recovery_settings" {
  type = map({
    name     = string
    priority = number
  })
}

variable "allow_admin_create_user_only" {
  type = bool
}
