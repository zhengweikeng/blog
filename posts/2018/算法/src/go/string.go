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

var squreMap = make(map[int]uint32)

func initSqureNum() {
	squreMap[0] = 1

	for i := 1; i <= 26; i++ {
		if i == 1 {
			squreMap[i] = 26
		} else {
			squreMap[i] = squreMap[i-1] * 26
		}
	}
}

func hash(str string) uint32 {
	var result uint32
	strLen := len(str)

	for index, rune := range str {
		result += (uint32(rune%97) + 1) * squreMap[strLen-index-1]
	}
	return result
}

func rk(mainString, patternString string) bool {
	mainStrLen := len(mainString)
	patternStrLen := len(patternString)

	if mainStrLen < patternStrLen {
		return false
	}

	var subStrHash uint32
	patternStrHash := hash(patternString)

	for i := 0; i < mainStrLen-patternStrLen+1; i++ {
		subStr := mainString[i : patternStrLen+i]
		subStrHash = hash(subStr)

		if subStrHash == patternStrHash {
			return true
		}
	}

	return false
}

func main() {
	mainString := "acdefg"
	patternString := "cdeg"
	fmt.Println("bf result: ", bf(mainString, patternString))

	initSqureNum()
	fmt.Println("rk result: ", rk(mainString, patternString))

}
