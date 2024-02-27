package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/qredo/fusionchain/x/treasury/types"
)

func (k Keeper) GetSignRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignRequestCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetSignRequestCount(ctx sdk.Context, c uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignRequestCountKey)
	bz := sdk.Uint64ToBigEndian(c)
	store.Set(byteKey, bz)
}

func (k Keeper) GetSignRequest(ctx sdk.Context, id uint64) *types.SignRequest {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignRequestKey))
	signReqID := strconv.FormatUint(id, 10)
	bz := store.Get([]byte(signReqID))
	if bz == nil {
		return nil
	}
	var signRequest types.SignRequest
	k.cdc.MustUnmarshal(bz, &signRequest)
	return &signRequest
}

func (k Keeper) SetSignRequest(ctx sdk.Context, signRequest *types.SignRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignRequestKey))
	newValue := k.cdc.MustMarshal(signRequest)
	signReqID := strconv.FormatUint(signRequest.Id, 10)
	store.Set([]byte(signReqID), newValue)
}
