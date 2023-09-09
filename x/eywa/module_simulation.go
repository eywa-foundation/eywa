package eywa

import (
	"math/rand"

	"eywa/testutil/sample"
	eywasimulation "eywa/x/eywa/simulation"
	"eywa/x/eywa/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = eywasimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgRegisterUser = "op_weight_msg_register_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterUser int = 100

	opWeightMsgCreateHandshake = "op_weight_msg_create_handshake"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHandshake int = 100

	opWeightMsgCreateRelayServer = "op_weight_msg_create_relay_server"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRelayServer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	eywaGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&eywaGenesis)
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

	var weightMsgRegisterUser int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterUser, &weightMsgRegisterUser, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterUser = defaultWeightMsgRegisterUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterUser,
		eywasimulation.SimulateMsgRegisterUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateHandshake int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateHandshake, &weightMsgCreateHandshake, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHandshake = defaultWeightMsgCreateHandshake
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHandshake,
		eywasimulation.SimulateMsgCreateHandshake(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateRelayServer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRelayServer, &weightMsgCreateRelayServer, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRelayServer = defaultWeightMsgCreateRelayServer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRelayServer,
		eywasimulation.SimulateMsgCreateRelayServer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterUser,
			defaultWeightMsgRegisterUser,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				eywasimulation.SimulateMsgRegisterUser(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateHandshake,
			defaultWeightMsgCreateHandshake,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				eywasimulation.SimulateMsgCreateHandshake(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRelayServer,
			defaultWeightMsgCreateRelayServer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				eywasimulation.SimulateMsgCreateRelayServer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
