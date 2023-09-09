package cli

import (
	"strconv"

	"eywa/x/eywa/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateHandshake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-handshake [receiver] [room-id] [server-address]",
		Short: "Broadcast message create-handshake",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argRoomId := args[1]
			argServerAddress := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateHandshake(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argRoomId,
				argServerAddress,
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
