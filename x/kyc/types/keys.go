package types

const (
	// ModuleName defines the module name
	ModuleName = "kyc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_kyc"
)

var (
	ParamsKey = []byte("p_kyc")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	KycAddressKey      = "KycAddress/value/"
	KycAddressCountKey = "KycAddress/count/"
)
