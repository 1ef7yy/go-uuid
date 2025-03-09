package v1

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"

	"github.com/1ef7yy/go-uuid/types"
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

func GenerateUUIDv1() (types.UUID, error) {

	node, err := mac.GetNode()

	if err != nil {
		return types.UUID{}, err
	}

	time := utime.GetUUIDTime()

	var clockSeq uint16

	seq := make([]byte, 2)
	_, err = rand.Read(seq)
	if err != nil {
		return types.UUID{}, err
	}

	clockSeq = uint16(seq[0])<<8 | uint16(seq[1])

	timeLow := uint32(time & 0xffffffff)
	timeMid := uint16((time >> 32) & 0xffff)
	timeHigh := uint16((time >> 48) & 0x0fff)
	timeHigh = timeHigh | 0x1000 // set version number

	var buf bytes.Buffer

	buf.Grow(16)

	// unreadable af

	err = binary.Write(&buf, binary.BigEndian, timeLow)
	if err != nil {
		return types.UUID{}, err
	}
	err = binary.Write(&buf, binary.BigEndian, timeMid)
	if err != nil {
		return types.UUID{}, err
	}
	err = binary.Write(&buf, binary.BigEndian, timeHigh)
	if err != nil {
		return types.UUID{}, err
	}
	err = binary.Write(&buf, binary.BigEndian, clockSeq)
	if err != nil {
		return types.UUID{}, err
	}
	err = binary.Write(&buf, binary.BigEndian, node)
	if err != nil {
		return types.UUID{}, err
	}

	return types.UUID{
		Fields: [16]byte(buf.Bytes()),
	}, nil

}
