package main

import (
	"fmt"
)

func bf(mainString, patternString string) bool {
	mainStrLen := len(mainString)
	patternStrLen := len(patternString)

	if mainStrLen < patternStrLen {
		return false
	}

	for i := 0; i < mainStrLen-patternStrLen+1; i++ {
		subStr := mainString[i : patternStrLen+i]
		if subStr == patternString {
			return true
		}
	}

	return false
}

func main() {
	mainString := "acdefg"
	patternString := "cde"
	fmt.Println("bf result: ", bf(mainString, patternString))
}
