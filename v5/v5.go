package v5

import (
	"crypto/sha1"

	"github.com/1ef7yy/go-uuid/types"
)

var (
	DNSNamespace  = []byte("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	URLNamespace  = []byte("6ba7b811-9dad-11d1-80b4-00c04fd430c8")
	OIDNamespace  = []byte("6ba7b812-9dad-11d1-80b4-00c04fd430c8")
	X500Namespace = []byte("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
)

func GenerateUUIDv5(namespace, input []byte) (types.UUID, error) {
	hash := sha1.New()
	hash.Write(namespace[:])
	hash.Write(input)

	hashSum := hash.Sum(nil)

	hashSum[6] = (hashSum[6] & 0x0f) | 0x50
	hashSum[8] = (hashSum[8] & 0x3f) | 0x80

	var uuid types.UUID
	copy(uuid.Fields[:], hashSum[:16])
	return uuid, nil
}
