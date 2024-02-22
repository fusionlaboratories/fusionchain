// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

import "fmt"

func (a *Action) AddApprover(approver string) error {
	if a.Status != ActionStatus_ACTION_STATUS_PENDING {
		return fmt.Errorf("action already completed")
	}

	for _, a := range a.Approvers {
		if a == approver {
			return fmt.Errorf("approver already added")
		}
	}

	a.Approvers = append(a.Approvers, approver)
	return nil
}

func (a *Action) GetPolicyDataMap() map[string][]byte {
	m := make(map[string][]byte, len(a.PolicyData))
	for _, d := range a.PolicyData {
		m[d.Key] = d.Value
	}
	return m
}
