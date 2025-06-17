package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAsset{}

func NewMsgCreateAsset(creator string, address string, metadata string) *MsgCreateAsset {
	return &MsgCreateAsset{
		Creator:  creator,
		Address:  address,
		Metadata: metadata,
	}
}

func (msg *MsgCreateAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAsset{}

func NewMsgUpdateAsset(creator string, id uint64, address string, metadata string) *MsgUpdateAsset {
	return &MsgUpdateAsset{
		Id:       id,
		Creator:  creator,
		Address:  address,
		Metadata: metadata,
	}
}

func (msg *MsgUpdateAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAsset{}

func NewMsgDeleteAsset(creator string, id uint64) *MsgDeleteAsset {
	return &MsgDeleteAsset{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
