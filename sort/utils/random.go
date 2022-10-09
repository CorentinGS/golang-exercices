package utils

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
)

func GenerateRandomSeed() (int64, error) {
	// Generate a random seed
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		return -1, fmt.Errorf("cannot seed math/rand package with cryptographically secure random number generator")
	}
	return int64(binary.LittleEndian.Uint64(b[:])), nil
}
