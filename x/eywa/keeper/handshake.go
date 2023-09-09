package keeper

import (
	"eywa/x/eywa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateHandshakeService(ctx sdk.Context, handshake types.Handshake) string {
	handshake.HashId = k.GetHandshakeHashId(ctx, handshake)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HandshakeKey))
	appendedValue := k.cdc.MustMarshal(&handshake)
	store.Set(GetHandshakeHashIdByte(handshake.HashId), appendedValue)
	return handshake.HashId
}

func (k Keeper) GetHandshakeHashId(ctx sdk.Context, handshake types.Handshake) string {
	return handshake.Sender + handshake.Receiver
}

func GetHandshakeHashIdByte(hashId string) []byte {
	return []byte(hashId)
}
