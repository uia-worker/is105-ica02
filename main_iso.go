package main

import "./iso"

func main() {
	extendedASCIIStringLiteral := iso.GetExtendedASCIIStringLiteral()
	iso.IterateOverASCIIStringLiteral(extendedASCIIStringLiteral)
	iso.GreetingExtendedASCII()
}
