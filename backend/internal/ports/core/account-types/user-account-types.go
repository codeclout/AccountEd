package account_types

type UserAccountTypeCorePort interface {
	NewAccountType(id, name, timestamp *string) (*NewAccountTypeOutput, error)
	ListAccountTypes(accountTypes *[]byte) (*[]NewAccountTypeOutput, error)
	DeleteAccountType(in *[]byte) (*NewAccountTypeOutput, error)
	UpdateAccountType(count *int64, id *string) (*UpdatedAccountTypeOutput, error)
	FetchAccountType(id *[]byte) (*NewAccountTypeOutput, error)
}

type NewAccountTypeOutput struct {
	AccountType string `json:"account_type"`
	CreatedAt   string `json:"created_at"`
	ID          string `json:"_id"`
	ModifiedAt  string `json:"modified_at"`
}

type NewAccountTypeInput struct {
	AccountType string `json:"account_type"`
	CreatedAt   string `json:"created_at"`
	ModifiedAt  string `json:"modified_at"`
}

type UpdatedAccountTypeOutput struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
}
