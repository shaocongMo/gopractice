package leetcode

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	checkColor := image[sr][sc]
	if checkColor != newColor {
		image[sr][sc] = newColor
		// up
		if sr-1 >= 0 && image[sr-1][sc] == checkColor {
			floodFill(image, sr-1, sc, newColor)
		}
		// down
		if sr+1 < len(image) && image[sr+1][sc] == checkColor {
			floodFill(image, sr+1, sc, newColor)
		}
		//left
		if sc-1 >= 0 && image[sr][sc-1] == checkColor {
			floodFill(image, sr, sc-1, newColor)
		}
		//right
		if sc+1 < len(image[sr]) && image[sr][sc+1] == checkColor {
			floodFill(image, sr, sc+1, newColor)
		}
	}
	return image
}
