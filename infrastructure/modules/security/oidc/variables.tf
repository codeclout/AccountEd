variable "identity_provider_audience" {
  type = string
}

variable "identity_provider_url" {
  type = string
}

variable "tls_url" {
  type    = string
  default = ""
}

variable "use_tls_url" {
  type    = bool
  default = false
}
