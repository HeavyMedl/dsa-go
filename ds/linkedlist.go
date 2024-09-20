package ds

import "fmt"

// LinkedList represents a singly linked list with generic type T.
type LinkedList[T comparable] struct {
	head *Node[T] // Pointer to the first node
	tail *Node[T] // Pointer to the last node
	size int      // Number of nodes
}

// NewLinkedList initializes a LinkedList.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil, size: 0}
}

// InsertAtFront adds a new node with the given value at the start of the list.
func (list *LinkedList[T]) InsertAtFront(value T) {
	node := &Node[T]{value: value}
	if list.size > 0 {
		node.next = list.head
		list.head = node
	} else {
		list.head = node
		list.tail = node
	}
	list.size++
}

// InsertAtBack appends a new node with the given value to the end of the list.
func (list *LinkedList[T]) InsertAtBack(value T) {
	node := &Node[T]{value: value}
	if list.size > 0 {
		list.tail.next = node
		list.tail = node
	} else {
		list.head = node
		list.tail = node
	}
	list.size++
}

// Get returns the node at the specified index or an error if out of bounds.
func (list *LinkedList[T]) Get(index int) (*Node[T], error) {
	if index < 0 || index >= list.size {
		return nil, fmt.Errorf("index out of bounds")
	}
	if index == 0 {
		return list.head, nil
	}
	if index == list.size-1 {
		return list.tail, nil
	}
	currentNode := list.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode, nil
}

// InsertAt inserts a new node with the given value at the specified index.
func (list *LinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 || index > list.size {
		return fmt.Errorf("index out of bounds")
	}
	if index == 0 {
		list.InsertAtFront(value)
		return nil
	}
	if index == list.size {
		list.InsertAtBack(value)
		return nil
	}
	prevNode, _ := list.Get(index - 1)
	newNode := &Node[T]{value: value, next: prevNode.next}
	prevNode.next = newNode
	list.size++
	return nil
}

// DeleteFromFront removes the node at the start of the list.
func (list *LinkedList[T]) DeleteFromFront() {
	if list.size > 0 {
		list.head = list.head.next
		if list.size == 1 {
			list.tail = nil
		}
		list.size--
	}
}

// DeleteFromBack removes the node at the end of the list.
func (list *LinkedList[T]) DeleteFromBack() {
	if list.size > 0 {
		if list.size == 1 {
			list.head = nil
			list.tail = nil
		} else {
			secondToLast, _ := list.Get(list.size - 2)
			secondToLast.next = nil
			list.tail = secondToLast
		}
		list.size--
	}
}

// DeleteAt removes the node at the specified index.
func (list *LinkedList[T]) DeleteAt(index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("index out of bounds")
	}
	if index == 0 {
		list.DeleteFromFront()
		return nil
	}
	if index == list.size-1 {
		list.DeleteFromBack()
		return nil
	}
	prevNode, _ := list.Get(index - 1)
	prevNode.next = prevNode.next.next
	list.size--
	return nil
}

// Clear removes all nodes from the list.
func (list *LinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Update modifies the value of the node at the specified index.
func (list *LinkedList[T]) Update(index int, value T) (*Node[T], error) {
	if index < 0 || index >= list.size {
		return nil, fmt.Errorf("index out of bounds")
	}
	node, _ := list.Get(index)
	node.value = value
	return node, nil
}

// Search locates the first node with the specified value.
func (list *LinkedList[T]) Search(value T) *Node[T] {
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.value == value {
			return currentNode
		}
	}
	return nil
}

// ForEach applies the callback function to each node in the list.
func (list *LinkedList[T]) ForEach(callback func(*Node[T], int, *LinkedList[T])) {
	for index, currentNode := 0, list.head; currentNode != nil; currentNode = currentNode.next {
		callback(currentNode, index, list)
		index++
	}
}

// Print outputs the list in a readable format.
func (list *LinkedList[T]) Print() {
	formatted := ""
	list.ForEach(func(n *Node[T], i int, ll *LinkedList[T]) {
		formatted += fmt.Sprintf("(%v) -> ", n.Value())
	})
	if len(formatted) > 0 {
		fmt.Println(formatted[:len(formatted)-4]) // Trim the last " -> "
	} else {
		fmt.Println("Empty list")
	}
}

// Head returns the first node in the list.
func (list *LinkedList[T]) Head() *Node[T] {
	return list.head
}

// Tail returns the last node in the list.
func (list *LinkedList[T]) Tail() *Node[T] {
	return list.tail
}

// Size returns the total number of nodes in the list.
func (list *LinkedList[T]) Size() int {
	return list.size
}
