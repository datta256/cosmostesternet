package kyc

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "testernet/api/testernet/kyc"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "KycAddressAll",
					Use:       "list-kyc-address",
					Short:     "List all kyc-address",
				},
				{
					RpcMethod:      "KycAddress",
					Use:            "show-kyc-address [id]",
					Short:          "Shows a kyc-address by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateKycAddress",
					Use:            "create-kyc-address [address]",
					Short:          "Create kyc-address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "UpdateKycAddress",
					Use:            "update-kyc-address [id] [address]",
					Short:          "Update kyc-address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "DeleteKycAddress",
					Use:            "delete-kyc-address [id]",
					Short:          "Delete kyc-address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
