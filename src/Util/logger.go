package util

import "fmt"

/*
Going to print the data
*/
func LogPrintln(isprint bool, log ...any) {
	if !isprint {
		return
	}
	for _, val := range log {
		fmt.Printf("%v", val)
	}
	fmt.Print("\n")
}
