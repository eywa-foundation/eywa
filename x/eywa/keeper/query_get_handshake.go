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

func (k Keeper) GetHandshake(goCtx context.Context, req *types.QueryGetHandshakeRequest) (*types.QueryGetHandshakeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var handshakes []types.Handshake
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	handshakeStore := prefix.NewStore(store, types.KeyPrefix(types.HandshakeKey))

	pageRes, err := query.Paginate(handshakeStore, req.Pagination, func(key []byte, value []byte) error {
		var handshake types.Handshake
		if err := k.cdc.Unmarshal(value, &handshake); err != nil {
			return err
		}

		//if req.Receiver == handshake.Receiver
		if req.Receiver == handshake.Receiver {
			handshakes = append(handshakes, handshake)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetHandshakeResponse{Handshake: handshakes, Pagination: pageRes}, nil
}
