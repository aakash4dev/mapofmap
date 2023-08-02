package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mapofmap/x/mapofmap/types"
)

var _ = strconv.Itoa(0)

func CmdChainid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chainid [batchnumber]",
		Short: "Query chainid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqBatchnumber := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryChainidRequest{

				Batchnumber: reqBatchnumber,
			}

			res, err := queryClient.Chainid(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
