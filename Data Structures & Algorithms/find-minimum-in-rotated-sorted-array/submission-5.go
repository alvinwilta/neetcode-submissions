func findMin(nums []int) int {
	left := 0
	right := len(nums)-1
	res := nums[0]

	for left <= right {
		if nums[left] <= nums[right] {
			res = min(res, nums[left])
			break
		}

		mid := (left+right)/2
		res = min(res, nums[mid])
		if nums[left] > nums[mid] {
			right = mid-1
		} else {
			left = mid+1
		}
	}

	return res
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}