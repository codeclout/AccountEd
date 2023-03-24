package http

type UserAccountPort interface {
	HandleListAccountTypes(listLimit *int16) (interface{}, error)
	HandleFetchAccountType(accountTypeId *string) (interface{}, error)
}
