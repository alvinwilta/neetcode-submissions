func carFleet(target int, position []int, speed []int) int {
	// k: pos, v: original idx
	speedMapIdx := make(map[int]int)
	for i, pos := range position {
		speedMapIdx[pos] = i
	}

	sort.Ints(position)
	maxTime := float64(-1)
	fleet := 0
	
	for i:=len(position)-1;i>=0;i-- {
		// (target-position)/speed
		time := float64(target-position[i])/float64(speed[speedMapIdx[position[i]]])
		fmt.Println(time)
		if time > maxTime {
			maxTime = time
			fleet++
		}
	}

	return fleet
}
