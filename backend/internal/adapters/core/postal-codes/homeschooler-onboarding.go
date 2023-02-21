package postal_codes

import (
  "errors"

  port "github.com/codeclout/AccountEd/internal/ports/core/postal-codes"
)

type Adapter struct {
  log func(level, msg string)
}

func NewAdapter(l func(level, msg string)) *Adapter {
  return &Adapter{log: l}
}

func (a *Adapter) HandleFetchedPostalCode(geoInfo map[string]interface{}) (*port.GeoCodingOutput, error) {
  var localities []interface{}
  var out port.GeoCodingOutput

  if geoInfo["status"] == "ZERO_RESULTS" || geoInfo["status"] == "INVALID_REQUEST" {
    return &out, errors.New(geoInfo["status"].(string))
  }

  pl, ok := geoInfo["postcode_localities"].([]interface{})
  if ok {
    localities = pl
  }

  out = port.GeoCodingOutput{
    AddressComponents: geoInfo["address_components"].([]interface{}),
    FormattedAddress:  geoInfo["formatted_address"].(string),
    PlaceId:           geoInfo["place_id"].(string),
    PostalLocalities:  localities,
  }

  return &out, nil
}
