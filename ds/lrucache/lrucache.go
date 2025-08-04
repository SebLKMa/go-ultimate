package main

import (
	"container/list"
	"fmt"
)

// TODO:
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

func main() {
	// Create a doublylinkedlist and put nodes in it.
	l := list.New()
	n1 := Node{Data: 1}
	l.PushBack(n1)
	n2 := Node{Data: 2}
	l.PushBack(n2)
	//n3 := Node{Data: 3}
	//l.PushBack(n3)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("Remove")
	e := l.Front()
	l.Remove(e)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
