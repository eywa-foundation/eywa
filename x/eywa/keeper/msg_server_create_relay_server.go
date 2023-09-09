package keeper

import (
	"context"

	"eywa/x/eywa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRelayServer(goCtx context.Context, msg *types.MsgCreateRelayServer) (*types.MsgCreateRelayServerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateRelayServerResponse{}, nil
}
