package keeper

import (
	"context"

	"eywa/x/eywa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRelayServer(goCtx context.Context, req *types.QueryGetRelayServerRequest) (*types.QueryGetRelayServerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetRelayServerResponse{}, nil
}
