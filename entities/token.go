package entities

import (
	"bytes"

	"github.com/daoleno/uniswap-sdk-core/entities"
)

// SortsBefore returns true if the address of token a sorts before the address of the token b.
func SortsBefore(a, b *entities.Token) (bool, error) {
	if a.ChainId() != b.ChainId() {
		return false, entities.ErrDifferentChain
	}
	if a.Address == b.Address {
		return false, entities.ErrSameAddress
	}
	return bytes.Compare(a.Address[:], b.Address[:]) < 0, nil
}
