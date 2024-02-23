// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package api

// Module represents a simple interface for sub-processes within
// a service.
type Module interface {
	Start() error
	Stop() error
	Healthcheck() *HealthResponse
}
