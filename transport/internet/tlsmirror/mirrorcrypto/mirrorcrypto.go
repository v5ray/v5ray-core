package mirrorcrypto

import "github.com/v4fly/v4ray-core/v0/common/crypto"

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

func generateInitialAEADNonce() crypto.BytesGenerator {
	return crypto.GenerateIncreasingNonce([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
}
