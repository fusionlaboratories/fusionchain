package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/qredo/fusionchain/x/treasury/types"
)

func (k Keeper) appendKey(ctx sdk.Context, key *types.Key, keyRequest *types.KeyRequest) {
	key.Id = keyRequest.Id

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyKey))
	newValue := k.cdc.MustMarshal(key)
	store.Set(sdk.Uint64ToBigEndian(key.Id), newValue)
}

func (k Keeper) GetKey(ctx sdk.Context, id uint64) (*types.Key, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyKey))
	keyID := strconv.FormatUint(id, 10)
	b := store.Get([]byte(keyID))
	if b == nil {
		return nil, false
	}

	var key types.Key
	k.cdc.MustUnmarshal(b, &key)

	return &key, true
}

func (k Keeper) SetKeyCount(ctx sdk.Context, c uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KeyCountKey)
	bz := sdk.Uint64ToBigEndian(c)
	store.Set(byteKey, bz)
}

func (k Keeper) GetKeyCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KeyCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetKey(ctx sdk.Context, key *types.Key) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyKey))
	newValue := k.cdc.MustMarshal(key)
	keyID := strconv.FormatUint(key.Id, 10)
	store.Set([]byte(keyID), newValue)
}
