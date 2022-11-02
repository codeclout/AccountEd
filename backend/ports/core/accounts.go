package ports

type AccountPort interface {
	NewAccountType(id interface{}, name string, timestamp string) (NewAccountTypeOutput, error)
}

type NewAccountTypeOutput struct {
	AccountType string      `json:"accountType"`
	CreatedAt   string      `json:"createdAt"`
	ID          interface{} `json:"id"`
	ModifiedAt  string      `json:"modifiedAt"`
}

type NewAccountTypeInput struct {
	AccountType string `json:"accountType"`
	CreatedAt   string `json:"createdAt"`
	ModifiedAt  string `json:"modifiedAt"`
}
