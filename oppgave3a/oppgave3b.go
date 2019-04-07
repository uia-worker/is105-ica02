package main

import(
	"fmt"
	"github.com/brisdalen/skole/is105-ica02/fileutils"
)

func main() {
	f1 := fileutils.FileToByteslice("../files/lang01.wl")
	fmt.Printf("% X\n", f1)
	f2 := fileutils.FileToByteslice("../files/lang02.wl")
	fmt.Printf("% X\n", f2)
	f3 := fileutils.FileToByteslice("../files/lang03.wl")
	fmt.Printf("% X\n", f3)
}

