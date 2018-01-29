package main

import(
	"fmt"
	"strings"
	"io"
)

func main() {
	r := strings.NewReader("hello, world!")
	b := make([]byte, 8)
	
	for{
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:%v] = %q\n", n, b[:n])

		if err == io.EOF{
			break
		}
	}
}

//output
$ ./io_reader_t1.exe
n = 8 err = <nil> b = [104 101 108 108 111 44 32 119]
b[:8] = "hello, w"
n = 5 err = <nil> b = [111 114 108 100 33 44 32 119]
b[:5] = "orld!"
n = 0 err = EOF b = [111 114 108 100 33 44 32 119]
b[:0] = ""