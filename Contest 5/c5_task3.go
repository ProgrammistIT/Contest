package main

import (
	"bufio"
	"fmt"
	"os"
)

func heapUp(heap []int) {
	for i := len(heap) - 1; ; {
		if i == 0 {
			return
		}
		p := int((i - 1) / 2)

		if heap[p] < heap[i] {
			heap[p], heap[i] = heap[i], heap[p]
			i = p
		} else {
			return
		}
	}
}

func maxChildIndex(heap []int, index int) int {
	if index+1 == len(heap) || heap[index] > heap[index+1] {
		return index
	}
	return index + 1
}

func heapDown(heap []int) {
	for parent := 0; ; {
		index := (parent * 2) + 1
		if index >= len(heap) {
			return
		}

		index = maxChildIndex(heap, index)

		if heap[index] > heap[parent] {
			heap[index], heap[parent] = heap[parent], heap[index]
			parent = index
			continue
		}
		return
	}
}

func Insert(heap *[]int, value int) {
	*heap = append(*heap, value)
	heapUp(*heap)
}

func Extract(heap *[]int) {
	last := (*heap)[len(*heap)-1]
	last, (*heap)[0] = (*heap)[0], last
	fmt.Println(last)
	(*heap) = (*heap)[0 : len(*heap)-1]
	heapDown(*heap)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	var heap []int
	for i := 0; i < n; i++ {
		var command int
		fmt.Fscan(in, &command)
		if command == 0 {
			var value int
			fmt.Fscan(in, &value)
			Insert(&heap, value)
		} else {
			Extract(&heap)
		}
	}
}
