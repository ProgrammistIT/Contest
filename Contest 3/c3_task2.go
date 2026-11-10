package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	A, B int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, count int
	fmt.Fscan(in, &n, &count)
	uniqueMap := make(map[Pair]bool)
	for i := 0; i < count; i++ {
		var first, second int
		fmt.Fscan(in, &first, &second)
		uniqueMap[Pair{first, second}] = true
	}
	if len(uniqueMap) == int(n*(n-1)/2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
