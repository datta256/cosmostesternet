package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateIdentity{}

func NewMsgCreateIdentity(creator string, address string, metadata string, powerlevels string) *MsgCreateIdentity {
	return &MsgCreateIdentity{
		Creator:     creator,
		Address:     address,
		Metadata:    metadata,
		Powerlevels: powerlevels,
	}
}

func (msg *MsgCreateIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateIdentity{}

func NewMsgUpdateIdentity(creator string, id uint64, address string, metadata string, powerlevels string) *MsgUpdateIdentity {
	return &MsgUpdateIdentity{
		Id:          id,
		Creator:     creator,
		Address:     address,
		Metadata:    metadata,
		Powerlevels: powerlevels,
	}
}

func (msg *MsgUpdateIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteIdentity{}

func NewMsgDeleteIdentity(creator string, id uint64) *MsgDeleteIdentity {
	return &MsgDeleteIdentity{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
