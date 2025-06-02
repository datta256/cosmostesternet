package identity

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "testernet/api/testernet/identity"
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
					RpcMethod: "IdentityAll",
					Use:       "list-identity",
					Short:     "List all identity",
				},
				{
					RpcMethod:      "Identity",
					Use:            "show-identity [id]",
					Short:          "Shows a identity by id",
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
					RpcMethod:      "CreateIdentity",
					Use:            "create-identity [address] [metadata] [powerlevels]",
					Short:          "Create identity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "metadata"}, {ProtoField: "powerlevels"}},
				},
				{
					RpcMethod:      "UpdateIdentity",
					Use:            "update-identity [id] [address] [metadata] [powerlevels]",
					Short:          "Update identity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "address"}, {ProtoField: "metadata"}, {ProtoField: "powerlevels"}},
				},
				{
					RpcMethod:      "DeleteIdentity",
					Use:            "delete-identity [id]",
					Short:          "Delete identity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
