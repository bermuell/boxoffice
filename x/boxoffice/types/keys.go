package types

const (
	// ModuleName defines the module name
	ModuleName = "boxoffice"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_boxoffice"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo/value/"
)

const (
	ConcertCreatedEventType         = "new-concert-created" // Indicates what event type to listen to
	ConcertCreatedEventCreator      = "creator"             // Subsidiary information
	ConcertCreatedEventConcertIndex = "concert-index"       // What concert is relevant
)
