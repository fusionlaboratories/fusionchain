// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

// revenue events
const (
	EventTypeRegisterRevenue      = "register_revenue"
	EventTypeCancelRevenue        = "cancel_revenue"
	EventTypeUpdateRevenue        = "update_revenue"
	EventTypeDistributeDevRevenue = "distribute_dev_revenue"

	AttributeKeyContract          = "contract"
	AttributeKeyWithdrawerAddress = "withdrawer_address"
)
