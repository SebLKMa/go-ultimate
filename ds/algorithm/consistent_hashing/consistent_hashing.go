package main

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"sort"
)

// ConsistentHashing represents the hash ring structure
type ConsistentHashing struct {
	numReplicas int
	ring        map[string]string // Maps the string representation of big.Int hash to server string
	sortedKeys  []*big.Int        // Sorted slice of server hashes
	servers     map[string]bool   // Set keeping track of unique active servers
}

// NewConsistentHashing initializes the hash ring with a list of servers
func NewConsistentHashing(servers []string, numReplicas int) *ConsistentHashing {
	ch := &ConsistentHashing{
		numReplicas: numReplicas,
		ring:        make(map[string]string),
		sortedKeys:  make([]*big.Int, 0),
		servers:     make(map[string]bool),
	}
	for _, server := range servers {
		ch.AddServer(server)
	}
	return ch
}

// hash is the MD5 internal function that returns a 128-bit big.Int
func (ch *ConsistentHashing) hash(key string) *big.Int {
	h := md5.Sum([]byte(key))
	return new(big.Int).SetBytes(h[:])
}

// AddServer adds a server and its virtual replicas to the hash ring
func (ch *ConsistentHashing) AddServer(server string) {
	fmt.Println("Adding servers...")
	ch.servers[server] = true
	for i := 0; i < ch.numReplicas; i++ {
		hashVal := ch.hash(fmt.Sprintf("%s-%d", server, i))
		ch.ring[hashVal.String()] = server
		fmt.Printf("hash:%s -> server:%s\n", hashVal.String(), server)

		// Emulate Python's bisect.insort (bisect_right behavior)
		idx := sort.Search(len(ch.sortedKeys), func(j int) bool {
			return ch.sortedKeys[j].Cmp(hashVal) > 0
		})

		// Insert hashVal into the sorted slice at the found index
		ch.sortedKeys = append(ch.sortedKeys, nil)
		copy(ch.sortedKeys[idx+1:], ch.sortedKeys[idx:])
		ch.sortedKeys[idx] = hashVal
	}
}

// RemoveServer removes a server and all its replicas from the hash ring
func (ch *ConsistentHashing) RemoveServer(server string) {
	if _, exists := ch.servers[server]; exists {
		delete(ch.servers, server)
		for i := 0; i < ch.numReplicas; i++ {
			hashVal := ch.hash(fmt.Sprintf("%s-%d", server, i))
			delete(ch.ring, hashVal.String())

			// Find exact matching hash position using binary search
			idx := sort.Search(len(ch.sortedKeys), func(j int) bool {
				return ch.sortedKeys[j].Cmp(hashVal) >= 0
			})
			if idx < len(ch.sortedKeys) && ch.sortedKeys[idx].Cmp(hashVal) == 0 {
				// Remove the element from the slice
				ch.sortedKeys = append(ch.sortedKeys[:idx], ch.sortedKeys[idx+1:]...)
			}
		}
	}
}

// GetServer finds the closest server on the ring for the given key
func (ch *ConsistentHashing) GetServer(key string) string {
	if len(ch.ring) == 0 {
		return ""
	}
	hashVal := ch.hash(key)
	fmt.Printf("key hash:%s\n", hashVal.String())

	// Emulate Python's bisect.bisect (bisect_right behavior)
	idx := sort.Search(len(ch.sortedKeys), func(j int) bool {
		return ch.sortedKeys[j].Cmp(hashVal) > 0
	})

	// Wrap around to the beginning of the ring if necessary
	idx = idx % len(ch.sortedKeys)
	return ch.ring[ch.sortedKeys[idx].String()]
}

func main() {
	// Initialize with servers
	servers := []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	ch := NewConsistentHashing(servers, 3)

	testKeys := []string{"user_101", "user_202", "request_abc", "data_99"}

	fmt.Println("Initial Lookups:")
	for _, key := range testKeys {
		fmt.Printf("  %s -> %s\n", key, ch.GetServer(key))
	}

	// Remove a server and verify redistributed keys
	fmt.Println("\nAfter removing S2:")
	ch.RemoveServer("S2")
	for _, key := range testKeys {
		fmt.Printf("  %s -> %s\n", key, ch.GetServer(key))
	}

	fmt.Println("\nTest matching keys to servers again...")
	testKeysAgain := []string{"user_202", "user_505"}
	for _, key := range testKeysAgain {
		fmt.Printf("  %s -> %s\n", key, ch.GetServer(key))
	}

}
