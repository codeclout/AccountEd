package server

import (
	"encoding/json"
	"os"
)

type environment struct {
	EmailProcessorAPIKey string `json:"email_processor_api_key"`
	EmailProcessorDomain string `json:"email_processor_domain"`
	Port                 string `json:"port"`
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
		Port:                 os.Getenv("PORT"),
	}

	env, _ := json.Marshal(envConfig)
	_ = json.Unmarshal(env, &out)

	return &out

}
