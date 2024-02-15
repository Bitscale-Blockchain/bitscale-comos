package types

const (
	// ModuleName defines the module name
	ModuleName = "bitscale"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bitscale"

	// Version defines the current version the IBC module supports
	Version = "bitscale-1"

	// PortID is the default port id that module binds to
	PortID = "bitscale"
)

var (
	ParamsKey = []byte("p_bitscale")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("bitscale-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
