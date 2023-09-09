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

func (k Keeper) GetChat(goCtx context.Context, req *types.QueryGetChatRequest) (*types.QueryGetChatResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var chats []types.Chat
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	chatStore := prefix.NewStore(store, types.KeyPrefix(types.ChatKey))

	pageRes, err := query.Paginate(chatStore, req.Pagination, func(key []byte, value []byte) error {
		var chat types.Chat
		if err := k.cdc.Unmarshal(value, &chat); err != nil {
			return err
		}

		//if req.Receiver == handshake.Receiver
		if req.RoomId == chat.RoomId {
			chats = append(chats, chat)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}
	return &types.QueryGetChatResponse{Chat: chats, Pagination: pageRes}, nil
}
