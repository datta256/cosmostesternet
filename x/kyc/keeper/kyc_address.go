package keeper

import (
	"context"
	"encoding/binary"

	"testernet/x/kyc/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKycAddressCount get the total number of kycAddress
func (k Keeper) GetKycAddressCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.KycAddressCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetKycAddressCount set the total number of kycAddress
func (k Keeper) SetKycAddressCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.KycAddressCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendKycAddress appends a kycAddress in the store with a new id and update the count
func (k Keeper) AppendKycAddress(
	ctx context.Context,
	kycAddress types.KycAddress,
) uint64 {
	// Create the kycAddress
	count := k.GetKycAddressCount(ctx)

	// Set the ID of the appended value
	kycAddress.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.KycAddressKey))
	appendedValue := k.cdc.MustMarshal(&kycAddress)
	store.Set(GetKycAddressIDBytes(kycAddress.Id), appendedValue)

	// Update kycAddress count
	k.SetKycAddressCount(ctx, count+1)

	return count
}

// SetKycAddress set a specific kycAddress in the store
func (k Keeper) SetKycAddress(ctx context.Context, kycAddress types.KycAddress) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.KycAddressKey))
	b := k.cdc.MustMarshal(&kycAddress)
	store.Set(GetKycAddressIDBytes(kycAddress.Id), b)
}

// GetKycAddress returns a kycAddress from its id
func (k Keeper) GetKycAddress(ctx context.Context, id uint64) (val types.KycAddress, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.KycAddressKey))
	b := store.Get(GetKycAddressIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKycAddress removes a kycAddress from the store
func (k Keeper) RemoveKycAddress(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.KycAddressKey))
	store.Delete(GetKycAddressIDBytes(id))
}

// GetAllKycAddress returns all kycAddress
func (k Keeper) GetAllKycAddress(ctx context.Context) (list []types.KycAddress) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.KycAddressKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KycAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// IsKycApproved checks if an address exists in the KYC list
func (k Keeper) IsKycApproved(ctx sdk.Context, addr string) bool {
	addresses := k.GetAllKycAddress(ctx)
	for _, kyc := range addresses {
		if kyc.Address == addr {
			return true
		}
	}
	return false
}

// GetKycAddressIDBytes returns the byte representation of the ID
func GetKycAddressIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.KycAddressKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
