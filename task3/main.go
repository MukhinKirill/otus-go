package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := "qwe\\\\4"
	writeUnpackedString(s)
	os.Exit(0)
}

func writeUnpackedString(str string) {
	var currentRune rune
	for i := 0; i < len(str); i++ {
		var alphaCountStr string
		if isAlpha(rune(str[i])) {
			currentRune = rune(str[i])
		} else if isEscape(rune(str[i])) {
			if i < len(str) {
				i++
				currentRune = rune(str[i])
			}
		} else {
			for j := i; j < len(str); j++ {
				if isNumber(rune(str[j])) {
					alphaCountStr = alphaCountStr + string(str[j])
					i++
				} else {
					i--
					break
				}
			}
		}
		count, err := strconv.Atoi(alphaCountStr)
		if err != nil {
			fmt.Printf("%c", currentRune)
		} else {
			for k := 1; k < count; k++ {
				fmt.Printf("%c", currentRune)
			}
		}
	}
}

func isEscape(char rune) bool {
	return '\\' == char
}

func isNumber(char rune) bool {
	if '0' < char && char < '9' {
		return true
	}
	return false
}

func isAlpha(char rune) bool {
	if ('a' < char && char < 'z') || ('A' < char && char < 'Z') {
		return true
	}
	return false
}
