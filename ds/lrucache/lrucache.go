package main

import (
	"container/list"
	"fmt"
)

// Least Recently Used (LRU) cache
/*
A Least Recently Used (LRU) cache, when implemented using a hash map,
leverages the strengths of both data structures to achieve efficient caching.
This combination allows for fast lookups and updates while effectively managing cache eviction based on usage frequency.

Core Mechanism:

Hash Map (or Dictionary/Map):
The hash map stores the actual key-value pairs of the cached data.
Its primary role is to provide O(1) average time complexity for get (lookup) and put (insertion/update) operations,
allowing quick access to cached items based on their keys.

Doubly Linked List:
A separate doubly linked list is used to maintain the order of data items based on their recency of use.
The head of the linked list represents the most recently used item.
The tail of the linked list represents the least recently used item.

How it Works:
get(key) operation:
The hash map is used to quickly locate the corresponding node in the doubly linked list.
If the item is found, it is considered "recently used" and is moved to the head of the linked list.
The value associated with the key is returned.

put(key, value) operation:
If the key already exists in the cache, its value is updated, and the corresponding node in the linked list is moved to the head.
If the key does not exist and the cache is not yet full, a new node is created, added to the head of the linked list,
and the key-value pair is added to the hash map.
If the key does not exist and the cache is full, the least recently used item (the node at the tail of the linked list)
is removed from both the linked list and the hash map to make space.
The new item is then added to the head of the linked list and to the hash map.

Benefits:
O(1) average time complexity: for get and put operations due to the hash map.
Efficient eviction: of least recently used items using the doubly linked list, which allows for O(1) removal from the tail.
*/

type Node struct {
	Data interface{} // Can hold any data type
}

func simualateLRUcache() {
	maxCacheSix := 2

	m := make(map[int]Node, maxCacheSix)

	// Create a doublylinkedlist and put nodes in it.
	l := list.New()
	value := 1
	n1 := Node{Data: value}
	l.PushBack(n1)
	m[value] = n1
	value = 2
	n2 := Node{Data: value}
	l.PushBack(n2)
	m[value] = n2

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Printf("%v\n", m)

	if l.Len() >= maxCacheSix {
		fmt.Println("Remove")
		e := l.Front()
		n, ok := e.Value.(Node)
		if ok {
			delete(m, n.Data.(int))
		}
		l.Remove(e)

		value = 3
		n3 := Node{Data: value}
		l.PushBack(n3)
		m[value] = n3
	}

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Printf("%v\n", m)
}

func main() {
	simualateLRUcache()
}
