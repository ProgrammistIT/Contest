package main

import "fmt"

func F(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	answer := 0
	prev := 0
	curr := 1
	for i := 2; i <= n; i++ {
		answer = (prev + curr) % 10
		prev = curr
		curr = answer
	}
	return answer
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(F(n))
}
