package utime

import "time"

// UUID time by the RFC 9562
// is a 60 bit timestamp from 15 Oct 1582 UTC
// with 100 ns intervals

var (
	GregOffset = 122192928000000000 * time.Nanosecond
)

func ConvertUnixToGregorian(now time.Duration) (time.Duration, error) {
	gregorianTime := (now / 100) + GregOffset
	return gregorianTime, nil
}
