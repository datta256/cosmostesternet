package kyc

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"testernet/testutil/sample"
	kycsimulation "testernet/x/kyc/simulation"
	"testernet/x/kyc/types"
)

// avoid unused import issue
var (
	_ = kycsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateKycAddress = "op_weight_msg_kyc_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKycAddress int = 100

	opWeightMsgUpdateKycAddress = "op_weight_msg_kyc_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKycAddress int = 100

	opWeightMsgDeleteKycAddress = "op_weight_msg_kyc_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKycAddress int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	kycGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		KycAddressList: []types.KycAddress{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		KycAddressCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&kycGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateKycAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateKycAddress, &weightMsgCreateKycAddress, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKycAddress = defaultWeightMsgCreateKycAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKycAddress,
		kycsimulation.SimulateMsgCreateKycAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateKycAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateKycAddress, &weightMsgUpdateKycAddress, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateKycAddress = defaultWeightMsgUpdateKycAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateKycAddress,
		kycsimulation.SimulateMsgUpdateKycAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteKycAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteKycAddress, &weightMsgDeleteKycAddress, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteKycAddress = defaultWeightMsgDeleteKycAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteKycAddress,
		kycsimulation.SimulateMsgDeleteKycAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateKycAddress,
			defaultWeightMsgCreateKycAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				kycsimulation.SimulateMsgCreateKycAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateKycAddress,
			defaultWeightMsgUpdateKycAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				kycsimulation.SimulateMsgUpdateKycAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteKycAddress,
			defaultWeightMsgDeleteKycAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				kycsimulation.SimulateMsgDeleteKycAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
