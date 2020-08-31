package algorithms

import "fmt"

func sum_in_array(arr []int, sum int) {
	m := make(map[int]bool)
	for _, v := range arr {
		diff := sum - v
		_, ok := m[diff] // check for existence
		if ok {
			fmt.Printf("%d + %d\n", v, diff)
		}
		m[v] = true
	}
}
