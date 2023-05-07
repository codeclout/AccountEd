package protocols

type TransferBounds struct {
  DatabaseConnectionTimeout int64 `hcl:"database_connection_timeout"`
  MaxListItems              int64 `hcl:"default_list_count_limit"`
}
