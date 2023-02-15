// https://leetcode.com/problems/add-to-array-form-of-integer/

package leetcode

func addToArrayForm(num []int, k int) []int {
	ln := len(num)
	for i := 0; i < ln && k > 0; i++ {
		k += num[ln-i-1]
		num[ln-i-1] = k % 10
		k = k / 10
	}
	for k > 0 {
		num = append([]int{k % 10}, num...)
		k = k / 10
	}
	return num
}
