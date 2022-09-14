package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{}
	s = make([]int, 3, 5)
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = append(s, 999)
	fmt.Printf("len %d cap %d %d\n", len(s), cap(s), s)
	subSlice()
}
func subSlice() {
	arr := make([]int, 5, 5)
	var sub = arr[0:2]
	fmt.Printf("%p,%p", &arr, &sub[0])
	sub[1] = 9
	fmt.Println(sub, arr)
	sub = append(sub, 999)
	fmt.Println(sub, arr)
	sub = append(sub, 111)
	sub = append(sub, 111)
	fmt.Println(sub, arr)
	sub = append(sub, 111)
	fmt.Printf("%v %v len sub=%d len arr=%d  cap sub=%d    cap arr=%d  sub-ptr=%p,arr-ptr=%p,arr[0]=%p", sub, arr, len(sub), len(arr), cap(sub), cap(arr), &sub, &arr, &arr[0])

}
