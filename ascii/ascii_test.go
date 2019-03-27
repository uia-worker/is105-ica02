package ascii

import(
	"testing"
)

func TestGreetingASCII(t *testing.T) {
	s := GreetingASCII()
	isAscii := withinAscii(s)

	if !isAscii {
		t.Fatal("Not ASCII")
	}
}

func withinAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
		if s[i] < 128 && s[i] > 255 {
			return false
		}
	}
	return true
}

