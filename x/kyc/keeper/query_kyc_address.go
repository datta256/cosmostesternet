package keeper

import (
	"context"

	"testernet/x/kyc/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) KycAddressAll(ctx context.Context, req *types.QueryAllKycAddressRequest) (*types.QueryAllKycAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var kycAddresss []types.KycAddress

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	kycAddressStore := prefix.NewStore(store, types.KeyPrefix(types.KycAddressKey))

	pageRes, err := query.Paginate(kycAddressStore, req.Pagination, func(key []byte, value []byte) error {
		var kycAddress types.KycAddress
		if err := k.cdc.Unmarshal(value, &kycAddress); err != nil {
			return err
		}

		kycAddresss = append(kycAddresss, kycAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKycAddressResponse{KycAddress: kycAddresss, Pagination: pageRes}, nil
}

func (k Keeper) KycAddress(ctx context.Context, req *types.QueryGetKycAddressRequest) (*types.QueryGetKycAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	kycAddress, found := k.GetKycAddress(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetKycAddressResponse{KycAddress: kycAddress}, nil
}
