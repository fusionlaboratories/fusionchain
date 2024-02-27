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
	count := store.Get(byteKey)
	if count == nil {
		return 0
	}
	return sdk.BigEndianToUint64(count)
}

func (k Keeper) SetKeyRequestCount(ctx sdk.Context, c uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KeyRequestCountKey)
	count := sdk.Uint64ToBigEndian(c)
	store.Set(byteKey, count)
}

func (k Keeper) GetKeyRequest(ctx sdk.Context, id uint64) *types.KeyRequest {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyRequestKey))
	keyReqID := strconv.FormatUint(id, 10)
	count := store.Get([]byte(keyReqID))
	if count == nil {
		return nil
	}
	var keyRequest types.KeyRequest
	k.cdc.MustUnmarshal(count, &keyRequest)
	return &keyRequest
}

func (k Keeper) SetKeyRequest(ctx sdk.Context, keyRequest *types.KeyRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyRequestKey))
	newValue := k.cdc.MustMarshal(keyRequest)
	keyReqID := strconv.FormatUint(keyRequest.Id, 10)
	store.Set([]byte(keyReqID), newValue)
}
