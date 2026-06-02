package graph_test

import (
	"reflect"
	"testing"

	"github.com/seblkma/go-ultimate/ds/graph"
)

// Graph used in most tests (undirected, 5 nodes):
//
//	0 --1-- 1
//	|      /|
//	4    2/ 5
//	|   /   |
//	2 --1-- 3 --2-- 4
//
// Edges: 0-1(1), 0-2(4), 1-2(2), 1-3(5), 2-3(1), 3-4(2)
// Shortest distances from 0: {0:0, 1:1, 2:3, 3:4, 4:6}
func buildGraph() *graph.Graph {
	g := graph.New()
	g.AddUndirectedEdge(0, 1, 1)
	g.AddUndirectedEdge(0, 2, 4)
	g.AddUndirectedEdge(1, 2, 2)
	g.AddUndirectedEdge(1, 3, 5)
	g.AddUndirectedEdge(2, 3, 1)
	g.AddUndirectedEdge(3, 4, 2)
	return g
}

func TestDijkstraAllDistances(t *testing.T) {
	g := buildGraph()
	got := g.Dijkstra(0)
	want := map[int]int{0: 0, 1: 1, 2: 3, 3: 4, 4: 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Dijkstra(src=0) = %v, want %v", got, want)
	}
}

func TestDijkstraFromMiddleNode(t *testing.T) {
	g := buildGraph()
	got := g.Dijkstra(2)
	want := map[int]int{0: 3, 1: 2, 2: 0, 3: 1, 4: 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Dijkstra(src=2) = %v, want %v", got, want)
	}
}

// To run just this test,
// GOWORK=off go test -v -run TestShortestPath
func TestShortestPathDebug(t *testing.T) {
	g := buildGraph()
	g.Print()
	dist, path := g.ShortestPath(0, 4)
	if dist != 6 {
		t.Errorf("ShortestPath distance = %d, want 6", dist)
	}
	want := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(path, want) {
		t.Errorf("ShortestPath path = %v, want %v", path, want)
	}
}

func TestShortestPathSameNode(t *testing.T) {
	g := buildGraph()
	dist, path := g.ShortestPath(2, 2)
	if dist != 0 {
		t.Errorf("distance src→src = %d, want 0", dist)
	}
	if !reflect.DeepEqual(path, []int{2}) {
		t.Errorf("path src→src = %v, want [2]", path)
	}
}

func TestShortestPathUnreachable(t *testing.T) {
	// Directed graph: 0 → 1, but not 1 → 0.
	g := graph.New()
	g.AddEdge(0, 1, 5)
	dist, path := g.ShortestPath(1, 0)
	if dist != graph.Infinite {
		t.Errorf("expected Inf for unreachable node, got %d", dist)
	}
	if path != nil {
		t.Errorf("expected nil path for unreachable node, got %v", path)
	}
}

func TestDirectedGraph(t *testing.T) {
	// Directed: 0→1(1), 0→2(10), 1→2(1) — shortest 0→2 goes through 1.
	g := graph.New()
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 10)
	g.AddEdge(1, 2, 1)

	dist, path := g.ShortestPath(0, 2)
	if dist != 2 {
		t.Errorf("distance = %d, want 2", dist)
	}
	want := []int{0, 1, 2}
	if !reflect.DeepEqual(path, want) {
		t.Errorf("path = %v, want %v", path, want)
	}
}
