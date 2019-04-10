package main

import(
	"bytes"
	"os"
	//"fmt"
	"io"
	"github.com/brisdalen/skole/is105-ica02/fileutils"
	"golang.org/x/text/encoding/charmap"
)

var(
	encodings = []*charmap.Charmap {
		charmap.ISO8859_1,	//0
		charmap.ISO8859_2,	//1
		charmap.ISO8859_3,	//2
		charmap.ISO8859_4,	//3
		charmap.ISO8859_5,	//4
		charmap.ISO8859_6,	//5
		charmap.ISO8859_7,	//6
		charmap.ISO8859_8,	//7
		charmap.ISO8859_9,	//8
		charmap.ISO8859_10,	//9
		charmap.ISO8859_13,	//10
		charmap.ISO8859_14,	//11
		charmap.ISO8859_15,	//12
		charmap.ISO8859_16}	//13
)

func main() {

	pathLang01 := "../files/lang01.wl"
	b1 := fileutils.FileToByteslice(pathLang01)
	EncodeISO8859Pages(pathLang01, b1)

	pathLang02 := "../files/lang02.wl"
	b2 := fileutils.FileToByteslice(pathLang02)
	EncodeISO8859Pages(pathLang02, b2)

	pathLang03 := "../files/lang03.wl"
	b3 := fileutils.FileToByteslice(pathLang03)
	EncodeISO8859Pages(pathLang03, b3)
}

func EncodeISO8859Pages(path string, b []byte) {
	input := b
	var out *os.File
	var r io.Reader
	newPath := path[9:len(path)-3]
	// Kanskje lage en mappe for resultat med os.Mkdir?
	err := os.Mkdir(newPath, 0755)
	if err != nil {
		panic(err)
	}

	for _, e := range encodings {
		out, _ = os.Create(newPath + "/" + newPath + " " + e.String() +".txt")
		r = e.NewDecoder().Reader(bytes.NewReader(input))
		io.Copy(out, r)
		out.Close()
	}
}

