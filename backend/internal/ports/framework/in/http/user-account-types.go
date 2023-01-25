package ports

type UserAccountPort interface {
	HandleCreateAccountType(createAccountTypeInput *string) (interface{}, error)
	HandleListAccountTypes(listLimit *int16) (interface{}, error)
	HandleRemoveAccountType(accountType *string) (interface{}, error)
	HandleUpdateAccountType(newAccountType, accountTypeId *string) (interface{}, error)
	HandleFetchAccountType(accountTypeId *string) (interface{}, error)
}
