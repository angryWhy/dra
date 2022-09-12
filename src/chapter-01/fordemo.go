package chapter_01

import (
	"fmt"
	"os"
)

const (
	Hello string = "hello"
)

func fordemo() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*
 *
 *
 */
