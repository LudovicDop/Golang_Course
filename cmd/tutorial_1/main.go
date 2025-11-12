package main

import (
	"fmt"
)

func main() {
	var tab [3]int32
	tab2 := [...]int32{2, 6}
	var tab3 = []int32{3, 2}
	var inc []int32 = []int32{7, 9}

	tab3 = append(tab3, 4222)
	fmt.Print(tab3)
	tab[0] = 42
	tab[1] = 12

	fmt.Println(tab[0])

	fmt.Printf("Content: %v, %v, %v\n", &tab[0], &tab[1], &tab[2])

	fmt.Printf("Content2: %v", tab[1:3])

	tab2 = append(tab2, inc...)
	fmt.Printf("Tab2: %v", tab2)

}
