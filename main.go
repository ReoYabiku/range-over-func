package main

import (
	"iter"
)

func NormalTopologicalSort(graph [][]int) []int {
	inputs := make([]int, len(graph))
	for _, nxts := range graph {
		for _, nxt := range nxts {
			inputs[nxt]++
		}
	}

	var stack []int
	for i, input := range inputs {
		if input == 0 {
			stack = append(stack, i)
		}
	}

	var sorted []int
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		sorted = append(sorted, top)

		for _, nxt := range graph[top] {
			inputs[nxt]--
			if inputs[nxt] == 0 {
				stack = append(stack, nxt)
			}
		}
	}

	return sorted
}

func IterTopologicalSort(graph [][]int) iter.Seq[int] {
	inputs := make([]int, len(graph))
	for _, nxts := range graph {
		for _, nxt := range nxts {
			inputs[nxt]++
		}
	}

	var stack []int
	for i, input := range inputs {
		if input == 0 {
			stack = append(stack, i)
		}
	}

	return func(yield func(int) bool) {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if !yield(top) {
				return
			}

			for _, nxt := range graph[top] {
				inputs[nxt]--
				if inputs[nxt] == 0 {
					stack = append(stack, nxt)
				}
			}
		}
	}
}
