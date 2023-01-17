package main

func reverseWords(s string) string {
	curr := 0
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		isEnd := i == len(b)-1
		if b[i] == ' ' || isEnd {
			j := curr
			c := i
			if !isEnd {
				c--
			}

			for j < c {
				b[j], b[c] = b[c], b[j]
				j++
				c--
			}

			curr = i + 1
		}
	}

	return string(b)
}
