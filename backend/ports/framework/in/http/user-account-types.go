package ports

type UserAccountPort interface {
	HandleCreateAccountType(i interface{}) error
	HandleListAccountTypes(i interface{}) error
}
