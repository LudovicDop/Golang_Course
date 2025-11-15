package main

import (
	"fmt"
)

func main() {
	var ptr *int32 = new(int32)
	var ptr2
	*ptr = 42
	fmt.Printf("address ptr: %v and value %v", ptr, *ptr)

}
