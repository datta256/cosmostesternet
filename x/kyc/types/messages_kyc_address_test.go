package types

import (
	"testing"

	"testernet/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateKycAddress_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateKycAddress
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateKycAddress{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateKycAddress{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateKycAddress_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateKycAddress
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateKycAddress{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateKycAddress{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteKycAddress_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteKycAddress
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteKycAddress{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteKycAddress{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
