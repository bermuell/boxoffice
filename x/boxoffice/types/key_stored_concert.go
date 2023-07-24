package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredConcertKeyPrefix is the prefix to retrieve all StoredConcert
	StoredConcertKeyPrefix = "StoredConcert/value/"
)

// StoredConcertKey returns the store key to retrieve a StoredConcert from the index fields
func StoredConcertKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
