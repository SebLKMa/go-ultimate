package graph

import (
	"container/heap"
	"fmt"
	"math"
)

// Inf is the sentinel distance for unreachable nodes.
const Infinite = math.MaxInt

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

func (g *Graph) Print() {
	fmt.Println("Graph: map{src: [dst weight] [dst weight] ...}")
	fmt.Printf("%v\n", g)
}

// AddEdge adds a directed edge from → to with the given weight.
// Adding an edge automatically registers both endpoints in the graph.
func (g *Graph) AddEdge(from, to, weight int) {
	g.adj[from] = append(g.adj[from], Edge{To: to, Weight: weight})
	// Adds 'to' if it does not exist, with no adjacent node yet.
	if _, ok := g.adj[to]; !ok {
		g.adj[to] = nil
	}
}

// AddUndirectedEdge adds edges in both directions.
func (g *Graph) AddUndirectedEdge(u, v, weight int) {
	g.AddEdge(u, v, weight) // u to v with weight
	g.AddEdge(v, u, weight) // v to u with same weight
}

// item is a min-heap entry: (node, accumulated distance from source).
type item struct {
	node, dist int
}

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

// Dijkstra returns the shortest distance from src to every reachable node.
// Unreachable nodes are absent from the returned map.
func (g *Graph) Dijkstra(src int) map[int]int {
	// dist holds the best known distance to each node seen so far.
	// Starting node costs 0 to reach itself.
	dist := map[int]int{src: 0}

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

// ShortestPath returns the distance and node sequence of the shortest path
// from src to dst. Returns Infinite and nil when dst is unreachable.
func (g *Graph) ShortestPath(src, dst int) (int, []int) {
	dist := map[int]int{src: 0}
	// prev records, for each node, which node led to it on the cheapest route.
	// Used after the search finishes to reconstruct the full path.
	prev := map[int]int{}

	h := &minHeap{{node: src, dist: 0}}
	heap.Init(h)

	for h.Len() > 0 {
		fmt.Printf("h.Len: %d\n", h.Len())
		// Always process the node with the smallest known distance first.
		cur := heap.Pop(h).(item)
		fmt.Printf("cur: %v\n", cur)

		// Stale entry — a shorter path was already settled, skip.
		if cur.dist > dist[cur.node] {
			continue
		}

		// Relax each outgoing edge from the current node.
		for _, e := range g.adj[cur.node] {
			d := dist[cur.node] + e.Weight
			if best, seen := dist[e.To]; !seen || d < best {
				dist[e.To] = d
				// Remember that the cheapest way to reach e.To is via cur.node.
				prev[e.To] = cur.node
				heap.Push(h, item{node: e.To, dist: d})
			}
		}
		fmt.Printf("dist: %v\n", dist)
	}
	fmt.Printf("prev: %v\n", prev)

	d, reachable := dist[dst]
	if !reachable {
		return Infinite, nil
	}

	// Trace backwards from dst to src using the prev breadcrumbs,
	// then reverse into a src→dst order.
	path := []int{}
	for at := dst; at != src; at = prev[at] {
		path = append([]int{at}, path...)
	}
	return d, append([]int{src}, path...)
}
