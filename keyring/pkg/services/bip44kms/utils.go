// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package kms

import "time"

const (
	keyIDLength = 64
	pkPrefix    = "pk"

	rateLimitPerSecond = 5
)

func requeueKeyItemWithTimeout(c chan *keyRequestQueueItem, item *keyRequestQueueItem, timeout time.Duration) {
	item.retries++
	go func() {
		time.Sleep(timeout)
		c <- item
	}()
}

func requeueSigItemWithTimeout(c chan *signatureRequestQueueItem, item *signatureRequestQueueItem, timeout time.Duration) {
	item.retries++
	go func() {
		time.Sleep(timeout)
		c <- item
	}()
}

func makeThreads(n int) chan struct{} {
	t := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		t <- struct{}{}
	}
	return t
}
