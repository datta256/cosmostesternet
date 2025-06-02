package ante

import (
	"encoding/binary"
	"fmt"
	"testernet/x/kyc/keeper"

	"github.com/cosmos/cosmos-sdk/runtime"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type KycAnteDecorator struct {
	KycKeeper keeper.Keeper
}

func NewKycAnteDecorator(k keeper.Keeper) KycAnteDecorator {
	return KycAnteDecorator{KycKeeper: k}
}

func (kad KycAnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	store := runtime.KVStoreAdapter(kad.KycKeeper.StoreService().OpenKVStore(ctx))
	bz := store.Get([]byte("bypass_kyc_tx_count"))
	var count uint64 = 0
	if bz != nil {
		count = binary.BigEndian.Uint64(bz)
	}
	if count < 3 { // Allow only the first 3 txs to bypass KYC
		count++
		bz := make([]byte, 8)
		binary.BigEndian.PutUint64(bz, count)
		store.Set([]byte("bypass_kyc_tx_count"), bz)
		return next(ctx, tx, simulate)
	}

	msgs := tx.GetMsgs()
	var foundSigner bool
	for _, msg := range msgs {
		// Check "from" address(es)
		if msgWithSigners, ok := msg.(interface{ GetSigners() []sdk.AccAddress }); ok {
			signers := msgWithSigners.GetSigners()
			if len(signers) > 0 {
				foundSigner = true
				for _, addr := range signers {
					if !kad.KycKeeper.IsKycApproved(ctx, addr.String()) {
						return ctx, fmt.Errorf("from address %s is not KYC approved: %w", addr.String(), sdkerrors.ErrUnauthorized)
					}
				}
			}
		}

		// Check "to" address for common message types (e.g., bank MsgSend)
		switch m := msg.(type) {
		case *banktypes.MsgSend:
			if !kad.KycKeeper.IsKycApproved(ctx, m.ToAddress) {
				return ctx, fmt.Errorf("to address %s is not KYC approved: %w", m.ToAddress, sdkerrors.ErrUnauthorized)
			}
			// Add more cases here for other message types with "to" addresses if needed
		}
	}
	if !foundSigner {
		return ctx, fmt.Errorf("transaction is not from a KYC address: %w", sdkerrors.ErrUnauthorized)
	}

	return next(ctx, tx, simulate)
}
