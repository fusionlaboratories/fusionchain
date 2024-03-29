// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package policy

import (
	"fmt"

	"gitlab.qredo.com/edmund/blackbird/verifier/golang/protobuf"
	bbird "gitlab.qredo.com/edmund/blackbird/verifier/golang/simple"
	"google.golang.org/protobuf/proto"
)

// AnyInGroupPolicy is a simple policy where any member of a group can verify it.
type AnyInGroupPolicy struct {
	group  []string
	policy *protobuf.Policy
}

var _ Policy = &AnyInGroupPolicy{}

func (*AnyInGroupPolicy) Validate() error { return nil }

func (p *AnyInGroupPolicy) AddressToParticipant(addr string) (string, error) {
	for _, s := range p.group {
		if s == addr {
			return addr, nil
		}
	}
	return "", fmt.Errorf("address not a participant of this policy")
}

func (p *AnyInGroupPolicy) Verify(approvers ApproverSet, _ PolicyPayload, _ map[string][]byte) error {
	policyBz, err := proto.Marshal(p.policy)
	if err != nil {
		return err
	}

	return bbird.Verify(policyBz, nil, nil, nil, approvers)
}

func NewAnyInGroupPolicy(group []string) *AnyInGroupPolicy {
	policy := &protobuf.Policy{
		Tag:           protobuf.PolicyTag_POLICY_ANY,
		Threshold:     1,
		AddressPrefix: "",
		Subpolicies:   make([]*protobuf.Policy, len(group)),
	}
	for i, s := range group {
		policy.Subpolicies[i] = &protobuf.Policy{
			Tag:           protobuf.PolicyTag_POLICY_SIGNATURE,
			AddressPrefix: "",
			Address: &protobuf.Policy_CookedAddress{
				CookedAddress: s,
			},
		}
	}

	return &AnyInGroupPolicy{
		group:  group,
		policy: policy,
	}
}
