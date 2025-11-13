package main

import (
	"fmt"
)

func main() {
	var str [] rune = []rune("Résumé")


	for i, v := range str{
		fmt.Printf("%v, %v\n", i, v)
	}

}
