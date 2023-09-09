package keeper

import (
	"context"

	"eywa/x/eywa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateHandshake(goCtx context.Context, msg *types.MsgCreateHandshake) (*types.MsgCreateHandshakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateHandshakeResponse{}, nil
}
