package main

import "fmt"

var v = "1, 2, 3"

func main() {
	v := []int{1, 2, 3}
	if v != nil {
		v := 123
		fmt.Println("v3 : ", v)
	}
	fmt.Println("v2 : ", v)
}

func init() {
	fmt.Println("v1 : ", v)
}
