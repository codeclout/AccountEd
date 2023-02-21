package postal_codes

type PostalCodeCorePort interface {
  HandleFetchedPostalCode(geoInfo map[string]interface{}) (*GeoCodingOutput, error)
}

type GeoCodingOutput struct {
  AddressComponents []interface{} `json:"address_components"`
  FormattedAddress  string        `json:"formatted_address"`
  PlaceId           string        `json:"place_id"`
  PostalLocalities  []interface{} `json:"postcode_localities"`
}
