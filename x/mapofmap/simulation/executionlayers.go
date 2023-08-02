package simulation

import (
	"math/rand"
	"strconv"

	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"mapofmap/x/mapofmap/keeper"
	"mapofmap/x/mapofmap/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateExecutionlayers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateExecutionlayers{
			Creator: simAccount.Address.String(),
			Index:   strconv.Itoa(i),
		}

		_, found := k.GetExecutionlayers(ctx, msg.Index)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Executionlayers already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateExecutionlayers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount         = simtypes.Account{}
			executionlayers    = types.Executionlayers{}
			msg                = &types.MsgUpdateExecutionlayers{}
			allExecutionlayers = k.GetAllExecutionlayers(ctx)
			found              = false
		)
		for _, obj := range allExecutionlayers {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				executionlayers = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "executionlayers creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Index = executionlayers.Index

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteExecutionlayers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount         = simtypes.Account{}
			executionlayers    = types.Executionlayers{}
			msg                = &types.MsgUpdateExecutionlayers{}
			allExecutionlayers = k.GetAllExecutionlayers(ctx)
			found              = false
		)
		for _, obj := range allExecutionlayers {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				executionlayers = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "executionlayers creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Index = executionlayers.Index

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
