package chapter_01

import (
	"fmt"
	"os"
	"strings"
)

func hello() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
func demo() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
