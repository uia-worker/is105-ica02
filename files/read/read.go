package read

import(
	"fmt"
	"github.com/brisdalen/skole/is105-ica02/fileutils"

)

func Translate(filename string) {
	fmt.Printf("% b", fileutils.FileToByteslice(filename))
}