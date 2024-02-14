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
package keeper_test

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/treasury"
	"github.com/qredo/fusionchain/x/treasury/types"
)

func TestKeeper_Keys(t *testing.T) {

	type args struct {
		keys []types.Key
		req  *types.QueryKeysRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QueryKeysResponse
		wantErr bool
	}{
		{
			name: "PASS: ecdsa - return key requests for a workspace and a keyring",
			args: args{
				keys: []types.Key{defaultECDSAKey},
				req:  &types.QueryKeysRequest{},
			},
			want: &types.QueryKeysResponse{
				Pagination: &query.PageResponse{Total: 1},
				Keys: []*types.KeyResponse{{Key: &defaultECDSAKey, Wallets: []*types.WalletKeyResponse{
					{Address: "qredo18wvrcug8acpwn3py30cyjmmtfgqxh7n04d4ygs", Type: types.WalletType_WALLET_TYPE_FUSION},
					{Address: "0x185Ac2b596EB99f2a31Ad637320746354e7dC3f9", Type: types.WalletType_WALLET_TYPE_ETH},
					{Address: "celestia18wvrcug8acpwn3py30cyjmmtfgqxh7n0mszfg0", Type: types.WalletType_WALLET_TYPE_CELESTIA},
				}}},
			},
		},
		{
			name: "PASS: ecdsa - return keys for fusion addresses",
			args: args{
				keys: []types.Key{defaultECDSAKey},
				req: &types.QueryKeysRequest{
					Type: types.WalletType_WALLET_TYPE_FUSION,
				},
			},
			want: &types.QueryKeysResponse{
				Pagination: &query.PageResponse{Total: 1},
				Keys: []*types.KeyResponse{{Key: &defaultECDSAKey, Wallets: []*types.WalletKeyResponse{
					{Address: "qredo18wvrcug8acpwn3py30cyjmmtfgqxh7n04d4ygs", Type: types.WalletType_WALLET_TYPE_FUSION},
				}}},
			},
		},
		{
			name: "PASS: ecdsa - return keys for eth addresses",
			args: args{
				keys: []types.Key{defaultECDSAKey},
				req: &types.QueryKeysRequest{
					Type: types.WalletType_WALLET_TYPE_ETH,
				},
			},
			want: &types.QueryKeysResponse{
				Pagination: &query.PageResponse{Total: 1},
				Keys: []*types.KeyResponse{{Key: &defaultECDSAKey, Wallets: []*types.WalletKeyResponse{
					{Address: "0x185Ac2b596EB99f2a31Ad637320746354e7dC3f9", Type: types.WalletType_WALLET_TYPE_ETH},
				}}},
			},
		},
		{
			name: "PASS: ecdsa - return keys for celestia addresses",
			args: args{
				keys: []types.Key{defaultECDSAKey},
				req: &types.QueryKeysRequest{
					Type: types.WalletType_WALLET_TYPE_CELESTIA,
				},
			},
			want: &types.QueryKeysResponse{
				Pagination: &query.PageResponse{Total: 1},
				Keys: []*types.KeyResponse{{Key: &defaultECDSAKey, Wallets: []*types.WalletKeyResponse{
					{Address: "celestia18wvrcug8acpwn3py30cyjmmtfgqxh7n0mszfg0", Type: types.WalletType_WALLET_TYPE_CELESTIA},
				}}},
			},
		},
		{
			name: "PASS: eddsa - return keys for sui addresses",
			args: args{
				keys: []types.Key{defaultEdDSAKey},
				req: &types.QueryKeysRequest{
					Type: types.WalletType_WALLET_TYPE_SUI,
				},
			},
			want: &types.QueryKeysResponse{
				Pagination: &query.PageResponse{Total: 1},
				Keys: []*types.KeyResponse{{Key: &defaultEdDSAKey, Wallets: []*types.WalletKeyResponse{
					{Address: "0x9061c905900bb96457ac7d73832697686bc64e2a102c38f6fe6bff8ba8002bf0", Type: types.WalletType_WALLET_TYPE_SUI},
				}}},
			},
		},
		{
			name: "PASS: eddsa - return keys for all addresses",
			args: args{
				keys: []types.Key{defaultEdDSAKey},
				req: &types.QueryKeysRequest{
					Type: types.WalletType_WALLET_TYPE_UNSPECIFIED,
				},
			},
			want: &types.QueryKeysResponse{
				Pagination: &query.PageResponse{Total: 1},
				Keys: []*types.KeyResponse{{Key: &defaultEdDSAKey, Wallets: []*types.WalletKeyResponse{
					{Address: "0x9061c905900bb96457ac7d73832697686bc64e2a102c38f6fe6bff8ba8002bf0", Type: types.WalletType_WALLET_TYPE_SUI},
				}}},
			},
		},
		{
			name: "PASS: ecdsa - return keys for all addresses from a specific workspace",
			args: args{
				keys: []types.Key{defaultECDSAKey},
				req: &types.QueryKeysRequest{
					WorkspaceAddr: defaultECDSAKey.WorkspaceAddr,
				},
			},
			want: &types.QueryKeysResponse{
				Pagination: &query.PageResponse{Total: 1},
				Keys: []*types.KeyResponse{{Key: &defaultECDSAKey, Wallets: []*types.WalletKeyResponse{
					{Address: "qredo18wvrcug8acpwn3py30cyjmmtfgqxh7n04d4ygs", Type: types.WalletType_WALLET_TYPE_FUSION},
					{Address: "0x185Ac2b596EB99f2a31Ad637320746354e7dC3f9", Type: types.WalletType_WALLET_TYPE_ETH},
					{Address: "celestia18wvrcug8acpwn3py30cyjmmtfgqxh7n0mszfg0", Type: types.WalletType_WALLET_TYPE_CELESTIA},
				}}},
			},
		},
		{
			name: "FAIL: request is empty",
			args: args{
				keys: []types.Key{defaultECDSAKey},
				req:  nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "FAIL: ecdsa - keyId not found",
			args: args{
				keys: []types.Key{defaultECDSAKey},
				req: &types.QueryKeysRequest{
					KeyId:         2,
					WorkspaceAddr: defaultECDSAKey.WorkspaceAddr,
				},
			},
			wantErr: true,
		},
		{
			name: "FAIL: ecdsa - return keys for all addresses from a specific workspace",
			args: args{
				keys: []types.Key{defaultECDSAKey},
				req: &types.QueryKeysRequest{
					WorkspaceAddr: "anotherWorkspace",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			tk := keepers.TreasuryKeeper
			ctx := keepers.Ctx
			goCtx := sdk.WrapSDKContext(ctx)

			genesis := types.GenesisState{
				Keys: tt.args.keys,
			}
			treasury.InitGenesis(ctx, *tk, genesis)

			got, err := tk.Keys(goCtx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() got = %v, want %v", got, tt.want)
			}
		})
	}
}
