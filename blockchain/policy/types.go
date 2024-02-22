// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package policy

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/gogoproto/proto"
)

type ApproverSet map[string]bool

func BuildApproverSet(approvers []string) ApproverSet {
	approverSet := make(ApproverSet, len(approvers))
	for _, a := range approvers {
		approverSet[a] = true
	}
	return approverSet
}

type PolicyPayload struct {
	cdc codec.BinaryCodec
	any *cdctypes.Any
}

type PolicyPayloadI any

func NewPolicyPayload(cdc codec.BinaryCodec, any *cdctypes.Any) PolicyPayload {
	return PolicyPayload{
		cdc: cdc,
		any: any,
	}
}

func EmptyPolicyPayload() PolicyPayload {
	return NewPolicyPayload(nil, nil)
}

func UnpackPayload[P PolicyPayloadI](p PolicyPayload) (*P, error) {
	var payload P

	if p.any != nil && p.cdc == nil {
		return nil, fmt.Errorf("codec is nil")
	}

	if p.any == nil || p.cdc == nil {
		return nil, nil
	}

	err := p.cdc.UnpackAny(p.any, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

type Policy interface {
	// Validate checks that the policy is valid (well formed).
	// The returned error is nil if the policy is valid.
	Validate() error

	// AddressToParticipant returns the participant shorthand for the given
	// address.
	//
	// The Verify() method will receive the list of approvers as shorthands.
	AddressToParticipant(addr string) (string, error)

	// Verify tries to verify the current policy. The returned error is nil if
	// the policy is valid.
	Verify(approvers ApproverSet, payload PolicyPayload, policyData map[string][]byte) error
}

type PolicyMetadata interface {
	// Metadata returns the metadata associated with the policy. This is used
	// to return additional information about the policy in query responses.
	Metadata() (proto.Message, error)
}
