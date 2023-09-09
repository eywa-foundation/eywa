package keeper

import (
	"eywa/x/eywa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateRelayServerService(ctx sdk.Context, relayServer types.RelayServer) string {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RelayServerKey))
	appendedValue := k.cdc.MustMarshal(&relayServer)
	store.Set(GetRelayServerRelayAccountBytes(relayServer.RelayAccount), appendedValue)
	return relayServer.RelayAccount
}

func GetRelayServerRelayAccountBytes(relayAccount string) []byte {
	return []byte(relayAccount)
}
