package other

func Permute(nums []int) [][]int {
	leng := len(nums)
	res := make([][]int, 6)
	current := leng - 2
	next := leng - 1
	for i := 0; i < 6; i++ {
		tmp := nums
		tmp[current], tmp[next] = tmp[next], tmp[current]
		res[i] = tmp
		if next == leng-1 {
			next = 0
		} else {
			next++
		}
		if current == leng-1 {
			current = 0
		} else {
			current++
		}
	}
	return res
}
