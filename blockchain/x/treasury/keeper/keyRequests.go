package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/qredo/fusionchain/x/treasury/types"
)

func (k Keeper) GetKeyRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KeyRequestCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetKeyRequestCount(ctx sdk.Context, c uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KeyRequestCountKey)
	bz := sdk.Uint64ToBigEndian(c)
	store.Set(byteKey, bz)
}

func (k Keeper) GetKeyRequest(ctx sdk.Context, id uint64) *types.KeyRequest {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyRequestKey))
	keyReqId := strconv.FormatUint(id, 10)
	bz := store.Get([]byte(keyReqId))
	if bz == nil {
		return nil
	}
	var keyRequest types.KeyRequest
	k.cdc.MustUnmarshal(bz, &keyRequest)
	return &keyRequest
}

func (k Keeper) SetKeyRequest(ctx sdk.Context, keyRequest *types.KeyRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyRequestKey))
	newValue := k.cdc.MustMarshal(keyRequest)
	keyReqId := strconv.FormatUint(keyRequest.Id, 10)
	store.Set([]byte(keyReqId), newValue)
}
