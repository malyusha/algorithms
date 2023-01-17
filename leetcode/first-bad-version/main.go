package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	min, max := 1, n

	for min < max {
		mid := min + ((max - min) / 2)

		if isBadVersion(mid) {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return min
}

func isBadVersion(version int) bool {
	return true
}
