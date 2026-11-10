package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(adj [][]int, visited []bool, vertex int, component *[]int) {
	visited[vertex] = true
	*component = append(*component, vertex)

	for _, neighbour := range adj[vertex] {
		if !visited[neighbour] {
			dfs(adj, visited, neighbour, component)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var from, to int
		fmt.Fscan(in, &from, &to)
		adj[from] = append(adj[from], to)
		if from != to {
			adj[to] = append(adj[to], from)
		}
	}

	var answer [][]int
	visited := make([]bool, n+1)

	for vertex := 1; vertex <= n; vertex++ {
		if !visited[vertex] {
			var component []int
			dfs(adj, visited, vertex, &component)
			answer = append(answer, component)
		}
	}
	fmt.Println(len(answer))
	for _, component := range answer {
		fmt.Println(len(component))
		for _, vertex := range component {
			fmt.Print(vertex, " ")
		}
		fmt.Println()
	}
}
