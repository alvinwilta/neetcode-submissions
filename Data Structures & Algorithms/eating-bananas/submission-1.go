func minEatingSpeed(piles []int, h int) int {
	// search for max
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	min := 1
	res := max
	for min <= max {
		mid := (max+min)/2
		if getEatingTime(piles,mid) <= h {
			res = mid
			max = mid-1
		} else {
			min = mid+1
		} 
	}

	return res
}

func getEatingTime(piles []int, k int) int {
	hours := 0
	for _, pile := range piles {
		hours += pile/k
		if pile%k > 0 {
			hours += 1
		}
	}

	return hours
}