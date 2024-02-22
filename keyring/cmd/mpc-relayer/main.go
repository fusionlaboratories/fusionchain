// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/qredo/fusionchain/keyring/pkg/common"
	service "github.com/qredo/fusionchain/keyring/pkg/services/mpcrelayer"
)

const envPrefix = "MPCRELAYER"

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
	var config service.ServiceConfig

	if err := common.ParseYAMLConfig(configFilePath, &config, envPrefix); err != nil {
		log.Fatal(fmt.Errorf("parse config error: %v", err))
	}
	mpcRelayer, err := service.BuildService(config)
	if err != nil {
		log.Fatal(fmt.Errorf("build service error: %v", err))
	}

	if err := mpcRelayer.Start(); err != nil {
		log.Fatal(fmt.Errorf("start service error: %v", err))
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan
	if err := mpcRelayer.Stop(sig); err != nil {
		log.Fatal(err)
	}
}
