package main

import (
	"fmt"
)

func main() {
	fmt.Print(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1

	for start < end {
		if !isAlphaNumeric(s[start]) {
			start++
			continue
		}

		if !isAlphaNumeric(s[end]) {
			end--
			continue
		}

		if transform(s[start]) != transform(s[end]) {
			return false
		}

		start++
		end--
	}

	return true
}

func isAlphaNumeric(l byte) bool {
	return (transform(l) >= 97 && transform(l) <= 122) || (l >= 48 && l <= 57)
}

func transform(l byte) byte {
	if l <= 90 && l >= 65 {
		l += 32
	}

	return l
}
