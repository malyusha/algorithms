package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	baseColor := image[sr][sc]
	if baseColor != color {
		fill(image, sr, sc, baseColor, color)
	}

	return image
}

func fill(image [][]int, sr int, sc int, color int, newColor int) {
	if image[sr][sc] == color {
		image[sr][sc] = newColor

		if sr >= 1 {
			fill(image, sr-1, sc, color, newColor)
		}

		if sc >= 1 {
			fill(image, sr, sc-1, color, newColor)
		}

		if sr+1 < len(image) {
			fill(image, sr+1, sc, color, newColor)
		}

		if sc+1 < len(image[0]) {
			fill(image, sr, sc+1, color, newColor)
		}
	}
}
