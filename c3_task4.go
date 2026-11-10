package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(adj [][]int, visited []bool, vertex int) {
	visited[vertex] = true

	for _, neighbour := range adj[vertex] {
		if !visited[neighbour] {
			dfs(adj, visited, neighbour)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, edges int
	fmt.Fscan(in, &n, &edges)

	adj := make([][]int, n+1)

	for i := 0; i < edges; i++ {
		var from, to int
		fmt.Fscan(in, &from, &to)
		adj[from] = append(adj[from], to)
		if from != to {
			adj[to] = append(adj[to], from)
		}
	}

	visited := make([]bool, n+1)
	dfs(adj, visited, 1)

	answer := make([]int, 0)
	for i, value := range visited {
		if value {
			answer = append(answer, i)
		}
	}

	fmt.Println(len(answer))
	for _, value := range answer {
		fmt.Print(value, " ")
	}
}
