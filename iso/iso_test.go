package iso

import (
	"testing"
)

func TestGreetingExtendedASCII(t *testing.T) {
	s := GreetingExtendedASCII()
	isExtendedAscii := withinExtendedAscii(s)

	if !isExtendedAscii {
		t.Fatal("Not ExtendedASCII")
	}
}

func withinExtendedAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 255 {
			return false
		}

	}
	return true
}
