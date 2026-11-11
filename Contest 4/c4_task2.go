package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const INF = math.MaxInt

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

	visited := make([]bool, n+1)
	dist := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dist[i] = INF
	}
	dist[s] = 0

	for step := 0; step < n; step++ {
		minVertex := -1

		for i := 1; i <= n; i++ {
			if visited[i] {
				continue
			}

			if minVertex == -1 || dist[i] < dist[minVertex] {
				minVertex = i
			}
		}

		if minVertex == -1 || dist[minVertex] == INF {
			break
		}

		visited[minVertex] = true
		for j := 1; j <= n; j++ {
			weight := matrix[minVertex][j]

			if visited[j] || weight < 0 {
				continue
			}

			if dist[minVertex]+weight < dist[j] {
				dist[j] = dist[minVertex] + weight
			}
		}
	}

	if dist[f] == INF {
		fmt.Print(-1)
	} else {
		fmt.Print(dist[f])
	}
}
