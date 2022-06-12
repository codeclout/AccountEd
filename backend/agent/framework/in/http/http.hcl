application "http" {
  def_read_timeout          = "3s"
  def_write_timeout         = "3s"
  disable_third_party_ascii = true
  name                      = "AccountEd"
  network                   = "tcp4"
  printRoutes               = true
  response_encoding_suffix   = ".gz"
  stream_response_body      = false
  trustedProxies            = [""]
  withETag                  = true
  withTrustedProxies        = true
}