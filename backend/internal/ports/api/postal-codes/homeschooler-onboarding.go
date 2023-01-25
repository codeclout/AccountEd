package postal_codes

type PostalCodeApiPort interface {
	FetchPostalCodeDetails(address *string) (interface{}, error)
}
