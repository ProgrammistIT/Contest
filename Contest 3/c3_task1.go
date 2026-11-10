package main

import (
	"bufio"
	"fmt"
	"os"
)

func getNeighbors(graph [][]int, vertex int) []int {
	answer := make([]int, 0)
	for j := range graph {
		if graph[vertex][j] == 1 {
			answer = append(answer, j)
		}
	}
	return answer
}

func dfs(graph [][]int, vertex int, parents []int, parent int, visited []bool) []int {
	visited[vertex] = true
	parents[vertex] = parent
	for _, neighbour := range getNeighbors(graph, vertex) {
		if !visited[neighbour] {
			cycle := dfs(graph, neighbour, parents, vertex, visited)
			if cycle != nil {
				return cycle
			}
		} else if neighbour != parent {
			cycle := make([]int, 0)
			current := vertex

			for current != neighbour {
				cycle = append(cycle, current)
				current = parents[current]
			}
			cycle = append(cycle, neighbour)

			return cycle
		}
	}

	return nil
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	graph := make([][]int, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &graph[i][j])
		}
	}

	visited := make([]bool, len(graph))
	parents := make([]int, len(graph))

	for i := range parents {
		parents[i] = -1
	}

	for vertex := range graph {
		if !visited[vertex] {
			cycle := dfs(graph, vertex, parents, -1, visited)
			if cycle != nil {
				fmt.Println("YES")
				fmt.Println(len(cycle))
				for _, vertex := range cycle {
					fmt.Print(vertex+1, " ")
				}
				return
			}
		}
	}

	fmt.Println("NO")
}
