package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var x byte
			fmt.Fscan(reader, &x)
			if x == 1 {
				fmt.Println(i+1, j+1)
			}
		}
	}
}
