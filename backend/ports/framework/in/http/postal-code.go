package ports

type PostalCodePort interface {
  HandleFetchPostalCodeDetails(address *string) (interface{}, error)
}
