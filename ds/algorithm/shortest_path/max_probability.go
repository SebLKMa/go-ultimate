package main

import "fmt"

// https://leetcode.com/problems/path-with-maximum-probability/solutions/5696030/10000easy-solutionwith-explanation-by-mr-bwgt/?envType=problem-list-v2&envId=shortest-path
/*
Find the path between two nodes in an undirected graph that maximizes the product of edge probabilities (instead distances on edges).
The Bellman-Ford algorithm, which is typically used to find the shortest paths in graphs with negative weights,
can be adapted to solve this problem.
Instead of minimizing distances, we will maximize probabilities by updating the maximum probability to reach each node iteratively.

Approach 🚀
1️⃣Initialize an array dist where dist[i] holds the maximum probability to reach node i from the start node.
Set dist[start] = 1 since the probability of starting at the start node is 1.
2️⃣Perform up to n-1 iterations, where n is the number of nodes.
In each iteration, check each edge and update the probability of reaching the neighboring nodes.
3️⃣For each edge (u, v), if the probability of reaching v through u (i.e., dist[u] * succProb[i]) is greater than
the current known probability to reach v (dist[v]), update dist[v].
Similarly, update dist[u] if the probability of reaching u through v is greater.
4️⃣After completing the iterations, dist[end] will contain the maximum probability of reaching the end node from the start node. If there's no path, it will remain 0.

Time complexity:⏲️
The algorithm runs in O(n×E), where n is the number of nodes and E is the number of edges.
This is because we perform n-1 iterations over all the edges.
Space complexity:🛰️
The space complexity is O(n) since we are using an array dist of size n to store the maximum probability for each node.
*/

func MaxProbability(n int, edges [][]int, edges_probabilities []float64, start_node int, end_node int) float64 {
	max_probabilities := make([]float64, n)
	max_probabilities[start_node] = 1.0

	for i := 0; i < n-1; i++ {
		updated := false
		for j := 0; j < len(edges); j++ {
			u := edges[j][0]
			v := edges[j][1]
			probability := edges_probabilities[j]

			if max_probabilities[u]*probability > max_probabilities[v] {
				max_probabilities[v] = max_probabilities[u] * probability
				updated = true
			}
			if max_probabilities[v]*probability > max_probabilities[u] {
				max_probabilities[u] = max_probabilities[v] * probability
				updated = true
			}
			if !updated {
				break
			}
		}
	}

	return max_probabilities[end_node]
}

func TestMaxProbability() {
	/*
		Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2
		Output: 0.25000
		Explanation: There are two paths from start to end, one having a probability of success = 0.2 and
		the other has 0.5 * 0.5 = 0.25.
	*/
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}        // array of u,v
	edges_probabilities := []float64{0.5, 0.5, 0.2} // corresponding u,v probabilities
	start_node := 0
	end_node := 2
	max_probability := MaxProbability(3, edges, edges_probabilities, start_node, end_node)
	fmt.Println("\nTest case 1")
	fmt.Printf("From node %d to node %d, the highest probability is %f\n", start_node, end_node, max_probability)

	/*
		Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2
		Output: 0.30000
	*/
	edges = [][]int{{0, 1}, {1, 2}, {0, 2}}        // array of u,v
	edges_probabilities = []float64{0.5, 0.5, 0.3} // corresponding u,v probabilities
	start_node = 0
	end_node = 2
	max_probability = MaxProbability(3, edges, edges_probabilities, start_node, end_node)
	fmt.Println("\nTest case 2a")
	fmt.Printf("From node %d to node %d, the highest probability is %f\n", start_node, end_node, max_probability)

	/*
		Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.8,0.5,0.3], start = 0, end = 2
		Output: 0.40000
	*/
	edges = [][]int{{0, 1}, {1, 2}, {0, 2}}        // array of u,v
	edges_probabilities = []float64{0.8, 0.5, 0.3} // 0,1 probability chabged to higher probability of 0.8
	start_node = 0
	end_node = 2
	max_probability = MaxProbability(3, edges, edges_probabilities, start_node, end_node)
	fmt.Println("\nTest case 2b")
	fmt.Printf("From node %d to node %d, the highest probability is %f\n", start_node, end_node, max_probability)

	/*
		Input: n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2
		Output: 0.00000
		Explanation: There is no path between 0 and 2.
	*/
	edges = [][]int{{0, 1}}              // array of u,v
	edges_probabilities = []float64{0.5} // corresponding u,v probabilities
	start_node = 0
	end_node = 2
	max_probability = MaxProbability(3, edges, edges_probabilities, start_node, end_node)
	fmt.Println("\nTest case 3")
	fmt.Printf("From node %d to node %d, the highest probability is %f\n", start_node, end_node, max_probability)
}

// https://leetcode.com/problems/find-edges-in-shortest-paths/description/
/*
	TODO
*/

func main() {
	TestMaxProbability()
}
