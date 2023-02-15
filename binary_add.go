package leetcode

func add(a byte, b byte) (byte, byte) {
	if a == '1' && b == '1' {
		return '0', '1'
	} else if a == '1' || b == '1' {
		return '1', '0'
	} else {
		return '0', '0'
	}
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	l := max(la, lb)
	res := make([]byte, l+1)
	var rem byte = '0'
	for i := 0; i < l; i++ {
		var ai byte = '0'
		var bi byte = '0'
		if i < la {
			ai = a[la-i-1]
		}
		if i < lb {
			bi = b[lb-i-1]
		}
		ri, rem2 := add(ai, bi)
		ri2, rem3 := add(ri, rem)
		res[l-i] = ri2
		rem, _ = add(rem2, rem3)
	}

	if rem == '1' {
		res[0] = '1'
		return string(res)
	}

	return string(res[1:])
}
