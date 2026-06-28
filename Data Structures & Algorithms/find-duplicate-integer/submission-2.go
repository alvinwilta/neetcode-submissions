func findDuplicate(nums []int) int {
    // phase 1, find starting
	slow := 0
	fast := 0
	for slow != fast || slow == 0 {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fmt.Println(fast)

	// 2nd phase , find duplicate
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
