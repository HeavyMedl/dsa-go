package ds

// Node represents an individual element in a singly linked list.
// Each node holds a value of type T and a pointer to the next node in the list.
type Node[T interface{}] struct {
	value T        // The value stored in the node, of generic type T.
	next  *Node[T] // Pointer to the next node in the list, or nil if this is the last node.
}

// Value returns the value stored in the node.
// It provides access to the node's data, allowing it to be retrieved from outside the struct.
func (node *Node[T]) Value() T {
	return node.value
}
