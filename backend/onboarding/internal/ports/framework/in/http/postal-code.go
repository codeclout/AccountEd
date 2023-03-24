package http

type PostalCodePort interface {
  HandleFetchPostalCodeDetails(address *string) (interface{}, error)
}
