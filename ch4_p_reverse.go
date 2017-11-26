package main

import (
	"fmt"
)

func reverse(pInt *[5]int, len int) {
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		pInt[i], pInt[j] = pInt[j], pInt[i]
	}
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	reverse(&arr, 5)
	fmt.Printf("%v\n", arr) //[5 4 3 2 1]
}
