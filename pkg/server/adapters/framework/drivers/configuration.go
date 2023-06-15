package drivers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

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

// Load reads the configuration file, decodes it into a TransferBounds struct, and then converts it into a map[string]interface{} for
// easier manipulation. It gets the working directory, constructs the file path, and decodes the file content using hclsimple.DecodeFile. If there are
// any errors during decoding, it logs an error message and panics. After decoding, the TransferBounds struct is marshaled into a JSON and then
// unmarshaled into a map[string]interface{}. The resulting map is returned.
func (a *Adapter) Load() (*map[string]interface{}, error) {
	var bounds protocols.TransferBounds
	var out map[string]interface{}

	workingDirectory, _ := os.Getwd()
	fileLocation := filepath.Join(workingDirectory, path)

	e := hclsimple.DecodeFile(fileLocation, nil, &bounds)
	if e != nil {
		x, ok := e.(hcl.Diagnostics)

		if ok {
			a.log.Error(fmt.Sprintf("Failed to load runtime staticConfig: %s", x[0].Summary))
			return nil, errors.New(x[0].Error())
		} else {
			a.log.Error(fmt.Sprintf("Failed to get runtime staticConfig: %s", e.Error()))
			return nil, errors.New(e.Error())
		}
	}

	val := reflect.ValueOf(&bounds).Elem()

	for i := 0; i < val.NumField(); i++ {
		out[val.Type().Field(i).Name] = val.Field(i).Interface()
	}

	return &out, nil
}
