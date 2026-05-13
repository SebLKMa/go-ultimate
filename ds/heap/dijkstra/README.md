# Shortest Path

  dijkstra.go — two exported functions:  
  - Dijkstra(g, src) — returns shortest distance from src to all nodes  
  - ShortestPath(g, src, dst) — returns distance + reconstructed node path  
  
  The min-heap holds (node, dist) items ordered by dist. Stale heap entries (where a shorter path was already settled) are skipped with a cur.dist > dist[cur.node] guard — the standard lazy-deletion trick that avoids the need for a DecreaseKey operation.

  dijkstra_test.go — 5 tests covering: all-distances from source, distances from a middle node, path reconstruction, unreachable node, and zero-distance self-path.

  Run with: `GOWORK=off go test -v ./... from inside ds/heap/dijkstra/`.  

## References
https://dev.to/douglasmakey/implementation-of-dijkstra-using-heap-in-go-6e3 
https://leetcode.com/discuss/post/6125318/shortest-path-in-weighted-undirected-gra-hpoa/  