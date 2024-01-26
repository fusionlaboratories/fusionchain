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
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	CliName        string
	ChainID        string
	KeyringBackend string
	Node           string
	SendDenom      string
	AccountName    string
	Mnemonic       string
	HDPath         string
	GasPrices      string
	OtherFlags     string
}

func ConfigFromEnv() Config {
	return Config{
		CliName:        envOrDefault("CLI_NAME", "fusiond"),
		ChainID:        envOrDefault("CHAIN_ID", "qredofusiontestnet_257-1"),
		KeyringBackend: envOrDefault("KEYRING_BACKEND", "test"),
		Node:           envOrDefault("NODE", "http://localhost:26657"),
		SendDenom:      envOrDefault("DENOM", "10000000000nQRDO"),
		AccountName:    envOrDefault("ACCOUNT_NAME", "shulgin"),
		Mnemonic:       envOrDefault("MNEMONIC", ""),
		HDPath:         envOrDefault("HD_PATH", "m/44'/60'/0'/0/0"),
		GasPrices:      envOrDefault("GAS_PRICES", "1000000000nQRDO"),
		OtherFlags:     envOrDefault("OTHER_FLAGS", ""),
	}
}

type Client struct {
	cfg Config
}

func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	c := &Client{
		cfg: cfg,
	}

	if err := c.setupConfig(ctx); err != nil {
		return nil, err
	}

	if cfg.Mnemonic != "" {
		if err := c.setupNewAccount(ctx); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) baseCmd() string {
	// Build a string like this:
	// fusiond --node tcp://localhost:26657 --fees 20nQRDO
	return strings.Join([]string{
		c.cfg.CliName,
		"--node",
		c.cfg.Node,
		"--gas-prices",
		c.cfg.GasPrices,
		"--from",
		c.cfg.AccountName,
	}, " ")
}

func (c *Client) setupNewAccount(ctx context.Context) error {
	// echo $mnemonic | $baseCmd keys add $SK1 --recover
	cmd := strings.Join([]string{
		"echo",
		c.cfg.Mnemonic,
		"|",
		c.baseCmd(),
		"keys",
		"add",
		c.cfg.AccountName,
		"--recover",
	}, " ")
	return e(ctx, cmd)
}

func (c *Client) setupConfig(ctx context.Context) error {
	// fusiond config keyring-backend $KEYRING
	cmd := strings.Join([]string{
		c.baseCmd(),
		"config",
		"keyring-backend",
		c.cfg.KeyringBackend,
	}, " ")
	if err := e(ctx, cmd); err != nil {
		return err
	}

	// fusiond config chain-id $CHAINID
	cmd = strings.Join([]string{
		c.baseCmd(),
		"config",
		"chain-id",
		c.cfg.ChainID,
	}, " ")
	if err := e(ctx, cmd); err != nil {
		return err
	}

	return nil
}

func (c *Client) Send(ctx context.Context, dest string, amount *big.Int) error {
	// $baseCmd tx bank send bridge qredo1f6zkpwezlw58mssh0qat8d0dvwu3qpw63c64lm 100000000nQRDO --yes
	cmd := strings.Join([]string{
		c.baseCmd(),
		"tx",
		"bank",
		"send",
		c.cfg.AccountName,
		dest,
		fmt.Sprintf("%snQRDO", amount.String()),
		"--yes",
	}, " ")
	return e(ctx, cmd)
}

func e(ctx context.Context, cmd string) error {
	cccc := exec.CommandContext(ctx, "sh", "-c", cmd)
	output, err := cccc.Output()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		log.Println(string(output), string(exitErr.Stderr))
	}
	return err
}

func envOrDefault(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
