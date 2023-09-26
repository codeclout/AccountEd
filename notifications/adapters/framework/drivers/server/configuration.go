package server

import (
	"fmt"
	"os"
	"reflect"

	"github.com/codeclout/AccountEd/pkg/monitoring"
)

type environment struct {
	EmailProcessorAPIKey string
	EmailProcessorDomain string
	EmailVerifierApiPath string
	Port                 string
	Region               string
	RuntimeEnvironment   string
	SLARoutePerformance  string
}

type Adapter struct {
	monitor monitoring.Adapter
}

func NewAdapter(monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		monitor: monitor,
	}
}

func (a *Adapter) LoadNotificationsConfig() *map[string]interface{} {
	var out = make(map[string]interface{})
	var s string

	envConfig := environment{
		EmailProcessorAPIKey: os.Getenv("EMAIL_PROCESSOR_API_KEY"),
		EmailProcessorDomain: os.Getenv("EMAIL_PROCESSOR_DOMAIN"),
		EmailVerifierApiPath: os.Getenv("EMAIL_VERIFIER_API_PATH"),
		Port:                 os.Getenv("NOTIFICATIONS_PORT"),
		Region:               os.Getenv("AWS_REGION"),
		RuntimeEnvironment:   os.Getenv("ENVIRONMENT"),
		SLARoutePerformance:  os.Getenv("PERFORMANCE_SLA"),
	}

	val := reflect.ValueOf(&envConfig).Elem()

	for i := 0; i < val.NumField(); i++ {
		out[val.Type().Field(i).Name] = val.Field(i).Interface()
	}

	for k, v := range out {
		switch x := v.(type) {
		case string:
			if x == (s) {
				a.monitor.LogGenericError(fmt.Sprintf("Notification:%s is not defined in the environment", k))
				os.Exit(1)
			}
		default:
			panic("invalid Notification configuration type")
		}
	}

	return &out

}
