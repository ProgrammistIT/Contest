package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const INF = math.MaxFloat64

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, s, e int
	fmt.Fscan(in, &n, &m, &s, &e)

	matrix := make([][]float64, n+1)

	for i := 0; i <= n; i++ {
		matrix[i] = make([]float64, n+1)
		for j := range matrix[i] {
			matrix[i][j] = INF
		}

		matrix[i][i] = 0
	}

	for i := 0; i < m; i++ {
		var from, to, weight int
		fmt.Fscan(in, &from, &to, &weight)
		matrix[from][to] = float64(weight) / 100
		matrix[to][from] = float64(weight) / 100
	}

	for k := 1; k <= n; k++ {
		for from := 1; from <= n; from++ {
			for to := 1; to <= n; to++ {
				if matrix[from][k] == INF || matrix[k][to] == INF {
					continue
				}

				robChance := 1 - (1-matrix[from][k])*(1-matrix[k][to])
				if robChance < matrix[from][to] {
					matrix[from][to] = robChance
				}
			}
		}
	}

	fmt.Print(matrix[s][e])
}
