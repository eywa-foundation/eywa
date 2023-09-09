package keeper

import (
	"crypto/sha256"
	"eywa/x/eywa/types"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateChatService(ctx sdk.Context, chat types.Chat) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChatKey))

	chat.ChatIndex = k.GetChatIndex(ctx, chat)
	appendedValue := k.cdc.MustMarshal(&chat)
	store.Set(GetChatIndexByte(chat.ChatIndex), appendedValue)
	return chat.ChatIndex
}

func (k Keeper) GetChatIndex(ctx sdk.Context, chat types.Chat) string {
	chatIndex := chat.RoomId + chat.Message + chat.Sender + chat.Receiver + fmt.Sprintf("%d", chat.Time)
	chatIndexHash := sha256.Sum256([]byte(chatIndex))
	return string(chatIndexHash[:])
}

func GetChatIndexByte(chatIndex string) []byte {
	return []byte(chatIndex)
}
