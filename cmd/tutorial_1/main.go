package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var mutex = sync.Mutex{}
var fakeDb []string = []string{"user1", "user2", "user3", "user4", "user5", "user6"}
var result []string

func main() {

	t0 := time.Now()

	for i := range fakeDb {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("The result is: %v", result)
	fmt.Printf("\n\x1b[33mTotal execution time: %v\x1b[0m", time.Since(t0))
}

func dbCall(i int) {
	// var delay float32 = rand.Float32() * 3

	// fmt.Printf("delay: %v\n", delay)
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("The result is: %v\n", i)
	mutex.Lock()
	result = append(result, fakeDb[i])
	mutex.Unlock()
	wg.Done()

}
