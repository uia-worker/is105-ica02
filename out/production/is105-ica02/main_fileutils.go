package main

import (
	"fmt"
	"github.com/Nosp1/Is-105/is105-ica02/fileutils"
)

func main ()  {
	a1 := fileutils.FileToByteslice("files/lang01.wl")
	a2 := fileutils.FileToByteslice("files/lang02.wl")
	a3 := fileutils.FileToByteslice("files/lang03.wl")
	fmt.Println("lang01.wl")
	fmt.Printf("%  x\n", a1)
	fmt.Println("lang02.wl")
	fmt.Printf("% x\n", a2)
	fmt.Println("lang03.wl")
	fmt.Printf("% x\n", a3)
}