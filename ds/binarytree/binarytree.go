package main

import "fmt"

type BNode struct {
	keyValue int
	left     *BNode
	right    *BNode
}

type BTree struct {
	root *BNode
}

func (t *BTree) Insert(newKey int) {
	if t.root != nil {
		t.insert(newKey, t.root)
		return
	}

	newNode := BNode{
		keyValue: newKey,
		left:     nil,
		right:    nil,
	}
	t.root = &newNode
}

func (t *BTree) Search(key int) *BNode {

	return t.search(key, t.root)
}

func (t *BTree) insert(newKey int, node *BNode) {
	if newKey < node.keyValue {
		if node.left == nil {
			newNode := BNode{
				keyValue: newKey,
				left:     nil,
				right:    nil,
			}
			node.left = &newNode
		} else {
			t.insert(newKey, node.left)
		}
	}
	if newKey >= node.keyValue {
		if node.right == nil {
			newNode := BNode{
				keyValue: newKey,
				left:     nil,
				right:    nil,
			}
			node.right = &newNode
		} else {
			t.insert(newKey, node.right)
		}
	}
}

func (t *BTree) search(key int, node *BNode) *BNode {
	if node == nil {
		return nil // nothing found
	}

	if key == node.keyValue {
		return node // found
	}

	if key < node.keyValue {
		return t.search(key, node.left)
	}

	return t.search(key, node.right)
}

func main() {
	btree := BTree{}
	btree.Insert(67)
	btree.Insert(42)
	btree.Insert(33)
	btree.Insert(86)
	fmt.Printf("My tree:\n%#v\n", btree)

	key := 42
	bnode := btree.Search(key)
	if bnode != nil {
		fmt.Printf("Search found %d My node:\n%#v\n", key, bnode)
	} else {
		fmt.Printf("Search not found %d\n", key)
	}

	key = 1942
	bnode = btree.Search(key)
	if bnode != nil {
		fmt.Printf("Search found %d My node:\n%#v\n", key, bnode)
	} else {
		fmt.Printf("Search not found %d\n", key)
	}
}
