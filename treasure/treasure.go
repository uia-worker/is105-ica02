package treasure

import (
	"bytes"
	"fmt"
)
// PrintTreasureUTF8 tar en "string literal" som INN-data
// og skriver ut en korrekt text på med UTF-8 koding
// Koden er for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasureUTF8(treasureString string) {
	treasureBytes := []byte(treasureString)

	treasureBytes = bytes.Replace(treasureBytes, []byte("\xe5"), []byte(""), -1)
	treasureBytes = bytes.Replace(treasureBytes, []byte("\xe6"), []byte(""), -1)
	treasureBytes = bytes.Replace(treasureBytes, []byte("\xf8"), []byte(""), -1)

	fmt.Printf("%q", treasureBytes)
}