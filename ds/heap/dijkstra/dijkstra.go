package dijkstra

import (
	"container/heap"
	"math"
)

// Inf is the sentinel distance for unreachable nodes.
const Inf = math.MaxInt

// Edge is a directed weighted edge to a neighbour node.
type Edge struct {
	To     int
	Weight int
}

// Graph is an adjacency-list weighted directed graph.
type Graph [][]Edge

// NewGraph creates a Graph with n nodes (0 … n-1).
func NewGraph(n int) Graph {
	return make(Graph, n)
}

// AddEdge adds a directed edge u → v.
func (g Graph) AddEdge(u, v, weight int) {
	g[u] = append(g[u], Edge{To: v, Weight: weight})
}

// AddUndirectedEdge adds edges in both directions.
func (g Graph) AddUndirectedEdge(u, v, weight int) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// item is a min-heap entry: (node, accumulated distance from source).
type item struct {
	node int
	dist int
}

type minHeap []item

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) { *h = append(*h, x.(item)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Dijkstra returns the shortest distance from src to every node in g.
// Unreachable nodes have distance Inf.
func Dijkstra(g Graph, src int) []int {
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = Inf
	}
	dist[src] = 0

	h := &minHeap{{node: src, dist: 0}}
	heap.Init(h)

	for h.Len() > 0 {
		cur := heap.Pop(h).(item)
		// Stale entry — a shorter path was already settled.
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range g[cur.node] {
			if d := dist[cur.node] + e.Weight; d < dist[e.To] {
				dist[e.To] = d
				heap.Push(h, item{node: e.To, dist: d})
			}
		}
	}
	return dist
}

// ShortestPath returns the shortest distance from src to dst and the node
// sequence of that path. Returns Inf and nil when dst is unreachable.
func ShortestPath(g Graph, src, dst int) (int, []int) {
	n := len(g)
	dist := make([]int, n)
	prev := make([]int, n)
	for i := range dist {
		dist[i] = Inf
		prev[i] = -1
	}
	dist[src] = 0

	h := &minHeap{{node: src, dist: 0}}
	heap.Init(h)

	for h.Len() > 0 {
		cur := heap.Pop(h).(item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range g[cur.node] {
			if d := dist[cur.node] + e.Weight; d < dist[e.To] {
				dist[e.To] = d
				prev[e.To] = cur.node
				heap.Push(h, item{node: e.To, dist: d})
			}
		}
	}

	if dist[dst] == Inf {
		return Inf, nil
	}

	path := []int{}
	for at := dst; at != -1; at = prev[at] {
		path = append([]int{at}, path...)
	}
	return dist[dst], path
}
