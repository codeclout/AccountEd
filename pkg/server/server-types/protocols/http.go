package protocols

type ServerProtocolHttpMetadata struct {
	RoutePrefix      string
	ServerName       string
	UseOnlyGetRoutes bool
}

type TransferBounds struct {
	DatabaseConnectionTimeout int64 `hcl:"database_connection_timeout"`
	MaxListItems              int64 `hcl:"default_list_count_limit"`
}
