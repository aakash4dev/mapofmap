package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"mapofmap/x/mapofmap/types"
)

var _ = strconv.Itoa(0)

func CmdAddbatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addbatch [chainid] [batch]",
		Short: "Broadcast message addbatch",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChainid := args[0]
			argBatch := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddbatch(
				clientCtx.GetFromAddress().String(),
				argChainid,
				argBatch,
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
