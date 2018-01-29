package main

import (
	"fmt"
	"io"
	"os"
)

type MyReader struct{}

func Validate(r io.Reader) {
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ {
		n, err := r.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
				return
			}
		}

		o += n
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			return
		}
	}

	if o == 0 {
		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
		return
	}
	fmt.Println("OK!")
}

//实现一个Reader类型， 它不断生成ASCII字符‘A’的流
// add a Read([]byte) (int, error) method to MyReader
func (mr MyReader) Read(b []byte) (n int, err error) {
	i := 0
	for ; i < len(b); i++ {
		b[i] = 'A'

		if i == 10 {
			b[i] = 'B'
		}
	}
	return i, nil
}

func main() {
	Validate(MyReader{})
}

//output
// $ ./io_reader_t2.exe
// got byte 42 at offset 10, want 'A'
