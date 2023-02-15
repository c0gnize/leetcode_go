package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	res := 0
	cur := 0
	next := 0
	last := len(nums) - 1
	for i := 0; i < last; i++ {
		next = max(next, i+nums[i])
		if i == cur {
			res++
			cur = next
			if cur >= last {
				return res
			}
		}
	}
	return res
}
