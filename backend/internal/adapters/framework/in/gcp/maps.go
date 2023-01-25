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
  mapApiKey     string
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
    mapApiKey:     mapKey,
    runtimeConfig: runtimeConfig,
  }
}

func (a *Adapter) GetAddressDetails(address *string) (map[string]interface{}, error) {
  var data map[string]interface{}
  var host = "https://maps.googleapis.com/maps/api/geocode/json?"
  var query = "address=" + url.QueryEscape(*address) + "&" + "key=" + a.mapApiKey

  resp, e := http.Get(host + query)
  if e != nil {
    a.log("error", e.Error())
    return nil, e
  }

  defer func(Body io.ReadCloser) {
    e := Body.Close()
    if e != nil {
      a.log("error", e.Error())
    }
  }(resp.Body)

  b, e := io.ReadAll(resp.Body)
  if e != nil {
    a.log("error", e.Error())
    return nil, e
  }

  _ = json.Unmarshal(b, &data)
  results := data["results"].([]interface{})
  geoInfo := results[0].(map[string]interface{})

  return geoInfo, nil
}
