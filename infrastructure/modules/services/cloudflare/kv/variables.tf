variable "step_one" {
  type = map(string)

  validation {
    condition     = length(var.step_one) >= 2
    error_message = "Onboarding step 1 requires at least 2 user flow options"
  }
}

variable "step_two" {
  type = map(string)
}

variable "step_one_header" {
  type = string
}

variable "step_two_header" {
  type = string
}

variable "step_three_header" {
  type = string
}

variable "worker_kv_ns" {
  type = string
}
