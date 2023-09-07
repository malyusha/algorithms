package main

import (
	"fmt"
)

func maximalSquare(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	dp := make([][]int, rows+1)
	for i := 0; i < rows+1; i++ {
		dp[i] = make([]int, cols+1)
	}

	max := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if grid[i-1][j-1] == '0' {
				continue
			}

			dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}
	}

	return max * max
}

func maximalSquareCache(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	cache := make(map[string]int)

	memoMaxSquare(matrix, cache, rows, cols, 0, 0)

	max := 0

	for _, maxSquare := range cache {
		if maxSquare > max {
			max = maxSquare
		}
	}

	return max * max
}

func memoMaxSquare(matrix [][]byte, cache map[string]int, rows, cols, r, c int) int {
	if r >= rows || c >= cols {
		return 0
	}

	cacheKey := fmt.Sprintf("%d_%d", r, c)
	computed, ok := cache[cacheKey]
	if !ok {
		// compute
		down := memoMaxSquare(matrix, cache, rows, cols, r+1, c)
		right := memoMaxSquare(matrix, cache, rows, cols, r, c+1)
		diagRight := memoMaxSquare(matrix, cache, rows, cols, r+1, c+1)

		cache[cacheKey] = 0
		if matrix[r][c] == 1 {
			cache[cacheKey] = 1 + min(down, right, diagRight)
		}

		computed = cache[cacheKey]
	}

	return computed
}

func min(args ...int) int {
	if len(args) == 0 {
		panic("no args")
	}

	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}

	return min
}

func maximalSquareSimple(matrix [][]byte) int {
	var maxSquare int
	var cols int
	var one = byte(1)
	rows := len(matrix)
	if rows != 0 {
		cols = len(matrix[0])
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var curSq int
			isSquare := true

			if matrix[i][j] != one {
				continue
			}

			curSq = 1
			// start to go diagonally to the right and check if we can form a square
			for curSq+i < rows && curSq+j < cols && isSquare {
				for k := j; k <= curSq+j; k++ {
					if matrix[i+curSq][k] != one {
						isSquare = false
						break
					}
				}
				for k := i; k <= curSq+i; k++ {
					if matrix[k][curSq+j] != one {
						isSquare = false
						break
					}
				}

				if isSquare {
					curSq++
				}
			}

			if curSq > maxSquare {
				maxSquare = curSq
			}
		}
	}

	return maxSquare * maxSquare
}
