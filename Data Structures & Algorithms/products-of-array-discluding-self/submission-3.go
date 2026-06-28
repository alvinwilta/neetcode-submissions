func productExceptSelf(nums []int) []int {
	prefixSum := map[int]int{}
	suffixSum := map[int]int{}
	prefixSum[0]=1
	suffixSum[len(nums)-1]=1

	// calculate sums
	for i:=1;i<len(nums);i++ {
		prefixSum[i] = prefixSum[i-1]*nums[i-1]
		suffixSum[len(nums)-i-1] = suffixSum[len(nums)-i]*nums[len(nums)-i]
	}

	// calculate solution
	solution := []int{}
	for i:=0;i<len(nums);i++ {
		solution = append(solution, prefixSum[i]*suffixSum[i])
	}

	return solution
}
