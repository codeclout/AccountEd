package gcp

import (
  "encoding/json"
  "io"
  "net/http"
  "net/url"

  "googlemaps.github.io/maps"
)

type Adapter struct {
  client        *maps.Client
  log           func(level, msg string)
  runtimeConfig map[string]interface{}
}

func NewAdapter(runtimeConfig map[string]interface{}, logger func(level, msg string)) *Adapter {
  mapKey, ok := runtimeConfig["MapKey"].(string)
  if !ok {
    logger("fatal", "maps api key is not a string")
  }

  client, e := maps.NewClient(maps.WithAPIKey(mapKey))
  if e != nil {
    logger("fatal", e.Error())
  }

  return &Adapter{
    client:        client,
    log:           logger,
    runtimeConfig: runtimeConfig,
  }
}

func (a *Adapter) GetAddressDetails(address *string) (map[string]interface{}, error) {
  var data map[string]interface{}
  var uri = "https://maps.googleapis.com/maps/api/geocode/json?" + "address=" + url.QueryEscape(*address) + "&" + "key=" + a.runtimeConfig["MapKey"].(string)

  resp, _ := http.Get(uri)

  defer func(Body io.ReadCloser) {
    e := Body.Close()
    if e != nil {
      a.log("error", e.Error())
    }
  }(resp.Body)

  b, _ := io.ReadAll(resp.Body)

  _ = json.Unmarshal(b, &data)
  results := data["results"].([]interface{})
  geoInfo := results[0].(map[string]interface{})

  return geoInfo, nil
}
