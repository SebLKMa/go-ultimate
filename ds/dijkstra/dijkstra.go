package main

import (
	"container/heap"
	"fmt"
	"math"
	"reflect"
)

// Inf is the sentinel distance for unreachable nodes.
const Infinite = math.MaxInt

// Edge is a weighted directed edge.
type Edge struct {
	To     string
	Weight int
}

// Graph is a weighted directed graph backed by a map-based adjacency list.
// Node IDs are arbitrary integers — no pre-declaration of node count needed.
type Graph struct {
	adj map[string][]Edge
}

// New returns an empty Graph.
func New() *Graph {
	return &Graph{adj: make(map[string][]Edge)}
}

func (g *Graph) Print() {
	fmt.Println("Graph: map{src: [dst weight] [dst weight] ...}")
	fmt.Printf("%v\n", g)
}

// AddEdge adds a directed edge from → to with the given weight.
// Adding an edge automatically registers both endpoints in the graph.
func (g *Graph) AddEdge(from, to string, weight int) {
	g.adj[from] = append(g.adj[from], Edge{To: to, Weight: weight})
	// Adds 'to' if it does not exist, with no adjacent node yet.
	if _, ok := g.adj[to]; !ok {
		g.adj[to] = nil
	}
}

// AddUndirectedEdge adds edges in both directions.
func (g *Graph) AddUndirectedEdge(u, v string, weight int) {
	g.AddEdge(u, v, weight) // u to v with weight
	g.AddEdge(v, u, weight) // v to u with same weight
}

// item is a min-heap entry: (node, accumulated distance from source).
type item struct {
	node string
	dist int
}

// See: https://pkg.go.dev/container/heap
type minHeap []item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(item)) } // "cast" any to item
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1] // zero index
	*h = old[:n-1]
	return x
}

// See: https://www.youtube.com/watch?v=pVfj6mxhdMw
// Dijkstra returns the shortest distance from src to every reachable node.
// Unreachable nodes are absent from the returned map.
func (g *Graph) Dijkstra(src string) map[string]int {
	// dist holds the best known distance to each node seen so far.
	// Starting node costs 0 to reach itself.
	dist := map[string]int{src: 0}

	// Seed the min-heap with the starting node.
	h := &minHeap{{node: src, dist: 0}}
	heap.Init(h)

	for h.Len() > 0 {
		// Always process the node with the smallest known distance first.
		cur := heap.Pop(h).(item)

		// Skip stale heap entries: a shorter path to this node was already
		// settled, so this entry is outdated and can be ignored.
		if cur.dist > dist[cur.node] {
			continue
		}

		// Relax each outgoing edge from the current node.
		for _, e := range g.adj[cur.node] {
			d := dist[cur.node] + e.Weight
			// If this route is cheaper than anything seen before, record it
			// and push the neighbour onto the heap for future exploration.
			if best, seen := dist[e.To]; !seen || d < best {
				dist[e.To] = d
				heap.Push(h, item{node: e.To, dist: d})
			}
		}
	}
	return dist
}

// See: https://www.youtube.com/watch?v=pVfj6mxhdMw
// ShortestPath returns the distance and node sequence of the shortest path
// from src to dst. Returns Infinite and nil when dst is unreachable.
func (g *Graph) ShortestPath(src, dst string) (int, []string) {
	// Distances from source
	dist := map[string]int{src: 0}

	// Previous node that led to it on the cheapest route.
	// Used after the search finishes to reconstruct the full path.
	prev := map[string]string{}

	// Starts the heap with src item distance 0
	h := &minHeap{{node: src, dist: 0}}
	heap.Init(h)

	for h.Len() > 0 {

		// minHeap Pop returns the smallest distance.
		// Always process the node with the smallest known distance first.
		cur := heap.Pop(h).(item)
		fmt.Printf("cur: %v\n", cur)

		// Stale entry — a shorter path was already settled, skip.
		if cur.dist > dist[cur.node] {
			continue
		}

		// For each adjacent node, update distance from source if better
		for _, e := range g.adj[cur.node] {
			d := dist[cur.node] + e.Weight
			// if best, seen := dist[e.To]; !seen || d < best {
			best, visited := dist[e.To]
			if !visited || d < best {
				dist[e.To] = d
				// Remember that the cheapest way to reach e.To is via cur.node.
				prev[e.To] = cur.node
				heap.Push(h, item{node: e.To, dist: d})
			}
		}
		fmt.Printf("dist: %v\n", dist)
		fmt.Printf("prev: %v\n", prev)
		fmt.Printf("h.Len:%d h:%v\n", h.Len(), h)
	}
	fmt.Printf("prev: %v\n", prev)

	d, reachable := dist[dst]
	if !reachable {
		return Infinite, nil
	}

	// Trace backwards from dst to src using the prev breadcrumbs,
	// then reverse into a src→dst order.
	path := []string{}
	for at := dst; at != src; at = prev[at] {
		path = append([]string{at}, path...)
	}
	return d, append([]string{src}, path...)
}

// Edges: A-B(6), A-D(1), B-E(2), D-E(1), B-C(5), E-C(5)
// Shortest distances from 0: {0:0, 1:1, 2:3, 3:4, 4:6}
func buildGraph() *Graph {
	g := New()
	g.AddUndirectedEdge("A", "B", 6)
	g.AddUndirectedEdge("A", "D", 1)
	g.AddUndirectedEdge("B", "E", 2)
	g.AddUndirectedEdge("D", "E", 1)
	g.AddUndirectedEdge("B", "C", 5)
	g.AddUndirectedEdge("E", "C", 5)
	return g
}

func main() {
	g := buildGraph()
	g.Print()
	dist, path := g.ShortestPath("A", "C")
	expectedDist := 7
	fmt.Printf("Expected shortest distance = %d\n", expectedDist)
	fmt.Printf("Computed shortest distance = %d\n", dist)
	if dist != expectedDist {
		fmt.Println("Shortest distance wrong!")
	}
	expectedPath := []string{"A", "D", "E", "C"}
	fmt.Printf("Expected path: %v\n", expectedPath)
	fmt.Printf("Computed path: %v\n", path)
	if !reflect.DeepEqual(path, expectedPath) {
		fmt.Println("Shortest path wrong!")
	}
}
