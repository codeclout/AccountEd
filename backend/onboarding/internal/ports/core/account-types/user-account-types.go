package account_types

type UserAccountTypeCorePort interface {
	ListAccountTypes(accountTypes *[]byte) (*[]NewAccountTypeOutput, error)
	FetchAccountType(id *[]byte) (*NewAccountTypeOutput, error)
}

type NewAccountTypeOutput struct {
	AccountType string `json:"account_type"`
	CreatedAt   string `json:"created_at"`
	ID          string `json:"_id"`
	ModifiedAt  string `json:"modified_at"`
}
