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
)

const path = "./config.hcl"

type Server struct {
	GetOnlyConstraint bool    `hcl:"is_app_get_only" json:"is_app_get_only"`
	Name              string  `hcl:"application_name" json:"application_name"`
	RoutePrefix       string  `hcl:"route_prefix" json:"route_prefix"`
	SLARoutes         float64 `hcl:"sla_routes" json:"sla_routes"`
}

type Adapter struct {
	log *slog.Logger
}

func NewAdapter(log *slog.Logger) *Adapter {
	return &Adapter{
		log: log,
	}
}

// LoadMemberConfig reads and decodes the configuration file located by the "path" constant, and
// populates a Server instance. The Server instance is then converted to a JSON object which is
// unmarshalled into a map[string]interface{}. This method returns a pointer to the map.
// In case of any errors, it logs the error message and panics with the corresponding error.
func (a *Adapter) LoadMemberConfig() *map[string]interface{} {
	var configuration Server
	var out map[string]interface{}
	var s string

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

	for k, v := range out {
		switch x := v.(type) {
		case string:
			if x == (s) {
				a.log.Error(fmt.Sprintf("Members:%s is not defined in the environment", k))
				os.Exit(1)
			}
		case bool:
		default:
			panic("invalid Members configuration type")
		}
	}

	return &out
}
