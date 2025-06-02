package keeper

import (
	"context"
	"fmt"

	"testernet/x/kyc/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateKycAddress(goCtx context.Context, msg *types.MsgCreateKycAddress) (*types.MsgCreateKycAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if there are any KYCed addresses yet
	count := k.GetKycAddressCount(ctx)
	if count > 0 {
		// Get the first KYCed address (the owner)
		first, found := k.GetKycAddress(ctx, 0)
		if !found {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "owner not found")
		}
		// Only the owner can add new KYC addresses
		if msg.Creator != first.Address {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the owner can KYC new addresses")
		}
	}

	var kycAddress = types.KycAddress{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	id := k.AppendKycAddress(
		ctx,
		kycAddress,
	)

	return &types.MsgCreateKycAddressResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateKycAddress(goCtx context.Context, msg *types.MsgUpdateKycAddress) (*types.MsgUpdateKycAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var kycAddress = types.KycAddress{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
	}

	// Checks that the element exists
	val, found := k.GetKycAddress(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetKycAddress(ctx, kycAddress)

	return &types.MsgUpdateKycAddressResponse{}, nil
}

func (k msgServer) DeleteKycAddress(goCtx context.Context, msg *types.MsgDeleteKycAddress) (*types.MsgDeleteKycAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetKycAddress(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveKycAddress(ctx, msg.Id)

	return &types.MsgDeleteKycAddressResponse{}, nil
}
