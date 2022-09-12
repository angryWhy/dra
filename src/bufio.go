package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	test()
}
func str_spl() {
	s := "hello go"
	arr := strings.Contains(s, "o")
	fmt.Println(arr)
}
func str_spl2() {
	s := "hello go"
	arr := strings.HasPrefix(s, "e")
	fmt.Println(arr)
}
func str_spl3() {
	s := "hello go"
	arr := strings.HasSuffix(s, "o")
	fmt.Println(arr)
}
func str_spl4() {
	s := "hello go"
	arr := strings.Index(s, "o")
	fmt.Println(arr)
}
func add() {
	s := "hello"
	t := "ni"
	p := "hao"
	sub1 := s + "" + t + "" + p
	sub2 := fmt.Sprintf("%s %s %s", s, t, p)
	sub3 := strings.Join([]string{s, t, p}, "")
	sub4 := strings.Builder{}
	sub4.WriteString(s)
	sub4.WriteString("")
	sub4.WriteString(t)
	sub4.WriteString("")
	sub4.WriteString(p)
	sub5 := sub4.String()
	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println(sub3)
	fmt.Println(sub5)
}
func test() {
	var ua uint64 = math.MaxUint64
	fmt.Printf("ua=%d %b\n", ua, ua)
	ue := uint8(ua)
	fmt.Printf("ue=%d %b\n", ue, ue)
	ub := uint64((ue))
	fmt.Printf("ub=%d %b\n", ub, ub)
}
