package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // n数组 长度 容量
	if s == nil {
		fmt.Println("nil!")
	}
}
