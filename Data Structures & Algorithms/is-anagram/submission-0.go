func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	charsMap := map[rune]int{}

	for _, char := range s {
		charsMap[char] = charsMap[char] + 1
	}

	for _, char := range t {
		if charsMap[char] == 0 {
			return false
		}
		charsMap[char] = charsMap[char] - 1
	}

	for _, val := range charsMap{
		if val != 0{
			return false
		}
	}




	return true
}
