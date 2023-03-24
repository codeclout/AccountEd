package http

type UserAccountFrameworkAdminPort interface {
	HandleCreateAccountType(createAccountTypeInput *string) (interface{}, error)
	HandleRemoveAccountType(accountType *string) (interface{}, error)
	HandleUpdateAccountType(newAccountType, accountTypeId *string) (interface{}, error)
}
