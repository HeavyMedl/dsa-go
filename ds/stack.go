package ds

import "fmt"

// Stack represents a generic stack data structure implemented using a linked list.
type Stack[T comparable] struct {
	list *LinkedList[T] // LinkedList used to store stack elements
}

// NewStack creates and initializes a new empty Stack.
// It returns a pointer to the newly created Stack.
func NewStack[T comparable]() *Stack[T] {
	list := NewLinkedList[T]()   // Create a new linked list
	return &Stack[T]{list: list} // Initialize and return the Stack
}

// Size returns the number of elements in the stack.
// It delegates to the Size method of the underlying linked list.
func (stack *Stack[T]) Size() int {
	return stack.list.Size() // Return the size of the linked list
}

// IsEmpty checks if the stack is empty.
// It returns true if the stack has no elements, otherwise false.
func (stack *Stack[T]) IsEmpty() bool {
	return stack.list.Size() == 0 // Check if the linked list is empty
}

// Peek returns the top element of the stack without removing it.
// It returns the head node of the linked list.
func (stack *Stack[T]) Peek() *Node[T] {
	return stack.list.Head() // Return the head node of the linked list
}

// Push adds a new element to the top of the stack.
// It inserts the value at the front of the linked list.
func (stack *Stack[T]) Push(value T) {
	stack.list.InsertAtFront(value) // Insert the value at the front of the linked list
}

// Pop removes and returns the top element of the stack.
// It deletes the head node of the linked list and returns it.
func (stack *Stack[T]) Pop() *Node[T] {
	if !stack.IsEmpty() { // Check if the stack is not empty
		top := stack.list.Head()     // Get the top element
		stack.list.DeleteFromFront() // Remove the top element from the linked list
		top.next = nil               // Disconnect the node from the list
		return top                   // Return the removed node
	}
	return nil // Return nil if the stack is empty
}

// Print outputs the stack elements in a readable format.
func (stack *Stack[T]) Print() {
	var formatted string
	stack.list.ForEach(func(n *Node[T], i int, ll *LinkedList[T]) {
		formatted += fmt.Sprintf("| (%v) |\n", n.Value()) // Format each node's value
	})
	if len(formatted) > 0 {
		formatted += " ----- " // Add separator if stack is not empty
		fmt.Println(formatted)
	} else {
		fmt.Println("Empty Stack") // Indicate if the stack is empty
	}
}
