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

func connectedComponent(adj [][]int, visited []bool, vertex int) ([]int, []int) {
	oldVisited := make([]bool, len(visited))
	copy(oldVisited, visited)

	dfs(adj, visited, vertex)
	var component, notVisited []int
	for i, value := range visited {
		if value && !oldVisited[i] {
			component = append(component, i)
		} else if i != 0 && !value {
			notVisited = append(notVisited, i)
		}
	}
	return component, notVisited
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

	var component, notVisited []int
	visited := make([]bool, n+1)
	component, notVisited = connectedComponent(adj, visited, 1)

	var answer [][]int
	answer = append(answer, component)

	for _, vertex := range notVisited {
		if visited[vertex] {
			continue
		}
		component, notVisited = connectedComponent(adj, visited, vertex)
		answer = append(answer, component)
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
