package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

var m sync.Mutex

func calc(salaries map[string]float32) {
	m.Lock()
	var all, average float32

	for _, salary := range salaries {
		all += salary
	}
	runtime.Gosched()
	average = all / float32(len(salaries))
	fmt.Println(all, average)
	m.Unlock()
}

func add(salaries map[string]float32, name string, salary float32) {
	m.Lock()
	salaries[name] = salary
	fmt.Println(salaries)
	m.Unlock()
}

func main() {
	runtime.GOMAXPROCS(8)

	salaries := map[string]float32{
		"John": 100.0,
		"Jack": 200.0,
	}
	go calc(salaries)
	go add(salaries, "Jane", 300.0)
	time.Sleep(time.Second * 1)
}
