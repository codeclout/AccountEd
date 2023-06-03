package server

import (
	"encoding/json"
	"os"
)

type environment struct {
	EmailProcessorAPIKey string `json:"email_processor_api_key"`
	EmailProcessorDomain string `json:"email_processor_domain"`
	EmailVerifierApiPath string `json:"email_verifier_api_path"`
	Port                 string `json:"port"`
	SLARoutePerformance  string `json:"sla_route_performance"`
}

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) LoadNotificationsConfig() *map[string]interface{} {
	var out map[string]interface{}

	envConfig := environment{
		EmailProcessorAPIKey: os.Getenv("EMAIL_PROCESSOR_API_KEY"),
		EmailProcessorDomain: os.Getenv("EMAIL_PROCESSOR_DOMAIN"),
		EmailVerifierApiPath: os.Getenv("EMAIL_VERIFIER_API_PATH"),
		Port:                 os.Getenv("PORT"),
		SLARoutePerformance:  os.Getenv("PERFORMANCE_SLA"),
	}

	env, _ := json.Marshal(envConfig)
	_ = json.Unmarshal(env, &out)

	return &out

}
