package iso

import (
	"fmt"
	"testing"
)

var greetingtestExtendedASCII = []struct {
	n1       string
	expected string
}{
	{"Salut, ça va °-", ""},
	{"Salut, ça va °-) Κοστίζει €50", ""},
	{"Salut, ça va °-) Κοστίζει ​€50 Forstår du?​",""},
}
//test will fail as the expected string input - redundant test from inital deliverable.
func TestGreetingtestextendedASCII(t *testing.T) {
	for _, v := range greetingtestExtendedASCII{
		if val := GreetingExtendedASCII(); val != v.expected{
			t.Errorf(" greetingExtendedASCII( %q) returned %q expected %q",v.n1, val, v.expected)

		}
	}

}
// test that test the string to see if its within the extended ASCII characters.
func TestGreetingExtendedASCIIONLY(t *testing.T) {
	val := GreetingExtendedASCII()
	for _, i := range val {
		if i < 0x7f && i > 0xff {
			t.Errorf("greetingExtendedASCII() returns non-extended ASCII value: %q - %v", i, i)
			fmt.Printf("%c", i)
		}
		fmt.Printf("%c", i)
	}

}
