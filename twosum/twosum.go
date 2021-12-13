package twosum

func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		if kmap, ok := mp[target-v]; ok {
			return []int{kmap, i}
		}
		mp[v] = i
	}
	return []int{}
}
