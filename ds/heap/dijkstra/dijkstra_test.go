package dijkstra_test

import (
	"reflect"
	"testing"

	"github.com/seblkma/go-ultimate/ds/heap/dijkstra"
)

// Graph used in all tests (undirected):
//
//	0 --1-- 1
//	|       |
//	4       2
//	|       |
//	2 --5-- 3 --1-- 4  (wait, see edges below)
//
// Edges: 0-1(1), 0-2(4), 1-2(2), 1-3(5), 2-3(1)
// Shortest distances from 0: {0:0, 1:1, 2:3, 3:4}
// Shortest path 0→3: 0→1→2→3
func buildGraph() dijkstra.Graph {
	g := dijkstra.NewGraph(4)
	g.AddUndirectedEdge(0, 1, 1)
	g.AddUndirectedEdge(0, 2, 4)
	g.AddUndirectedEdge(1, 2, 2)
	g.AddUndirectedEdge(1, 3, 5)
	g.AddUndirectedEdge(2, 3, 1)
	return g
}

func TestDijkstraDistances(t *testing.T) {
	g := buildGraph()
	got := dijkstra.Dijkstra(g, 0)
	want := []int{0, 1, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Dijkstra(src=0) = %v, want %v", got, want)
	}
}

func TestDijkstraFromMiddleNode(t *testing.T) {
	g := buildGraph()
	got := dijkstra.Dijkstra(g, 2)
	want := []int{3, 2, 0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Dijkstra(src=2) = %v, want %v", got, want)
	}
}

func TestShortestPath(t *testing.T) {
	g := buildGraph()
	dist, path := dijkstra.ShortestPath(g, 0, 3)
	if dist != 4 {
		t.Errorf("ShortestPath distance = %d, want 4", dist)
	}
	want := []int{0, 1, 2, 3}
	if !reflect.DeepEqual(path, want) {
		t.Errorf("ShortestPath path = %v, want %v", path, want)
	}
}

func TestShortestPathUnreachable(t *testing.T) {
	// Directed graph where node 1 is isolated (no incoming edges from 0).
	g := dijkstra.NewGraph(2)
	g.AddEdge(1, 0, 1) // only 1→0
	dist, path := dijkstra.ShortestPath(g, 0, 1)
	if dist != dijkstra.Inf {
		t.Errorf("expected Inf for unreachable node, got %d", dist)
	}
	if path != nil {
		t.Errorf("expected nil path for unreachable node, got %v", path)
	}
}

func TestShortestPathSameNode(t *testing.T) {
	g := buildGraph()
	dist, path := dijkstra.ShortestPath(g, 2, 2)
	if dist != 0 {
		t.Errorf("distance src→src = %d, want 0", dist)
	}
	if !reflect.DeepEqual(path, []int{2}) {
		t.Errorf("path src→src = %v, want [2]", path)
	}
}
