package v4

import (
	"crypto/rand"
	"fmt"

	"github.com/1ef7yy/go-uuid/types"
	// mrand "math/rand"
)

// TODO: make standard compliant
func GenerateUUIDv4() (types.UUID, error) {
	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		return types.UUID{}, fmt.Errorf("error generating uuidv4: %s", err.Error())
	}

	bytes[6] = (bytes[6] & 0x0f) | 0x40
	bytes[8] = (bytes[8] & 0x3f) | 0x80

	return types.UUID{
		Fields: [16]byte(bytes),
	}, nil
}
