package main

import(
	"fmt"
)

const byteString = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
const newByte = "\xc2\xbd\x3f\x3d\x3f\x20\xe2\x8c\x98"

func main() {
	fmt.Printf("%s\n", byteString)
	fmt.Printf("%q\n", byteString)
	fmt.Printf("%+q\n", byteString)

	for i := 0; i < len(byteString); i++ {
		fmt.Printf("%c", byteString[i])
	}
	fmt.Println("\n-----------")

	fmt.Printf("%s\n", newByte)
}