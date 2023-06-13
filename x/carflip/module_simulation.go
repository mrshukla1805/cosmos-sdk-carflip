package carflip

import (
	"math/rand"

	"carflip/testutil/sample"
	carflipsimulation "carflip/x/carflip/simulation"
	"carflip/x/carflip/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = carflipsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateCar = "op_weight_msg_car"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCar int = 100

	opWeightMsgUpdateCar = "op_weight_msg_car"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCar int = 100

	opWeightMsgDeleteCar = "op_weight_msg_car"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCar int = 100

	opWeightMsgCreateRequest = "op_weight_msg_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRequest int = 100

	opWeightMsgUpdateRequest = "op_weight_msg_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRequest int = 100

	opWeightMsgDeleteRequest = "op_weight_msg_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRequest int = 100

	opWeightMsgUpdateRequestStatus = "op_weight_msg_update_request_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRequestStatus int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	carflipGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		CarList: []types.Car{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		CarCount: 2,
		RequestList: []types.Request{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		RequestCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&carflipGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateCar int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCar, &weightMsgCreateCar, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCar = defaultWeightMsgCreateCar
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCar,
		carflipsimulation.SimulateMsgCreateCar(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCar int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCar, &weightMsgUpdateCar, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCar = defaultWeightMsgUpdateCar
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCar,
		carflipsimulation.SimulateMsgUpdateCar(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCar int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCar, &weightMsgDeleteCar, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCar = defaultWeightMsgDeleteCar
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCar,
		carflipsimulation.SimulateMsgDeleteCar(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRequest, &weightMsgCreateRequest, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRequest = defaultWeightMsgCreateRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRequest,
		carflipsimulation.SimulateMsgCreateRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateRequest, &weightMsgUpdateRequest, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRequest = defaultWeightMsgUpdateRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRequest,
		carflipsimulation.SimulateMsgUpdateRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteRequest, &weightMsgDeleteRequest, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRequest = defaultWeightMsgDeleteRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRequest,
		carflipsimulation.SimulateMsgDeleteRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRequestStatus int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateRequestStatus, &weightMsgUpdateRequestStatus, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRequestStatus = defaultWeightMsgUpdateRequestStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRequestStatus,
		carflipsimulation.SimulateMsgUpdateRequestStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
