package main

import "fmt"

func main() {
	fmt.Println(gcd(16, 24))
	fmt.Println(fib(100))
}

// GCD 两个整数的最大公约数
// greatest common divisor
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Fibonacci 斐波那契数列
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
