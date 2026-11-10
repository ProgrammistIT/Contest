package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func CountingSort(array []int) []int {

	if len(array) == 0 {
		return array
	}
	min := math.MaxInt
	max := math.MinInt
	for _, value := range array {
		// search min value
		if min >= value {
			min = value
		}
		// search max value
		if max <= value {
			max = value
		}
	}

	temporary := make([]int, (max - min + 1))
	for _, v := range array {
		temporary[v-min]++
	}

	idx := 0
	for i, count := range temporary {
		for ; count > 0; count-- {
			array[idx] = i + min
			idx++
		}
	}
	return array
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &array[i])
	}

	array = CountingSort(array)
	for i, value := range array {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, value)
	}
	fmt.Fprintln(writer)
}
