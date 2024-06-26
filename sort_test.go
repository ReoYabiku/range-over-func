package main

import (
	"fmt"
	"io"
	"testing"
)

var graph [][]int

func TestMain(m *testing.M) {
	graph = [][]int{
		{},
		{6},
		{5},
		{1, 2},
		{1, 6},
		{1},
		{7},
		{8},
		{},
	}

	m.Run()
}

func BenchmarkNormalTopSort(b *testing.B) {
	for range b.N {
		for v := range NormalTopologicalSort(graph) {
			fmt.Fprintf(io.Discard, "%d\n", v)
		}
	}
}

func BenchmarkTopSort(b *testing.B) {
	for range b.N {
		for v := range IterTopologicalSort(graph) {
			fmt.Fprintf(io.Discard, "%d\n", v)
		}
	}
}
