package main

import (
	"fmt"
	"github.com/hidoudou/go-recipes/intset"
)

func Example_one() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))
}

func Example_two() {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)
	fmt.Println(x.String())
	fmt.Println(x)
}

func main() {
	Example_one()

	Example_two()
}