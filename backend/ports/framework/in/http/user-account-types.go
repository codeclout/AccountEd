package ports

type UserAccountPort interface {
	HandleCreateAccountType(i interface{}) error
	HandleListAccountTypes(limit int64, i interface{}) error
	HandleRemoveAccountType(i interface{}) error
	HandleUpdateAccountType(id []byte, i interface{}) error
	HandleFetchAccountType(i interface{}) error
}
