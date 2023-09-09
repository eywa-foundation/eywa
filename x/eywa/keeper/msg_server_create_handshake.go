package keeper

import (
	"context"

	"eywa/x/eywa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateHandshake(goCtx context.Context, msg *types.MsgCreateHandshake) (*types.MsgCreateHandshakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var handshake = types.Handshake{
		Sender:        msg.Creator,
		Receiver:      msg.Receiver,
		RoomId:        msg.RoomId,
		ServerAddress: msg.ServerAddress,
	}

	_ = k.CreateHandshakeService(ctx, handshake)

	return &types.MsgCreateHandshakeResponse{}, nil
}
