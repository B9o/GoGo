package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	guess := 1.0
	for count := 0; count <= 10; count += 1 {
		if guess*guess > x {
			guess = guess * 2
		} else {
			guess -= (guess*guess - x) / (2 * guess)
		}
	}
	return guess

}

func main() {
	fmt.Println(Sqrt(2))
}

// 牛顿法求平方根
