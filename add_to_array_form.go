// https://leetcode.com/problems/add-to-array-form-of-integer/

package leetcode

import (
	"strconv"
)

func addToArrayForm(num []int, k int) []int {
	sk := strconv.Itoa(k)
	num2 := make([]int, len(sk))
	for i, v := range sk {
		num2[i] = int(v - '0')
	}
	l1 := len(num)
	l2 := len(num2)
	l := max(l1, l2)
	res := make([]int, l+1)
	rem := 0
	for i := 0; i < l; i++ {
		ai := 0
		bi := 0
		if i < l1 {
			ai = num[l1-i-1]
		}
		if i < l2 {
			bi = num2[l2-i-1]
		}
		ri := ai + bi + rem
		res[l-i] = ri % 10
		rem = ri / 10
	}
	if rem != 0 {
		res[0] = rem
		return res
	}
	return res[1:]
}
