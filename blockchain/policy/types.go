package policy

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
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

type PolicyPayloadI interface{}

func NewPolicyPayload(cdc codec.BinaryCodec, any *cdctypes.Any) PolicyPayload {
	return PolicyPayload{
		cdc: cdc,
		any: any,
	}
}

func EmptyPolicyPayload() PolicyPayload {
	return NewPolicyPayload(nil, nil)
}

func UnpackPayload[P PolicyPayloadI](p PolicyPayload) (P, error) {
	var (
		empty   P
		payload PolicyPayloadI
	)

	err := p.cdc.UnpackAny(p.any, &payload)
	if err != nil {
		return empty, err
	}

	if payload == nil {
		return empty, nil
	}
	return payload.(P), nil
}

type Policy interface {
	// Verify tries to verify the current policy. The returned error is nil if
	// the policy is valid.
	Verify(approvers ApproverSet, payload PolicyPayload) error
}
