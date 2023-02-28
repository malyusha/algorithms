package main

import (
	"strconv"
	"strings"
)

func main() {
	isIsomorphicUsingDicts("egg", "add")
}

func isIsomorphicBest(s string, t string) bool {
	dictS, dictT := [256]byte{}, [256]byte{}

	for i := 0; i < len(s); i++ {
		if dictS[s[i]] != dictT[t[i]] {
			return false
		}

		dictS[s[i]] = byte(i + 1)
		dictT[t[i]] = byte(i + 1)
	}

	return true
}

func isIsomorphicUsingDicts(s string, t string) bool {
	dict1, dict2 := [256]rune{}, [256]rune{}
	fillDict(&dict1)
	fillDict(&dict2)

	sr, tr := []rune(s), []rune(t)
	for i := 0; i < len(sr); i++ {
		ch1, ch2 := sr[i], tr[i]

		if dict1[ch1] == -1 && dict2[ch2] == -1 {
			dict1[ch1] = ch2
			dict2[ch2] = ch1
		} else if !(dict1[ch1] == ch2 && dict2[ch2] == ch1) {
			return false
		}
	}

	return true
}

func fillDict(dict *[256]rune) {
	for ix := range dict {
		dict[ix] = -1
	}
}

func isIsomorphicUsingBuilder(s string, t string) bool {
	return transform(s) == transform(t)
}

func transform(s string) string {
	m := make(map[rune]int)
	var b strings.Builder
	for i, ch := range s {
		if _, ok := m[ch]; !ok {
			m[ch] = i
		}

		b.WriteString(strconv.Itoa(m[ch]))
		b.WriteString(" ")
	}

	return b.String()
}
