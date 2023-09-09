package keeper

import (
	"context"

	"eywa/x/eywa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRelayServer(goCtx context.Context, req *types.QueryGetRelayServerRequest) (*types.QueryGetRelayServerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var relayServers []types.RelayServer
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	relayServerStore := prefix.NewStore(store, types.KeyPrefix(types.RelayServerKey))

	pageRes, err := query.Paginate(relayServerStore, req.Pagination, func(key []byte, value []byte) error {
		var relayServer types.RelayServer
		if err := k.cdc.Unmarshal(value, &relayServer); err != nil {
			return err
		}

		relayServers = append(relayServers, relayServer)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetRelayServerResponse{RelayServer: relayServers, Pagination: pageRes}, nil
}
