package bitscale

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"bitscale/testutil/sample"
	bitscalesimulation "bitscale/x/bitscale/simulation"
	"bitscale/x/bitscale/types"
)

// avoid unused import issue
var (
	_ = bitscalesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateDenom = "op_weight_msg_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDenom int = 100

	opWeightMsgUpdateDenom = "op_weight_msg_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDenom int = 100

	opWeightMsgDeleteDenom = "op_weight_msg_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDenom int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bitscaleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		DenomList: []types.Denom{
			{
				Owner: sample.AccAddress(),
				Denom: "0",
			},
			{
				Owner: sample.AccAddress(),
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bitscaleGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateDenom, &weightMsgCreateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDenom = defaultWeightMsgCreateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDenom,
		bitscalesimulation.SimulateMsgCreateDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateDenom, &weightMsgUpdateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDenom = defaultWeightMsgUpdateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDenom,
		bitscalesimulation.SimulateMsgUpdateDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteDenom, &weightMsgDeleteDenom, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDenom = defaultWeightMsgDeleteDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDenom,
		bitscalesimulation.SimulateMsgDeleteDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateDenom,
			defaultWeightMsgCreateDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bitscalesimulation.SimulateMsgCreateDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateDenom,
			defaultWeightMsgUpdateDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bitscalesimulation.SimulateMsgUpdateDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteDenom,
			defaultWeightMsgDeleteDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bitscalesimulation.SimulateMsgDeleteDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
