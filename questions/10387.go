package questions

import "fmt"

func FirstUniqChar(s string) int {
	var x int
	var firstMap = make(map[rune]int)
	for k, v := range s {
		if (1<<(v-97))&x == 0 {
			firstMap[v] = k
			x |= 1 << (v - 97)
		} else {
			delete(firstMap, v)
		}
	}
	fmt.Println(firstMap)
	fmt.Printf("%028b\n", x)
	min := len(s)
	for _, v := range firstMap {
		if v < min {
			min = v
		}
	}

	return min
}
