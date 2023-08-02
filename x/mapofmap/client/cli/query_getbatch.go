package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mapofmap/x/mapofmap/types"
)

var _ = strconv.Itoa(0)

func CmdGetbatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getbatch [chainid] [batchnumber]",
		Short: "Query getbatch",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqChainid := args[0]
			reqBatchnumber := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetbatchRequest{

				Chainid:     reqChainid,
				Batchnumber: reqBatchnumber,
			}

			res, err := queryClient.Getbatch(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
