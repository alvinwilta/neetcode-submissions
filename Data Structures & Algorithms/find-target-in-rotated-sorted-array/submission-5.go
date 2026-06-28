func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1

	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return mid
		}

		// check where to go
		if nums[left] < nums[right] {
			if target < nums[mid] {
				right = mid-1
			} else {
				left = mid+1
			}
		} else {
			if nums[left] <= nums[mid] {
				// ascending on the left
				if target >= nums[left] && target < nums[mid] {
					right = mid-1
				} else {
					left = mid+1
				}
			} else {
				// ascending on the right
				if target > nums[mid] && target <= nums[right] {
					left = mid+1
				} else {
					right = mid-1
				}
			}
		}
	}

	return -1
}
