package graph

import (
	"container/heap"
	"math"
)

// Inf is the sentinel distance for unreachable nodes.
const Inf = math.MaxInt

// Edge is a weighted directed edge.
type Edge struct {
	To     int
	Weight int
}

// Graph is a weighted directed graph backed by a map-based adjacency list.
// Node IDs are arbitrary integers — no pre-declaration of node count needed.
type Graph struct {
	adj map[int][]Edge
}

// New returns an empty Graph.
func New() *Graph {
	return &Graph{adj: make(map[int][]Edge)}
}

// AddEdge adds a directed edge from → to with the given weight.
// Adding an edge automatically registers both endpoints in the graph.
func (g *Graph) AddEdge(from, to, weight int) {
	g.adj[from] = append(g.adj[from], Edge{To: to, Weight: weight})
	if _, ok := g.adj[to]; !ok {
		g.adj[to] = nil
	}
}

// AddUndirectedEdge adds edges in both directions.
func (g *Graph) AddUndirectedEdge(u, v, weight int) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// item is a min-heap entry: (node, accumulated distance from source).
type item struct {
	node, dist int
}

type minHeap []item

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(item)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Dijkstra returns the shortest distance from src to every reachable node.
// Unreachable nodes are absent from the returned map.
func (g *Graph) Dijkstra(src int) map[int]int {
	dist := map[int]int{src: 0}
	h := &minHeap{{node: src, dist: 0}}
	heap.Init(h)

	for h.Len() > 0 {
		cur := heap.Pop(h).(item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range g.adj[cur.node] {
			d := dist[cur.node] + e.Weight
			if best, seen := dist[e.To]; !seen || d < best {
				dist[e.To] = d
				heap.Push(h, item{node: e.To, dist: d})
			}
		}
	}
	return dist
}

// ShortestPath returns the distance and node sequence of the shortest path
// from src to dst. Returns Inf and nil when dst is unreachable.
func (g *Graph) ShortestPath(src, dst int) (int, []int) {
	dist := map[int]int{src: 0}
	prev := map[int]int{}
	h := &minHeap{{node: src, dist: 0}}
	heap.Init(h)

	for h.Len() > 0 {
		cur := heap.Pop(h).(item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range g.adj[cur.node] {
			d := dist[cur.node] + e.Weight
			if best, seen := dist[e.To]; !seen || d < best {
				dist[e.To] = d
				prev[e.To] = cur.node
				heap.Push(h, item{node: e.To, dist: d})
			}
		}
	}

	d, reachable := dist[dst]
	if !reachable {
		return Inf, nil
	}

	path := []int{}
	for at := dst; at != src; at = prev[at] {
		path = append([]int{at}, path...)
	}
	return d, append([]int{src}, path...)
}
