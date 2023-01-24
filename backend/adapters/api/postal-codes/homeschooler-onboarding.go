package postal_codes

import (
  ports "github.com/codeclout/AccountEd/ports/core/postal-codes"
  "github.com/codeclout/AccountEd/ports/framework/in/gcp"
)

type Adapter struct {
  log            func(level, msg string)
  mapFramework   gcp.PostalCodeFrameworkIn
  postalCodeCore ports.PostalCodeCorePort
}

func NewAdapter(pcc ports.PostalCodeCorePort, mp gcp.PostalCodeFrameworkIn, logger func(level, msg string)) *Adapter {
  return &Adapter{
    log:            logger,
    mapFramework:   mp,
    postalCodeCore: pcc,
  }
}

func (a *Adapter) FetchPostalCodeDetails(address *string) (interface{}, error) {
  details, e := a.mapFramework.GetAddressDetails(address)
  if e != nil {
    a.log("error", e.Error())
    return nil, e
  }

  result, e := a.postalCodeCore.HandleFetchedPostalCode(details)
  if e != nil {
    return nil, e
  }

  return result, nil
}
