import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	// maps count to element, k: count, v: arr of el for that count
	countMapNumArr := make(map[int][]int)
	// store count for each element, k: num, v: count
	numCountMap := make(map[int]int)

	// calculate count
	for _, num := range nums {
		numCountMap[num] += 1
	}

	// map count
	countArr := []int{}
	for num, count := range numCountMap {
		if _, ok := countMapNumArr[count]; !ok {
			countMapNumArr[count] = []int{num}
			countArr = append(countArr,count)
		} else {
			countMapNumArr[count] = append(countMapNumArr[count], num)
		}
	}

	// sort the count
	slices.Sort(countArr)

	// build sorted k frequent element
	result := []int{}
	for _, count := range countArr {
		result = append(countMapNumArr[count], result...)
	}
	return result[:k]
}
