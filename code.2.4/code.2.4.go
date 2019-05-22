package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s[1] = 15
	fmt.Println(s[1])

	var s1 []int
	s1 = make([]int, 15, 25)
	s1 = append(s1, 1)
	s1 = append(s1, s...)

	s2 := s1[1:3:3]
	copy(s2, s1[14:20])
	fmt.Println(s1, "len is:", len(s1), "cap is:", cap(s1))
	fmt.Println(s2)

	a := [3]int{1, 2, 3}
	fmt.Println(a)
}
