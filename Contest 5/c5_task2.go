package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Edge struct {
	from, to, weight int
}

const INF = math.MaxInt32

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k, s, f int
	fmt.Fscan(in, &n, &m, &k, &s, &f)

	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &edges[i].from, &edges[i].to, &edges[i].weight)
	}

	dist := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dist[i] = INF
	}
	dist[s] = 0

	for i := 1; i <= k; i++ {
		prev := make([]int, n+1)
		copy(prev, dist)

		for _, edge := range edges {
			from := edge.from
			to := edge.to
			weight := edge.weight

			if prev[from] != INF && prev[from]+weight < dist[to] {
				dist[to] = prev[from] + weight
			}
		}
	}

	if dist[f] == INF {
		fmt.Print(-1)
	} else {
		fmt.Print(dist[f])
	}
}
