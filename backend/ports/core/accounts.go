package ports

type AccountPort interface {
	NewAccountType(insertId interface{}, name string) (NewAccountTypeOutput, error)
}

type NewAccountTypeOutput struct {
	AccountType string
	ID          interface{}
}
