package main

import (
	"fmt"
	"github.com/Nosp1/Is-105/is105-ica02/unicode"
)

func main()  {
	jp := unicode.Translate("nord og sør", "jp")
	is := unicode.Translate("nord og sør", "is")

	fmt.Println(jp)
	fmt.Println(is)

}
