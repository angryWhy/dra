package main

import "fmt"

func main() {
	v := 1
	fmt.Println(inc(&v))
	fmt.Println(gc(20, 50))
	fmt.Println(fli(10))
}
func inc(p *int) int {
	*p++
	return *p
}
func gc(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
func fli(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
