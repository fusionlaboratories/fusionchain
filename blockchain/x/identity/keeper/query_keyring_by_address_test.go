// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper_test

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/identity"
	"github.com/qredo/fusionchain/x/identity/types"
)

func TestKeeper_KeyringByAddress(t *testing.T) {

	var defaultKr = types.Keyring{
		Address:     "qredokeyring1ph63us46lyw56vrzgaq",
		Creator:     "testCreator",
		Description: "testDescription",
		Admins:      []string{"testCreator"},
		Fees:        &types.KeyringFees{KeyReq: 0, SigReq: 0},
		IsActive:    true,
	}

	type args struct {
		keyring *types.Keyring
		req     *types.QueryKeyringByAddressRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QueryKeyringByAddressResponse
		wantErr bool
	}{
		{
			name: "PASS: get a keyring by address",
			args: args{
				keyring: &defaultKr,
				req: &types.QueryKeyringByAddressRequest{
					Address: "qredokeyring1ph63us46lyw56vrzgaq",
				},
			},
			want: &types.QueryKeyringByAddressResponse{Keyring: &types.Keyring{
				Address:     "qredokeyring1ph63us46lyw56vrzgaq",
				Creator:     "testCreator",
				Description: "testDescription",
				Admins:      []string{"testCreator"},
				Parties:     nil,
				Fees:        &types.KeyringFees{KeyReq: 0, SigReq: 0},
				IsActive:    true,
			}},
		},
		{
			name: "FAIL: keyring by address not found",
			args: args{
				keyring: &defaultKr,
				req: &types.QueryKeyringByAddressRequest{
					Address: "qredokeyring10kjg2u5s22lezv8dahk",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "FAIL: invalid request",
			args: args{
				keyring: &defaultKr,
				req:     nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			ctx := keepers.Ctx
			goCtx := sdk.WrapSDKContext(ctx)

			genesis := types.GenesisState{
				Keyrings: []types.Keyring{*tt.args.keyring},
			}
			identity.InitGenesis(ctx, *ik, genesis)

			got, err := ik.KeyringByAddress(goCtx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyringByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyringByAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
