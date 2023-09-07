package main

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	if len(s) > 12 {
		return []string{}
	}

	res := make([]string, 0)

	backtrack(s, 0, 0, "", &res)

	return res
}

func backtrack(s string, i int, dots int, curIp string, res *[]string) {
	if dots == 4 && i == len(s) {
		*res = append(*res, curIp[:len(curIp)-1])
		return
	}
	if dots > 4 {
		return
	}

	for j := i; j < min(i+3, len(s)); j++ {
		intPart, err := strconv.Atoi(s[i : j+1])
		if err != nil {
			return
		}

		if intPart <= 255 && (i == j || s[i] != '0') {
			backtrack(s, j+1, dots+1, curIp+s[i:j+1]+".", res)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
