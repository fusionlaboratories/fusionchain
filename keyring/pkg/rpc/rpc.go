// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package rpc

import (
	"context"
	"net/http"
)

type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Client interface {
	Do(*http.Request) (*http.Response, error)
}
