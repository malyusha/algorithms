package main

var romanMap = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result := 0
	start := len(s) - 1
	for i := start; i >= 0; i-- {
		r := s[i]
		addition := romanMap[r]
		result += addition

		if i == start {
			continue
		}

		prev := s[i+1]
		if (r == 'I' && (prev == 'V' || prev == 'X')) ||
			(r == 'X' && (prev == 'L' || prev == 'C')) ||
			(r == 'C' && (prev == 'D' || prev == 'M')) {
			result -= addition * 2
		}
	}

	return result
}
