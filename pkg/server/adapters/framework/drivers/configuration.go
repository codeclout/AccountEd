package drivers

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"

  "github.com/hashicorp/hcl/v2"
  "github.com/hashicorp/hcl/v2/hclsimple"
  "golang.org/x/exp/slog"

  "github.com/codeclout/AccountEd/pkg/server/server-types/protocols"
)

const path = "./config.hcl"

type Adapter struct {
  log *slog.Logger
}

func NewAdapter(log *slog.Logger) *Adapter {
  return &Adapter{
    log: log,
  }
}

func (a *Adapter) Load() *map[string]interface{} {
  var bounds protocols.TransferBounds
  var out map[string]interface{}

  workingDirectory, _ := os.Getwd()
  fileLocation := filepath.Join(workingDirectory, path)

  e := hclsimple.DecodeFile(fileLocation, nil, &bounds)
  if e != nil {
    x, ok := e.(hcl.Diagnostics)

    if ok {
      a.log.Error("fatal", fmt.Sprintf("Failed to load runtime staticConfig: %s", e.(hcl.Diagnostics)[0].Summary))
      panic(e)
    } else {
      a.log.Error("fatal", fmt.Sprintf("Failed to get runtime staticConfig: %v", x))
      panic(e)
    }
  }

  b, _ := json.Marshal(&bounds)
  _ = json.Unmarshal(b, &out)

  return &out
}
