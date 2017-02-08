package core

import "crypto/sha256"

func ContactKey(contactBytes []byte) []byte {
	sum := sha256.Sum256(contactBytes)
	return sum[:]
}
