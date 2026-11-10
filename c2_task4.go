package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	inD := make(map[int]int)
	outD := make(map[int]int)

	for i := 0; i < m; i++ {
		var first, second int
		fmt.Fscan(reader, &first, &second)

		outD[first]++
		inD[second]++
	}

	for v := 1; v <= n; v++ {
		fmt.Fprintln(writer, inD[v])
		fmt.Fprintln(writer, outD[v])
	}
}
