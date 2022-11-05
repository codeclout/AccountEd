package ports

type AccountPort interface {
	NewAccountType(id interface{}, name string, timestamp string) (NewAccountTypeOutput, error)
	ListAccountTypes(accountTypes []byte) ([]NewAccountTypeOutput, error)
}

type NewAccountTypeOutput struct {
	AccountType string      `json:"account_type"`
	CreatedAt   string      `json:"created_at"`
	ID          interface{} `json:"_id"`
	ModifiedAt  string      `json:"modified_at"`
}

type NewAccountTypeInput struct {
	AccountType string `json:"account_type"`
	CreatedAt   string `json:"created_at"`
	ModifiedAt  string `json:"modified_at"`
}
