package algorithms

import "fmt"

func mmmap() {
	s := map[int]bool{5: true, 2: true}
	_, ok := s[5] // check for existence
	fmt.Println(ok)
	s[8] = true // add element
	delete(s, 2) // remove element
	fmt.Println(s)
}