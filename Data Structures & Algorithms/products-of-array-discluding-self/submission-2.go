func productExceptSelf(nums []int) []int {
	leftProduct := make(map[int]int)
	rightProduct := make(map[int]int)
	result := make([]int,len(nums))
	
	leftProduct[0] = 1
	rightProduct[len(nums)-1] = 1

	for i:=1;i<len(nums);i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
		rightProduct[len(nums)-i-1] = rightProduct[len(nums)-i] * nums[len(nums)-i]

	}



	
	for i := range nums {
		result[i] = leftProduct[i] * rightProduct[i]
	}

	return result
}
