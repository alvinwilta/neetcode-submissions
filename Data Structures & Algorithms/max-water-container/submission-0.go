func maxArea(heights []int) int {
	left := 0
	right := len(heights)-1
	maxArea := 0

	for left < right {
		// calculate area
		area := min(heights[left], heights[right])*(right-left)
		if area > maxArea {
			maxArea = area
		}

		// traverse
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}