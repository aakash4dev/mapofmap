package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"mapofmap/x/mapofmap/types"
)

func CmdCreateExecutionlayers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-executionlayers [index] [name] [info] [chainid]",
		Short: "Create a new executionlayers",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argName := args[1]
			argInfo := args[2]
			argChainid := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateExecutionlayers(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argName,
				argInfo,
				argChainid,
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

func CmdUpdateExecutionlayers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-executionlayers [index] [name] [info] [chainid]",
		Short: "Update a executionlayers",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argName := args[1]
			argInfo := args[2]
			argChainid := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateExecutionlayers(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argName,
				argInfo,
				argChainid,
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

func CmdDeleteExecutionlayers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-executionlayers [index]",
		Short: "Delete a executionlayers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteExecutionlayers(
				clientCtx.GetFromAddress().String(),
				indexIndex,
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
