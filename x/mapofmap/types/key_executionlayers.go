package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExecutionlayersKeyPrefix is the prefix to retrieve all Executionlayers
	ExecutionlayersKeyPrefix = "Executionlayers/value/"
)

// ExecutionlayersKey returns the store key to retrieve a Executionlayers from the index fields
func ExecutionlayersKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
