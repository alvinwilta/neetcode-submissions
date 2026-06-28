type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(fmt.Sprintf("%d#", len(str)))
		sb.WriteString(str)
	}

	return sb.String()
}

// 4#Hello5#World

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	decoded := []string{}
	for i:=0;i<len(encoded);i++ {
		// get length
		lengthStr := ""
		for ;i<len(encoded);i++ {
			if string(encoded[i]) == "#" {
				break
			}
			lengthStr += string(encoded[i])
		}
		length, _ := strconv.Atoi(lengthStr)
		decoded = append(decoded, encoded[i+1:length+i+1])
		i = i+length
	}

	return decoded
}
