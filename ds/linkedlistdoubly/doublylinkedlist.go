package linkedlistdoubly

// Use
// https://pkg.go.dev/container/list

type Node struct {
	data interface{} // Can hold any data type
	next *Node
	prev *Node
}
