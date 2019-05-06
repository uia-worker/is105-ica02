package main

import (
	"./iso"
	"fmt"
)

func main() {

	extendedASCIIStringLiteral := iso.GetExtendedASCIIStringLiteral()
	iso.IterateOverASCIIStringLiteral(extendedASCIIStringLiteral)
	fmt.Println("Under kommer stringen: \"Salut, ça va °-) Κοστίζει €50\"")
	iso.GreetingExtendedASCII()
	iso.GreetingExtendedASCIIwithHex()
	iso.GreetGreet()
}
