func threeSum(nums []int) [][]int {
	result := make([][]int,0)
	sort.Ints(nums)

	for i:=0;i<len(nums)-2;i++ {
		if i > 0 && nums[i] == nums[i-1] {
            continue
        }
		left :=i+1
		right := len(nums)-1
		for left < right {
			sum := nums[i]+nums[left]+nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i],nums[left],nums[right]})
				left++
				right--
				for nums[right+1] == nums[right] && left < right {
					right--
				}
				for nums[left-1] == nums[left] && left < right {
					left++
				}
			} else if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			}
		}
	}

	return result
}
