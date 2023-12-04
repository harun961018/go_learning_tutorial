package main

import (
	"fmt"
	"time"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func PrintTime(label string, t *time.Time) {
	// Printfln("%s: Day: %v: Month: %v Year: %v",
	// 	label, t.Day(), t.Month(), t.Year())
	// layout := "Day: 02 Month: Jan Year: 2006"
	// fmt.Println(label, t.Format(layout))
	fmt.Println(label, t.Format(time.RFC822Z))
}

func time_methods() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
}

func main() {
	time_methods()
}
