func searchMatrix(matrix [][]int, target int) bool {
	top := 0
	bot := len(matrix)-1
	end := len(matrix[0])-1

	for top <= bot {
		mid := (top+bot)/2
		if matrix[mid][0] <= target && target <= matrix[mid][end] {
			return binarySearch(matrix[mid], target)
		} else if target > matrix[mid][0] {
			top = mid+1
		} else if target < matrix[mid][end] {
			bot = mid-1
		}
	}

	return false
}

// for log(n) searches
func binarySearch(n []int, target int) bool {
	left := 0
	right := len(n)-1

	for left <= right {
		mid := (left+right)/2
		if target == n[mid] {
			return true
		} else if target > n[mid]  {
			left = mid+1
		} else if target < n[mid] {
			right = mid-1
		}
	}

	return false
}