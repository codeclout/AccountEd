package postal_codes

import port "github.com/codeclout/AccountEd/ports/core/postal-codes"

type Adapter struct {
  log func(level, msg string)
}

func NewAdapter(l func(level, msg string)) *Adapter {
  return &Adapter{log: l}
}

func (a *Adapter) HandleFetchedPostalCode(geoInfo map[string]interface{}) (*port.GeoCodingInput, error) {
  var localities []interface{}
  var out port.GeoCodingInput

  pl, ok := geoInfo["postcode_localities"].([]interface{})
  if ok {
    localities = pl
  }

  out = port.GeoCodingInput{
    AddressComponents: geoInfo["address_components"].([]interface{}),
    FormattedAddress:  geoInfo["formatted_address"].(string),
    PlaceId:           geoInfo["place_id"].(string),
    PostalLocalities:  localities,
  }

  return &out, nil
}
