package leetcode

const alen = 26

var exists = struct{}{}

func distinctNames(ideas []string) int64 {
	var res int64 = 0
	a := "a"[0]

	gr := make([]map[string]struct{}, alen)
	for i := 0; i < alen; i++ {
		gr[i] = make(map[string]struct{})
	}
	for i := 0; i < len(ideas); i++ {
		idea := ideas[i]
		gr[idea[0]-a][idea[1:]] = exists
	}

	// gr := make(map[int]map[string]struct{})
	// a := "a"[0]
	// for i := 0; i < len(ideas); i++ {
	// 	idea := ideas[i]
	// 	gi := int(idea[0] - a)
	// 	_, ok := gr[gi]
	// 	if !ok {
	// 		gr[gi] = make(map[string]struct{})
	// 	}
	// 	gr[int(idea[0]-a)][idea[1:]] = exists
	// }

	for i := 0; i < alen-1; i++ {
		for j := i + 1; j < alen; j++ {
			mutual := 0
			for k := range gr[i] {
				_, ok := gr[j][k]
				if ok {
					mutual++
				}
			}
			res += int64((len(gr[i]) - mutual) * (len(gr[j]) - mutual) * 2)
		}
	}
	return res
}

// func main() {
// 	fmt.Println(distinctNames([]string{"coffee", "donuts", "time", "toffee"})) // 6
// }
