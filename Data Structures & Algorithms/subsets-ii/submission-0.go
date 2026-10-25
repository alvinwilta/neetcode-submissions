
func subsetsWithDup(nums []int) [][]int {
	// sort nums
	sort.Ints(nums)
	return recSubsetsWithDup([]int{}, nums)
}

func recSubsetsWithDup(curr, nums []int) [][]int {
	res := [][]int{curr}
	if len(nums) == 0 {
		return res
	}

	dupl := map[int]bool{}
	for i, num := range nums {
		if dupl[num] {
			continue
		}

		currCpy := append([]int{}, curr...)
		currCpy = append(currCpy, num)
		numsCpy := append([]int{}, nums[i+1:]...)
		res = append(res, recSubsetsWithDup(currCpy, numsCpy)...)
		dupl[num] = true
	}

	return res
}