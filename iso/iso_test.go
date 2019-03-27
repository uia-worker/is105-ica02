package iso

import "testing"

var greetingtestExtendedASCII = []struct {
	n1       string
	expected string
}{
	{"Salut, ça va °-", ""},
	{"Salut, ça va °-) Κοστίζει €50", ""},
}

func TestGreetingtestextendedASCII(t *testing.T) {
	for _, v := range greetingtestExtendedASCII{
		if val := GreetingExtendedASCII(); val != v.expected{
			t.Errorf(" greetingExtendedASCII( %q) returned %q expected %q",v.n1, val, v.expected)

		}
	}

}
func TestGreetingExtendedASCIIONly(t *testing.T) {
	val := GreetingExtendedASCII()
	for _, i := range val {
		if !(i > 0x80 && i < 0xff) {
			t.Errorf("greetingExtendedASCII() returns non-extended ASCII value: %q - %v", i, i)
		}
	}
}
