func dailyTemperatures(temperatures []int) []int {
	tempIndexStack := make([]int,0)
	result := make([]int,len(temperatures))

	for i:=len(temperatures)-1;i>=0;i-- {
		if len(tempIndexStack) == 0 {
			// calculate distance
			result[i] = 0
		} else {
			// cleanup stack
			var j int
			for j=len(tempIndexStack)-1;j>=0;j-- {
				if temperatures[i] >= temperatures[tempIndexStack[j]] {
					tempIndexStack = tempIndexStack[:j]
				} else {
					break
				}
			}

			// calculate distance
			if len(tempIndexStack) == 0 {
				result[i] = 0
			} else {
				result[i] = tempIndexStack[len(tempIndexStack)-1]-i
			}
		}
		tempIndexStack = append(tempIndexStack,i)
		
	}

	return result
}
