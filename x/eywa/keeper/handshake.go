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
	store.Set(GetHandshakeReceiverBytes(handshake.Receiver), appendedValue)
	return handshake.HashId
}

func (k Keeper) GetHandshakeHashId(ctx sdk.Context, handshake types.Handshake) string {
	return handshake.Sender + handshake.Receiver
}

func GetHandshakeReceiverBytes(receiver string) []byte {
	return []byte(receiver)
}
