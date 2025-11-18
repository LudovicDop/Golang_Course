package main

import (
	"fmt"
	"time"
)

func main() {
	
}

func test(ch chan int) {
	defer close(ch)
	ch <- 104
	ch <- 105
	fmt.Printf("Exiting process\n")
}
