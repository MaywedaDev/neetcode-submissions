
func topKFrequent(nums []int, k int) []int {
	//Use a map to get the frequency of each number
	frequencyMap := map[int]int{}

	for _, num := range nums {
		frequencyMap[num] = frequencyMap[num] + 1
	}
	// Use the map to create a slice of map with key(number) and value(frequency)
	frequencySlices := []map[string]int{}

	for key, val := range frequencyMap{
		frequencySlices = append(frequencySlices, map[string]int{"num": key, "freq": val})
	}
	// Sort by frequency
	sort.Slice(frequencySlices, func(a, b int) bool {
		return (frequencySlices[a]["freq"] < frequencySlices[b]["freq"])
	})
	// Iterate and return the K most frequent

	outputArr := []int{}
	numsLength := len(frequencySlices)
	for i := 1; i <= k; i++ {
		outputArr = append(outputArr, frequencySlices[numsLength - i]["num"])
	}

	return outputArr
}
