package gcp

type PostalCodeFrameworkIn interface {
	GetAddressDetails(address *string) (map[string]interface{}, error)
}
