package main

import (
	"sync"
	"time"
	"strconv"
)

var m *sync.RWMutex

var wg sync.WaitGroup

func main() {
	m = new(sync.RWMutex)

	//for i := 1; i < 5; i++ {
	//	go read(i)
	//}

	//for j := 5; j < 7; j++ {
	//	go write(j)
	//}
	wg.Add(5)

	go read("Read1")
	go read("Read2")
	go write("Write3")
	go write("Write4")
	go read("Read5")

	wg.Wait()
}


func read(i string) {
	defer wg.Done()
	println(i, "\t-- read enter")
	m.RLock()
	println(i, "\t-- read start")
	for a := 0; a < 3; a++ {
		println(i, "\t-- reading..." + strconv.Itoa(a))
		time.Sleep(500 * time.Millisecond)
	}
	println(i, "\t-- read over")
	m.RUnlock()
}

func write(i string) {
	defer wg.Done()
	println(i, "\t-- write enter")
	m.Lock()
	println(i, "\t-- write start")
	for a := 0; a < 3; a++ {
		println(i, "\t-- writeing..." + strconv.Itoa(a))
		time.Sleep(500 * time.Millisecond)
	}
	println(i, "\t-- write over")
	m.Unlock()
}
