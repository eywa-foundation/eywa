package keeper

import (
	"context"

	"eywa/x/eywa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRelayServer(goCtx context.Context, msg *types.MsgCreateRelayServer) (*types.MsgCreateRelayServerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var relayServer = types.RelayServer{
		RelayAccount: msg.Creator,
		Location:     msg.Location,
		Nickname:     msg.Nickname,
	}

	_ = k.CreateRelayServerService(ctx, relayServer)

	return &types.MsgCreateRelayServerResponse{}, nil
}
