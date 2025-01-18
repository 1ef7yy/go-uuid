package utime

import (
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	expectations := map[time.Duration]time.Duration{
		1645557742000000000 * time.Nanosecond: time.Nanosecond * 138648505420000000,
	}

	for k, v := range expectations {
		val, err := ConvertUnixToGregorian(k)

		if err != nil {
			t.Fatalf("got error: %s", err.Error())
		}
		if val != v {
			t.Fatalf("test failed, expected: %s, got: %s", v.String(), val.String())
		}

		t.Logf("[success] expected: %d, got: %d", v, val)
	}
}
