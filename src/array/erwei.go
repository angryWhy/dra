package main

import "fmt"

func a() {

}
func main() {
	arr1 := [5][3]int{{1, 23, 3}, {1, 2, 3}}

	for i, ele := range arr1 {
		for j, n := range ele {
			fmt.Printf("%d %d %d \n ", i, j, n)
		}
	}
}
