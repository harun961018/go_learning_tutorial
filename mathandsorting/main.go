package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// math_funcs()
	// rand_number()
	sort_methods()
}

func math_funcs() {
	val1 := 279.00
	val2 := 48.95
	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val2))
	Printfln("Copysign: %v", math.Copysign(val1, -5))
	Printfln("Floor: %v", math.Floor(val2))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, 2))
	Printfln("Round: %v", math.Round(val2))
	Printfln("RoundToEven: %v", math.RoundToEven(val2))
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func rand_number() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, rand.Int())
	}
}

func sort_methods() {
	ints := []int{9, 4, 2, -1, 10}
	Printfln("Ints: %v", ints)
	sort.Ints(ints)
	Printfln("Ints Sorted: %v", ints)
	floats := []float64{279, 48.95, 19.50}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats Sorted: %v", floats)
	strings := []string{"Kayak", "Lifejacket", "Stadium"}
	Printfln("Strings: %v", strings)
	if !sort.StringsAreSorted(strings) {
		sort.Strings(strings)
		Printfln("Strings Sorted: %v", strings)
	} else {
		Printfln("Strings Already Sorted: %v", strings)
	}
}
