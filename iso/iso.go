package iso

import "fmt"

// Oppgave 2a
// Lag selv en "string literal" med utvidet ASCII
// Blir det kun en slik "string literal" eller trenger man flere
// for å representere utvidet ASCII?
// Her er consten for exstended ASCII
const extendedASCII = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f" +
	"\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f" +
	"\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf" +
	"\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf" +
	"\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf" +
	"\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf" +
	"\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef" +
	"\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"

//metode for å kjøre ASCII const
func RunASCIIConst() {
	IterateOverASCIIStringLiteral(extendedASCII)
}

// IterateOverASCIIStringLiteral tar en "string literal" som INN-data
//%X  base 16 loweer-case letters for a-f
//%q double-qouted string safely escaped with go Syntax
func IterateOverASCIIStringLiteral(sl string) {
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%X %q %b\n", sl[i], sl[i], sl[i])
	}
}

//funskjon for å hente constructen extended ascii
func GetExtendedASCIIStringLiteral() string {
	return extendedASCII
}

// GreetingExtendedASCII returnerer en tekst-streng i
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtendedASCII() string {
	s := []byte("\"Salut, ça va °-) Κοστίζει €50\"")
	//fmt.Printf("%X\n %q %b\n", s,s,s )
	for i := 0; i < len(s); i++{
		fmt.Printf("%X %q %b\n", s[i], s[i], s[i])
	}
	return string(s[:])


}
