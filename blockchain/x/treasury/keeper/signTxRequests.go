package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/qredo/fusionchain/x/treasury/types"
)

func (k Keeper) GetSignTxRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignTransactionRequestCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetSignTxRequestCount(ctx sdk.Context, c uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignTransactionRequestCountKey)
	bz := sdk.Uint64ToBigEndian(c)
	store.Set(byteKey, bz)
}

func (k Keeper) GetSignTxRequest(ctx sdk.Context, id uint64) *types.SignTransactionRequest {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignTransactionRequestKey))
	signTxReqId := strconv.FormatUint(id, 10)
	bz := store.Get([]byte(signTxReqId))
	if bz == nil {
		return nil
	}
	var signTxRequest types.SignTransactionRequest
	k.cdc.MustUnmarshal(bz, &signTxRequest)
	return &signTxRequest
}

func (k Keeper) SetSignTxRequest(ctx sdk.Context, signTxRequest *types.SignTransactionRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignTransactionRequestKey))
	newValue := k.cdc.MustMarshal(signTxRequest)
	signTxReqId := strconv.FormatUint(signTxRequest.Id, 10)
	store.Set([]byte(signTxReqId), newValue)
}
