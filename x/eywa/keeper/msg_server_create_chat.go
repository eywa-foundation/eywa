package keeper

import (
	"context"

	"eywa/x/eywa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateChat(goCtx context.Context, msg *types.MsgCreateChat) (*types.MsgCreateChatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var chat = types.Chat{
		RoomId:   msg.RoomId,
		Sender:   msg.Creator,
		Receiver: msg.Receiver,
		Message:  msg.Message,
		Time:     msg.Time,
	}

	_ = k.CreateChatService(ctx, chat)

	return &types.MsgCreateChatResponse{}, nil
}
