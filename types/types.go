package types

import "fmt"

type UUID struct {
	Fields [16]byte
}

func (u UUID) String() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		u.Fields[0:4],
		u.Fields[4:6],
		u.Fields[6:8],
		u.Fields[8:10],
		u.Fields[10:],
	)
}
