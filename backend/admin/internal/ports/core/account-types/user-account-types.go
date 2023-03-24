package account_types

type UserAccountTypeCoreAdminPort interface {
	NewAccountType(id, name, timestamp *string) (*NewAccountTypeOutput, error)
	DeleteAccountType(in *[]byte) (*NewAccountTypeOutput, error)
	UpdateAccountType(count *int64, id *string) (*UpdatedAccountTypeOutput, error)
}

type NewAccountTypeInput struct {
	AccountType string `json:"account_type"`
	CreatedAt   string `json:"created_at"`
	ModifiedAt  string `json:"modified_at"`
}

type NewAccountTypeOutput struct {
	AccountType string `json:"account_type"`
	CreatedAt   string `json:"created_at"`
	ID          string `json:"_id"`
	ModifiedAt  string `json:"modified_at"`
}

type UpdatedAccountTypeOutput struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
}
