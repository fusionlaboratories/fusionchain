// Copyright 2023 Qredo Ltd.
// This file is part of the Fusion library.
//
// The Fusion library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Fusion library. If not, see https://github.com/qredo/fusionchain/blob/main/LICENSE
package bridge

import (
	"bytes"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	serviceName = "qrdo-bridge"
)

var (
	// Config vars

	defaultFusionURL     = "localhost:9090"
	defaultFusionChainID = "qredofusiontestnet_257-1"

	defaultQredochainURL = "localhost:26657"

	defaultHandlerTimeout = 60 * time.Second
	defaultQueryTimeout   = 5 * time.Second

	defaultMaxRetries    int64 = 10
	defaultQueryInterval int64 = 5

	defaultRetryTimeout = 5 * time.Second

	defaultChanSize = 1000

	defaultPageLimit uint64 = 10

	defaultThreads = 2
)

// ServiceConfig represents the main application configuration struct.
type ServiceConfig struct {
	QredochainURL string `yaml:"qredochainurl"`
	LogLevel      string `yaml:"loglevel"`
	LogFormat     string `yaml:"logformat"`
	LogToFile     bool   `yaml:"logtofile"`
	RetrySleep    int64  `yaml:"retrySleep"`
}

var emptyConfig = ServiceConfig{}

var defaultConfig = ServiceConfig{
	LogLevel:      "info",
	LogFormat:     "plain",
	LogToFile:     false,
	QredochainURL: defaultQredochainURL,
}

func isEmpty(c ServiceConfig) bool {
	b, _ := yaml.Marshal(c)
	e, _ := yaml.Marshal(emptyConfig)
	return bytes.Equal(b, e)
}

// sanitizeConfig Partially empty configs will be sanitized with default values.
func sanitizeConfig(config ServiceConfig) (cfg ServiceConfig, defaultUsed bool) {
	if isEmpty(config) {
		defaultUsed = true
		cfg = defaultConfig
		return
	}
	cfg = config

	if config.LogLevel == "" {
		cfg.LogLevel = "info"
	}

	if config.LogFormat == "" {
		cfg.LogFormat = "plain"
	}

	return
}
