package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var from, to int
		fmt.Fscan(in, &from, &to)

		adj[from] = append(adj[from], to)
	}

	color := make([]int, n+1)

	for i := range color {
		color[i] = -1
	}

	for i := 1; i <= n; i++ {
		if color[i] != -1 {
			continue
		}

		color[i] = 0
		var queue []int
		queue = append(queue, i)
		for len(queue) != 0 {
			parent := queue[0]
			queue = queue[1:]

			for _, v := range adj[parent] {
				if color[v] == -1 {
					queue = append(queue, v)
					color[v] = 1 - color[parent]
					continue
				} else if color[parent] == color[v] {
					fmt.Print("NO")
					return
				}
			}
		}
	}

	fmt.Println("YES")

	for i := 1; i <= n; i++ {
		if color[i] == 0 {
			fmt.Print(i, " ")
		}
	}
}
