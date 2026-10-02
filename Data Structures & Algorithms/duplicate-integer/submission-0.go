func hasDuplicate(nums []int) bool {

    numsMap := map[int]int{}

    for _, num := range(nums) {
        numsMap[num] = numsMap[num] + 1
        if numsMap[num] > 1 {
            return true
        }
    }

    return false
}
