package chapter_01

import (
	"fmt"
	"os"
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
