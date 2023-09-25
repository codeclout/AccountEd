package drivers

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"

	"github.com/codeclout/AccountEd/pkg/monitoring"
)

//go:embed config.hcl
var PackageConfiguration []byte

type MetadataAndSettings struct {
	Settings Settings `hcl:"Settings,block"`
}

type Settings struct {
	GRPCClientConnectionTimeout float64 `hcl:"grpc_service_client_connection_timeout"`
}

type Adapter struct {
	monitor monitoring.Adapter
}

func NewAdapter(monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		monitor: monitor,
	}
}

func (a *Adapter) LoadServerConfiguration() *map[string]interface{} {
	var metadataAndSettings MetadataAndSettings
	var out = make(map[string]interface{})
	var s string

	e := hclsimple.Decode("config.hcl", PackageConfiguration, nil, &metadataAndSettings)
	if e != nil {
		var x *hcl.Diagnostics

		if errors.As(e, &x) {
			for _, x := range *x {
				if x.Severity == hcl.DiagError {
					a.monitor.LogGenericError(fmt.Sprintf("Failed to load server runtime staticConfig: %s", x))
				}
			}
			return nil
		}
	}

	settings := reflect.ValueOf(&metadataAndSettings.Settings).Elem()

	for i := 0; i < settings.NumField(); i++ {
		out[settings.Type().Field(i).Name] = settings.Field(i).Interface()
	}

	for k, v := range out {
		switch x := v.(type) {
		case string:
			if x == (s) {
				a.monitor.LogGenericError(fmt.Sprintf("Server Package:%s is not defined in the environment", k))
				os.Exit(1)
			}
		case bool:
			continue
		case float64:
			continue
		default:
			panic("invalid Server Package configuration type")
		}
	}

	return &out
}
