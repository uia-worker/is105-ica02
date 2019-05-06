package main

import (
	"./iso"
	"fmt"
)

func main() {

	extendedASCIIStringLiteral := iso.GetExtendedASCIIStringLiteral()
	fmt.Println("Her kommer ASCII og uvitdet ascii fra 0x80 til  FF")
	iso.IterateOverASCIIStringLiteral(extendedASCIIStringLiteral)
	fmt.Println("Under kommer stringen: \"Salut, ça va °-) Κοστίζει €50\"")
	iso.GreetingExtendedASCII()
	fmt.Println("Under kommer stringen  med Hexadesimale kodepunkt")
	iso.GreetingExtendedASCIIWithHex()
	fmt.Println("Under kommer strengen: \"Salut, ça va °-) Ça coute ​€50\"")
	iso.GreetGreet()
}
