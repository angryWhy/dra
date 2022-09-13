package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	add(&arr)
	fmt.Println(arr)
}
func add(arr *[5]int) {
	arr[0] += 10
}
