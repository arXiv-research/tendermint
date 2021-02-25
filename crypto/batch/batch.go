package batch

import (
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/sr25519"
)

// CheckBatch checks if a key type implements the batch verifier interface.
// Currently only ed25519 & sr25519 supports batch verification.
func CheckBatch(pk crypto.PubKey) (crypto.BatchVerifier, bool) {

	switch pk.Type() {
	case ed25519.KeyType:
		return ed25519.NewBatchVerifier(), true
	case sr25519.KeyType:
		return sr25519.NewBatchVerifier(), true
	}

	// case where the key does not support batch verification
	return nil, false
}
