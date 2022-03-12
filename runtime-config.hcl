organization = "cosis-io"

acls {
  read  = "service:read"
  write = "service:write"
}

cloudProvider {
  aws {
    role_duration = 3600
  }
}

healthcheck {
  interval = "3s"
  retries  = 3
  timeout  = "7s"
}

service "http" {
  address              = "0.0.0.0"
  cache_driver         = "redis"
  db_driver            = "mongodb"
  db_connection_string = ""
  port                 = 8088
  use_cache            = true
}
// workspaces { name = "accountEd-${env.ENV}-org-management" }