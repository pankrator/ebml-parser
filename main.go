package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pankrator/ebml-parser/tools"
)

func main() {
	// var s uint64 = 0xffffffffffffff
	// var s2 uint64 = 0b11111111111111111111111111111111111111111111111111111111
	// fmt.Println(s, s2)
	// os.Exit(0)
	// b, _ := hex.DecodeString(s)
	// fmt.Println(binary.BigEndian.Uint64(b))
	reader := open("sample.webm")

	parser := tools.Parser{}
	elPipe := parser.Parse(reader)
	for el := range elPipe {
		fmt.Printf("%s[%s][%s]=%d\n", el.Element.Name, el.Element.Hex, string(el.Element.Typ), el.DataSize)
	}
}

func open(name string) io.Reader {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	return f
}
