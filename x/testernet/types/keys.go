package types

const (
	// ModuleName defines the module name
	ModuleName = "testernet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_testernet"
)

var (
	ParamsKey = []byte("p_testernet")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
