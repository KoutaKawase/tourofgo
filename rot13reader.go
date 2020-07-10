package main

// import (
// 	"io"
// 	"os"
// 	"strings"
// )

// type rot13Reader struct {
// 	r io.Reader
// }

// func (rot13 rot13Reader) Read(b []byte) (int, error) {
// 	n, err := rot13.r.Read(b)
// 	for i, v := range b {
// 		switch {
// 		case v >= 'A' && v < 'N', v >= 'a' && v < 'n':
// 			b[i] += 13
// 		case v >= 'N' && v <= 'Z', v <= 'n' && v <= 'z':
// 			b[i] -= 13
// 		}
// 	}
// 	return n, err
// }

// func main() {
// 	s := strings.NewReader("hello!")
// 	r := rot13Reader{s}
// 	io.Copy(os.Stdout, &r)
// }
