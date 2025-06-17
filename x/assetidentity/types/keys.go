package types

const (
	// ModuleName defines the module name
	ModuleName = "assetidentity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_assetidentity"
)

var (
	ParamsKey = []byte("p_assetidentity")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AssetKey      = "Asset/value/"
	AssetCountKey = "Asset/count/"
)
