package unicode

import (
	"fmt"
)


// Kode for Oppgave 4a
func Translate(expression string, language string) string {

	uni := ""
	if expression == "nord og sør" {
		if language == "jp" {
			uni = "\x22" +  expression +  "\x22" + " på japansk er: " + "\x22\xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97\x22"
		} else if language == "is" {
			uni =  "\x22 " + expression  + "\x22 " +  " på islandsk er: " + "\x22\x6E\x6f\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\x22"
		}
	}
	return uni
}

// Kode for Oppgave 4b
func UnicodeCodePointDemo() {
	// Hva er dette?
	// Er det likt på MS Windows og macOS?
	fmt.Println("\xf0\x9F\x98\x80")
	fmt.Println("\xf0\x9F\x98\x97")
	// Demonstrerer at deler av et tegn representert med flere bytes
	// kan ikke tolkes innenfor koden (unicode)
	fmt.Println("\xf0\x9F\x98")
	fmt.Println("\xf0\x9F")
	fmt.Println("\xf0")
}
