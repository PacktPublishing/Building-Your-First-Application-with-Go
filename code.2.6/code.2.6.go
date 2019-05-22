package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	m := map[int]int{1: 2, 2: 3}
	fmt.Println(m)

	var m2 map[int]string
	m2 = make(map[int]string, 10)
	fmt.Println(m2)

	m[1] = 15
	delete(m, 1)
	fmt.Println(m)

	m3 := map[string]string{"name": "Jack"}
	res, _ := json.Marshal(m3)
	fmt.Println(res)
}
