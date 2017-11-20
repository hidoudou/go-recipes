package main
import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	b := a
	reverse(a[:])
	fmt.Println(a)
	fmt.Println(b)

	s := []int{0, 1, 2, 3, 4, 5}
	t := s
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)

	fmt.Println(t)

	//两个slice不能使用==做比较，只能遍历元素进行比较
	//slice可以和nil用==做比较
	//一个nil值的slice的len和cap都是0
	//一个非nil值得slice的len和cap也有可能是0，如[]int{}
	var ss []int
	ss = nil
	ss = []int(nil)
	//上面的ss的len=0，且ss==nil
	ss = []int{}//len(ss) = 0, ss != nil
	fmt.Println(ss)

	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)
}