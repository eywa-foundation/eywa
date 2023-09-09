package keeper

import (
	"context"

	"eywa/x/eywa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateChat(goCtx context.Context, msg *types.MsgCreateChat) (*types.MsgCreateChatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateChatResponse{}, nil
}
