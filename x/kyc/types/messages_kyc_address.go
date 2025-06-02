package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateKycAddress{}

func NewMsgCreateKycAddress(creator string, address string) *MsgCreateKycAddress {
	return &MsgCreateKycAddress{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateKycAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateKycAddress{}

func NewMsgUpdateKycAddress(creator string, id uint64, address string) *MsgUpdateKycAddress {
	return &MsgUpdateKycAddress{
		Id:      id,
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateKycAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteKycAddress{}

func NewMsgDeleteKycAddress(creator string, id uint64) *MsgDeleteKycAddress {
	return &MsgDeleteKycAddress{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteKycAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
