package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				for k := 0; k < n; k++ {
					if matrix[j][k] == 1 {
						if matrix[i][k] != 1 {
							fmt.Print("NO")
							return
						}
					}
				}
			}
		}
	}
	fmt.Print("YES")
}
