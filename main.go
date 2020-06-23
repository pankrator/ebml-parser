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
	// b := encoder.WriteVInt(s)
	// fmt.Println(tools.ReadVint(b, 0))
	b := tools.UInt64ToByte(1000)
	reverseArr(b)
	fmt.Println(tools.ToUint64(b), len(b))
	os.Exit(0)

	reader := open("sample.webm")

	parser := tools.Parser{}
	elPipe := parser.Parse(reader)
	for el := range elPipe {
		fmt.Printf("%s[%s][%s]=%d\n", el.Element.Name, el.Element.Hex, string(el.Element.Typ), el.DataSize)
		if el.Element.Name == "Timestamp" {
			number := tools.ToUint64(el.Data)
			fmt.Printf("%+v\n", number)
		}
	}
}

func reverseArr(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func open(name string) io.Reader {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	return f
}
