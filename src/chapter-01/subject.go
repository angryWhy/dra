package chapter_01

import (
	"fmt"
	"os"
	"strings"
)

// 打印 `os.Args[0]`，即被执行命令本身的名字。
func subject1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

// 打印每个参数的索引和值，每个一行。
func subject2() {
	for s, arg := range os.Args[1:] {
		fmt.Println(s, arg)
	}
}
