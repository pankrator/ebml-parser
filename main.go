package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pankrator/ebml-parser/tools"
)

func main() {
	// s := 0b00000010
	// fmt.Println(s)
	reader := open("sample.webm")

	parser := tools.Parser{}
	elPipe := parser.Parse(reader)
	for el := range elPipe {
		fmt.Printf("%s[%s][%s]=%d\n", el.Element.Name, el.Element.Hex, string(el.Element.Typ), len(el.Data))
	}
}

func open(name string) io.Reader {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	return f
}
