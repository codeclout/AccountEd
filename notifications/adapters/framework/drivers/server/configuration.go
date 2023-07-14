package server

import (
	"fmt"
	"os"
	"reflect"

	"golang.org/x/exp/slog"
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
	log *slog.Logger
}

func NewAdapter(log *slog.Logger) *Adapter {
	return &Adapter{
		log: log,
	}
}

// LoadNotificationsConfig checks and loads environment variables for the adapter configuration such as EMAIL_PROCESSOR_API_KEY, EMAIL_PROCESSOR_DOMAIN,
// EMAIL_VERIFIER_API_PATH, PORT, and PERFORMANCE_SLA. It returns a pointer to a map containing these environment variables as key-value pairs.
// If any string environment variable is not set, the method will log an error and forcefully exit the program. If the configuration value is of an unexpected type,
// the method will panic with a "invalid AWS configuration type" message.
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
				a.log.Error(fmt.Sprintf("Notification:%s is not defined in the environment", k))
				os.Exit(1)
			}
		default:
			panic("invalid Notification configuration type")
		}
	}

	return &out

}
