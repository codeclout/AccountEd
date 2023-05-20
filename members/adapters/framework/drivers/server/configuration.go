package server

import (
  "encoding/json"
  "errors"
  "fmt"
  "os"
  "path/filepath"

  "github.com/hashicorp/hcl/v2"
  "github.com/hashicorp/hcl/v2/hclsimple"
  "golang.org/x/exp/slog"

  "github.com/codeclout/AccountEd/pkg/monitoring"
)

const path = "./config.hcl"

type Server struct {
  GetOnlyConstraint bool   `hcl:"isAppGetOnly"`
  Name              string `hcl:"applicationName"`
  RoutePrefix       string `hcl:"routePrefix"`
}

type Adapter struct {
  log *slog.Logger
}

func NewAdapter(monitor *monitoring.Adapter) *Adapter {
  return &Adapter{
    log: monitor.Logger,
  }
}

func (a *Adapter) LoadMemberConfig() *map[string]interface{} {
  var configuration Server
  var out map[string]interface{}

  workingDirectory, _ := os.Getwd()
  fileLocation := filepath.Join(workingDirectory, path)

  e := hclsimple.DecodeFile(fileLocation, nil, &configuration)
  if e != nil {
    var x hcl.Diagnostics
    ok := errors.Is(e, x)

    if ok {
      a.log.Error("fatal", fmt.Sprintf("Failed to load runtime staticConfig: %s", e.(hcl.Diagnostics)[0].Summary))
      panic(e)
    } else {
      a.log.Error("fatal", fmt.Sprintf("Failed to get runtime staticConfig: %v", x))
      panic(e)
    }
  }

  b, _ := json.Marshal(&configuration)
  _ = json.Unmarshal(b, &out)

  return &out
}
