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
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/vrischmann/envconfig"
	"gopkg.in/yaml.v3"

	"github.com/qredo/fusionchain/bridge"
)

const envPrefix = "BRIDGE"

var (
	configFilePath string
	configFilePtr  = flag.String("config", "config.yml", "path to config file")
)

// go run main.go --config ./config.yml
// go run main.go --config {path_to_config_file}

func init() {
	// Parse flag containing path to config file
	flag.Parse()
	if configFilePtr != nil {
		configFilePath = *configFilePtr
	}
}

func main() {
	var config bridge.ServiceConfig

	if err := ParseYAMLConfig(configFilePath, &config, envPrefix); err != nil {
		log.Fatal(fmt.Errorf("parse config error: %v", err))
	}
	kms, err := bridge.BuildService(config)
	if err != nil {
		log.Fatal(fmt.Errorf("build service error: %v", err))
	}

	if err := kms.Start(); err != nil {
		log.Fatal(fmt.Errorf("start service error: %v", err))
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	if err := kms.Stop(); err != nil {
		log.Fatal(err)
	}
}

// ParseYAMLConfig parse configuration file or environment variables, receiver must be a pointer
func ParseYAMLConfig(configFile string, receiver any, prefix string) error {
	b, err := os.ReadFile(filepath.Clean(configFile))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if b != nil {
		if err := yaml.Unmarshal(b, receiver); err != nil {
			return err
		}
	}
	// environment variables supersede config yaml files
	if err := envconfig.InitWithOptions(receiver, envconfig.Options{Prefix: prefix, AllOptional: true}); err != nil {
		return err
	}
	return nil
}
