package main

func minFlipsMonoIncr(s string) int {
	counter, flips := 0, 0
	for i := range s {
		if s[i] == '1' {
			counter += 1
		} else {
			flips += 1
			flips = min(flips, counter)
		}
	}

	return flips
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
