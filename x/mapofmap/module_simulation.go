package mapofmap

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"mapofmap/testutil/sample"
	mapofmapsimulation "mapofmap/x/mapofmap/simulation"
	"mapofmap/x/mapofmap/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = mapofmapsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateExecutionlayers = "op_weight_msg_executionlayers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateExecutionlayers int = 100

	opWeightMsgUpdateExecutionlayers = "op_weight_msg_executionlayers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateExecutionlayers int = 100

	opWeightMsgDeleteExecutionlayers = "op_weight_msg_executionlayers"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteExecutionlayers int = 100

	opWeightMsgAddbatch = "op_weight_msg_addbatch"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddbatch int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	mapofmapGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ExecutionlayersList: []types.Executionlayers{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&mapofmapGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateExecutionlayers int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateExecutionlayers, &weightMsgCreateExecutionlayers, nil,
		func(_ *rand.Rand) {
			weightMsgCreateExecutionlayers = defaultWeightMsgCreateExecutionlayers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateExecutionlayers,
		mapofmapsimulation.SimulateMsgCreateExecutionlayers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateExecutionlayers int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateExecutionlayers, &weightMsgUpdateExecutionlayers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateExecutionlayers = defaultWeightMsgUpdateExecutionlayers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateExecutionlayers,
		mapofmapsimulation.SimulateMsgUpdateExecutionlayers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteExecutionlayers int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteExecutionlayers, &weightMsgDeleteExecutionlayers, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteExecutionlayers = defaultWeightMsgDeleteExecutionlayers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteExecutionlayers,
		mapofmapsimulation.SimulateMsgDeleteExecutionlayers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddbatch int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddbatch, &weightMsgAddbatch, nil,
		func(_ *rand.Rand) {
			weightMsgAddbatch = defaultWeightMsgAddbatch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddbatch,
		mapofmapsimulation.SimulateMsgAddbatch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateExecutionlayers,
			defaultWeightMsgCreateExecutionlayers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				mapofmapsimulation.SimulateMsgCreateExecutionlayers(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateExecutionlayers,
			defaultWeightMsgUpdateExecutionlayers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				mapofmapsimulation.SimulateMsgUpdateExecutionlayers(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteExecutionlayers,
			defaultWeightMsgDeleteExecutionlayers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				mapofmapsimulation.SimulateMsgDeleteExecutionlayers(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddbatch,
			defaultWeightMsgAddbatch,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				mapofmapsimulation.SimulateMsgAddbatch(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
