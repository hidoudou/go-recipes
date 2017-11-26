package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	fmt.Printf("nonempty2, %q\n", out)
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	data_2 := []string{"one", "", "three"}
	data_2 = nonempty(data_2)
	data_2_ret := nonempty(data_2)

	fmt.Printf("%q\n", data_2_ret)
	fmt.Printf("%q\n", data_2)

	data_3 := []string{"hello", "", "world"}
	fmt.Printf("%q\n", nonempty2(data_3))

	//output
	// ["one" "three"]
	// ["one" "three" "three"]
	// ["one" "three"]
	// ["one" "three"]
	// nonempty2, []
	// ["hello" "world"]
}
