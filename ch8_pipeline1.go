package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			fmt.Println("send data to naturals.....")
			naturals <- x
		}

		fmt.Println("func 1 exit......")
	}()

	go func() {
		for {
			fmt.Println("read data from naturals.....")
			x := <-naturals
			fmt.Println("send data to squares.....")
			squares <- x * x
		}
	}()

	for {
		fmt.Println("read data from squares.....")
		fmt.Println(<-squares)
	}
}
