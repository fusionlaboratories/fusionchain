package keeper_test

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"

	// idtypes "github.com/qredo/fusionchain/x/identity/types"

	"github.com/qredo/fusionchain/x/treasury/keeper"
	"github.com/qredo/fusionchain/x/treasury/types"
)

func Test_msgServer_NewSignatureRequest(t *testing.T) {

	type args struct {
		// msgWs *idtypes.MsgNewWorkspace
		// msgKr *idtypes.MsgNewKeyring
		key *types.Key
	}

	tests := []struct {
		name    string
		args    args
		want    *types.MsgNewSignatureRequestResponse
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				key: &types.Key{
					Id:            1,
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					Type:          types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
					PublicKey:     []byte("AzHYAR5QO+fNBO0kZU8Jdx5yE0byLFtf7hT1xcFWFngj"),
				},
				// msgWs:  idtypes.NewMsgNewWorkspace("testCreator", 0, 0),
				// msgKr:  idtypes.NewMsgNewKeyring("testCreator", "test", 0, 0, 0),
				// msgKey: types.NewMsgNewKeyRequest("testCreator", "qredoworkspace14a2hpadpsy9h5m6us54", "qredokeyring1ph63us46lyw56vrzgaq", types.KeyType_KEY_TYPE_ECDSA_SECP256K1, 1000),
				// msgSig: types.NewMsgNewSignatureRequest("testCreator", 1, []byte("778f572f33acfab831365d52e563a0ddd2829ddd7060bec69719b7e41f6ef91c"), 1000),
			},
			want:    &types.MsgNewSignatureRequestResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ik, ctx := keepertest.IdentityKeeper(t)
			// goCtx := sdk.WrapSDKContext(ctx)
			// msgSerId := idkeeper.NewMsgServerImpl(*ik)
			// wsRes, err := msgSerId.NewWorkspace(goCtx, tt.args.msgWs)
			// if err != nil {
			// 	t.Fatalf("NewWorkspace() error = %v", err)
			// }

			// krRes, err := msgSerId.NewKeyring(goCtx, tt.args.msgKr)
			// if err != nil {
			// 	t.Fatalf("NewWorkspace() error = %v", err)
			// }

			tk, ctx := keepertest.TreasuryKeeper(t)
			goCtx := sdk.WrapSDKContext(ctx)
			msgSerTk := keeper.NewMsgServerImpl(*tk)

			// keyRes, err := msgSerTk.NewKeyRequest(goCtx, &types.MsgNewKeyRequest{"testCreator", wsRes.GetAddress(), krRes.GetAddress(), 1, 1000})
			// if err != nil {
			// 	t.Fatalf("NewKeyRequest() error = %v", err)
			// }
			// fmt.Println(keyRes)

			got, err := msgSerTk.NewSignatureRequest(goCtx, &types.MsgNewSignatureRequest{
				Creator:        "testCreator",
				KeyId:          tt.args.key.Id,
				DataForSigning: []byte("778f572f33acfab831365d52e563a0ddd2829ddd7060bec69719b7e41f6ef91c"),
				Btl:            1000,
			})
			if err != nil {
				t.Fatalf("NewSignatureRequest() error = %v", err)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewSignatureRequest() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
