package v1

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"github.com/1ef7yy/go-uuid/utils/mac"
	"github.com/1ef7yy/go-uuid/utils/utime"
)

// UUIDv1
// 128 bit long (16 bytes)
// is a 32 char hex string in a format such as follows:
// xxxxxxxx-xxxx-1xxx-xxxx-xxxxxxxxxxxx
//
// first 60 bits are reserved for timestamp from 1582, Oct 15th
// in a 100 nanosecond increment
//
// first 8 chars are time_low
// next 4 chars are time_mid
// next 4 chars are time_high in which the first char is reserved for version number (1 in this case)
// next 3 chars are clock_seq and are needed for randomness
// one char is reserved
// last 12 chars are the node (MAC)

func GenerateUUIDv1() (string, error) {

	node, err := mac.GetNode()

	if err != nil {
		return "", err
	}

	time := utime.GetUUIDTime()

	var clockSeq uint16

	seq := make([]byte, 2)
	_, err = rand.Read(seq)
	if err != nil {
		return "", err
	}

	clockSeq = uint16(seq[0])<<8 | uint16(seq[1])

	timeLow := uint32(time & 0xffffffff)
	timeMid := uint16((time >> 32) & 0xffff)
	timeHigh := uint16((time >> 48) & 0x0fff)
	timeHigh = timeHigh | 0x1000

	var buf bytes.Buffer

	buf.Grow(36)

	buf.WriteString(fmt.Sprintf("%08x-", timeLow))
	buf.WriteString(fmt.Sprintf("%04x-", timeMid))
	buf.WriteString(fmt.Sprintf("%04x-", timeHigh))

	buf.WriteString(fmt.Sprintf("%04x-", clockSeq))

	buf.WriteString(fmt.Sprintf("%012x", node))

	return buf.String(), nil

}
