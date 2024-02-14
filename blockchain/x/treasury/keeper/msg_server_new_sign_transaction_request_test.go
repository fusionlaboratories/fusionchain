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

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/identity"
	idTypes "github.com/qredo/fusionchain/x/identity/types"
	"github.com/qredo/fusionchain/x/treasury"
	"github.com/qredo/fusionchain/x/treasury/keeper"
	"github.com/qredo/fusionchain/x/treasury/types"
)

func Test_msgServer_NewSignTransactionRequest(t *testing.T) {
	t.SkipNow()
	unsignedTx1 := []byte{2, 65, 62, 38, 30, 38, 35, 30, 34, 61, 38, 31, 37, 63, 38, 30, 30, 38, 32, 35, 32, 30, 38, 39, 34, 39, 39, 33, 66, 34, 35, 36, 36, 36, 62, 32, 61, 37, 38, 34, 33, 34, 37, 31, 31, 64, 31, 61, 32, 30, 64, 32, 61, 39, 37, 33, 33, 63, 30, 37, 61, 35, 33, 31, 38, 38, 37, 30, 65, 33, 35, 66, 61, 39, 33, 31, 61, 30, 30, 30, 30, 38, 30, 38, 30, 38, 30, 38, 30}
	// unsignedTx2 := []byte("0xf86c8004a817c80082520894993f45666b2a78434711d1a20d2a9733c07a53188038d7ea4c68000")

	var metadataAny *cdctypes.Any

	metadata := types.MetadataEthereum{
		ChainId: 11155111,
	}
	metadataAny, _ = cdctypes.NewAnyWithValue(&metadata)

	type args struct {
		keyring   *idTypes.Keyring
		workspace *idTypes.Workspace
		key       *types.Key
		msg       *types.MsgNewSignTransactionRequest
	}
	tests := []struct {
		name              string
		args              args
		wantSignTxRequest *types.SignTransactionRequest
		want              *types.MsgNewSignTransactionRequestResponse
		wantErr           bool
	}{
		{
			name: "PASS: valid signTransactionRequest",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				key: &types.Key{
					Id:            1,
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					Type:          types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
					PublicKey:     []byte{3, 49, 216, 1, 30, 80, 59, 231, 205, 4, 237, 36, 101, 79, 9, 119, 30, 114, 19, 70, 242, 44, 91, 95, 238, 20, 245, 197, 193, 86, 22, 120, 35},
				},
				msg: types.NewMsgNewSignTransactionRequest("testCreator", 1, types.WalletType_WALLET_TYPE_ETH, unsignedTx1, 1000, metadataAny),
			},
			wantSignTxRequest: &types.SignTransactionRequest{
				Id:                  1,
				Creator:             "testCreator",
				KeyId:               1,
				WalletType:          types.WalletType_WALLET_TYPE_ETH,
				UnsignedTransaction: unsignedTx1,
				SignRequestId:       1,
			},
			want: &types.MsgNewSignTransactionRequestResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			ctx := keepers.Ctx

			idGenesis := idTypes.GenesisState{
				Keyrings:   []idTypes.Keyring{*tt.args.keyring},
				Workspaces: []idTypes.Workspace{*tt.args.workspace},
			}
			identity.InitGenesis(ctx, *ik, idGenesis)

			tk := keepers.TreasuryKeeper
			tGenesis := types.GenesisState{
				Keys: []types.Key{*tt.args.key},
			}
			treasury.InitGenesis(ctx, *tk, tGenesis)

			goCtx := sdk.WrapSDKContext(ctx)
			msgSer := keeper.NewMsgServerImpl(*tk)

			got, err := msgSer.NewSignTransactionRequest(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewSignTransactionRequest() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {

				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("NewSignTransactionRequest() got = %v, want %v", got, tt.want)
				}

				gotKeyReq, bool := tk.KeyRequestsRepo().Get(ctx, got.Id)
				if !bool {
					t.Fatalf("SignTransactionRequestRepo().Get failed, error = %v", bool)
				}

				if !reflect.DeepEqual(gotKeyReq, tt.wantSignTxRequest) {
					t.Fatalf("NewSignTransactionRequest() got = %v, want %v", gotKeyReq, tt.wantSignTxRequest)
				}
			}
		})
	}
}
