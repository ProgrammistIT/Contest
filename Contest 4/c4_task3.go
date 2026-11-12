package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const INF = math.MaxInt

func dijkstraParentList(matrix [][]int, n int, s int) []int {
	dist := make([]int, n+1)
	visited := make([]bool, n+1)
	parents := make([]int, n+1)

	for i := 0; i < len(dist); i++ {
		dist[i] = INF
		parents[i] = -1
	}

	parents[s] = 0
	dist[s] = 0
	for step := 0; step < n; step++ {
		minVertex := -1

		for j := 1; j <= n; j++ {
			if visited[j] {
				continue
			}

			if minVertex == -1 || dist[j] < dist[minVertex] {
				minVertex = j
			}
		}

		if minVertex == -1 || dist[minVertex] == INF {
			break
		}

		visited[minVertex] = true
		for i := 1; i <= n; i++ {
			weight := matrix[minVertex][i]
			if visited[i] || weight < 0 {
				continue
			}

			if dist[i] > weight+dist[minVertex] {
				dist[i] = weight + dist[minVertex]
				parents[i] = minVertex
			}
		}
	}
	return parents
}

func makeOrder(parents []int, s int, f int) []int {
	var order []int
	order = append(order, f)
	pointer := parents[f]
	for i := 0; i < len(parents); i++ {
		if parents[pointer] == 0 {
			order = append(order, s)
			break
		}

		order = append(order, pointer)
		pointer = parents[pointer]
	}
	return order
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, s, f int
	fmt.Fscan(in, &n, &s, &f)

	matrix := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		matrix[i] = make([]int, n+1)

		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}

	parents := dijkstraParentList(matrix, n, s)

	if parents[f] == -1 {
		fmt.Print(parents[f])
		return
	}

	order := makeOrder(parents, s, f)

	for i := len(order) - 1; i >= 0; i-- {
		fmt.Print(order[i], " ")
	}
}
