func permute(nums []int) [][]int {
	return permuteRec([]int{}, nums)
}

func permuteRec(curr, nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{curr}
	}

	res := [][]int{}
	for i, num := range nums {
		currCpy := append([]int{}, curr...)
		currCpy = append(currCpy, num)
		newNums := append([]int{},nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		res = append(res, permuteRec(currCpy, newNums)...)
	}

	return res
}
