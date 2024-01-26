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
	"context"

	"github.com/qredo/assets/libs/logger"
	"github.com/qredo/assets/libs/nodeconnector"
)

// BuildService constructs the main application based on supplied config parameters
func BuildService(config ServiceConfig) (*Bridge, error) {
	cfg, _ := sanitizeConfig(config) // set default values is none supplied

	logger, err := logger.NewLogger(logger.Level(cfg.LogLevel), logger.Format(cfg.LogFormat), cfg.LogToFile, serviceName)
	if err != nil {
		return nil, err
	}

	log = logger.Logger

	tmClient, err := nodeconnector.NewDNSSafe(config.QredochainURL)
	if err != nil {
		return nil, err
	}

	fusion, err := NewClient(context.Background(), ConfigFromEnv())
	if err != nil {
		return nil, err
	}

	if err := tmClient.Start(); err != nil {
		return nil, err
	}

	return NewBridge(nodeconnector.TendermintKeySearcher{Client: tmClient}, NewRetriever(tmClient), fusion), nil
}
