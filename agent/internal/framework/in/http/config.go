package http

import (
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type ApplicationConfig struct {
	App Application `hcl:"application,block"`
}

type Application struct { //nolint:maligned
	ETag                    bool `hcl:"withEtag"`
	ReadTimeout             time.Duration
	WriteTimeout            time.Duration
	CompressedFileSuffix    string   `hcl:"response_encoding_suffix"`
	DisableStartupMessage   bool     `hcl:"disable_third_party_ascii"`
	AppName                 string   `hcl:"name"`
	StreamRequestBody       bool     `hcl:"stream_response_body"`
	Network                 string   `hcl:"network"`
	EnableTrustedProxyCheck bool     `hcl:"withTrustedProxies"`
	TrustedProxies          []string `hcl:"trustedProxies"`
	EnablePrintRoutes       bool     `hcl:"printRoutes"`
}

type ApplicationOption func(*ApplicationConfig)

func WithEtag() ApplicationOption {
	return func(a *ApplicationConfig) {
		a.App.ETag = true
	}
}

func NewApplicationConfig() (*ApplicationConfig, error) {
	readTimeout, _ := time.ParseDuration(`hcl:"def_read_timeout"`)
	writeTimeout, _ := time.ParseDuration(`hcl:"def_write_timeout"`)

	wd, wde := os.Getwd()
	if wde != nil {
		return nil, wde
	}

	a := &ApplicationConfig{
		App: Application{
			ReadTimeout:  time.Duration(readTimeout.Seconds()),
			WriteTimeout: time.Duration(writeTimeout.Seconds()),
		},
	}

	configFileLocation := filepath.Join(wd, "http.hcl")
	e := hclsimple.DecodeFile(configFileLocation, nil, a)

	if e != nil {
		return nil, e
	}

	return a, nil
}
