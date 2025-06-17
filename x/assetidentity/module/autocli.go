package assetidentity

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "testernet/api/testernet/assetidentity"
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
					RpcMethod: "AssetAll",
					Use:       "list-asset",
					Short:     "List all asset",
				},
				{
					RpcMethod:      "Asset",
					Use:            "show-asset [id]",
					Short:          "Shows a asset by id",
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
					RpcMethod:      "CreateAsset",
					Use:            "create-asset [address] [metadata]",
					Short:          "Create asset",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "UpdateAsset",
					Use:            "update-asset [id] [address] [metadata]",
					Short:          "Update asset",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "address"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "DeleteAsset",
					Use:            "delete-asset [id]",
					Short:          "Delete asset",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
