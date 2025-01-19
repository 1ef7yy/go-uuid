package utime

import "time"

// UUID time by the RFC 9562 standard
// is a 60 bit timestamp from 15 Oct 1582 UTC
// with 100 ns intervals

// the problem with the "time" package in golang is
// the timestamp nanoseconds are stored as a int64
// so you cannot directly sub from now to 1582
// you have to first sub unix, then greg

func GetUUIDTime() uint64 {
	var timestamp uint64
	now := time.Now().UTC()
	greg := time.Date(1582, time.October, 15, 0, 0, 0, 0, time.UTC)

	unix := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

	gregToUnix := unix.Sub(greg).Nanoseconds()

	timestamp += uint64(gregToUnix)

	timestamp += uint64(now.Sub(unix).Nanoseconds())

	return timestamp
}
