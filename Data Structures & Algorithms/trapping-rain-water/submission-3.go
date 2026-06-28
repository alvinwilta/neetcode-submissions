func trap(height []int) int {
	// solution: O(n) space and O(n) time
	n := len(height)
	if n < 3 {
		// cannot make container
		return 0
	}

	// get the max height from left/right side, starting from that index
	maxLeftHeight := make(map[int]int)
	maxRightHeight := make(map[int]int)

	for i:=1;i<n;i++{
		// compare to the prev height (height[i-1])
		// because the current index is the container, not the wall
		maxLeftHeight[i] = max(maxLeftHeight[i-1],height[i-1])
	}
	for i:=n-2;i>=0;i--{
		maxRightHeight[i] = max(maxRightHeight[i+1],height[i+1])
	}
	
	// calculate water area
	water := 0
	for i:=1;i<n-1;i++ {
		area := min(maxLeftHeight[i],maxRightHeight[i])-height[i]
		if area > 0 {
			water += area
		}
	}

	return water
}

func max(a,b int) int {
	if a>b {return a}
	return b
}

func min(a,b int) int {
	if a>b{return b}
	return a
}
