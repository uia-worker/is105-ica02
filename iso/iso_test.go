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

/*
Tester om metoden GreetingExstendedASCII() er innenfor 0x7f og 0xff
 */
func TestGreetingExtendedASCIIONLY(t *testing.T) {
	val := GreetingExtendedASCII()
	for _, i := range val {
		if i < 0x7f && i > 0xff {
			t.Errorf("greetingExtendedASCII() returns non-extended ASCII value: %q - %v", i, i)
			fmt.Printf("%c", i)

		}

	}
	return true
}
