package account_types

type UserAccountTypeCorePort interface {
	NewAccountType(id *string, name *string, timestamp *string) (*NewAccountTypeOutput, error)
	ListAccountTypes(accountTypes []byte) (*[]NewAccountTypeOutput, error)
	DeleteAccountType(in []byte) (*NewAccountTypeOutput, error)
	UpdateAccountType(in []byte) (NewAccountTypeOutput, error)
	FetchAccountType(id []byte) (*NewAccountTypeOutput, error)
}

type NewAccountTypeOutput struct {
	AccountType string `json:"account_type"`
	CreatedAt   string `json:"created_at"`
	ID          string `json:"id"`
	ModifiedAt  string `json:"modified_at"`
}

type NewAccountTypeInput struct {
	AccountType string `json:"account_type"`
	CreatedAt   string `json:"created_at"`
	ModifiedAt  string `json:"modified_at"`
}
