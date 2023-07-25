package server

import (
	"fmt"
	"os"
	"reflect"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type environment struct {
	EmailProcessorAPIKey string `json:"email_processor_api_key"`
	EmailProcessorDomain string `json:"email_processor_domain"`
	EmailVerifierApiPath string `json:"email_verifier_api_path"`
	Port                 string `json:"port"`
	Region               string `json:"awsRegion"`
	SLARoutePerformance  string `json:"sla_route_performance"`
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
		Port:                 os.Getenv("PORT"),
		Region:               os.Getenv("AWS_REGION"),
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
