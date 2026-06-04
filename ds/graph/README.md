# Graph — Dijkstra's Shortest Path

These youtube videos have good visual explanations.  
https://youtu.be/j0OUwduDOS0?si=iPaI4WlMOVpjgMdG  
https://www.youtube.com/watch?v=pVfj6mxhdMw  
https://www.youtube.com/watch?v=bZkzH5x0SKU  
https://www.youtube.com/watch?v=EFg3u_E6eHU  
https://www.youtube.com/watch?v=CmIQ29cUGiE  


## What is a graph?

Think of a graph as a map of cities connected by roads. Each city is a **node** and each road is an **edge**. Roads can have a **weight** — the distance or cost to travel along them.

This package models that kind of map and answers two questions:
- How far is every city from a given starting city?
- What is the exact route to get from city A to city B as cheaply as possible?

## How the graph is stored (`graph.go`)

Each node is just an integer ID (0, 1, 2, …). The graph keeps a lookup table (`map[int][]Edge`) where, for any node, you can instantly get the list of roads leaving it and their costs.

```
Node 0 → [{To:1, Weight:1}, {To:2, Weight:4}]
Node 1 → [{To:0, Weight:1}, {To:2, Weight:2}, ...]
...
```

You don't need to declare how many nodes exist upfront — just start adding edges and both endpoints are registered automatically.

**Two ways to add a road:**
- `AddEdge(from, to, weight)` — one-way road (directed)
- `AddUndirectedEdge(u, v, weight)` — two-way road (calls AddEdge in both directions)

## How Dijkstra works (plain English)

Dijkstra's algorithm finds the shortest path the same way a cautious traveller would:

1. Start at the source city. Your distance to yourself is 0; everyone else is unknown.
2. Always visit the closest unvisited city next (this is where the **min-heap** helps — it's like a priority queue that always hands you the nearest city).
3. From that city, check all its roads. If going through this city gets you to a neighbour cheaper than any previously known route, update the neighbour's best-known distance.
4. Repeat until no cities remain in the queue.

The **min-heap** is the engine that makes step 2 fast. Without it you'd scan every city each time to find the nearest one; with it, the nearest city is always at the top.

**Stale entry skip:** When a shorter path to a city is discovered, the old (worse) entry is left in the heap rather than removed — removing from the middle of a heap is expensive. Instead, when that stale entry eventually pops out, we check: "is this distance worse than what we already know?" If yes, throw it away and move on.

## The two functions

### `Dijkstra(src) map[int]int`

Returns the shortest distance from `src` to every reachable node. Nodes that cannot be reached simply don't appear in the result map.

```
g.Dijkstra(0)  →  {0:0, 1:1, 2:3, 3:4, 4:6}
```

### `ShortestPath(src, dst) (distance, []int)`

Returns the total cost and the actual sequence of nodes to walk. If `dst` is unreachable, returns `Inf` and `nil`.

```
g.ShortestPath(0, 4)  →  6, [0, 1, 2, 3, 4]
```

Path reconstruction works by remembering, for each node, which node led to it with the best distance (`prev` map). After the algorithm finishes, we trace backwards from `dst` to `src` using those breadcrumbs.

## The test graph (`graph_test.go`)

Most tests use this 5-node undirected graph:

```
0 --1-- 1
|      /|
4    2/ 5
|   /   |
2 --1-- 3 --2-- 4
```

Numbers on edges are weights (travel costs). The `/` diagonal is the edge between nodes 1 and 2 (weight 2). The shortest path from 0 to 4 is **0→1→2→3→4** with total cost **6** — the direct road from 0 to 2 (weight 4) is skipped because going via node 1 first (0→1 costs 1, then 1→2 costs 2) reaches node 2 for only 3.

## What each test checks

| Test | What it verifies |
|------|-----------------|
| `TestDijkstraAllDistances` | All distances from node 0 are correct |
| `TestDijkstraFromMiddleNode` | Starting from a middle node (2) still gives correct distances |
| `TestShortestPath` | The path 0→4 has cost 6 and follows nodes [0,1,2,3,4] |
| `TestShortestPathSameNode` | Distance from a node to itself is 0, path is just [node] |
| `TestShortestPathUnreachable` | A one-way road that blocks return gives Inf and nil path |
| `TestDirectedGraph` | On a directed graph, the algorithm picks the cheaper multi-hop route over a direct but expensive road |

## Run the tests

```sh
GOWORK=off go test -v ./...
```

`GOWORK=off` is needed because this module has its own `go.mod` and is not part of the workspace.

## Summary

`graph.go` — the core implementation:  
- Graph struct backed by map[int][]Edge rather than the slice-of-slices approach in ds/heap/dijkstra/. This means node IDs
can be arbitrary integers with no need to declare a node count upfront, and AddEdge automatically registers both endpoints.  
- Dijkstra(src) returns map[int]int — only reachable nodes appear in the result (absent = unreachable), which fits the
map-based graph naturally.  
- ShortestPath(src, dst) returns distance + reconstructed path, using the same lazy-deletion pattern (skip stale heap entries
where cur.dist > dist[cur.node]).  

`graph_test.go` — 6 tests covering: all distances from source, distances from a middle node, path reconstruction, src→src zero
distance, unreachable node, and a directed graph where the shorter path routes through an intermediate node.  

Run with: GOWORK=off go test -v ./... from inside ds/graph/.  
