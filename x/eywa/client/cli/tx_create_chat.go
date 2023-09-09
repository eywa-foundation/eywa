package cli

import (
	"strconv"

	"eywa/x/eywa/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateChat() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chat [room-id] [receiver] [message] [time]",
		Short: "Broadcast message create-chat",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRoomId := args[0]
			argReceiver := args[1]
			argMessage := args[2]
			argTime, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateChat(
				clientCtx.GetFromAddress().String(),
				argRoomId,
				argReceiver,
				argMessage,
				argTime,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
