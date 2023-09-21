package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/qredo/fusionchain/repo"
	"github.com/qredo/fusionchain/x/identity/types"
	policy "github.com/qredo/fusionchain/x/policy/keeper"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeKey     storetypes.StoreKey
		memKey       storetypes.StoreKey
		paramstore   paramtypes.Subspace
		policyKeeper *policy.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	policyKeeper *policy.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		policyKeeper: policyKeeper,
	}
}

func (k Keeper) KeyringsRepo() repo.ObjectRepo[*types.Keyring] {
	return repo.ObjectRepo[*types.Keyring]{
		Constructor: func() *types.Keyring { return &types.Keyring{} },
		StoreKey:    k.storeKey,
		Cdc:         k.cdc,
		CountKey:    types.KeyPrefix(types.KeyringCountKey),
		ObjKey:      types.KeyPrefix(types.KeyringKey),
	}
}

func (Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
