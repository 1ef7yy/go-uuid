package v7

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/1ef7yy/go-uuid/types"
)

func GenerateUUIDv7() (types.UUID, error) {
	var uuid types.UUID

	_, err := rand.Read(uuid.Fields[:])
	if err != nil {
		return uuid, err
	}

	timestamp := big.NewInt(time.Now().UnixMilli())

	timestamp.FillBytes(uuid.Fields[0:6])

	uuid.Fields[6] = (uuid.Fields[6] & 0x0f) | 0x70
	uuid.Fields[8] = (uuid.Fields[8] & 0x3f) | 0x80

	return uuid, nil
}
