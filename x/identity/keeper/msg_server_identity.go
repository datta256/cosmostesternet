package keeper

import (
	"context"
	"fmt"

	"testernet/x/identity/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateIdentity(goCtx context.Context, msg *types.MsgCreateIdentity) (*types.MsgCreateIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var identity = types.Identity{
		Creator:     msg.Creator,
		Address:     msg.Address,
		Metadata:    msg.Metadata,
		Powerlevels: msg.Powerlevels,
	}

	id := k.AppendIdentity(
		ctx,
		identity,
	)

	return &types.MsgCreateIdentityResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateIdentity(goCtx context.Context, msg *types.MsgUpdateIdentity) (*types.MsgUpdateIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var identity = types.Identity{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Address:     msg.Address,
		Metadata:    msg.Metadata,
		Powerlevels: msg.Powerlevels,
	}

	// Checks that the element exists
	val, found := k.GetIdentity(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetIdentity(ctx, identity)

	return &types.MsgUpdateIdentityResponse{}, nil
}

func (k msgServer) DeleteIdentity(goCtx context.Context, msg *types.MsgDeleteIdentity) (*types.MsgDeleteIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetIdentity(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveIdentity(ctx, msg.Id)

	return &types.MsgDeleteIdentityResponse{}, nil
}
