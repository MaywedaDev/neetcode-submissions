import "slices"

func groupAnagrams(strs []string) [][]string {
	// Create map to hold anagrams
	anagramsMap := make(map[string][]string)
	// Output array
	anagramsArr := [][]string{}
	for _,str := range strs{
		// Convert string to runes array. Then sort
		runes := []rune(str)
		slices.Sort(runes)
		// Convert runes back to string
		key := string(runes)
		// Add to slice of anagrams with same key
		anagramsMap[key] = append(anagramsMap[key], str)
	}

	// Add all slices in values of anagramMap to outpur array
	for _, slice := range anagramsMap{
		anagramsArr = append(anagramsArr, slice)
	}

	return anagramsArr
	}
