type Solution struct{
}

func (s *Solution) Encode(strs []string) string {
	// handle existing slashes
	var result string
	for _, str := range strs {
		result += fmt.Sprintf("#%d#%s", len(str), str)
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	for i:=0;i<len(encoded);i++ {
		// 0 = #
		// 1 = n
		if encoded[i] == '#' {
			// find length
			var strLength string
			for j:=i+1;j<len(encoded);j++{
				if encoded[j] == '#' {
					strLength = encoded[i+1:j]
					i=j
					break
				}
			}
			fmt.Println(strLength)
			length, _ := strconv.Atoi(strLength)
			result = append(result, encoded[i+1:i+length+1])
			i = i+length
		}
	}

	return result
}
