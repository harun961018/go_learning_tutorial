package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

type Product struct {
	Name, Category string
	Price          float64
}

var Kayak = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    279,
}
var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

func readData(reader io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func writeData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	// r := strings.NewReader("Kayak")
	// var builder strings.Builder
	// writeData(r, &builder)
	// Printfln("String builder contents: %s", builder.String())
	var b bool = true
	var str string = "Hello"
	var fval float64 = 99.99
	var ival int = 200
	var pointer *int = &ival
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	for _, val := range []interface{}{b, str, fval, ival, pointer} {
		encoder.Encode(val)
	}
	fmt.Print(writer.String())
}
