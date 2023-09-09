package cli

import (
	"strconv"

	"eywa/x/eywa/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetHandshake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-handshake [receiver]",
		Short: "Query get-handshake",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqReceiver := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetHandshakeRequest{

				Receiver: reqReceiver,
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			res, err := queryClient.GetHandshake(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
